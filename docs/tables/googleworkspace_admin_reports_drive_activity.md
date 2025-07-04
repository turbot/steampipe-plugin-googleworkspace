---
title: "Steampipe Table: googleworkspace_admin_reports_drive_activity - Query Google Workspace Admin Reports Drive Activity Events using SQL"
description: "Allows users to query Drive activity events from the Google Workspace Admin Reports API, providing insights into file operations and user interactions on Google Drive."
folder: "Cloud Admin Reports"
---

# Table: googleworkspace_admin_reports_drive_activity - Query Google Workspace Admin Reports Drive Activity Events using SQL

The Google Workspace Admin Reports Drive Activity API captures detailed events related to Google Drive operations—such as file views, edits, downloads, creations, and deletions—performed by users in your Workspace domain. These logs help you audit file access patterns and investigate unauthorized operations.

## Table Usage Guide

The `googleworkspace_admin_reports_drive_activity` table enables security and operations teams to:

* Track who accessed, created, or modified files.
* Investigate data exfiltration or anomalous behaviors.
* Correlate IP addresses and event types for compliance reporting.

> :point_right: Note that `event_name` values appear in arrays when multiple events occur in a single activity, e.g., `[edit change_user_access add_to_folder upload]`.
>
> :exclamation: For performance, always use the `time` key column to restrict the query to a specific time window.

## Examples

### Basic info

Retrieve Drive activity events in the last 24 hours, showing user and file information.

```sql
select
  time,
  actor_email,
  file_name,
  event_name
from
  googleworkspace_admin_reports_drive_activity
where
  time > now() - interval '1 day';
```

### Show events for a specific file

List edits and views on `Passwords.txt` during the last week.

```sql
select
  time,
  actor_email,
  event_name,
  ip_address
from
  googleworkspace_admin_reports_drive_activity
where
  file_name = 'Passwords.txt'
  and event_name in ('[edit]', '[view]')
  and time > now() - interval '7 days';
```

### Find activities from a specific IP

Identify all Drive operations from `8.8.8.8`.

```sql
select
  time,
  actor_email,
  event_name,
  file_name
from
  googleworkspace_admin_reports_drive_activity
where
  ip_address = '8.8.8.8';
```

### Activities in a custom time range

Query Drive events between two timestamps.

```sql
select
  time,
  actor_email,
  event_name,
  file_name
from
  googleworkspace_admin_reports_drive_activity
where
  time between '2025-06-15T00:00:00Z' and '2025-06-16T23:59:59Z';
```

### Top users by Drive events

Rank users by total Drive activities in the past 5 hours.

```sql
select
  actor_email,
  count(*) as total_events
from
  googleworkspace_admin_reports_drive_activity
where
  time >= now() - interval '5 hours'
group by
  actor_email
order by
  total_events desc
limit 10;
```
