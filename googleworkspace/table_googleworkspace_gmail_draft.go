package googleworkspace

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"google.golang.org/api/gmail/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceGmailDraft(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_gmail_draft",
		Description: "Retrieves draft messages in the specified user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listGmailDrafts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Required,
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"draft_id", "user_id"}),
			Hydrate:    getGmailDraft,
		},
		Columns: []*plugin.Column{
			{
				Name:        "draft_id",
				Description: "The immutable ID of the draft.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Id"),
			},
			{
				Name:        "message_id",
				Description: "The immutable ID of the message.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Message.Id"),
			},
			{
				Name:        "message_thread_id",
				Description: "The ID of the thread the message belongs to.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Message.ThreadId"),
			},
			{
				Name:        "user_id",
				Description: "User's email address. If not specified, indicates the current authenticated user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("user_id"),
			},
			{
				Name:        "message_history_id",
				Description: "The ID of the last history record that modified this message.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.HistoryId"),
			},
			{
				Name:        "message_internal_date",
				Description: "The internal message creation timestamp which determines ordering in the inbox.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.InternalDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "message_raw",
				Description: "The entire email message in an RFC 2822 formatted and base64url encoded string.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.Raw").NullIfZero(),
			},
			{
				Name:        "message_size_estimate",
				Description: "Estimated size in bytes of the message.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.SizeEstimate"),
			},
			{
				Name:        "message_snippet",
				Description: "A short part of the message text.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.Snippet").NullIfZero(),
			},
			{
				Name:        "query",
				Description: "A string to filter messages matching the specified query.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "message_label_ids",
				Description: "A list of IDs of labels applied to this message.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.LabelIds"),
			},
			{
				Name:        "message_payload",
				Description: "The parsed email structure in the message parts.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailDraft,
				Transform:   transform.FromField("Message.Payload"),
			},
		},
	}
}

//// LIST FUNCTION

func listGmailDrafts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var userID string
	if d.EqualsQuals["user_id"] != nil {
		userID = d.EqualsQuals["user_id"].GetStringValue()
	}

	var queryFilter, query string
	var filter []string

	if d.Quals["message_internal_date"] != nil {
		for _, q := range d.Quals["message_internal_date"].Quals {
			tsSecs := q.Value.GetTimestampValue().GetSeconds()
			switch q.Operator {
			case "=":
				filter = append(filter, fmt.Sprintf("after:%s before:%s", strconv.Itoa(int(tsSecs)), strconv.Itoa(int(tsSecs+1))))
			case ">=":
				filter = append(filter, fmt.Sprintf("after:%s", strconv.Itoa(int(tsSecs))))
			case ">":
				filter = append(filter, fmt.Sprintf("after:%s", strconv.Itoa(int(tsSecs))))
			case "<=":
				filter = append(filter, fmt.Sprintf("before:%s", strconv.Itoa(int(tsSecs)+1)))
			case "<":
				filter = append(filter, fmt.Sprintf("before:%s", strconv.Itoa(int(tsSecs))))
			}
		}
	}

	// Only return messages matching the specified query. Supports the same query format as the Gmail search box.
	// For example, "from:someuser@example.com is:unread"
	if d.EqualsQuals["query"] != nil {
		queryFilter = d.EqualsQuals["query"].GetStringValue()
	}

	if queryFilter != "" {
		query = queryFilter
	} else if len(filter) > 0 {
		query = strings.Join(filter, " and ")
	}

	// Setting the maximum number of messages, API can return in a single page
	maxResults := int64(500)

	limit := d.QueryContext.Limit
	if d.QueryContext.Limit != nil {
		if *limit < maxResults {
			maxResults = *limit
		}
	}

	resp := service.Users.Drafts.List(userID).Q(query).MaxResults(maxResults)
	if err := resp.Pages(ctx, func(page *gmail.ListDraftsResponse) error {
		for _, draft := range page.Drafts {
			d.StreamListItem(ctx, draft)

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

func getGmailDraft(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var userID string
	if d.EqualsQuals["user_id"] != nil {
		userID = d.EqualsQuals["user_id"].GetStringValue()
	}

	var draftID string
	if h.Item != nil {
		draftID = h.Item.(*gmail.Draft).Id
	} else {
		draftID = d.EqualsQuals["draft_id"].GetStringValue()
	}

	// Return nil, if no input provided
	if draftID == "" || userID == "" {
		return nil, nil
	}

	resp, err := service.Users.Drafts.Get(userID, draftID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
