package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceGmailMySettings(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_gmail_my_settings",
		Description: "Retrieves settings for the current authenticated user account.",
		List: &plugin.ListConfig{
			Hydrate:           listGmailMyUser,
			ShouldIgnoreError: isNotFoundError([]string{"403"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "user_email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("EmailAddress"),
			},
			{
				Name:        "display_language",
				Description: "Specifies the language settings for the specified account.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getGmailMyLanguage,
			},
			{
				Name:        "auto_forwarding",
				Description: "Describes the auto-forwarding setting for the specified account.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyAutoForwardingSetting,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "delegates",
				Description: "A list of delegates for the specified account.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     listGmailMyDelegateSettings,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "imap",
				Description: "Describes the IMAP setting for the specified account.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyImapSetting,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "pop",
				Description: "Describes the POP settings for the specified account.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyPopSetting,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "vacation",
				Description: "Describes the vacation responder settings for the specified account.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getGmailMyVacationSetting,
				Transform:   transform.FromValue(),
			},
		},
	}
}

//// LIST FUNCTION

func listGmailMyUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.GetProfile("me").Do()
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, resp)

	return nil, nil
}

//// HYDRATE FUNCTIONS

// Lists the delegates for the current authenticated user's account.
// NOTE: This method is only available to service account clients that have been delegated domain-wide authority.
func listGmailMyDelegateSettings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.Delegates.List("me").Do()
	if err != nil {
		return nil, err
	}

	return resp.Delegates, nil
}

// Gets the auto-forwarding setting for the current authenticated user's account.
func getGmailMyAutoForwardingSetting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetAutoForwarding("me").Do()
	if err != nil {
		return nil, err
	}

	// If the property is set with default value, it doesn't show in response
	if resp != nil {
		result := map[string]interface{}{
			"disposition":  resp.Disposition,
			"emailAddress": resp.EmailAddress,
			"enabled":      resp.Enabled,
		}
		return result, nil
	}

	return nil, nil
}

// Gets IMAP settings.
func getGmailMyImapSetting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetImap("me").Do()
	if err != nil {
		return nil, err
	}

	// If the property is set with default value, it doesn't show in response
	if resp != nil {
		result := map[string]interface{}{
			"autoExpunge":     resp.AutoExpunge,
			"enabled":         resp.Enabled,
			"expungeBehavior": resp.ExpungeBehavior,
			"maxFolderSize":   resp.MaxFolderSize,
		}
		return result, nil
	}

	return nil, nil
}

// Gets language settings.
func getGmailMyLanguage(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetLanguage("me").Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Gets POP settings.
func getGmailMyPopSetting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetPop("me").Do()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Gets vacation responder settings.
func getGmailMyVacationSetting(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := GmailService(ctx, d)
	if err != nil {
		return nil, err
	}

	resp, err := service.Users.Settings.GetVacation("me").Do()
	if err != nil {
		return nil, err
	}

	// If the property is set with default value, it doesn't show in response
	if resp != nil {
		result := map[string]interface{}{
			"enableAutoReply":    resp.EnableAutoReply,
			"responseSubject":    resp.ResponseSubject,
			"restrictToContacts": resp.RestrictToContacts,
			"restrictToDomain":   resp.RestrictToDomain,
		}
		return result, nil
	}

	return nil, nil
}