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
  calendar_id = 'user@domain.com';
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
  googleworkspace_calendar_event
where
  calendar_id = 'user@domain.com'
  and start_time > now()::timestamp
  and end_time < ('now'::timestamp + interval '3 days');
```

### List all out-of-office(OOO) events for upcoming 2 days

```sql
select
  calendar_id,
  start_time
from
  googleworkspace_calendar_event
where
  calendar_id = 'company-calendar@domain.com';
```
