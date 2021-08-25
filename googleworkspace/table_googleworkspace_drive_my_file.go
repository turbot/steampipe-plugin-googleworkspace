package googleworkspace

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/plugin"

	"google.golang.org/api/drive/v3"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceDriveMyFile(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_drive_my_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate: listDriveMyFiles,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "name",
					Require: plugin.Optional,
				},
				{
					Name:    "query",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDriveMyFile,
		},
		Columns: driveFileColumns(),
	}
}

//// LIST FUNCTION

func listDriveMyFiles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := DriveService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Query string for searching files. Refer https://developers.google.com/drive/api/v3/search-files
	// For example, "name contains 'steampipe'", returns all the files containing the word 'steampipe'
	var query string
	if d.KeyColumnQuals["name"] != nil {
		name := d.KeyColumnQuals["name"].GetStringValue()
		query = "name = " + fmt.Sprintf("\"%s\" ", name)
	}
	if d.KeyColumnQuals["query"] != nil {
		query = d.KeyColumnQuals["query"].GetStringValue()
	}

	// Use "*" to return all fields
	resp := service.Files.List().Fields("nextPageToken, files(*)").Q(query)
	if err := resp.Pages(ctx, func(page *drive.FileList) error {
		for _, file := range page.Files {
			d.StreamListItem(ctx, file)
		}
		return nil
	}); err != nil {
		if IsAPIDisabledError(err) {
			return nil, nil
		}
		return nil, err
	}

	return nil, err
}

//// HYDRATE FUNCTIONS

func getDriveMyFile(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getDriveMyFile")

	// Create service
	service, err := DriveService(ctx, d)
	if err != nil {
		return nil, err
	}
	fileID := d.KeyColumnQuals["id"].GetStringValue()

	// Return nil, if no input provided
	if fileID == "" {
		return nil, nil
	}

	// Use "*" to return all fields
	resp, err := service.Files.Get(fileID).Fields("*").Do()
	if err != nil {
		if IsAPIDisabledError(err) {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil
}
