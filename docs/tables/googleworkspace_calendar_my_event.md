---
title: "Steampipe Table: googleworkspace_calendar_my_event - Query Google Workspace Calendar Events using SQL"
description: "Allows users to query Google Workspace Calendar Events, specifically the events of the authenticated user, providing insights into event details and schedules."
---

# Table: googleworkspace_calendar_my_event - Query Google Workspace Calendar Events using SQL

Google Workspace Calendar is a core service within Google Workspace that allows users to schedule events, invite people, and customize their calendars to suit their needs. It provides a centralized way to manage schedules, meetings, and appointments, helping users stay organized and informed about their upcoming events. Google Workspace Calendar helps you stay updated about your schedule and take necessary actions when needed.

## Table Usage Guide

The `googleworkspace_calendar_my_event` table provides insights into Google Workspace Calendar Events. As an administrator or a user, explore event-specific details through this table, including event start and end times, attendees, and event status. Utilize it to uncover information about your events, such as those with conflicting schedules, attendees' responses to event invitations, and details about recurring events.

## Examples

### Basic info
Gain insights into upcoming events from your Google Workspace Calendar. This query allows you to plan and prioritize by providing a snapshot of the next 10 events, including their summaries and associated hangout links.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
order by start_time
limit 10;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
order by start_time
limit 10;
```

### List events scheduled for tomorrow
Gain insights into your upcoming events by pinpointing the specific ones scheduled for tomorrow, allowing for effective planning and time management.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= (current_date + interval '1 day')
  and start_time < (current_date + interval '2 days')
order by start_time;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date('now', '+1 day')
  and start_time < date('now', '+2 day')
order by start_time;
```

### List events scheduled in next 4 days
Discover the segments that have events scheduled in the coming four days. This is useful for planning and managing your schedule effectively.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= current_date
  and start_time <= (current_date + interval '4 days')
order by start_time;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date('now')
  and start_time <= date('now', '+4 days')
order by start_time;
```

### List events scheduled in current month
Explore which events are scheduled for the current month to manage your time and plan accordingly. This allows you to gain insights into your schedule, helping to avoid clashes and ensure efficient time management.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date_trunc('month', current_date)
  and start_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date('now','start of month')
  and start_time <= date('now','start of month','+1 month')
order by start_time;
```

### List events scheduled in current week
Explore the schedule for the current week to understand your upcoming commitments and plan accordingly. This helps in efficiently managing your time by gaining insights into the events lined up for the week.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date_trunc('week', current_date)
  and start_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date('now', 'weekday 0', '-7 days')
  and start_time < date('now', 'weekday 0')
order by start_time;
```

### List upcoming events scheduled on every Tuesday and Thursday
Discover the segments that have upcoming events scheduled on Tuesdays and Thursdays. This is useful for planning and organizing your week ahead with a focus on those specific days.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time,
  day
from
  googleworkspace_calendar_my_event
where
  extract(dow from start_time) in (2, 4)
  and start_time >= current_date
order by start_time
limit 10;
```

```sql+sqlite
select
  summary,
  hangout_link,
  start_time,
  end_time,
  day
from
  googleworkspace_calendar_my_event
where
  strftime('%w', start_time) in ('2', '4')
  and date(start_time) >= date('now')
order by start_time
limit 10;
```