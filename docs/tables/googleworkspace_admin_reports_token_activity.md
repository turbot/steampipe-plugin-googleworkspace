---
title: "Steampipe Table: googleworkspace_admin_reports_token_activity - Query Google Workspace Admin Reports Token Activity Events using SQL"
description: "Allows users to query token activity events from the Google Workspace Admin Reports API, providing insights into OAuth and API token usage and revocation events."
---

# Table: googleworkspace_admin_reports_token_activity - Query Google Workspace Admin Reports Token Activity Events using SQL

The Google Workspace Admin Reports Token Activity API captures events related to OAuth and API tokens—such as token authorizations, grant events, and revocations—performed in your Workspace domain. These logs help you audit third-party app integrations and detect unauthorized token usage.

## Table Usage Guide

Use the `googleworkspace_admin_reports_token_activity` table to:

* Monitor OAuth consent and token grant events.
* Detect and investigate token revocations or suspicions of token misuse.

## Examples

### Basic info

Retrieve token activity events in the last 24 hours.

```sql
select
  time,
  actor_email,
  event_name,
  app_name
from
  googleworkspace_admin_reports_token_activity
where
  time > now() - interval '1 day';
```

### Token events for a specific application

Identify all events related to the "Google Chrome" app over the last week.

```sql
select
  time,
  actor_email,
  event_name,
  app_name
from
  googleworkspace_admin_reports_token_activity
where
  app_name = 'Google Chrome'
  and time > now() - interval '7 days';
```

### App usage

List all apps that the members of your organization connected to through Google authentication last month.

```sql
select distinct
  app_name
from
  googleworkspace_admin_reports_token_activity
where
  time > now() - interval '1 month';
```

### Custom time window audit

Query token activity between two specific timestamps.

```sql
select
  time,
  actor_email,
  event_name,
  app_name
from
  googleworkspace_admin_reports_token_activity
where
  time between '2025-06-01T00:00:00Z' and '2025-06-07T23:59:59Z';
```
