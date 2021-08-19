# Table: googleworkspace_calendar_my_event

Query information about previous and upcoming events scheduled in a google calendar of current logged in user.

## Examples

### Basic info

```sql
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

### List of events scheduled for tomorrow

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= (current_date + interval '1 day')
  and end_time < (current_date + interval '2 days')
order by start_time;
```

### List of events scheduled in next 4 days

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= current_date
  and end_time <= (current_date + interval '4 days')
order by start_time;
```

### List of events scheduled in current month

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date_trunc('month', current_date)
  and end_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

### List of events scheduled in current week

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time >= date_trunc('week', current_date)
  and end_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```

### List of upcoming events scheduled on every tuesday and thursday

```sql
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
