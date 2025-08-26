package googleworkspace

import (
	"context"
	"fmt"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/query_cache"
	admin "google.golang.org/api/admin/reports/v1"
)

//// TABLE DEFINITION

func tableGoogleworkspaceActivityReport(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_activity_report",
		Description: "Google Workspace Activity Report",

		List: &plugin.ListConfig{
			Hydrate: listGoogleworkspaceAdminReportsActivities,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "application_name", Require: plugin.Required},
				{Name: "time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
				{Name: "actor_email", Require: plugin.Optional},
				{Name: "ip_address", Require: plugin.Optional},
				{Name: "event_name", Require: plugin.Optional, CacheMatch: query_cache.CacheMatchExact},
			},
			Tags: map[string]string{"service": "admin", "product": "reports", "action": "activities.list"},
		},
		Columns: []*plugin.Column{
			{
				Name:        "time",
				Description: "Time of occurrence of the activity.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Id.Time"),
			},
			{
				Name:        "actor_email",
				Description: "Email address of the actor.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Actor.Email"),
			},
			{
				Name:        "event_name",
				Description: "The name of the event (if queried).",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("event_name"),
			},
			{
				Name:        "event_names",
				Description: "List of event names for this activity.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Events").Transform(extractEventNames),
			},
			{
				Name:        "unique_qualifier",
				Description: "Unique qualifier ID for this activity.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id.UniqueQualifier"),
			},
			{
				Name:        "application_name",
				Description: "Application name to which the event belongs.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id.ApplicationName"),
			},
			{
				Name:        "ip_address",
				Description: "IP address associated with the activity.",
				Type:        proto.ColumnType_IPADDR,
			},
			{
				Name:        "actor_profile_id",
				Description: "The unique Google Workspace profile ID of the actor.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Actor.ProfileId"),
			},
			{
				Name:        "customer_id",
				Description: "The unique ID of the customer to retrieve data for.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id.CustomerId"),
			},
			{
				Name:        "events",
				Description: "Activity events in the report.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listGoogleworkspaceAdminReportsActivities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// https://developers.google.com/admin-sdk/reports/v1/reference/activities/list#authorization-scopes
	service, err := ReportsServiceWithScope(ctx, d, admin.AdminReportsAuditReadonlyScope)
	if err != nil {
		plugin.Logger(ctx).Error("googleworkspace_activity_report.listGoogleworkspaceAdminReportsActivities", "service_error", err)
		return nil, err
	}

	// Required application_name qualifier
	appName := d.EqualsQualString("application_name")

	// Validate application_name
	valid := map[string]bool{
		"access_transparency":      true,
		"admin":                    true,
		"calendar":                 true,
		"chat":                     true,
		"drive":                    true,
		"gcp":                      true,
		"gplus":                    true,
		"groups":                   true,
		"groups_enterprise":        true,
		"jamboard":                 true,
		"login":                    true,
		"meet":                     true,
		"mobile":                   true,
		"rules":                    true,
		"saml":                     true,
		"token":                    true,
		"user_accounts":            true,
		"context_aware_access":     true,
		"chrome":                   true,
		"data_studio":              true,
		"keep":                     true,
		"vault":                    true,
		"gemini_in_workspace_apps": true,
	}

	if !valid[appName] {
		return nil, fmt.Errorf("unsupported application_name: %q", appName)
	}

	// Setting the maximum number of activities, API can return in a single page
	maxResults := int64(1000)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < maxResults {
			maxResults = *limit
		}
	}

	// Determine userKey: default to "all", or use the actor_email if provided
	userKey := "all"
	if ae := d.EqualsQualString("actor_email"); ae != "" {
		userKey = ae
	}

	// Build API call with the chosen category
	resp := service.Activities.List(userKey, appName).MaxResults(maxResults)

	if quals := d.Quals["time"]; quals != nil {
		var startTime, endTime time.Time
		for _, q := range quals.Quals {
			if ts := q.Value.GetTimestampValue(); ts != nil {
				t := ts.AsTime()
				switch q.Operator {
				case "=":
					startTime, endTime = t, t
				case ">", ">=":
					startTime = t.Add(time.Nanosecond)
				case "<", "<=":
					endTime = t
				}
			}
		}
		if !startTime.IsZero() {
			resp.StartTime(startTime.Format(time.RFC3339))
		}
		if !endTime.IsZero() {
			resp.EndTime(endTime.Format(time.RFC3339))
		}
	}

	if qual := d.EqualsQuals["ip_address"]; qual != nil {
		address := qual.GetInetValue().GetAddr()
		resp = resp.ActorIpAddress(address)
	}

	if qual := d.EqualsQualString("event_name"); qual != "" {
		resp = resp.EventName(qual)
	}

	err = resp.Pages(ctx, func(page *admin.Activities) error {
		// rate limit
		d.WaitForListRateLimit(ctx)

		for _, activity := range page.Items {
			d.StreamListItem(ctx, activity)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				page.NextPageToken = ""
				break
			}
		}
		return nil
	})
	if err != nil {
		plugin.Logger(ctx).Error("googleworkspace_activity_report.listGoogleworkspaceAdminReportsActivities", "api_error", err)
		return nil, err
	}

	return nil, nil
}

//// TRANSFORM FUNCTIONS

func extractEventNames(_ context.Context, d *transform.TransformData) (interface{}, error) {
	activity, ok := d.HydrateItem.(*admin.Activity)
	if !ok {
		return nil, nil
	}
	if activity.Events == nil {
		return nil, nil
	}
	names := []string{}
	for _, e := range activity.Events {
		if e.Name != "" {
			names = append(names, e.Name)
		}
	}
	return names, nil
}
