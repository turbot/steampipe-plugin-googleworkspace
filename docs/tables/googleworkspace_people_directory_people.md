# Table: googleworkspace_people_directory_people

List domain contacts in the authenticated user's domain directory.

## Examples

### Basic info

```sql
select
  display_name,
  primary_email_address,
  case
    when org -> 'metadata' ->> 'primary' = 'true' then org ->> 'title'
  end as job_title,
  case
    when ph -> 'metadata' ->> 'primary' = 'true' then ph ->> 'value'
  end as primary_contact
from
  googleworkspace_people_directory_people
  left join jsonb_array_elements(organizations) as org on true
  left join jsonb_array_elements(phone_numbers) as ph on true;
```
