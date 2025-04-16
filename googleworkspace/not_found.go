package googleworkspace

import (
	"slices"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"google.golang.org/api/googleapi"
)

// function which returns an IsNotFoundErrorPredicate for Google Workspace API calls
func isNotFoundError(notFoundErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		if gerr, ok := err.(*googleapi.Error); ok {
			return slices.Contains(notFoundErrors, types.ToString(gerr.Code))
		}
		return false
	}
}
