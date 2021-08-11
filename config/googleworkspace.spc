connection "googleworkspace" {
  plugin = "googleworkspace"

  # `credential_file` (required) - The path to a JSON credential file that contains service account credentials.
  # If `credential_file` is not specified in a connection, credentials will be loaded from:
  #  - The path specified in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable, if set
  #credential_file = "/path/to/<public_key_fingerprint>-privatekey.json"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate one of those users to access the workspace.
  #impersonated_user_email = "username@domain.com"
}