# Table: googleworkspace_people_contact_group

List contact groups owned by the authenticated user.

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

### List deleted contact groups

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

### List members in each contact group

```sql
select
  cg.name as contact_group_name,
  c.given_name as member_name,
  c.primary_email_address as member_primary_email
from
  googleworkspace_people_contact_group as cg,
  jsonb_array_elements_text(member_resource_names) as m_name,
  googleworkspace_people_contact as c
where
  c.resource_name = m_name;
```

### Get member count for each contact group

```sql
select
  resource_name,
  name,
  formatted_name,
  member_count
from
  googleworkspace_people_contact_group;
```
