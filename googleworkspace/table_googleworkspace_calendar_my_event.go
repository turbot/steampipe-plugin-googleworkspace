package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"

	"google.golang.org/api/calendar/v3"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceCalendarMyEvents(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_calendar_my_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate:           listCalendarMyEvents,
			ShouldIgnoreError: isNotFoundError([]string{"404", "400"}),
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "query",
					Require: plugin.Optional,
				},
				{
					Name:    "start_time",
					Require: plugin.Optional,
				},
				{
					Name:    "end_time",
					Require: plugin.Optional,
				},
			},
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listCalendarMyEvents(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := CalendarService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp := service.Events.List("primary")

	// Additional filter functions, invokes when optional filter columns are specified
	if d.KeyColumnQuals["query"] != nil {
		query := d.KeyColumnQuals["query"].GetStringValue()
		resp.Q(query)
	}
	if d.Quals["start_time"] != nil {
		startTime := d.KeyColumnQuals["start_time"].GetTimestampValue().AsTime().Format("2006-01-02T15:04:05.000Z")
		resp.TimeMax(startTime)
	}
	if d.Quals["end_time"] != nil {
		endTime := d.KeyColumnQuals["end_time"].GetTimestampValue().AsTime().Format("2006-01-02T15:04:05.000Z")
		resp.TimeMin(endTime)
	}

	if err := resp.Pages(ctx, func(page *calendar.Events) error {
		for _, event := range page.Items {
			d.StreamListItem(ctx, calendarEvent{*event, page.Summary})
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, err
}
