package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"

	"google.golang.org/api/drive/v3"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceDriveMyFiles(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_drive_my_file",
		Description: "Retrieves file's metadata or content owned by an user.",
		List: &plugin.ListConfig{
			Hydrate: listDriveMyFiles,
			KeyColumns: []*plugin.KeyColumn{
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
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "Specifies the name of the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "mime_type",
				Description: "The MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploaded content if no value is provided.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "drive_id",
				Description: "ID of the shared drive the file resides in.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "owned_by_me",
				Description: "Indicates whether the user owns the file, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "shared",
				Description: "Indicates whether the file has been shared, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "copy_requires_writer_permission",
				Description: "Indicates whether the options to copy, print, or download this file, should be disabled for readers and commenters, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "created_time",
				Description: "The time at which the file was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "description",
				Description: "A short description of the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "explicitly_trashed",
				Description: "Indicates whether the file has been explicitly trashed, as opposed to recursively trashed from a parent folder.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "file_extension",
				Description: "The final component of fullFileExtension.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "folder_color_rgb",
				Description: "The color for a folder or shortcut to a folder as an RGB hex string.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "full_file_extension",
				Description: "The full file extension extracted from the name field.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "has_augmented_permissions",
				Description: "Indicates whether there are permissions directly on this file, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "has_thumbnail",
				Description: "Indicates whether this file has a thumbnail, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "head_revision_id",
				Description: "The ID of the file's head revision.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "icon_link",
				Description: "A static, unauthenticated link to the file's icon.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "is_app_authorized",
				Description: "Indicates whether the file was created or opened by the requesting app, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "md5_checksum",
				Description: "The MD5 checksum for the content of the file.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "modified_by_me",
				Description: "Indicates whether the file has been modified by this user, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "modified_by_me_time",
				Description: "The last time the file was modified by the use.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "modified_time",
				Description: "The last time the file was modified by anyone.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "original_file_name",
				Description: "The original filename of the uploaded content if available, or else the original value of the name field.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("OriginalFilename").NullIfZero(),
			},
			{
				Name:        "query",
				Description: "A search query combining one or more search terms to [filter](https://developers.google.com/drive/api/v3/search-files) the file results.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("query"),
			},
			{
				Name:        "quota_bytes_used",
				Description: "The number of storage quota bytes used by the file.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "resource_key",
				Description: "A key needed to access the item via a shared link.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "shared_with_me_time",
				Description: "The time at which the file was shared with the user.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "size",
				Description: "The size of the file's content in bytes.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "starred",
				Description: "Indicates whether the user has starred the file, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "thumbnail_link",
				Description: "A short-lived link to the file's thumbnail, if available.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "thumbnail_version",
				Description: "The thumbnail version for use in thumbnail cache invalidation.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "trashed",
				Description: "Indicates whether the file has been trashed, either explicitly or from a trashed parent folder, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "trashed_time",
				Description: "The time that the item was trashed.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "version",
				Description: "A monotonically increasing version number for the file.",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "viewed_by_me",
				Description: "Indicates whether the the file has been viewed by this user, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "viewed_by_me_time",
				Description: "The last time the file was viewed by the user.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "web_content_link",
				Description: "A link for downloading the content of the file in a browser.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "web_view_link",
				Description: "A link for opening the file in a relevant Google editor or viewer in a browser.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "writers_can_share",
				Description: "Indicates whether users with only writer permission can modify the file's permissions, or not.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "app_properties",
				Description: "A collection of arbitrary key-value pairs which are private to the requesting app.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "capabilities",
				Description: "Describes capabilities the current user has on this file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "content_hints",
				Description: "Additional information about the content of the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "content_restrictions",
				Description: "Restrictions for accessing the content of the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "export_links",
				Description: "Links for exporting Docs Editors files to specific formats.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "image_media_metadata",
				Description: "Additional metadata about image media, if available.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "last_modifying_user",
				Description: "The last user to modify the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "link_share_metadata",
				Description: "Contains details about the link URLs that clients are using to refer to this item.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "owners",
				Description: "The owner of this file. Only certain legacy files may have more than one owner.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "parents",
				Description: "The IDs of the parent folders which contain the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "permission_ids",
				Description: "List of permission IDs for users with access to this file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "permissions",
				Description: "The full list of permissions for the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "properties",
				Description: "A collection of arbitrary key-value pairs which are visible to all apps.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "sharing_user",
				Description: "The user who shared the file with the requesting user, if applicable.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "shortcut_details",
				Description: "Shortcut file details. Only populated for shortcut files, which have the mimeType field set to application/vnd.google-apps.shortcut.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "spaces",
				Description: "The list of spaces which contain the file.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "trashing_user",
				Description: "Specifies the user who trashed the file explicitly.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "video_media_metadata",
				Description: "Additional metadata about video media.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDriveMyFiles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := DriveService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Query string for searching shared drives. Refer https://developers.google.com/drive/api/v3/search-files
	// For example, "name contains 'steampipe'", returns all the files containing the word 'steampipe'
	var query string
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
