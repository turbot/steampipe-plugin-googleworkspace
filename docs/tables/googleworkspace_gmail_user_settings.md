# Table: googleworkspace_gmail_user_settings

Query information about user's settings for the specified account.

**NOTE:**

- To list the `delegates` for the specified account, use service account clients that have been delegated domain-wide authority.

## Examples

### Basic info

```sql
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_user_settings;
```

### List users can delegate access to their mailbox to other users in domain

```sql
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_user_settings
where
  delegates is not null;
```

### List users with IMAP access enabled

```sql
select
  user_email,
  display_language,
  (imap ->> 'enabled')::boolean as imap_enabled
from
  googleworkspace_gmail_user_settings
where
  (imap ->> 'enabled')::boolean;
```

### List users with POP access enabled

```sql
select
  user_email,
  display_language,
  pop ->> 'accessWindow' as pop_access_window
from
  googleworkspace_gmail_user_settings
where
  pop ->> 'accessWindow' = 'enabled';
```

### List users with automatic forwarding option enabled

```sql
select
  user_email,
  display_language,
  (auto_forwarding ->> 'enabled')::boolean as auto_forwarding_enabled
from
  googleworkspace_gmail_user_settings
where
  (auto_forwarding ->> 'enabled')::boolean;
```
