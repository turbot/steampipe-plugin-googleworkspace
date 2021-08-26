# Table: googleworkspace_gmail_draft

List draft messages in a specific user's mailbox.

The `googleworkspace_gmail_draft` table can be used to query draft messages from any mailbox, if you have access; and **you must specify user's email address** in the where or join clause (`where user_id=`, `join googleworkspace_gmail_draft on user_id=`).

To list all of **your** draft messages use the `googleworkspace_gmail_my_draft` table instead.

## Examples

### Basic info

```sql
select
  draft_id,
  message_id,
  message_thread_id,
  message_internal_date,
  message_size_estimate,
  message_snippet
from
  googleworkspace_gmail_draft
where
  user_id = 'user@domain.com';
```

### List unread draft messages

```sql
select
  draft_id,
  message_id,
  message_thread_id,
  message_internal_date,
  message_size_estimate,
  message_snippet
from
  googleworkspace_gmail_draft
where
  user_id = 'user@domain.com'
  and query = 'is:unread';
```

### List draft messages older than 30 days

```sql
select
  draft_id,
  message_id,
  message_thread_id,
  message_internal_date,
  message_size_estimate,
  message_snippet
from
  googleworkspace_gmail_draft
where
  user_id = 'user@domain.com'
  and message_internal_date <= (current_date - interval '30' day);
```

### List draft messages without a body

```sql
select
  draft_id,
  message_id,
  message_thread_id,
  message_internal_date,
  message_size_estimate,
  message_snippet
from
  googleworkspace_gmail_draft
where
  user_id = 'user@domain.com'
  and message_snippet is null;
```
