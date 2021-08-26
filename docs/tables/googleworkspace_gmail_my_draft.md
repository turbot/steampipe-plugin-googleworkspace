# Table: googleworkspace_gmail_my_draft

List draft messages in your mailbox.

To query messages in any mailbox, use the `googleworkspace_gmail_draft` table.

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
  googleworkspace_gmail_my_draft;
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
  googleworkspace_gmail_my_draft
where
  query = 'is:unread';
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
  googleworkspace_gmail_my_draft
where
  message_internal_date <= (current_date - interval '30' day);
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
  googleworkspace_gmail_my_draft
where
  message_snippet is null;
```
