# Table: googleworkspace_gmail_my_settings

Get information about your email settings for IMAP, auto-forwarding, delegates, and more.

To query email settings information about any mailbox, use the `googleworkspace_gmail_settings` table.

## Examples

### Basic info

```sql
select
  user_email,
  display_language,
  delegates
from
  googleworkspace_gmail_my_settings;
```

### List users with delegated access to their mailbox

```sql
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

```sql
select
  user_email,
  display_language,
  (imap ->> 'enabled')::boolean as imap_enabled
from
  googleworkspace_gmail_my_settings
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
  googleworkspace_gmail_my_settings
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
  googleworkspace_gmail_my_settings
where
  (auto_forwarding ->> 'enabled')::boolean;
```
