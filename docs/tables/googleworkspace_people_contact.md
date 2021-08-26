# Table: googleworkspace_people_contact

List contacts for the authenticated user.

## Examples

### Basic info

```sql
select
  resource_name,
  display_name,
  given_name,
  primary_email_address,
  jsonb_pretty(organizations)
from
  googleworkspace_people_contact;
```

### List contacts by contact group

```sql
select
  cg.name as contact_group_name,
  c.given_name as member_name,
  c.primary_email_address as member_primary_email
from
  googleworkspace_people_contact as c,
  googleworkspace_people_contact_group as cg
where
  cg.member_resource_names ?| array[c.resource_name];
```

### List contacts belogning to the same organization

```sql
select
  display_name,
  primary_email_address,
  org ->> 'name' as organization_name,
  org ->> 'department' as department,
  org ->> 'title' as job_title
from
  googleworkspace_people_contact,
  jsonb_array_elements(organizations) as org
where
  org -> 'metadata' ->> 'primary' = 'true'
  and org ->> 'name' = 'Turbot';
```
