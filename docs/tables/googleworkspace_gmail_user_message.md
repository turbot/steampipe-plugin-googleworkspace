# Table: googleworkspace_gmail_user_message

Query information about messages in an user's mailbox.

## Examples

### Basic info

```sql
select
  id,
  thread_id,
  internal_date,
  size_estimate,
  snippet
from
  googleworkspace_gmail_user_message;
```

### List unread messages

```sql
select
  id,
  thread_id,
  internal_date,
  size_estimate,
  snippet
from
  googleworkspace_gmail_user_message
where
  query = 'is:unread';
```

### List messages from a specific user

```sql
select
  id,
  thread_id,
  internal_date,
  size_estimate,
  snippet
from
  googleworkspace_gmail_user_message
where
  query = 'from:someuser@example.com';
```

### List draft messages

```sql
select
  id,
  thread_id,
  internal_date,
  size_estimate,
  snippet
from
  googleworkspace_gmail_user_message
where
  label_ids ?& array['DRAFT'];
```
