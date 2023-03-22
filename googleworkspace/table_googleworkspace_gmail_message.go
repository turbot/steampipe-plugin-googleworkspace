package googleworkspace

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"google.golang.org/api/gmail/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceGmailMessage(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_gmail_message",
		Description: "Retrieves messages in the specified user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listGmailMessages,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Required,
				},
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
		},
		Get: &plugin.GetConfig{
			KeyColumns: 	plugin.AllColumns([]string{"id", "user_id"}),
			Hydrate:    	getGmailMessage,
			MaxConcurrency: 50,
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
				Name:        "user_id",
				Description: "User's email address. If not specified, indicates the current authenticated user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("user_id"),
			},
			{
				Name:        "history_id",
				Description: "The ID of the last history record that modified this message.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMessage,
			},
			{
				Name:        "sender_email",
				Description: "Specifies the email address of the sender.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMessage,
				Transform:   transform.From(extractMessageSender),
			},
			{
				Name:        "internal_date",
				Description: "The internal message creation timestamp which determines ordering in the inbox.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getGmailMessage,
				Transform:   transform.FromField("InternalDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "raw",
				Description: "The entire email message in an RFC 2822 formatted and base64url encoded string.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMessage,
			},
			{
				Name:        "size_estimate",
				Description: "Estimated size in bytes of the message.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getGmailMessage,
			},
			{
				Name:        "snippet",
				Description: "A short part of the message text.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMessage,
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
				Hydrate:     getGmailMessage,
			},
			{
				Name:        "payload",
				Description: "The parsed email structure in the message parts.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMessage,
			},
		},
	}
}

//// LIST FUNCTION

func listGmailMessages(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
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

	if d.EqualsQuals["sender_email"] != nil {
		filter = append(filter, fmt.Sprintf("%s = \"%s\"", "from", d.EqualsQuals["sender_email"].GetStringValue()))
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

	resp := service.Users.Messages.List(userID).Q(query).MaxResults(maxResults)
	if err := resp.Pages(ctx, func(page *gmail.ListMessagesResponse) error {
		for _, message := range page.Messages {
			d.StreamListItem(ctx, message)

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

func getGmailMessage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var userID string
	if d.EqualsQuals["user_id"] != nil {
		userID = d.EqualsQuals["user_id"].GetStringValue()
	}

	var messageID string
	if h.Item != nil {
		messageID = h.Item.(*gmail.Message).Id
	} else {
		messageID = d.EqualsQuals["id"].GetStringValue()
	}

	// Return nil, if no input provided
	if messageID == "" || userID == "" {
		return nil, nil
	}

	resp, err := service.Users.Messages.Get(userID, messageID).Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//// TRANSFORM FUNCTIONS

func extractMessageSender(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(*gmail.Message)
	if data.Payload == nil {
		return nil, nil
	}

	for _, payloadHeader := range data.Payload.Headers {
		if payloadHeader.Name == "From" {
			regexExp := regexp.MustCompile(`\<(.*?) *\>`)
			senderEmail := regexExp.FindStringSubmatch(payloadHeader.Value)
			if len(senderEmail) > 1 {
				return senderEmail[1], nil
			}
		}
	}

	return nil, nil
}
