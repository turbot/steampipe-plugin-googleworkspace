package googleworkspace

import (
    "context"
    "fmt"
    "time"

    "github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin"
    "github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
    "google.golang.org/api/admin/reports/v1"
)

//// TABLE DEFINITION

func tableGoogleworkspaceAdminReportsActivity(ctx context.Context) *plugin.Table {
    return &plugin.Table{
        Name:        "googleworkspace_admin_reports_activity",
        Description: "Google Workspace Admin Reports API",

        List: &plugin.ListConfig{
            Hydrate: listGoogleworkspaceAdminReportsActivities,
            KeyColumns: plugin.KeyColumnSlice{
                {Name: "application_name", Require: plugin.Required},
                {Name: "time", Require: plugin.Optional, Operators: []string{">", ">=", "<", "<=", "="}},
                {Name: "actor_email", Require: plugin.Optional},
                {Name: "ip_address", Require: plugin.Optional},
                {Name: "event_name", Require: plugin.Optional},
            },
            Tags: map[string]string{"service": "admin", "product": "reports", "action": "activities.list"},
        },
        Columns: []*plugin.Column{
            {
                Name:        "time",
                Description: "Timestamp of the activity (Id.Time) in RFC3339 format",
                Type:        proto.ColumnType_TIMESTAMP,
                Transform:   transform.FromField("Id.Time"),
            },
            {
                Name:        "actor_email",
                Description: "Email address of the actor (Actor.Email)",
                Type:        proto.ColumnType_STRING,
                Transform:   transform.FromField("Actor.Email"),
            },
            {
                Name:        "event_name",
                Description: "Event name (if queried)",
                Type:        proto.ColumnType_STRING,
                Hydrate:     getEventName,
                Transform:   transform.FromValue(),
            },
            {
                Name:        "event_names",
                Description: "List of event names for this activity",
                Type:        proto.ColumnType_JSON,
                Transform:   transform.FromField("Events").Transform(extractEventNames),
            },
            {
                Name:        "unique_qualifier",
                Description: "Unique qualifier ID for this activity",
                Type:        proto.ColumnType_STRING,
                Transform:   transform.FromField("Id.UniqueQualifier"),
            },
            {
                Name:        "application_name",
                Description: "Name of the report application (Id.ApplicationName)",
                Type:        proto.ColumnType_STRING,
                Transform:   transform.FromField("Id.ApplicationName"),
            },
            {
                Name:        "ip_address",
                Description: "IP address associated with the activity (IpAddress)",
                Type:        proto.ColumnType_STRING,
            },
            {
                 Name:        "actor_profile_id",
                 Type:        proto.ColumnType_STRING,
                 Transform:   transform.FromField("Actor.ProfileId"),
            },
            {
                Name:        "events",
                Description: "Full JSON array of detailed events (Events)",
                Type:        proto.ColumnType_JSON,
            },
        },
    }
}

//// LIST FUNCTION

func listGoogleworkspaceAdminReportsActivities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
    service, err := ReportsService(ctx, d)
    if err != nil {
        plugin.Logger(ctx).Error("googleworkspace_admin_reports_activity.list", "service_error", err)
        return nil, err
    }

    // Required application_name qualifier
    appName := d.EqualsQualString("application_name")

    // Validate application_name
    valid := map[string]bool{
        "access_transparency":          true,
        "admin":                        true,
        "calendar":                     true,
        "chat":                         true,
        "drive":                        true,
        "gcp":                          true,
        "gplus":                        true,
        "groups":                       true,
        "groups_enterprise":            true,
        "jamboard":                     true,
        "login":                        true,
        "meet":                         true,
        "mobile":                       true,
        "rules":                        true,
        "saml":                         true,
        "token":                        true,
        "user_accounts":                true,
        "context_aware_access":         true,
        "chrome":                       true,
        "data_studio":                  true,
        "keep":                         true,
        "vault":                        true,
        "gemini_in_workspace_apps":     true,
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
                case ">":
                    startTime = t.Add(time.Nanosecond)
                case ">=":
                    startTime = t
                case "<":
                    endTime = t
                case "<=":
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

    if qual := d.EqualsQualString("ip_address"); qual != "" {
        resp = resp.ActorIpAddress(qual)
    }

    if qual := d.EqualsQualString("event_name"); qual != "" {
        resp = resp.EventName(qual)
    }

    err = resp.Pages(ctx, func(page *admin.Activities) error {
        // rate limit
        d.WaitForListRateLimit(ctx)

        for _, activity := range page.Items {
            d.StreamListItem(ctx, activity)
            if d.RowsRemaining(ctx) == 0 {
                page.NextPageToken = ""
                break
            }
        }
        return nil
    })
    if err != nil {
        plugin.Logger(ctx).Error("googleworkspace_admin_reports_activity.list", "api_error", err)
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

func getEventName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
    // Did the user filter on event_name?
    if qual := d.EqualsQualString("event_name"); qual != "" {
        return qual, nil
    }

    // API response
    activity, ok := h.Item.(*admin.Activity)
    if !ok {
        return nil, nil
    }
    if activity.Events == nil {
        return nil, nil
    }
    names := []string{}
    for _, e := range activity.Events {
        if e.Name != "" {
            return e.Name, nil
        }
    }
    return names, nil
}