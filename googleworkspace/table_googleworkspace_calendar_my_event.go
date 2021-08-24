package googleworkspace

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/plugin"

	"google.golang.org/api/calendar/v3"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceCalendarMyEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_calendar_my_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate:           listCalendarMyEvents,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "query",
					Require: plugin.Optional,
				},
				{
					Name:      "start_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:      "end_time",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
			},
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listCalendarMyEvents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := CalendarService(ctx, d)
	if err != nil {
		return nil, err
	}

	// By default, API can return maximum 2500 records in a single page
	maxResult := int64(2500)
	// Reduce the basic request limit down if the user has only requested a small number of rows
	if d.QueryContext.Limit != nil {
		limit := d.QueryContext.Limit
		if *limit < maxResult {
			maxResult = *limit
		}
	}

	// Free text search terms to find events that match these terms in any field, except for extended properties
	var query string
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}

	resp := service.Events.List("primary").SingleEvents(true).Q(query).MaxResults(maxResult)
	// Additional filter queries using timestamp (if provided)
	quals := d.Quals
	if quals["start_time"] != nil {
		for _, q := range quals["start_time"].Quals {
			startTime := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			switch q.Operator {
			case ">=", ">", "=", "<", "<=":
				resp.TimeMin(startTime)
			}
		}
	}
	if quals["end_time"] != nil {
		for _, q := range quals["end_time"].Quals {
			endTime := q.Value.GetTimestampValue().AsTime().Format(time.RFC3339)
			switch q.Operator {
			case ">=", ">", "=", "<", "<=":
				resp.TimeMax(endTime)
			}
		}
	}
	if err := resp.Pages(ctx, func(page *calendar.Events) error {
		for _, event := range page.Items {
			d.StreamListItem(ctx, calendarEvent{*event, page.Summary})
		}
		return nil
	}); err != nil {
		if IsForbiddenError(err) {
			return nil, nil
		}
		return nil, err
	}

	return nil, err
}
