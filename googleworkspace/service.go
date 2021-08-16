package googleworkspace

import (
	"context"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	drive "google.golang.org/api/drive/v3"
)

func CalendarService(ctx context.Context, d *plugin.QueryData) (*calendar.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.calendar"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*calendar.Service), nil
	}

	// so it was not in cache - create service
	ts, err := getTokenSource(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := calendar.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func DocsService(ctx context.Context, d *plugin.QueryData) (*docs.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.docs"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*docs.Service), nil
	}

	// so it was not in cache - create service
	ts, err := getTokenSource(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := docs.NewService(ctx, option.WithTokenSource(ts))
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
	ts, err := getTokenSource(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := drive.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func SheetsService(ctx context.Context, d *plugin.QueryData) (*sheets.Service, error) {
	// have we already created and cached the service?
	serviceCacheKey := "googleworkspace.sheets"
	if cachedData, ok := d.ConnectionManager.Cache.Get(serviceCacheKey); ok {
		return cachedData.(*sheets.Service), nil
	}

	// so it was not in cache - create service
	ts, err := getTokenSource(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := sheets.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	// cache the service
	d.ConnectionManager.Cache.Set(serviceCacheKey, svc)

	return svc, nil
}

func getTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {
	// NOTE: based on https://developers.google.com/admin-sdk/directory/v1/guides/delegation#go

	// have we already created and cached the token?
	cacheKey := "googleworkspace.token_source"
	if ts, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialPath, impersonateUser string
	googleworkspaceConfig := GetConfig(d.Connection)
	if googleworkspaceConfig.CredentialFile != nil {
		credentialPath = *googleworkspaceConfig.CredentialFile
	}
	if googleworkspaceConfig.ImpersonatedUserEmail != nil {
		impersonateUser = *googleworkspaceConfig.ImpersonatedUserEmail
	}

	// If credential path not mentioned in steampipe config, search for env variable
	if credentialPath == "" {
		credentialPath = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}

	// Credentials not set
	if credentialPath == "" {
		return nil, errors.New("credential_file must be configured")
	}
	if impersonateUser == "" {
		return nil, errors.New("impersonated_user_email must be configured")
	}

	// Read credential file
	jsonCredentials, err := ioutil.ReadFile(credentialPath)
	if err != nil {
		return nil, err
	}

	// Authorize the request
	config, err := google.JWTConfigFromJSON(
		jsonCredentials,
		drive.DriveReadonlyScope,
		calendar.CalendarReadonlyScope,
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
