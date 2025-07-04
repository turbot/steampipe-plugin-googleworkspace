---

title: "Steampipe Table: googleworkspace_admin_reports_login_activity - Query Google Workspace Admin Reports Login Activity Events using SQL"
description: "Allows users to query login activity events from the Google Workspace Admin Reports API, providing insights into sign-in actions and related metadata."
folder: "Cloud Admin Reports"
---
# Table: googleworkspace_admin_reports_login_activity - Query Google Workspace Admin Reports Login Activity Events using SQL

The Google Workspace Admin Reports Login Activity API delivers audit logs for user sign-in activities, including successful and failed login attempts, MFA events, and recovery setting changes. These records help you detect unauthorized access, monitor authentication patterns, and investigate security incidents.

## Table Usage Guide

The `googleworkspace_admin_reports_login_activity` table is designed for tracking authentication events such as logins, OTP verifications, and recovery edits. Use it to identify anomalous sign-in behaviors and maintain compliance.

## Examples

### Basic info

Retrieve all login events in the last 24 hours.

```sql
select
  time,
  actor_email,
  event_name,
  ip_address
from
  googleworkspace_admin_reports_login_activity
where
  time > now() - interval '1 day';
```

### Filter by failed logins

Show all failed login attempts in the past week.

```sql
select
  time,
  actor_email,
  event_name,
  ip_address
from
  googleworkspace_admin_reports_login_activity
where
  event_name = 'login_failure'
  and time > now() - interval '7 days';
```

### User-specific events

Find all login events for a specific user.

```sql
select
  time,
  event_name,
  ip_address
from
  googleworkspace_admin_reports_login_activity
where
  actor_email = 'bob@example.com';
```

### Time window analysis

Query login activities between two timestamps.

```sql
select
  time,
  actor_email,
  event_name
from
  googleworkspace_admin_reports_login_activity
where
  time between '2025-06-10T00:00:00Z' and '2025-06-15T23:59:59Z';
```

### Top IP sources

Identify the top source IPs initiating login events in the last month.

```sql
select
  ip_address,
  count(*) as login_count
from
  googleworkspace_admin_reports_login_activity
where
  time > now() - interval '30 days'
group by
  ip_address
order by
  login_count desc
limit 10;
```
