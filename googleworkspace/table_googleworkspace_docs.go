package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceDocs(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_docs",
		Description: "Retrieves latest version of the specified document.",
		List: &plugin.ListConfig{
			Hydrate:           listDocs,
			KeyColumns:        plugin.SingleColumn("document_id"),
			ShouldIgnoreError: isNotFoundError([]string{"404", "400", "403"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "document_id",
				Description: "The ID of the document.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "title",
				Description: "The title of the document.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "suggestions_view_mode",
				Description: "The suggestions view mode applied to the document. Note: When editing a document, changes must be based on a document with SUGGESTIONS_INLINE. Possible values are: DEFAULT_FOR_CURRENT_ACCESS, SUGGESTIONS_INLINE, PREVIEW_SUGGESTIONS_ACCEPTED, and PREVIEW_WITHOUT_SUGGESTIONS.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "revision_id",
				Description: "The revision ID of the document. Can be used in update requests to specify which revision of a document to apply updates to and how the request should behave if the document has been edited since that revision.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "body",
				Description: "Describes the main body of the document.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "document_style",
				Description: "Describes the style of the document.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "footers",
				Description: "The footers in the document, keyed by footer ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "footnotes",
				Description: "The footnotes in the document, keyed by footnote ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "headers",
				Description: "The headers in the document, keyed by header ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "inline_objects",
				Description: "The inline objects in the document, keyed by object ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "lists",
				Description: "The lists in the document, keyed by list ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "named_ranges",
				Description: "The named ranges in the document, keyed by name.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "named_styles",
				Description: "The named styles. There is an entry for each of the possible named style types.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("NamedStyles.Styles"),
			},
			{
				Name:        "positioned_objects",
				Description: "The positioned objects in the document, keyed by object ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "suggested_document_style_changes",
				Description: "The suggested changes to the style of the document, keyed by suggestion ID.",
				Type:        proto.ColumnType_JSON,
			},
			{
				Name:        "suggested_named_style_changes",
				Description: "The suggested changes to he named styles of the document, keyed by suggestion ID.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listDocs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := DocsService(ctx, d)
	if err != nil {
		return nil, err
	}
	documentID := d.KeyColumnQuals["document_id"].GetStringValue()

	resp, err := service.Documents.Get(documentID).Do()
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, resp)

	return nil, nil
}
