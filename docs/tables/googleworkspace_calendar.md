# Table: googleworkspace_calendar

Get metadata information for a specific calendar.

You must specify the calendar ID in the where or join clause (`where id=`, `join googleworkspace_calendar on id=`).

## Examples

### Basic info

```sql
select
  summary,
  id,
  timezone
from
  googleworkspace_calendar
where
  id = 'user@domain.com';
```
