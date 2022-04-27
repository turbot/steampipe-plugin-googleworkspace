## v0.3.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#31](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/31))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#30](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/30))

## v0.2.1 [2022-04-14]

_Bug fixes_

- Fixed links in documentation for configuring OAuth client authentication.

## v0.2.0 [2022-01-31]

_What's new?_

- Added: The `credentials` argument can now be specified in the configuration file to pass in either the path to or the contents of a service account key file in JSON format ([#25](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/25))

_Deprecated_

- The `credential_file` argument in the configuration file is now deprecated and will be removed in the next major version. We recommend using the `credentials` argument instead, which can take the same file path as the `credential_file` argument. ([#25](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/25))

## v0.1.0 [2021-12-08]

_Enhancements_

- Recompiled plugin with Go version 1.17 ([#22](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/22))
- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#19](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/19))

## v0.0.3 [2021-10-20]

_Bug fixes_

- Fixed: All tables now return the service API disabled error directly instead of returning empty rows

## v0.0.2 [2021-09-17]

_What's new?_

- Added: Support for OAuth 2.0 authentication ([#11](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/11))
- Added: Additional optional key columns and better filtering capabilities to the following tables:
  - googleworkspace_calendar_event
  - googleworkspace_calendar_my_event
  - googleworkspace_drive
  - googleworkspace_drive_my_file
  - googleworkspace_gmail_draft
  - googleworkspace_gmail_my_draft
  - googleworkspace_gmail_my_message
  - googleworkspace_people_contact
  - googleworkspace_people_contact_group
  - googleworkspace_people_directory_people

_Enhancements_

- Updated: Improve context cancellation handling in all tables ([#11](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/11))

_Bug fixes_

- Fixed: Querying the `delegates` column in the `googleworkspace_gmail_my_settings` and `googleworkspace_gmail_settings` tables when using OAuth authentication now returns `null` instead of an error ([#11](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/11))

## v0.0.1 [2021-08-26]

_What's new?_

- New tables added

  - [googleworkspace_calendar](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_calendar)
  - [googleworkspace_calendar_event](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_calendar_event)
  - [googleworkspace_calendar_my_event](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_calendar_my_event)
  - [googleworkspace_drive](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_drive)
  - [googleworkspace_drive_my_file](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_drive_my_file)
  - [googleworkspace_gmail_draft](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_draft)
  - [googleworkspace_gmail_message](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_message)
  - [googleworkspace_gmail_my_draft](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_my_draft)
  - [googleworkspace_gmail_my_message](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_my_message)
  - [googleworkspace_gmail_my_settings](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_my_settings)
  - [googleworkspace_gmail_settings](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_gmail_settings)
  - [googleworkspace_people_contact](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_people_contact)
  - [googleworkspace_people_contact_group](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_people_contact_group)
  - [googleworkspace_people_directory_people](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables/googleworkspace_people_directory_people)
