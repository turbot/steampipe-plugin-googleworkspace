---
title: "Steampipe Table: googleworkspace_calendar_event - Query Google Workspace Calendar Events using SQL"
description: "Allows users to query Google Workspace Calendar Events, specifically the event details, providing insights into event schedules, attendees, and more."
---

# Table: googleworkspace_calendar_event - Query Google Workspace Calendar Events using SQL

Google Workspace Calendar Events is a feature within Google Workspace that allows you to schedule, manage, and track events. It provides a centralized way to manage events for various Google Workspace users, including details about the event, attendees, and more. Google Workspace Calendar Events helps you stay informed about the event schedules and take appropriate actions when needed.

## Table Usage Guide

The `googleworkspace_calendar_event` table provides insights into Calendar Events within Google Workspace. As an IT administrator, explore event-specific details through this table, including event schedules, attendees, and associated metadata. Utilize it to uncover information about events, such as those with multiple attendees, the status of the attendees, and the verification of event details.

**Important Notes**
- You must specify the `calendar_id` in the `where` or join clause (`where calendar_id=`, `join googleworkspace_calendar_event e on e.calendar_id=`) to query this table.

## Examples

### Basic info
Gain insights into upcoming events on a specific Google Workspace calendar. This query helps in planning and scheduling by providing details like the event summary, hangout link, start time, and end time of the ten soonest events.

```sql+postgres
select
  calendar_id,
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'user@domain.com'
order by start_time
limit 10;
```

```sql+sqlite
select
  calendar_id,
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'user@domain.com'
order by start_time
limit 10;
```

### List events scheduled in next 4 days
Identify upcoming events in your company's calendar for the next four days. This allows you to stay updated with the scheduled activities, their timings, and corresponding links, thereby aiding in effective time management and planning.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= current_date
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
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= date('now')
  and start_time <= date('now', '+4 days')
order by start_time;
```

### List events scheduled in current month
Explore the scheduled events for the current month to stay updated on important dates and activities. This is beneficial in managing your time and ensuring no important event is overlooked.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= date_trunc('month', current_date)
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
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= date('now','start of month')
  and start_time <= date('now','start of month','+1 month')
order by start_time;
```

### List events scheduled in current week
Explore which company events are scheduled for the upcoming week. This query is useful for keeping track of upcoming events and meetings in your organization.

```sql+postgres
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= date_trunc('week', current_date)
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
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and start_time >= date('now', 'weekday 0', '-7 days')
  and start_time < date('now', 'weekday 0')
order by start_time;
```

### List out of office (OOO) events in next 30 days
Discover the upcoming out of office events in the next month. This is useful for planning and coordinating team schedules and resources.

```sql+postgres
select
  summary,
  start_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and event_type = 'outOfOffice'
  and start_time >= current_date
  and start_time < current_date + interval '30 days'
order by start_time;
```

```sql+sqlite
select
  summary,
  start_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and event_type = 'outOfOffice'
  and start_time >= date('now')
  and start_time < date('now', '+30 days')
order by start_time;
```

### List upcoming Indian holidays in next 30 days
Discover the upcoming holidays in India within the next month. This can be useful for planning activities, scheduling events, or understanding potential business impacts due to national holidays.

```sql+postgres
select
  summary,
  start_time,
  day
from
  googleworkspace_calendar_event
where
  calendar_id = 'en.indian#holiday@group.v.calendar.google.com'
  and start_time >= current_date
  and start_time < current_date + interval '30 days'
order by start_time;
```

```sql+sqlite
select
  summary,
  start_time,
  day
from
  googleworkspace_calendar_event
where
  calendar_id = 'en.indian#holiday@group.v.calendar.google.com'
  and start_time >= date('now')
  and start_time < date('now', '+30 days')
order by start_time;
```