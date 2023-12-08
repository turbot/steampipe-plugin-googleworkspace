---
title: "Steampipe Table: googleworkspace_gmail_my_draft - Query Google Workspace Gmail Drafts using SQL"
description: "Allows users to query Gmail Drafts in Google Workspace, specifically the draft messages and their details, providing insights into the content, status, and metadata of drafts."
---

# Table: googleworkspace_gmail_my_draft - Query Google Workspace Gmail Drafts using SQL

Gmail Drafts in Google Workspace is a feature that allows users to save and manage draft messages before they are sent. These drafts include not only the content of the potential email but also metadata such as the draft's ID, message ID, and thread ID. Gmail Drafts serves as a useful tool for managing email communications and tracking unsent messages within a Google Workspace environment.

## Table Usage Guide

The `googleworkspace_gmail_my_draft` table provides insights into draft messages within Google Workspace's Gmail. As an IT administrator, explore draft-specific details through this table, including content, status, and associated metadata. Utilize it to uncover information about drafts, such as those that have been left unsent or abandoned, and the details of these drafts, to better manage email communications within your organization.

## Examples

### Basic info
Explore which drafts in your Google Workspace Gmail account have a large estimated size. This can help manage your storage space and identify drafts that may be too large to send.

```sql+postgres
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

```sql+sqlite
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
Explore which draft messages are still unread. This can help in prioritizing responses and ensuring important communications are not missed.

```sql+postgres
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

```sql+sqlite
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
Explore which draft messages have been left untouched for over a month. This query is useful in identifying stale drafts that might need attention or deletion.

```sql+postgres
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

```sql+sqlite
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
  message_internal_date <= date('now','-30 day');
```

### List draft messages without a body
Uncover the details of draft emails that lack content. This query is particularly useful when you want to clean up your drafts folder by identifying and removing empty draft messages.

```sql+postgres
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

```sql+sqlite
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