---
title: "Steampipe Table: googleworkspace_admin_reports_admin_activity - Query Google Workspace Admin Reports Admin Activity Events using SQL"
description: "Allows users to query admin activity events from the Google Workspace Admin Reports API, providing insights into administrative actions and login events."
folder: "Cloud Admin Reports"
---

# Table: googleworkspace_admin_reports_admin_activity - Query Google Workspace Admin Reports Admin Activity Events using SQL

The Google Workspace Admin Reports Admin Activity API provides audit logs for administrative actions and login events performed in your Workspace domain. These records enable you to detect privilege or user modifications, investigate security incidents, and monitor compliance.

## Table Usage Guide

The `googleworkspace_admin_reports_admin_activity` table is ideal for tracking admin operations such as user creation, password changes, and role assignments. Use it to monitor and audit critical administrative events.

> :point_right: Notice that the event_name are inside brackets, it's because we can have several events for the same entry, example : `[CHANGE_PASSWORD CHANGE_PASSWORD_ON_NEXT_LOGIN]`

## Examples

### Basic info

Retrieve events performed by administrators of your Google Workspace domain in the last 24 hours.

```sql
select
  time,
  actor_email,
  event_name,
  ip_address,
  events
from
  googleworkspace_admin_reports_admin_activity
where
  time > now() - interval '1 day';
```

### List all password change events

Show all password change operations performed by administrators.

```sql
select
  time,
  actor_email,
  event_name,
  user_email,
  ip_address
from
  googleworkspace_admin_reports_admin_activity
where
  event_name like '%CHANGE_PASSWORD%';
```

### List all user creation events

Find user creation events performed by admins originating from a specific IP range.

```sql
select
  time,
  actor_email,
  ip_address,
  user_email,
  event_name
from
  googleworkspace_admin_reports_admin_activity
where
  event_name = '[CREATE_USER]'
  and ip_address between '203.0.113.0' and '203.0.113.255';
```

### Get activities within a custom time window

Query admin activities between two timestamps.

```sql
select
  time,
  actor_email,
  event_name,
  events
from
  googleworkspace_admin_reports_admin_activity
where
  time between '2025-06-15T00:00:00Z' and '2025-06-20T23:59:59Z';
```
