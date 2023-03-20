package googleworkspace

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceDrive(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_drive",
		Description: "Drives defined user's shared drives in the Google Drive.",
		List: &plugin.ListConfig{
			Hydrate: listDrives,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:      "created_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:    "use_domain_admin_access",
					Require: plugin.Optional,
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDrive,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of this shared drive which is also the ID of the top level folder of this shared drive.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The name of this shared drive.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_time",
				Description: "The time at which the shared drive was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "hidden",
				Description: "Indicates whether the shared drive is hidden from default view, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "admin_managed_restrictions",
				Description: "Indicates whether administrative privileges on this shared drive are required to modify restrictions, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Restrictions.AdminManagedRestrictions"),
			},
			{
				Name:        "background_image_link",
				Description: "A short-lived link to this shared drive's background image.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "color_rgb",
				Description: "The color of this shared drive as an RGB hex string.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "copy_requires_writer_permission",
				Description: "Indicates whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and commenters, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Restrictions.CopyRequiresWriterPermission"),
			},
			{
				Name:        "domain_users_only",
				Description: "Indicates whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which this shared drive belongs.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Restrictions.DomainUsersOnly"),
			},
			{
				Name:        "drive_members_only",
				Description: "Indicates whether access to items inside this shared drive is restricted to its members, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Restrictions.DriveMembersOnly"),
			},
			{
				Name:        "theme_id",
				Description: "The ID of the theme from which the background image and color will be set.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "use_domain_admin_access",
				Description: "Issue the request as a domain administrator; if set to true, then all shared drives of the domain in which the requester is an administrator are returned. Please refer Refer https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromQual("use_domain_admin_access"),
			},
			{
				Name:        "query",
				Description: "Query string for [searching](https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties) shared drives.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "background_image_file",
				Description: "An image file and cropping parameters from which a background image for this shared drive is set.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "capabilities",
				Description: "Describes the capabilities the current user has on this shared drive.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDrives(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := DriveService(ctx, d)
	if err != nil {
		return nil, err
	}
	equalQuals := d.EqualsQuals

	var queryFilter, query string
	var filter []string

	if equalQuals["name"] != nil {
		filter = append(filter, fmt.Sprintf("%s = \"%s\"", "name", equalQuals["name"].GetStringValue()))
	}

	if d.Quals["created_time"] != nil {
		for _, q := range d.Quals["created_time"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime()
			beforeTime := givenTime.Add(time.Duration(-1) * time.Second).Format("2006-01-02T15:04:05.000Z")
			afterTime := givenTime.Add(time.Second * 1).Format("2006-01-02T15:04:05.000Z")

			// Since, the query filter matches the actual time
			switch q.Operator {
			case ">", "<":
				filter = append(filter, fmt.Sprintf("%s %s \"%s\"", "createdTime", q.Operator, givenTime.Format("2006-01-02T15:04:05.000Z")))
			case "=":
				filter = append(filter, fmt.Sprintf("createdTime > \"%s\" and createdTime < \"%s\"", beforeTime, afterTime))
			case ">=":
				filter = append(filter, fmt.Sprintf("%s > \"%s\"", "createdTime", beforeTime))
			case "<=":
				filter = append(filter, fmt.Sprintf("%s < \"%s\"", "createdTime", afterTime))
			}
		}
	}

	// Query string for searching shared drives. Refer https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties
	// For example, "hidden=true"
	if equalQuals["query"] != nil {
		queryFilter = equalQuals["query"].GetStringValue()
	}

	if queryFilter != "" {
		query = queryFilter
	} else if len(filter) > 0 {
		query = strings.Join(filter, " and ")
	}

	// Check for query context and requests only for queried columns
	givenColumns := d.QueryContext.Columns
	requiredFields := buildDriveRequestFields(ctx, givenColumns)

	// Set default as false
	// Need to set true for some of the query terms, i.e. when filtering using createdTime, memberCount, name, or organizerCount
	// Refer https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties
	var useDomainAdminAccess bool
	if equalQuals["use_domain_admin_access"] != nil {
		useDomainAdminAccess = equalQuals["use_domain_admin_access"].GetBoolValue()
	}

	// By default, API can return maximum 100 records in a single page
	pageSize := int64(100)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < pageSize {
			pageSize = *limit
		}
	}

	resp := service.Drives.List().Fields(requiredFields...).Q(query).UseDomainAdminAccess(useDomainAdminAccess).PageSize(pageSize)
	if err := resp.Pages(ctx, func(page *drive.DriveList) error {
		for _, data := range page.Drives {
			parsedTime, _ := time.Parse(time.RFC3339, data.CreatedTime)
			data.CreatedTime = parsedTime.Format(time.RFC3339)
			d.StreamListItem(ctx, data)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if plugin.IsCancelled(ctx) {
				page.NextPageToken = ""
				break
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDrive(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDrive")

	// Create service
	service, err := DriveService(ctx, d)
	if err != nil {
		return nil, err
	}
	id := d.EqualsQuals["id"].GetStringValue()

	// Return nil, if no input provided
	if id == "" {
		return nil, nil
	}

	// Check for query context and requests only for queried columns
	givenColumns := d.QueryContext.Columns
	requiredFields := buildDriveRequestFields(ctx, givenColumns)

	resp, err := service.Drives.Get(id).Fields(requiredFields...).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// buildDriveRequestFields :: Return columns passed in query context
func buildDriveRequestFields(ctx context.Context, queryColumns []string) []googleapi.Field {
	var fields []string
	var requestedFields []googleapi.Field

	// Since ID is unique, always add in the requested field
	if !helpers.StringSliceContains(queryColumns, "id") {
		queryColumns = append(queryColumns, "id")
	}

	for _, columnName := range queryColumns {
		// Optional columns
		if columnName == "query" || columnName == "use_domain_admin_access" || columnName == "_ctx" {
			continue
		}

		formattedColumnName := strcase.ToLowerCamel(columnName)
		switch columnName {
		case "admin_managed_restrictions", "copy_requires_writer_permission", "domain_users_only", "drive_members_only":
			fields = append(fields, "restrictions/"+formattedColumnName)
		default:
			fields = append(fields, formattedColumnName)
		}
	}

	givenFields := strings.Join(fields, ", ")
	requestedFields = append(requestedFields, googleapi.Field(fmt.Sprintf("nextPageToken, drives(%s)", givenFields)))

	return requestedFields
}
