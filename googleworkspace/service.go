package googleworkspace

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func CalendarService(ctx context.Context, d *plugin.QueryData) (*calendar.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.calendar"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*calendar.Service), nil
	}

	// so it was not in cache - create service
	opts, err := getSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := calendar.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func PeopleService(ctx context.Context, d *plugin.QueryData) (*people.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.people"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*people.Service), nil
	}

	// so it was not in cache - create service
	opts, err := getSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := people.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func DriveService(ctx context.Context, d *plugin.QueryData) (*drive.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.drive"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*drive.Service), nil
	}

	// so it was not in cache - create service
	opts, err := getSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := drive.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func GmailService(ctx context.Context, d *plugin.QueryData) (*gmail.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.gmail"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*gmail.Service), nil
	}

	// so it was not in cache - create service
	opts, err := getSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := gmail.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func getSessionConfig(ctx context.Context, d *plugin.QueryData) ([]option.ClientOption, error) {
	opts := []option.ClientOption{}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialContent, tokenPath string
	googleworkspaceConfig := GetConfig(d.Connection)

	// 'credential_file' in connection config is DEPRECATED, and will be removed in future release
	// use `credentials` instead
	if googleworkspaceConfig.Credentials != nil {
		credentialContent = *googleworkspaceConfig.Credentials
	} else if googleworkspaceConfig.CredentialFile != nil {
		credentialContent = *googleworkspaceConfig.CredentialFile
	}

	if googleworkspaceConfig.TokenPath != nil {
		tokenPath = *googleworkspaceConfig.TokenPath
	}

	// If credential path provided, use domain-wide delegation
	if credentialContent != "" {
		ts, err := getTokenSource(ctx, d)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	// If token path provided, authenticate using OAuth 2.0
	if tokenPath != "" {
		path, err := expandPath(tokenPath)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithCredentialsFile(path))
		return opts, nil
	}

	return nil, nil
}

// Returns a JWT TokenSource using the configuration and the HTTP client from the provided context.
func getTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {
	// Note: based on https://developers.google.com/admin-sdk/directory/v1/guides/delegation#go

	// have we already created and cached the token?
	cacheKey := "googleworkspace.token_source"
	if ts, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var impersonateUser string
	googleworkspaceConfig := GetConfig(d.Connection)

	// Read credential from JSON string, or from the given path
	// NOTE: 'credential_file' in connection config is DEPRECATED, and will be removed in future release
	// use `credentials` instead
	var creds string
	if googleworkspaceConfig.Credentials != nil {
		creds = *googleworkspaceConfig.Credentials
	} else if googleworkspaceConfig.CredentialFile != nil {
		creds = *googleworkspaceConfig.CredentialFile
	}

	// Read credential from JSON string, or from the given path
	credentialContent, err := pathOrContents(creds)
	if err != nil {
		return nil, err
	}

	if googleworkspaceConfig.ImpersonatedUserEmail != nil {
		impersonateUser = *googleworkspaceConfig.ImpersonatedUserEmail
	}

	// Return error, since impersonation required to authenticate using domain-wide delegation
	if impersonateUser == "" {
		return nil, errors.New("impersonated_user_email must be configured")
	}

	// Authorize the request
	config, err := google.JWTConfigFromJSON(
		[]byte(credentialContent),
		calendar.CalendarReadonlyScope,
		drive.DriveReadonlyScope,
		gmail.GmailReadonlyScope,
		people.ContactsOtherReadonlyScope,
		people.ContactsReadonlyScope,
		people.DirectoryReadonlyScope,
	)
	if err != nil {
		return nil, err
	}
	config.Subject = impersonateUser

	ts := config.TokenSource(ctx)

	// cache the token source
	d.ConnectionManager.Cache.Set(cacheKey, ts)

	return ts, nil
}
