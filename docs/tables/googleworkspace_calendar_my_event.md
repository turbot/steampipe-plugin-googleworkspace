# Table: googleworkspace_calendar_my_event

Query information about previous and upcoming events scheduled in a google calendar of current logged in user.

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
  googleworkspace_calendar_my_event;
```

### List of events scheduled for upcoming 3 days

```sql
select
  calendar_id,
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time > now()::timestamp
  and end_time < ('now'::timestamp + interval '3 days');
```
