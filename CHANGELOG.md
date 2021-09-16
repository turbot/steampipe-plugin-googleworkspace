## v0.0.2 [2021-09-16]

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

- Fixed: Querying the `delegates` column in the `googleworkspace_gmail_my_settings` and `googleworkspace_gmail_settings` tables when using OAuth authentication now returns null instead of an error ([#11](https://github.com/turbot/steampipe-plugin-googleworkspace/pull/11))

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
