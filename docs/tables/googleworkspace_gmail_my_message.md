# Table: googleworkspace_gmail_my_message

List messages in your mailbox.

To query messages in any mailbox, use the `googleworkspace_gmail_message` table.

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
  googleworkspace_gmail_my_message
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
  googleworkspace_gmail_my_message
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
  googleworkspace_gmail_my_message
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
  googleworkspace_gmail_my_message
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
  googleworkspace_gmail_my_message
where
  query = 'in:chats'
order by internal_date;
```
