---
title: "Steampipe Table: googleworkspace_gmail_my_message - Query Google Workspace Gmail Messages using SQL"
description: "Allows users to query Gmail Messages in Google Workspace, specifically the details of the user's messages, providing insights into email communication and potential anomalies."
---

# Table: googleworkspace_gmail_my_message - Query Google Workspace Gmail Messages using SQL

Gmail is a service within Google Workspace that provides a robust and secure platform for sending, receiving, and storing email. It offers an intuitive interface for users to manage their emails, including features such as spam filtering, conversation view, and powerful search. Gmail is designed to be accessed on any device, providing flexibility and continuity for users on the go.

## Table Usage Guide

The `googleworkspace_gmail_my_message` table provides insights into Gmail Messages within Google Workspace. As a system administrator, explore message-specific details through this table, including the sender, recipient, subject, and timestamp. Utilize it to uncover information about messages, such as those marked as spam, the communication patterns, and the verification of message headers.

## Examples

### Basic info
Explore your recent Gmail messages to gain insights into their content and size. This aids in managing your inbox by identifying large or outdated threads.

```sql+postgres
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

```sql+sqlite
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
Explore which unread messages have been received in the last two days. This helps to prioritize the most recent and potentially urgent communications that require your attention.

```sql+postgres
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

```sql+sqlite
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
Explore the communication history of a specific user to understand the context and frequency of their interactions. This is particularly useful in situations where you need to track the activity of a particular individual for audit or investigation purposes.

```sql+postgres
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

```sql+sqlite
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
Explore your draft messages in Gmail to understand their content and size. This can help in managing your drafts more effectively by identifying large drafts or older drafts that may need attention.

```sql+postgres
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

```sql+sqlite
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
Explore your chat history to gain insights into the frequency and content of your interactions. This can be useful for assessing communication patterns and identifying key conversations.

```sql+postgres
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

```sql+sqlite
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