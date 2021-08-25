connection "googleworkspace" {
  plugin = "googleworkspace"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated.
  # `impersonated_user_email` must be set, since the service account needs to impersonate a user with Admin API permissions to access the workspace resources.
  #impersonated_user_email = "username@domain.com"

  # `credential_file` (optional) - The path to a JSON credential file that contains service account credentials.
  # If `credential_file` is not specified in a connection, credentials will be loaded from the path specified in
  # the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
  #credential_file = "/path/to/my/creds.json"
}