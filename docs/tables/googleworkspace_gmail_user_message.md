# Table: googleworkspace_gmail_user_message

List messages in a user's mailbox.

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
  googleworkspace_gmail_user_message
order by internal_date
limit 10;
```

### List unread messages received in last 2 days

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
  query = 'is:unread newer_than:2d'
order by internal_date;
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
  query = 'from:someuser@example.com'
order by internal_date;
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
  query = 'in:draft'
order by internal_date;
```

### List chat messages

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
  query = 'in:chats'
order by internal_date;
```
