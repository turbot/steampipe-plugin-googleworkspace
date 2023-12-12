---
title: "Steampipe Table: googleworkspace_gmail_message - Query Google Workspace Gmail Messages using SQL"
description: "Allows users to query Gmail Messages in Google Workspace, specifically the detailed information about each message, providing insights into message metadata, labels, and thread details."
---

# Table: googleworkspace_gmail_message - Query Google Workspace Gmail Messages using SQL

A Gmail Message is a communication between two or more parties, typically in the form of an email. In Google Workspace, each message is associated with a unique ID and can contain various metadata, such as sender and recipient information, date and time stamps, and labels. Messages can be part of a thread, which groups together related messages for easier navigation and organization.

## Table Usage Guide

The `googleworkspace_gmail_message` table provides insights into Gmail Messages within Google Workspace. As a system administrator, explore message-specific details through this table, including metadata, labels, and thread information. Utilize it to uncover information about messages, such as those with specific labels, the relationships between messages and threads, and the verification of sender and recipient details.

**Important Notes**
- You must specify the `user_id` in the `where` or join clause (`where user_id=`, `join googleworkspace_gmail_my_message g on g.user_id=`) to query this table.

## Examples

### Basic info
Explore the basic information of your Gmail messages, such as their ID, thread ID, date, size estimate, and snippet. This is useful to gain insights into your Gmail activity and help manage your inbox effectively.

```sql+postgres
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

```sql+sqlite
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
Discover recent unread messages in your Gmail account. This query is useful for prioritizing your response to recent and unattended communications.

```sql+postgres
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

```sql+sqlite
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
Explore messages from a specific user in your Google Workspace Gmail account to gain insights into communication trends. This is particularly useful for understanding the frequency and content of interactions with specific individuals.

```sql+postgres
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

```sql+sqlite
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
Explore your draft messages in Gmail to gain insights into their content and size, and to determine their chronological order. This can be useful for managing and organizing your drafts effectively.

```sql+postgres
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

```sql+sqlite
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
Explore your Google Workspace Gmail chat messages to gain insights into the content and timing of your conversations. This could be useful in understanding communication patterns or tracking specific discussions.

```sql+postgres
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

```sql+sqlite
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