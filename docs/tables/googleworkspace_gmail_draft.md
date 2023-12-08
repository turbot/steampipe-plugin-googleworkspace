---
title: "Steampipe Table: googleworkspace_gmail_draft - Query Google Workspace Gmail Drafts using SQL"
description: "Allows users to query Gmail Drafts in Google Workspace, specifically the metadata and content of draft emails, providing insights into saved but unsent communications."
---

# Table: googleworkspace_gmail_draft - Query Google Workspace Gmail Drafts using SQL

Google Workspace's Gmail service offers a Drafts feature, where users can create, save, and manage draft emails before sending them. This feature provides a space for composing and editing emails, which can be saved for later completion and dispatch. The drafts can contain a variety of information, including recipients, subject lines, and body content.

## Table Usage Guide

The `googleworkspace_gmail_draft` table provides insights into draft emails within Google Workspace's Gmail service. As an IT administrator or security analyst, explore draft-specific details through this table, including metadata, message content, and associated user information. Utilize it to uncover information about unsent communications, such as those containing sensitive information, drafts saved by specific users, and the content of these saved but unsent messages.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join googleworkspace_gmail_draft g on g.user_id=`) to query this table.

## Examples

### Basic info
Explore which drafts in your Google Workspace Gmail account have a specific user ID. This can help you manage your drafts more effectively by identifying which drafts belong to a specific user, especially useful in large organizations where multiple users may be using the same account.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that contain unread draft messages in your Gmail account. This can be especially useful for managing your email workflow and ensuring important drafts don't get overlooked.

```sql+postgres
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

```sql+sqlite
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
Explore which draft messages have been left untouched for over 30 days. This could be useful for clearing out old drafts or identifying potential forgotten tasks.

```sql+postgres
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

```sql+sqlite
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
  and message_internal_date <= date('now','-30 day');
```

### List draft messages without a body
Discover the segments that consist of draft messages without any content. This can be useful for identifying and cleaning up unnecessary drafts, freeing up storage space and keeping your draft folder organized.

```sql+postgres
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

```sql+sqlite
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