package googleworkspace

import "regexp"

// IsForbiddenError: Returns true, if the service API is disabled.
func IsForbiddenError(err error) bool {
	regexExp := regexp.MustCompile(`googleapi: Error 403: [^\s]+ API has not been used in project [0-9]{12} before or it is disabled\.`)
	return regexExp.MatchString(err.Error())
}
