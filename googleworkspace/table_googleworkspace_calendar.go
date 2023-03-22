package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceCalendar(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_calendar",
		Description: "Metadata of the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate:           listCalendars,
			KeyColumns:        plugin.SingleColumn("id"),
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "Identifier of the calendar.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "summary",
				Description: "Title of the calendar.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "timezone",
				Description: "The time zone of the calendar.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("TimeZone"),
			},
			{
				Name:        "description",
				Description: "Description of the calendar.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "etag",
				Description: "ETag of the resource.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "location",
				Description: "Geographic location of the calendar as free-form text.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "conference_properties",
				Description: "Describes the conferencing properties for this calendar.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listCalendars(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := CalendarService(ctx, d)
	if err != nil {
		return nil, err
	}
	calendarID := d.EqualsQuals["id"].GetStringValue()

	resp, err := service.Calendars.Get(calendarID).Do()
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, resp)

	return nil, nil
}
