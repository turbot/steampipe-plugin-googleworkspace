---
title: "Steampipe Table: googleworkspace_calendar - Query Google Workspace Calendars using SQL"
description: "Allows users to query Google Workspace Calendars, specifically the details of each calendar such as the summary, description, location, timezone, and more."
---

# Table: googleworkspace_calendar - Query Google Workspace Calendars using SQL

Google Workspace Calendar is a time-management and scheduling service developed by Google. It enables users to create and edit events, manage multiple calendars, and share calendars with others. Google Workspace Calendar is fully integrated with other Google services, providing seamless scheduling and collaboration capabilities.

## Table Usage Guide

The `googleworkspace_calendar` table provides insights into calendars within Google Workspace. As a system administrator or IT professional, explore calendar-specific details through this table, including the summary, description, location, and timezone. Utilize it to manage and monitor the usage of calendars, such as those shared with many users, the timezone settings of each calendar, and the description and summary details.

**Important Notes**
- You must specify the `id` in the `where` or join clause (`where id=`, `join googleworkspace_calendar c on c.id=`) to query this table.

## Examples

### Basic info
Explore the basic details of a specific user's Google Workspace Calendar. This can help in understanding the user's time zone and other relevant information to enhance scheduling and coordination.

```sql+postgres
select
  summary,
  id,
  timezone
from
  googleworkspace_calendar
where
  id = 'user@domain.com';
```

```sql+sqlite
select
  summary,
  id,
  timezone
from
  googleworkspace_calendar
where
  id = 'user@domain.com';
```