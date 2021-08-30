package googleworkspace

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"google.golang.org/api/gmail/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceGmailMyMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_gmail_my_message",
		Description: "Retrieves messages in the current authenticated user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listGmailMyMessages,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "sender_email",
					Require: plugin.Optional,
				},
				{
					Name:      "internal_date",
					Require:   plugin.Optional,
					Operators: []string{">", ">=", "=", "<", "<="},
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"403"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getGmailMyMessage,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:           getGmailMyMessage,
				MaxConcurrency: 50,
			},
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The immutable ID of the message.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "thread_id",
				Description: "The ID of the thread the message belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "history_id",
				Description: "The ID of the last history record that modified this message.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMyMessage,
			},
			{
				Name:        "sender_email",
				Description: "Specifies the email address of the sender.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMyMessage,
				Transform:   transform.From(extractMessageSender),
			},
			{
				Name:        "internal_date",
				Description: "The internal message creation timestamp which determines ordering in the inbox.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getGmailMyMessage,
				Transform:   transform.FromField("InternalDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "raw",
				Description: "The entire email message in an RFC 2822 formatted and base64url encoded string.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMyMessage,
			},
			{
				Name:        "size_estimate",
				Description: "Estimated size in bytes of the message.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getGmailMyMessage,
			},
			{
				Name:        "snippet",
				Description: "A short part of the message text.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMyMessage,
			},
			{
				Name:        "query",
				Description: "A string to filter messages matching the specified query.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "label_ids",
				Description: "A list of IDs of labels applied to this message.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyMessage,
			},
			{
				Name:        "payload",
				Description: "The parsed email structure in the message parts.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyMessage,
			},
		},
	}
}

//// LIST FUNCTION

func listGmailMyMessages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var queryFilter, query string
	var filter []string

	if d.KeyColumnQuals["sender_email"] != nil {
		filter = append(filter, fmt.Sprintf("%s = \"%s\"", "from", d.KeyColumnQuals["sender_email"].GetStringValue()))
	}

	if d.Quals["internal_date"] != nil {
		for _, q := range d.Quals["internal_date"].Quals {
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
	// Note: Parameter cannot be used when accessing the api using the gmail.metadata scope.
	if d.KeyColumnQuals["query"] != nil {
		queryFilter = d.KeyColumnQuals["query"].GetStringValue()
	}

	if queryFilter != "" {
		query = queryFilter
	} else if len(filter) > 0 {
		query = strings.Join(filter, " and ")
	}

	// Setting the maximum number of messages, API can return in a single page
	maxResults := int64(500)

	resp := service.Users.Messages.List("me").Q(query).MaxResults(maxResults)
	if err := resp.Pages(ctx, func(page *gmail.ListMessagesResponse) error {
		for _, message := range page.Messages {
			d.StreamListItem(ctx, message)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getGmailMyMessage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var messageID string
	if h.Item != nil {
		messageID = h.Item.(*gmail.Message).Id
	} else {
		messageID = d.KeyColumnQuals["id"].GetStringValue()
	}

	// Return nil, if no input provided
	if messageID == "" {
		return nil, nil
	}

	resp, err := service.Users.Messages.Get("me", messageID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}
