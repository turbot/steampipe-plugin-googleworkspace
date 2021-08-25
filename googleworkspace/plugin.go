/*
Package googleworkspace implements a steampipe plugin for googleworkspace.

This plugin provides data that Steampipe uses to present foreign
tables that represent Google Workspace resources.
*/
package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
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
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googleworkspace_calendar":                tableGoogleWorkspaceCalendar(ctx),
			"googleworkspace_calendar_event":          tableGoogleWorkspaceCalendarEvent(ctx),
			"googleworkspace_calendar_my_event":       tableGoogleWorkspaceCalendarMyEvent(ctx),
			"googleworkspace_drive":                   tableGoogleWorkspaceDrive(ctx),
			"googleworkspace_drive_my_file":           tableGoogleWorkspaceDriveMyFile(ctx),
			"googleworkspace_gmail_user_draft":        tableGoogleWorkspaceGmailUserDraft(ctx),
			"googleworkspace_gmail_user_message":      tableGoogleWorkspaceGmailUserMessage(ctx),
			"googleworkspace_gmail_user_settings":     tableGoogleWorkspaceGmailUserSettings(ctx),
			"googleworkspace_people_connection":       tableGoogleWorkspacePeopleConnection(ctx),
			"googleworkspace_people_contact_group":    tableGoogleWorkspacePeopleContactGroup(ctx),
			"googleworkspace_people_directory_people": tableGoogleWorkspacePeopleDirectoryPeople(ctx),
		},
	}

	return p
}
