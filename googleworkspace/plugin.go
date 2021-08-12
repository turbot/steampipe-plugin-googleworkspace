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
			ShouldIgnoreError: isNotFoundError([]string{"404", "400"}),
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googleworkspace_calendar":          tableGoogleWorkspaceCalendars(ctx),
			"googleworkspace_calendar_event":    tableGoogleWorkspaceCalendarEvents(ctx),
			"googleworkspace_calendar_my_event": tableGoogleWorkspaceCalendarMyEvents(ctx),
			// "googleworkspace_drive":             tableGoogleWorkspaceDrive(ctx),
			// "googleworkspace_drive_my_file":     tableGoogleWorkspaceDriveMyFiles(ctx),
		},
	}

	return p
}
