# Table: googleworkspace_people_contact_group

Query information about contact groups owned by the current authenticated user.

## Examples

### Basic info

```sql
select
  resource_name,
  name,
  formatted_name,
  group_type
from
  googleworkspace_people_contact_group;
```

### List of deleted contact groups

```sql
select
  resource_name,
  name,
  formatted_name,
  deleted
from
  googleworkspace_people_contact_group
where
  deleted;
```

### List members in contact groups

```sql
select
  cg.name as contact_group_name,
  conn.given_name as member_name,
  conn.primary_email_address as member_primary_email
from
  googleworkspace_people_contact_group as cg,
  jsonb_array_elements_text(member_resource_names) as m_name,
  googleworkspace_people_connection as conn
where
  conn.resource_name = m_name;
```

### List contact groups with member count

```sql
select
  resource_name,
  name,
  formatted_name,
  member_count
from
  googleworkspace_people_contact_group;
```
