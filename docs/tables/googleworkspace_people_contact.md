---
title: "Steampipe Table: googleworkspace_people_contact - Query Google Workspace Contacts using SQL"
description: "Allows users to query Contacts in Google Workspace, specifically the contact details, providing insights into personal and professional information associated with Google Workspace users."
---

# Table: googleworkspace_people_contact - Query Google Workspace Contacts using SQL

Google Workspace Contacts is a resource within Google Workspace that allows users to save and organize contact information. It provides a centralized way to manage contact details for various Google Workspace users, including names, email addresses, phone numbers, and more. Google Workspace Contacts helps users stay connected and maintain a comprehensive directory of contacts within their organization.

## Table Usage Guide

The `googleworkspace_people_contact` table provides insights into contact details within Google Workspace. As a system administrator, explore contact-specific details through this table, including names, email addresses, phone numbers, and associated metadata. Utilize it to uncover information about contacts, such as their professional affiliations, communication details, and the verification of associated metadata.

## Examples

### Basic info
Explore the basic information of your Google Workspace contacts, such as their names and primary email addresses. This can help you assess and understand your contact organization structure better.

```sql+postgres
select
  resource_name,
  display_name,
  given_name,
  primary_email_address,
  jsonb_pretty(organizations)
from
  googleworkspace_people_contact;
```

```sql+sqlite
select
  resource_name,
  display_name,
  given_name,
  primary_email_address,
  organizations
from
  googleworkspace_people_contact;
```

### List contacts by contact group
Explore which contacts belong to specific groups in your Google Workspace. This can be useful for managing communication within teams or identifying groups for targeted outreach.

```sql+postgres
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

```sql+sqlite
Error: SQLite does not support array operations and the '?' operator used in PostgreSQL.
```

### List contacts belogning to the same organization
Explore which contacts in your Google Workspace belong to the same organization, allowing you to better categorize and manage your professional networks. This is particularly useful for identifying all contacts associated with a specific business entity, such as 'Turbot'.

```sql+postgres
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

```sql+sqlite
select
  display_name,
  primary_email_address,
  json_extract(org.value, '$.name') as organization_name,
  json_extract(org.value, '$.department') as department,
  json_extract(org.value, '$.title') as job_title
from
  googleworkspace_people_contact,
  json_each(organizations) as org
where
  json_extract(org.value, '$.metadata.primary') = 'true'
  and json_extract(org.value, '$.name') = 'Turbot';
```