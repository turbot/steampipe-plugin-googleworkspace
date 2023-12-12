/*
Package googleworkspace implements a steampipe plugin for googleworkspace.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Workspace resources.
*/
package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-googleworkspace"

// Plugin creates this (googleworkspace) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromCamel().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError([]string{"404"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		TableMap: map[string]*plugin.Table{
			"googleworkspace_calendar":                tableGoogleWorkspaceCalendar(ctx),
			"googleworkspace_calendar_event":          tableGoogleWorkspaceCalendarEvent(ctx),
			"googleworkspace_calendar_my_event":       tableGoogleWorkspaceCalendarMyEvent(ctx),
			"googleworkspace_drive":                   tableGoogleWorkspaceDrive(ctx),
			"googleworkspace_drive_my_file":           tableGoogleWorkspaceDriveMyFile(ctx),
			"googleworkspace_gmail_draft":             tableGoogleWorkspaceGmailDraft(ctx),
			"googleworkspace_gmail_message":           tableGoogleWorkspaceGmailMessage(ctx),
			"googleworkspace_gmail_my_draft":          tableGoogleWorkspaceGmailMyDraft(ctx),
			"googleworkspace_gmail_my_message":        tableGoogleWorkspaceGmailMyMessage(ctx),
			"googleworkspace_gmail_my_settings":       tableGoogleWorkspaceGmailMySettings(ctx),
			"googleworkspace_gmail_settings":          tableGoogleWorkspaceGmailSettings(ctx),
			"googleworkspace_people_contact":          tableGoogleWorkspacePeopleContact(ctx),
			"googleworkspace_people_contact_group":    tableGoogleWorkspacePeopleContactGroup(ctx),
			"googleworkspace_people_directory_people": tableGoogleWorkspacePeopleDirectoryPeople(ctx),
		},
	}

	return p
}
