package googleworkspace

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

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
	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < maxResult {
			maxResult = *limit
		}
	}

	// Free text search terms to find events that match these terms in any field, except for extended properties
	var query string
	if d.EqualsQuals["query"] != nil {
		query = d.EqualsQuals["query"].GetStringValue()
	}

	resp := service.Events.List("primary").ShowDeleted(false).SingleEvents(true).Q(query).MaxResults(maxResult)
	if d.Quals["start_time"] != nil {
		for _, q := range d.Quals["start_time"].Quals {
			givenTime := q.Value.GetTimestampValue().AsTime()
			beforeTime := givenTime.Add(time.Duration(-1) * time.Second).Format("2006-01-02T15:04:05.000Z")
			afterTime := givenTime.Add(time.Second * 1).Format("2006-01-02T15:04:05.000Z")

			switch q.Operator {
			case ">":
				resp.TimeMin(afterTime)
			case ">=":
				resp.TimeMin(givenTime.Format("2006-01-02T15:04:05.000Z"))
			case "=":
				resp.TimeMin(givenTime.Format("2006-01-02T15:04:05.000Z")).TimeMax(givenTime.Format("2006-01-02T15:04:05.000Z"))
			case "<=":
				resp.TimeMax(givenTime.Format("2006-01-02T15:04:05.000Z"))
			case "<":
				resp.TimeMax(beforeTime)
			}
		}
	}
	if err := resp.Pages(ctx, func(page *calendar.Events) error {
		for _, event := range page.Items {
			d.StreamListItem(ctx, calendarEvent{*event, page.Summary})

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
