# Table: googleworkspace_calendar_my_event

List previous and upcoming events scheduled in your calendar.

To query events in any calendar, use the `googleworkspace_calendar_event` table.

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

### List events scheduled for tomorrow

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
  and start_time < (current_date + interval '2 days')
order by start_time;
```

### List events scheduled in next 4 days

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
  and start_time <= (current_date + interval '4 days')
order by start_time;
```

### List events scheduled in current month

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
  and start_time <= date_trunc('month', current_date) + interval '1 month'
order by start_time;
```

### List events scheduled in current week

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
  and start_time < (date_trunc('week', current_date) + interval '7 days')
order by start_time;
```

### List upcoming events scheduled on every Tuesday and Thursday

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
