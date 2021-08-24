package googleworkspace

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
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
	var credentialPath, clientSecretPath string
	googleworkspaceConfig := GetConfig(d.Connection)
	if googleworkspaceConfig.CredentialFile != nil {
		credentialPath = *googleworkspaceConfig.CredentialFile
	}
	if googleworkspaceConfig.ClientSecretFile != nil {
		clientSecretPath = *googleworkspaceConfig.ClientSecretFile
	}

	// If credential path not mentioned in steampipe config, search for env variable
	if credentialPath == "" {
		credentialPath = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}

	// No credentials
	if credentialPath == "" && clientSecretPath == "" {
		return nil, errors.New("either credential_path, or client_secret_path must be configured")
	}

	// If credential path provided, use domain-wide delegation for authentication
	if credentialPath != "" {
		ts, err := getTokenSource(ctx, d)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	// Read the client secret, for authenticating using OAuth
	b, err := ioutil.ReadFile(clientSecretPath)
	if err != nil {
		return nil, errors.New("unable to read client secret file")
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(
		b,
		calendar.CalendarReadonlyScope,
		people.ContactsReadonlyScope,
		people.ContactsOtherReadonlyScope,
		drive.DriveReadonlyScope,
		gmail.GmailReadonlyScope,
	)
	if err != nil {
		return nil, errors.New("unable to parse client secret file to config")
	}

	clientToken, err := getHttpClientToken(d)
	if err != nil {
		return nil, err
	}
	httpClient := config.Client(context.Background(), clientToken)
	opts = append(opts, option.WithHTTPClient(httpClient))

	return opts, nil
}

// Returns a JWT TokenSource using the configuration and the HTTP client from the provided context.
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

	// Return error, since impersonation required to authenticate using domain-wide delegation
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
		calendar.CalendarReadonlyScope,
		drive.DriveReadonlyScope,
		gmail.GmailReadonlyScope,
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

// Read token file, and returns token used to authorize the requests to access protected resources on the OAuth 2.0 provider's backend.
func getHttpClientToken(d *plugin.QueryData) (*oauth2.Token, error) {
	// have we already created and cached the token?
	cacheKey := "googleworkspace.bearer_token"
	if ts, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return ts.(*oauth2.Token), nil
	}

	// Get the home dir path
	myself, err := user.Current()
	if err != nil {
		return nil, err
	}
	homedir := myself.HomeDir
	bearerTokenPath := homedir + "/.steampipe/token.json"

	// Read token file
	f, err := os.Open(bearerTokenPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	if err != nil {
		return nil, err
	}

	// cache the token source
	d.ConnectionManager.Cache.Set(cacheKey, tok)

	return tok, nil
}
