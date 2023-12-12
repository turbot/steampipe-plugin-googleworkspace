---
title: "Steampipe Table: googleworkspace_people_contact_group - Query Google Workspace People Contact Groups using SQL"
description: "Allows users to query People Contact Groups in Google Workspace, providing insights into contact group details and metadata."
---

# Table: googleworkspace_people_contact_group - Query Google Workspace People Contact Groups using SQL

Google Workspace People Contact Groups is a feature within Google Workspace that allows users to create and manage groups of contacts. It offers a centralized way to organize and manage contacts for various Google Workspace applications, including Gmail, Google Meet, and more. People Contact Groups help users efficiently manage communications and collaborations with groups of people.

## Table Usage Guide

The `googleworkspace_people_contact_group` table provides insights into People Contact Groups within Google Workspace. As an IT administrator or a Google Workspace user, you can explore group-specific details through this table, including group metadata, member count, and member resource names. Use it to manage and organize your Google Workspace contacts more efficiently, such as identifying large groups, finding groups without members, and understanding the structure of your contact groups.

## Examples

### Basic info
Explore the various contact groups in your Google Workspace, including their names and types, to better manage your organization's communication and collaboration. This can be particularly useful in large organizations where understanding group structures is key to efficient operations.

```sql+postgres
select
  resource_name,
  name,
  formatted_name,
  group_type
from
  googleworkspace_people_contact_group;
```

```sql+sqlite
select
  resource_name,
  name,
  formatted_name,
  group_type
from
  googleworkspace_people_contact_group;
```

### List deleted contact groups
Discover the segments that have been removed from your Google Workspace contact groups. This can be useful to track changes and manage your contacts more effectively.

```sql+postgres
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

```sql+sqlite
select
  resource_name,
  name,
  formatted_name,
  deleted
from
  googleworkspace_people_contact_group
where
  deleted = 1;
```

### List members in each contact group
Explore which members belong to each contact group in your Google Workspace, allowing you to better manage communication and collaboration within your organization. This query is particularly useful for gaining insights into group composition and identifying any necessary changes to group membership.

```sql+postgres
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

```sql+sqlite
select
  cg.name as contact_group_name,
  c.given_name as member_name,
  c.primary_email_address as member_primary_email
from
  googleworkspace_people_contact_group as cg,
  json_each(cg.member_resource_names) as m_name,
  googleworkspace_people_contact as c
where
  c.resource_name = m_name.value;
```

### Get member count for each contact group
Discover the segments that have varying membership within your contact groups. This query allows you to analyze the size of each group, helping you to better manage your resources and communications.

```sql+postgres
select
  resource_name,
  name,
  formatted_name,
  member_count
from
  googleworkspace_people_contact_group;
```

```sql+sqlite
select
  resource_name,
  name,
  formatted_name,
  member_count
from
  googleworkspace_people_contact_group;
```