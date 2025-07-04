---
title: "Steampipe Table: googleworkspace_admin_reports_mobile_activity - Query Google Workspace Admin Reports Mobile Activity Events using SQL"
description: "Allows users to query mobile activity events from the Google Workspace Admin Reports API, providing insights into device usage and mobile access patterns."
folder: "Cloud Admin Reports"
---

# Table: googleworkspace_admin_reports_mobile_activity - Query Google Workspace Admin Reports Mobile Activity Events using SQL

The Google Workspace Admin Reports Mobile Activity API captures events related to device actions—such as logins, syncs, OS updates, and new device enrollments—performed by users in your Workspace domain.

## Table Usage Guide


Use the `googleworkspace_admin_reports_mobile_activity` table to investigate how users interact with Google services from their devices. Track device enrollments, removals, sync operations, and any policy-related events.

## Examples

### Basic info

Retrieve mobile activity events in the last 24 hours.

```sql
select
  time,
  actor_email,
  event_name,
  device_model
from
  googleworkspace_admin_reports_mobile_activity
where
  time > now() - interval '1 day';
```

### Device sync events for a user

Show sync operations by `alice@example.com` over the past 3 days.

```sql
select
  time,
  actor_email,
  event_name,
  device_model
from
  googleworkspace_admin_reports_mobile_activity
where
  actor_email = 'alice@example.com'
  and event_name like '%DEVICE_SYNC_EVENT%'
  and time > now() - interval '3 days';
```

### Connections from a new device

Identify all new device registrations in the last week.

```sql
select
  time,
  actor_email,
  event_name,
  device_id,
  device_model
from
  googleworkspace_admin_reports_mobile_activity
where
  event_name = '[DEVICE_REGISTER_UNREGISTER_EVENT]'
  and time > now() - interval '7 days';
```

### Custom time window analysis

Query mobile activities between two timestamps.

```sql
select
  time,
  actor_email,
  event_name,
  device_id,
  device_model
from
  googleworkspace_admin_reports_mobile_activity
where
  time between '2025-06-10T00:00:00Z' and '2025-06-15T23:59:59Z';
```
