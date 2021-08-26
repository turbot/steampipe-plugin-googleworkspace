# Table: googleworkspace_gmail_message

List messages in a specific user's mailbox.

The `googleworkspace_gmail_message` table can be used to query user's messages from any mailbox, if you have access; and **you must specify user's email address** in the where or join clause (`where user_id=`, `join googleworkspace_gmail_message on user_id=`).

To list all of **your** messages use the `googleworkspace_gmail_my_message` table instead.

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
  googleworkspace_gmail_message
where
  user_id = 'user@domain.com'
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
  googleworkspace_gmail_message
where
  user_id = 'user@domain.com'
  and query = 'is:unread newer_than:2d'
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
  googleworkspace_gmail_message
where
  user_id = 'user@domain.com'
  and query = 'from:someuser@example.com'
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
  googleworkspace_gmail_message
where
  user_id = 'user@domain.com'
  and query = 'in:draft'
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
  googleworkspace_gmail_message
where
  user_id = 'user@domain.com'
  and query = 'in:chats'
order by internal_date;
```
