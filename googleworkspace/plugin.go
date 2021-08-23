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
			ShouldIgnoreError: isNotFoundError([]string{"404", "400", "403"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googleworkspace_calendar":                 tableGoogleWorkspaceCalendar(ctx),
			"googleworkspace_calendar_event":           tableGoogleWorkspaceCalendarEvent(ctx),
			"googleworkspace_calendar_my_event":        tableGoogleWorkspaceCalendarMyEvent(ctx),
			"googleworkspace_contact_connection":       tableGoogleWorkspaceContactConnection(ctx),
			"googleworkspace_contact_directory_people": tableGoogleWorkspaceContanctDirectoryPeople(ctx),
			"googleworkspace_contact_group":            tableGoogleWorkspaceContactGroup(ctx),
			"googleworkspace_docs":                     tableGoogleWorkspaceDocs(ctx),
			"googleworkspace_drive":                    tableGoogleWorkspaceDrive(ctx),
			"googleworkspace_drive_my_file":            tableGoogleWorkspaceDriveMyFiles(ctx),
			"googleworkspace_gmail_user_draft":         tableGoogleWorkspaceGmailUserDraft(ctx),
			"googleworkspace_gmail_user_message":       tableGoogleWorkspaceGmailUserMessage(ctx),
			"googleworkspace_gmail_user_setting":       tableGoogleWorkspaceGmailUserSetting(ctx),
			"googleworkspace_spreadsheet":              tableGoogleWorkspaceSpreadSheet(ctx),
		},
	}

	return p
}
