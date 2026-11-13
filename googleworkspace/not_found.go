package googleworkspace

import (
	"context"
	"slices"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"google.golang.org/api/googleapi"
)

// function which returns an IsNotFoundErrorPredicate for Google Workspace API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
		if gerr, ok := err.(*googleapi.Error); ok {
			return slices.Contains(notFoundErrors, types.ToString(gerr.Code))
		}
		return false
	}
}
