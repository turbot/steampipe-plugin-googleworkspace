---
title: "Steampipe Table: googleworkspace_gmail_settings - Query Google Workspace Gmail Settings using SQL"
description: "Allows users to query Gmail Settings in Google Workspace, providing insights into individual user settings and preferences within Gmail."
---

# Table: googleworkspace_gmail_settings - Query Google Workspace Gmail Settings using SQL

Google Workspace's Gmail service is a powerful email platform used by organizations globally. Its settings include various user preferences and configurations that govern the behavior of the Gmail interface for individual users. These settings encompass aspects such as display language, page size, keyboard shortcuts, and email forwarding rules.

## Table Usage Guide

The `googleworkspace_gmail_settings` table provides insights into individual user settings within Google Workspace's Gmail service. As a system administrator or IT professional, you can use this table to explore and manage user-specific settings and preferences in Gmail. This includes information on display language, email forwarding rules, keyboard shortcuts, and more, enabling efficient management and troubleshooting of user issues.

**Important Notes**
- You must specify the `user_email` in the `where` or join clause (`where user_email=`, `join googleworkspace_gmail_settings g on g.user_email=`) to query this table.
- To list delegated accounts, you must authenticate using a service account client that has been delegated domain-wide authority.

## Examples

### Basic info
Explore the language settings and delegates associated with a specific user's Gmail account. This can help in understanding the user's preferred language and who has access to their account.

```sql+postgres
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com';
```

```sql+sqlite
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com';
```

### List users with delegated access to their mailbox
Explore which users have granted others access to their mailbox, a useful feature for shared email accounts or teams managing a common inbox.

```sql+postgres
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and delegates is not null;
```

```sql+sqlite
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and delegates is not null;
```

### List users with IMAP access enabled
Analyze the settings to understand which users have enabled IMAP access in their Google Workspace Gmail settings. This can help in auditing user access and ensuring compliance with company email policies.

```sql+postgres
select
  user_email,
  display_language,
  (imap ->> 'enabled')::boolean as imap_enabled
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and (imap ->> 'enabled')::boolean;
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(imap, '$.enabled') as imap_enabled
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and json_extract(imap, '$.enabled');
```

### List users with POP access enabled
Explore which users have POP access enabled in their email settings. This is useful for identifying potential security risks or ensuring compliance with company policies regarding email access methods.

```sql+postgres
select
  user_email,
  display_language,
  pop ->> 'accessWindow' as pop_access_window
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and pop ->> 'accessWindow' = 'enabled';
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(pop, '$.accessWindow') as pop_access_window
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and json_extract(pop, '$.accessWindow') = 'enabled';
```

### List users with automatic forwarding enabled
Explore which users have automatic forwarding enabled in their email settings. This can be useful in maintaining data privacy and reducing the risk of sensitive information being inadvertently shared.

```sql+postgres
select
  user_email,
  display_language,
  (auto_forwarding ->> 'enabled')::boolean as auto_forwarding_enabled
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and (auto_forwarding ->> 'enabled')::boolean;
```

```sql+sqlite
select
  user_email,
  display_language,
  json_extract(auto_forwarding, '$.enabled') as auto_forwarding_enabled
from
  googleworkspace_gmail_settings
where
  user_email = 'user@domain.com'
  and json_extract(auto_forwarding, '$.enabled');
```