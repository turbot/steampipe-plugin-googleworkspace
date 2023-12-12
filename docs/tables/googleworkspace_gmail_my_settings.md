---
title: "Steampipe Table: googleworkspace_gmail_my_settings - Query Google Workspace Gmail Settings using SQL"
description: "Allows users to query Gmail Settings in Google Workspace, specifically the user's email settings including filters, forwarding rules, IMAP and POP settings, and send-as aliases."
---

# Table: googleworkspace_gmail_my_settings - Query Google Workspace Gmail Settings using SQL

Gmail Settings in Google Workspace is a feature that allows users to customize their Gmail experience according to their preferences. It includes options to manage filters, forwarding rules, IMAP and POP settings, and send-as aliases. These settings help users to manage their email communication effectively and efficiently.

## Table Usage Guide

The `googleworkspace_gmail_my_settings` table provides insights into the Gmail Settings in Google Workspace. As a system administrator or a user, explore your email settings through this table, including filters, forwarding rules, IMAP and POP settings, and send-as aliases. Utilize it to uncover information about your Gmail settings, such as those related to email forwarding, the filters applied to incoming emails, and the configuration of IMAP and POP settings.

**Important Notes**
- To list delegated accounts, you must authenticate using a service account client that has been delegated domain-wide authority.

## Examples

### Basic info
Explore the language settings and delegation details associated with your Google Workspace Gmail account. This can be useful for understanding user preferences and managing access rights.

```sql+postgres
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_my_settings;
```

```sql+sqlite
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_my_settings;
```

### List users with delegated access to their mailbox
Explore which users have delegated access to their mailbox in Google Workspace. This can be useful for assessing security and access control within your organization.

```sql+postgres
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_my_settings
where
  delegates is not null;
```

```sql+sqlite
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_my_settings
where
  delegates is not null;
```

### List users with IMAP access enabled
Explore which users have IMAP access enabled in their Gmail settings. This could be useful for administrators looking to manage or restrict certain types of email access.

```sql+postgres
select
  user_email,
  display_language,
  (imap ->> 'enabled')::boolean as imap_enabled
from
  googleworkspace_gmail_my_settings
where
  (imap ->> 'enabled')::boolean;
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(imap, '$.enabled') as imap_enabled
from
  googleworkspace_gmail_my_settings
where
  json_extract(imap, '$.enabled');
```

### List users with POP access enabled
Identify instances where users have enabled POP access in their Gmail settings. This is useful in understanding the email access preferences within your organization.

```sql+postgres
select
  user_email,
  display_language,
  pop ->> 'accessWindow' as pop_access_window
from
  googleworkspace_gmail_my_settings
where
  pop ->> 'accessWindow' = 'enabled';
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(pop, '$.accessWindow') as pop_access_window
from
  googleworkspace_gmail_my_settings
where
  json_extract(pop, '$.accessWindow') = 'enabled';
```

### List users with automatic forwarding enabled
Determine the areas in which users have enabled automatic forwarding in their email settings. This is useful for understanding the flow of information within your organization and ensuring compliance with communication policies.

```sql+postgres
select
  user_email,
  display_language,
  (auto_forwarding ->> 'enabled')::boolean as auto_forwarding_enabled
from
  googleworkspace_gmail_my_settings
where
  (auto_forwarding ->> 'enabled')::boolean;
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(auto_forwarding, '$.enabled') as auto_forwarding_enabled
from
  googleworkspace_gmail_my_settings
where
  json_extract(auto_forwarding, '$.enabled');
```