package googleworkspace

import (
	"context"
	"slices"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
	"google.golang.org/api/googleapi"
)

// isNotFoundError returns a plugin.ErrorPredicateWithContext that reports true when err is a *googleapi.Error whose HTTP status code is in notFoundErrors.
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(_ context.Context, _ *plugin.QueryData, _ *plugin.HydrateData, err error) bool {
		if gerr, ok := err.(*googleapi.Error); ok {
			return slices.Contains(notFoundErrors, types.ToString(gerr.Code))
		}
		return false
	}
}
