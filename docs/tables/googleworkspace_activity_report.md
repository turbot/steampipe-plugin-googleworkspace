---
title: "Steampipe Table: googleworkspace_activity_report - Query Google Workspace Admin Reports Activity using SQL"
description: "Allows users to query the Google Workspace Admin Reports API to retrieve detailed audit activity logs across various Google Workspace applications."
---

# Table: googleworkspace_activity_report - Query Google Workspace Admin Reports Activity using SQL

Google Workspace Activity Report provides visibility into user and administrator activity across your Google Workspace environment. You can query activity data from various Workspace applications—such as Drive, Gmail, and Login—to monitor usage patterns, security events, and administrative actions.

## Table Usage Guide

The `googleworkspace_activity_report` table in Steampipe provides a unified interface to query the Google Workspace Admin Reports API. It surfaces detailed audit logs across all Workspace applications (Drive, Calendar, Keep, Admin console, and more). You can use this table to investigate user actions, system events, and security-related activities within your Workspace environment.

**Important Notes**
- You must `application_name` in a `where` clause in order to use this table ([List of all applications](https://developers.google.com/workspace/admin/reports/reference/rest/v1/activities/list?hl=fr#applicationname)).
- You must have the `Super Admin` role to use this table.
- For improved performance, it is advised that you use the optional qual `time` to limit the result set to a specific time period.
- This table supports optional quals. Queries with optional quals are optimised to use Activity filters. Optional quals are supported for the following columns:
  - `actor_email`
  - `ip_address`
  - `event_name`

## Examples

### List all Drive events in the last hour
Retrieve audit events for Google Drive that occurred in the past hour.

```sql+postgres
select
  time,
  actor_email,
  event_names,
  param->>'value' as file_name,
  ip_address,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param
where
  application_name = 'drive'
  and param->>'name' = 'doc_title'
  and time > now() - interval '1 hour';
```

```sql+sqlite
select
  time,
  actor_email,
  event_names,
  param->>'value' as file_name,
  ip_address,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param
where
  application_name = 'drive'
  and param->>'name' = 'doc_title'
  and time > datetime('now', '-1 hour');
```

### List all password changes performed by administrators on users
Show all changes of password performed by administrators on users in the last month.

```sql+postgres
select
  time,
  actor_email,
  event_names,
  param->>'value' as user_email,
  ip_address,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param
where
  application_name = 'admin'
  and event_name = 'CHANGE_PASSWORD'
  and param->>'name' = 'USER_EMAIL'
  and time > now() - interval '1 month';
```

```sql+sqlite
select
  time,
  actor_email,
  event_names,
  param->>'value' as user_email,
  ip_address,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param
where
  application_name = 'admin'
  and event_name = 'CHANGE_PASSWORD'
  and param->>'name' = 'USER_EMAIL'
  and time > datetime('now', '-1 month');
```

### Show login failures by specific user
Show all failed login attempts by a specific user in the last week.

```sql+postgres
select
  time,
  event_names,
  ip_address
from
  googleworkspace_activity_report
where
  application_name = 'login'
  and actor_email = 'xxx@xxx.xxx'
  and event_name = 'login_failure'
  and time > now() - '1 week'::interval;
```

```sql+sqlite
select
  time,
  event_names,
  ip_address
from
  googleworkspace_activity_report
where
  application_name = 'login'
  and actor_email = 'xxx@xxx.xxx'
  and event_name = 'login_failure'
  and time > datetime('now', '-1 week');
```

### Show all connections from a new device
Identify all connections from a new device in the last week.

```sql+postgres
select
  time,
  actor_email,
  event_names,
  param1->>'value' as device_id,
  param2->>'value' as device_model,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param1
  cross join lateral jsonb_array_elements(evt->'parameters') as param2
where
  application_name = 'mobile'
  and event_name = 'DEVICE_REGISTER_UNREGISTER_EVENT'
  and param1->>'name' = 'DEVICE_ID'
  and param2->>'name' = 'DEVICE_MODEL'
  and time > now() - interval '1 day';
```

```sql+sqlite
select
  time,
  actor_email,
  event_names,
  param1->>'value' as device_id,
  param2->>'value' as device_model,
  events
from
  googleworkspace_activity_report as a
  cross join lateral jsonb_array_elements(a.events) as evt
  cross join lateral jsonb_array_elements(evt->'parameters') as param1
  cross join lateral jsonb_array_elements(evt->'parameters') as param2
where
  application_name = 'mobile'
  and event_names = 'DEVICE_REGISTER_UNREGISTER_EVENT'
  and param1->>'name' = 'DEVICE_ID'
  and param2->>'name' = 'DEVICE_MODEL'
  and time > datetime('now', '-1 day');
```
