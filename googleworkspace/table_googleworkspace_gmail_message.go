package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

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
					Name:    "query",
					Require: plugin.Optional,
				},
			},
			ShouldIgnoreError: isNotFoundError([]string{"403"}),
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "user_id"}),
			Hydrate:    getGmailMessage,
		},
		HydrateConfig: []plugin.HydrateConfig{
			{
				Func:           getGmailMessage,
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
	if d.KeyColumnQuals["user_id"] != nil {
		userID = d.KeyColumnQuals["user_id"].GetStringValue()
	}

	// Only return messages matching the specified query. Supports the same query format as the Gmail search box.
	// For example, "from:someuser@example.com is:unread"
	// Note: Parameter cannot be used when accessing the api using the gmail.metadata scope.
	var query string
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}

	// Setting the maximum number of messages, API can return in a single page
	maxResults := int64(500)

	resp := service.Users.Messages.List(userID).Q(query).MaxResults(maxResults)
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

func getGmailMessage(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	var userID string
	if d.KeyColumnQuals["user_id"] != nil {
		userID = d.KeyColumnQuals["user_id"].GetStringValue()
	}

	var messageID string
	if h.Item != nil {
		messageID = h.Item.(*gmail.Message).Id
	} else {
		messageID = d.KeyColumnQuals["id"].GetStringValue()
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
