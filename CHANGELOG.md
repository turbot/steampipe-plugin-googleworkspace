## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#82](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/82))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#82](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/82))

## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#77](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/77))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#77](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/77))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-googleworkspace/blob/main/docs/LICENSE). ([#77](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/77))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#76](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/76))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#56](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/56))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#51](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/51))
- Recompiled plugin with Go version `1.21`. ([#51](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/51))

## v0.6.0 [2023-03-23]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#40](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/40))

## v0.5.0 [2022-09-29]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#36](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/36))
- Recompiled plugin with Go version `1.19`. ([#36](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/36))

## v0.4.0 [2022-07-21]

_Bug fixes_

- Fixed the `GetConfig` max concurrency configuration in the `googleworkspace_gmail_message` and the `googleworkspace_gmail_my_message` tables to resolve the plugin validation errors. ([#34](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/34))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#34](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/34))

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
