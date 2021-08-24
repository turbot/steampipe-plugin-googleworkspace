package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"google.golang.org/api/gmail/v1"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceGmailUserDraft(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_gmail_user_draft",
		Description: "Retrieves draft messages in the user's mailbox.",
		List: &plugin.ListConfig{
			Hydrate: listGmailUserDrafts,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "user_id",
					Require: plugin.Optional,
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "draft_id",
					Require: plugin.Required,
				},
				{
					Name:    "user_id",
					Require: plugin.Optional,
				},
			},
			Hydrate: getGmailUserDraft,
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
				Name:        "message_history_id",
				Description: "The ID of the last history record that modified this message.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.HistoryId"),
			},
			{
				Name:        "message_internal_date",
				Description: "The internal message creation timestamp which determines ordering in the inbox.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.InternalDate").Transform(transform.UnixMsToTimestamp),
			},
			{
				Name:        "message_raw",
				Description: "The entire email message in an RFC 2822 formatted and base64url encoded string.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.Raw").NullIfZero(),
			},
			{
				Name:        "message_size_estimate",
				Description: "Estimated size in bytes of the message.",
				Type:        proto.ColumnType_INT,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.SizeEstimate"),
			},
			{
				Name:        "message_snippet",
				Description: "A short part of the message text.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.Snippet").NullIfZero(),
			},
			{
				Name:        "user_id",
				Description: "User's email address. If not specified, indicates the current authenticated user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("user_id"),
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
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.LabelIds"),
			},
			{
				Name:        "message_payload",
				Description: "The parsed email structure in the message parts.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailUserDraft,
				Transform:   transform.FromField("Message.Payload"),
			},
		},
	}
}

//// LIST FUNCTION

func listGmailUserDrafts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to me, to represent current logged-in user
	userID := "me"
	if d.KeyColumnQuals["user_id"] != nil {
		userID = d.KeyColumnQuals["user_id"].GetStringValue()
	}

	// Only return messages matching the specified query. Supports the same query format as the Gmail search box.
	// For example, "from:someuser@example.com is:unread"
	var query string
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}

	// Setting the maximum number of messages, API can return in a single page
	maxResults := int64(500)

	resp := service.Users.Drafts.List(userID).Q(query).MaxResults(maxResults)
	if err := resp.Pages(ctx, func(page *gmail.ListDraftsResponse) error {
		for _, draft := range page.Drafts {
			d.StreamListItem(ctx, draft)
		}
		return nil
	}); err != nil {
		if IsForbiddenError(err) {
			return nil, nil
		}
		return nil, err
	}

	return nil, nil
}

//// HYDRATE FUNCTIONS

func getGmailUserDraft(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Set default value to me, to represent current logged-in user
	userID := "me"
	if d.KeyColumnQuals["user_id"] != nil {
		userID = d.KeyColumnQuals["user_id"].GetStringValue()
	}

	var draftID string
	if h.Item != nil {
		draftID = h.Item.(*gmail.Draft).Id
	} else {
		draftID = d.KeyColumnQuals["draft_id"].GetStringValue()
	}

	// Return nil, if no input provided
	if draftID == "" {
		return nil, nil
	}

	resp, err := service.Users.Drafts.Get(userID, draftID).Do()
	if err != nil {
		if IsForbiddenError(err) {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil
}
