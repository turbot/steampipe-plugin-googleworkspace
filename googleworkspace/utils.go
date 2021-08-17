package googleworkspace

import "google.golang.org/api/googleapi"

// Function which returns an IsForbiddenError
func IsForbiddenError(err error) bool {
	if gerr, ok := err.(*googleapi.Error); ok {
		return gerr.Code == 403
	}
	return false
}
