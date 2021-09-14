# Table: googleworkspace_gmail_settings

Get information about specified user's email settings for IMAP, auto-forwarding, delegates, and more.

The `googleworkspace_gmail_settings` table can be used to query user's email settings from any user's mailbox, if you have access; and **you must specify user's email address** in the where or join clause (`where user_email=`, `join googleworkspace_gmail_settings on user_email=`).

To list all of **your** email settings use the `googleworkspace_gmail_my_settings` table instead.

**Note:** To list delegated accounts, you must authenticate using a service account client that has been delegated domain-wide authority.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

### List users with POP access enabled

```sql
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

### List users with automatic forwarding enabled

```sql
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
