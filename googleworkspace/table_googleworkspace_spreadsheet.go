package googleworkspace

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableGoogleWorkspaceSpreadSheet(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "googleworkspace_spreadsheet",
		Description: "Retrieves latest version of the specified document.",
		List: &plugin.ListConfig{
			Hydrate:           listSpreadsheets,
			KeyColumns:        plugin.AllColumns([]string{"id", "range"}),
			ShouldIgnoreError: isNotFoundError([]string{"404", "400", "403"}),
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Description: "The ID of the spreadsheet.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromQual("id"),
			},
			{
				Name:        "range",
				Description: "The range the values cover, in A1 notation. The range indicates the entire requested range, even though the values will exclude trailing rows and columns.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "major_dimension",
				Description: "The major dimension of the values. Possible values are: DIMENSION_UNSPECIFIED, ROWS, COLUMNS.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "values",
				Description: "The data that was read or to be written. This is an array of arrays, the outer array representing all the data and each inner array representing a major dimension.",
				Type:        proto.ColumnType_JSON,
			},
		},
	}
}

//// LIST FUNCTION

func listSpreadsheets(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service
	service, err := SheetsService(ctx, d)
	if err != nil {
		return nil, err
	}
	sheetID := d.KeyColumnQuals["id"].GetStringValue()
	sheetRange := d.KeyColumnQuals["range"].GetStringValue()

	resp, err := service.Spreadsheets.Values.Get(sheetID, sheetRange).Do()
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, resp)

	return nil, nil
}
