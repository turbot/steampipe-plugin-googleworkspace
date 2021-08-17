# Table: googleworkspace_gmail_user_setting

Query information about user's settings for the specified account.

## Examples

### Basic info

```sql
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_user_setting;
```

### List users can delegate access to their mailbox to other users in domain

```sql
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_user_setting
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
  googleworkspace_gmail_user_setting
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
  googleworkspace_gmail_user_setting
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
  googleworkspace_gmail_user_setting
where
  (auto_forwarding ->> 'enabled')::boolean;
```
