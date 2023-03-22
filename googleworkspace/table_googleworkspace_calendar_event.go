package googleworkspace

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"google.golang.org/api/calendar/v3"
)

type calendarEvent = struct {
	calendar.Event
	CalendarId string
}

func calendarEventColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "id",
			Description: "Opaque identifier of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "summary",
			Description: "Specifies the title of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "status",
			Description: "Status of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "calendar_id",
			Description: "Identifier of the calendar.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "start_time",
			Description: "Specifies the event start time.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromP(formatTimestamp, "StartTime").NullIfZero(),
		},
		{
			Name:        "end_time",
			Description: "Specifies the event end time.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromP(formatTimestamp, "EndTime").NullIfZero(),
		},
		{
			Name:        "day",
			Description: "Specifies the day of a week.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromP(formatTimestamp, "Day").NullIfZero(),
		},
		{
			Name:        "hangout_link",
			Description: "An absolute link to the Google Hangout associated with this event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "event_type",
			Description: "Specifies the type of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "html_link",
			Description: "An absolute link to this event in the Google Calendar Web UI.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "attendees_omitted",
			Description: "Indicates whether attendees may have been omitted from the event's representation, or not.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "color_id",
			Description: "The color of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "created_at",
			Description: "Creation time of the event.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Created").NullIfZero(),
		},
		{
			Name:        "description",
			Description: "A short user-defined description of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "end_time_unspecified",
			Description: "Indicates whether the end time is actually unspecified, or not.",
			Type:        proto.ColumnType_BOOL,
		},
		{
			Name:        "etag",
			Description: "ETag of the resource.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "guests_can_invite_others",
			Description: "Indicates whether attendees other than the organizer can invite others to the event, or not.",
			Type:        proto.ColumnType_BOOL,
			Default:     true,
		},
		{
			Name:        "guests_can_modify",
			Description: "Indicates whether attendees other than the organizer can modify the event, or not.",
			Type:        proto.ColumnType_BOOL,
			Default:     false,
		},
		{
			Name:        "guests_can_see_other_guests",
			Description: "Indicates whether attendees other than the organizer can modify the event, or not.",
			Type:        proto.ColumnType_BOOL,
			Default:     true,
		},
		{
			Name:        "ical_uid",
			Description: "Specifies the event unique identifier as defined in RFC5545. It is used to uniquely identify events accross calendaring systems and must be supplied when importing events via the import method.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromField("ICalUID"),
		},
		{
			Name:        "location",
			Description: "Geographic location of the event as free-form text.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "locked",
			Description: "Indicates whether this is a locked event copy where no changes can be made to the main event fields \"summary\", \"description\", \"location\", \"start\", \"end\" or \"recurrence\".",
			Type:        proto.ColumnType_BOOL,
			Default:     false,
		},
		{
			Name:        "private_copy",
			Description: "Indicates whether event propagation is disabled, or not.",
			Type:        proto.ColumnType_BOOL,
			Default:     false,
		},
		{
			Name:        "query",
			Description: "Filter string to filter events.",
			Type:        proto.ColumnType_STRING,
			Transform:   transform.FromQual("query"),
		},
		{
			Name:        "recurring_event_id",
			Description: "For an instance of a recurring event, this is the id of the recurring event to which this instance belongs.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "sequence",
			Description: "Sequence number as per iCalendar.",
			Type:        proto.ColumnType_INT,
		},
		{
			Name:        "timezone",
			Description: "The time zone of the calendar.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "transparency",
			Description: "Indicates whether the event blocks time on the calendar.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "updated_at",
			Description: "Last modification time of the event.",
			Type:        proto.ColumnType_TIMESTAMP,
			Transform:   transform.FromField("Updated").NullIfZero(),
		},
		{
			Name:        "visibility",
			Description: "Visibility of the event.",
			Type:        proto.ColumnType_STRING,
		},
		{
			Name:        "attachments",
			Description: "A list of file attachments for the event.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "attendees",
			Description: "A list of attendees of the event.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "conference_data",
			Description: "The conference-related information, such as details of a Google Meet conference.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "creator",
			Description: "Specifies the creator details of the event.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "extended_properties",
			Description: "A list of extended properties of the event.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "organizer",
			Description: "Specifies the organizer details of the event.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "original_start_time",
			Description: "For an instance of a recurring event, this is the time at which this event would start according to the recurrence data in the recurring event identified by recurringEventId.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "recurrence",
			Description: "A list of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "reminders",
			Description: "Information about the event's reminders for the authenticated user.",
			Type:        proto.ColumnType_JSON,
		},
		{
			Name:        "source",
			Description: "Source from which the event was created.",
			Type:        proto.ColumnType_JSON,
		},
	}
}

//// TABLE DEFINITION

func tableGoogleWorkspaceCalendarEvent(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_calendar_event",
		Description: "Events scheduled on the specified calendar.",
		List: &plugin.ListConfig{
			Hydrate:           listCalendarEvents,
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "calendar_id",
					Require: plugin.Required,
				},
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
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"calendar_id", "id"}),
			Hydrate:    getCalendarEvent,
		},
		Columns: calendarEventColumns(),
	}
}

//// LIST FUNCTION

func listCalendarEvents(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := CalendarService(ctx, d)
	if err != nil {
		return nil, err
	}
	calendarID := d.EqualsQuals["calendar_id"].GetStringValue()

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

	resp := service.Events.List(calendarID).ShowDeleted(false).SingleEvents(true).Q(query).MaxResults(maxResult)
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
			d.StreamListItem(ctx, calendarEvent{*event, calendarID})

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

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getCalendarEvent(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := CalendarService(ctx, d)
	if err != nil {
		return nil, err
	}
	calendarID := d.EqualsQuals["calendar_id"].GetStringValue()
	eventID := d.EqualsQuals["id"].GetStringValue()

	// Return nil, if no input provided
	if calendarID == "" || eventID == "" {
		return nil, nil
	}

	resp, err := service.Events.Get(calendarID, eventID).Do()
	if err != nil {
		return nil, err
	}

	return calendarEvent{*resp, calendarID}, err
}

//// TRANSFORM FUNCTIONS

func formatTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(calendarEvent)
	param := d.Param.(string)
	var startTime, endTime string

	// handling empty data
	if data.Start == nil || data.End == nil {
		return nil, nil
	}

	// If the event is an all-day event, response includes only
	// `Start.DateTime` and `End.DateTime`, and the value consists
	// only date, for example `2021-08-15`
	// Following transformation is used to parse the date in consistent format, i.e. RFC3339
	startTime = data.Start.DateTime
	if startTime == "" {
		startTime = parseTime(data.Start.Date)
	}

	endTime = data.End.DateTime
	if endTime == "" {
		endTime = parseTime(data.End.Date)
	}

	t, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return nil, err
	}

	formattedTime := map[string]string{
		"StartTime": startTime,
		"EndTime":   endTime,
		"Day":       t.Weekday().String(),
	}

	return formattedTime[param], nil
}

func parseTime(timeInString string) string {
	if timeInString == "" {
		return ""
	}
	parsedDate, err := time.Parse("2006-01-02", timeInString)
	if err != nil {
		return ""
	}
	return parsedDate.Format("2006-01-02T15:04:05.000Z")
}
