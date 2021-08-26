# Table: googleworkspace_gmail_user_settings

Get information about a user's email settings for IMAP, auto-forwarding, delegates, and more.

**Note:** To list the `delegates` for an account, the service account used for authentication requires domain-wide authority.

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

### List users with delegated access to their mailbox

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

### List users with automatic forwarding enabled

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
