connection "googleworkspace" {
  plugin = "googleworkspace"

  # You may connect to Google Workspace using more than one option:
  # 1. To authenticate using domain-wide delegation, specify  a service account credential file and the user email for impersonation
  # `credential_file` (optional) - The path to a JSON credential file that contains service account credentials.
  #credential_file         = "/path/to/my/creds.json"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate a user with Admin API permissions to access the workspace services.
  #impersonated_user_email = "username@domain.com"

  # 2. To authenticate using OAuth 2.0, specify a client secret file
  # `token_path` (optional) - The path to a JSON credential file that contains Google application credentials.
  # If `token_path` is not specified in a connection, credentials will be loaded from:
  #   - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set; otherwise
  #   - The standard location (`~/.config/gcloud/application_default_credentials.json`)
  #token_path = "~/.config/gcloud/application_default_credentials.json"
}
