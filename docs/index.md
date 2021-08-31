---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/googleworkspace.svg"
brand_color: "#ea4335"
display_name: "Google Workspace"
short_name: "googleworkspace"
description: "Steampipe plugin for querying users, groups, org units and more from your Google Workspace."
og_description: "Query Google Workspace with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/googleworkspace-social-graphic.png"
---

# Google Workspace + Steampipe

[Google Workspace](https://workspace.google.com) is a collection of cloud computing, productivity and collaboration tools, software and products developed and marketed by Google.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time > now()::timestamp
  and end_time < ('now'::timestamp + interval '1 day');
```

```
+----------------+--------------------------------------+---------------------+---------------------+
| summary        | hangout_link                         | start_time          | end_time            |
+----------------+--------------------------------------+---------------------+---------------------+
| Product Review | https://meet.google.com/ris-zooa-rxo | 2021-08-18 12:30:00 | 2021-08-18 13:00:00 |
+----------------+--------------------------------------+---------------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/googleworkspace/tables)**

## Get started

### Install

Download and install the latest Google Workspace plugin:

```bash
steampipe plugin install googleworkspace
```

### Credentials

| Item        | Description |
| :---------- | :-----------|
| Credentials | 1. To use **domain-wide delegation**, generate your [service account and credentials](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#create_the_service_account_and_credentials) and [delegate domain-wide authority to your service account](https://developers.google.com/admin-sdk/directory/v1/guides/delegation#delegate_domain-wide_authority_to_your_service_account). Enter the following OAuth 2.0 scopes for the services that the service account can access:<br />`https://www.googleapis.com/auth/calendar.readonly`,<br />`https://www.googleapis.com/auth/contacts.readonly`,<br />`https://www.googleapis.com/auth/contacts.other.readonly`,<br />`https://www.googleapis.com/auth/directory.readonly`,<br />`https://www.googleapis.com/auth/drive.readonly`,<br />`https://www.googleapis.com/auth/gmail.readonly`<br />2. To use **OAuth client**, configure your [credentials](#authenticate-using-oauth-client) |
| Radius      | Each connection represents a single Google Workspace account. |
| Resolution  | 1. Credentials from the JSON file specified by the `credential_file` parameter in your steampipe config.<br />2. Credentials from the JSON file specified by the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.<br />3. Credentials from the default json file location (~/.config/gcloud/application_default_credentials.json). |

### Configuration

Installing the latest googleworkspace plugin will create a config file (`~/.steampipe/config/googleworkspace.spc`) with a single connection named `googleworkspace`:

```hcl
connection "googleworkspace" {
  plugin = "googleworkspace"

  # You may connect to google workspace using more than one option
  # 1. To authenticate using domain-wide delegation, specify service account credential file, and the user email for impersonation
  # `credential_file` (optional) - The path to a JSON credential file that contains service account credentials.
  # If not specified in a connection, credentials will be loaded from the path specified in
  # the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
  # credential_file         = "/path/to/my/creds.json"

  # `impersonated_user_email` (required) - The email (string) of the user which should be impersonated. Needs permissions to access the Admin APIs.
  # `impersonated_user_email` must be set, since the service account needs to impersonate a user with Admin API permissions to access the workspace services.
  # impersonated_user_email = "username@domain.com"


  # 2. To authenticate OAuth 2.0 client, specify the default json file location
  # token_path              = "~/.config/gcloud/application_default_credentials.json"
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-googleworkspace
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)

## Advanced configuration options

### Authenticate using OAuth client

You can use client secret credential to protect the user's data by only granting tokens to authorized requestors. Use following steps to configure credentials:

- [Configure the OAuth consent screen](https://developers.google.com/workspace/guides/create-credentials#configure_the_oauth_consent_screen)
- [Create a OAuth client ID credential](https://developers.google.com/workspace/guides/create-credentials#create_a_oauth_client_id_credential), and download the client secret JSON file.
- Wherever you have the [Google Cloud SDK](https://cloud.google.com/sdk/docs/install) installed, run the following command with the
correct client secret JSON file parameters:
  
  ```sh
  gcloud auth application-default login \
    --client-id-file=client_secret.json \
    --no-launch-browser \
    --scopes="\
  https://www.googleapis.com/auth/calendar.readonly,\
  https://www.googleapis.com/auth/contacts.other.readonly,\
  https://www.googleapis.com/auth/contacts.readonly,\
  https://www.googleapis.com/auth/directory.readonly,\
  https://www.googleapis.com/auth/drive.readonly,\
  https://www.googleapis.com/auth/gmail.readonly"
  ```

- Copy the login link into your browser and authenticate as the user you would like to make
the API calls through.

- Review the output for the location of the **Application Default Credentials** file, which
usually appears following the text `Credentials saved to file`.

- Add the **Application Default Credentials** file path in steampipe config using `token_path`, or  pass the
path of this file to applications using the environment variable.

### Credentials from Environment Variables

By default, the plugin uses the credential file path provided in the connection config. You can also specify static credentials using environment variables:

```sh
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/my/creds.json
```
