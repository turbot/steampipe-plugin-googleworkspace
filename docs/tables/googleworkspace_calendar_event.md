# Table: googleworkspace_calendar_event

Query information about previous and upcoming events scheduled in a specified google calendar.

**Note:**

- A specific `calendar_id` must be defined in all queries to this table.

## Examples

### Basic info

```sql
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

### List of events scheduled in next 3 days

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'user@domain.com'
  and start_time >= now()::timestamp
  and end_time < ('now'::timestamp + interval '3 days')
order by start_time;
```

### List all out-of-office(OOO) events in next 30 days

```sql
select
  summary,
  start_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com'
  and event_type = 'outOfOffice'
  and start_time >= current_date
  and end_time < current_date + interval '30 days'
order by start_time;
```

### List upcoming indian holidays in next 30 days

```sql
select
  summary,
  start_time,
  day
from
  googleworkspace_calendar_event
where
  calendar_id = 'en.indian#holiday@group.v.calendar.google.com'
  and start_time >= current_date
  and end_time < current_date + interval '30 days'
order by start_time;
```
