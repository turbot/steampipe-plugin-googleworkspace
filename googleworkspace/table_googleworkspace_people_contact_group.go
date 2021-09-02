package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"google.golang.org/api/people/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspacePeopleContactGroup(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_people_contact_group",
		Description: "Contact groups owned by the authenticated user",
		List: &plugin.ListConfig{
			Hydrate: listPeopleContactGroups,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "max_members",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"404", "403"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "resource_name",
				Description: "The resource name for the contact group, assigned by the server.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The contact group name set by the group owner or a system provided name for system groups.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "group_type",
				Description: "The contact group type.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "formatted_name",
				Description: "The name translated and formatted in the viewer's account locale or the `Accept-Language` HTTP header locale for system groups names.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "deleted",
				Description: "Indicates whether the contact group resource has been deleted, or not.",
				Type:        proto.ColumnType_BOOL,
				Transform:   transform.FromField("Metadata.Deleted"),
			},
			{
				Name:        "max_members",
				Description: "Specifies the maximum number of members to return. Default is 2500, if no value provided.",
				Type:        proto.ColumnType_INT,
				Transform:   transform.FromQual("max_members"),
			},
			{
				Name:        "member_count",
				Description: "The total number of contacts in the group irrespective of max members in specified in the request.",
				Type:        proto.ColumnType_INT,
				Default:     0,
			},
			{
				Name:        "updated_time",
				Description: "The time the group was last updated.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Metadata.UpdateTime"),
			},
			{
				Name:        "client_data",
				Description: "The group's client data.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "member_resource_names",
				Description: "A list of contact person resource names that are members of the contact group.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listPeopleContactGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := PeopleService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default to 2500
	maxMembers := int64(2500)
	if d.KeyColumnQuals["max_members"] != nil {
		maxMembers = d.KeyColumnQuals["max_members"].GetInt64Value()
	}

	// `contactGroups.batchGet` can accept maximum of 200 resource names at a time, so make sure
	// `contactGroups.list` returns the same and append to this in chunks not more then 200.
	pageLimit := int64(200)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < pageLimit {
			pageLimit = *limit
		}
	}

	var count int64
	var contactGroupNames [][]string
	resp := service.ContactGroups.List().PageSize(pageLimit)
	if err := resp.Pages(ctx, func(page *people.ListContactGroupsResponse) error {
		var resourceNames []string
		// create a chunk of resourceNames of size 200
		for _, contactGroup := range page.ContactGroups {
			resourceNames = append(resourceNames, contactGroup.ResourceName)
			count++

			// Check if the context is cancelled for query
			// Break for loop if requested no of results achieved
			if plugin.IsCancelled(ctx) || (limit != nil && count >= *limit) {
				page.NextPageToken = ""
				break
			}
		}
		if len(resourceNames) > 0 {
			contactGroupNames = append(contactGroupNames, resourceNames)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// execute batchGet
	for _, contactGroups := range contactGroupNames {
		data, err := service.ContactGroups.BatchGet().ResourceNames(contactGroups...).MaxMembers(maxMembers).Do()
		if err != nil {
			return nil, err
		}
		if data.Responses != nil && len(data.Responses) > 0 {
			for _, i := range data.Responses {
				d.StreamListItem(ctx, i.ContactGroup)
			}
		}
	}

	return nil, nil
}
