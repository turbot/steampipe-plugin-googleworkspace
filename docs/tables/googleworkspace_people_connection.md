# Table: googleworkspace_people_connection

Query information about contacts owned by the current authenticated user.

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
  googleworkspace_people_connection;
```

### List connections by contact group

```sql
select
  cg.name as contact_group_name,
  conn.given_name as member_name,
  conn.primary_email_address as member_primary_email
from
  googleworkspace_people_connection as conn,
  googleworkspace_people_contact_group as cg
where
  cg.member_resource_names ?| array[conn.resource_name];
```

### List connections within same organization

```sql
select
  display_name,
  primary_email_address,
  org ->> 'name' as organization_name,
  org ->> 'department' as department,
  org ->> 'title' as job_title
from
  googleworkspace_people_connection,
  jsonb_array_elements(organizations) as org
where
  org -> 'metadata' ->> 'primary' = 'true'
  and org ->> 'name' = 'Turbot';
```
