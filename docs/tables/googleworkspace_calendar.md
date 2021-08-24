# Table: googleworkspace_calendar

Query information about metadata of the specified calendar.

**Note:**

- A specific `id` of the calendar must be defined in all queries to this table.

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
  id = 'subhajit@turbot.com';
```
