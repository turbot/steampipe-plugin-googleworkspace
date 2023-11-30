---
title: "Steampipe Table: googleworkspace_people_directory_people - Query Google Workspace Directory People using SQL"
description: "Allows users to query Directory People in Google Workspace, specifically the list of people in the directory, providing insights into user profiles and their metadata."
---

# Table: googleworkspace_people_directory_people - Query Google Workspace Directory People using SQL

Google Workspace Directory People is a resource within Google Workspace that allows you to manage and access user profiles in your organization's directory. It provides a centralized way to view and manage information about people in your organization, including their email addresses, phone numbers, and other profile details. Google Workspace Directory People helps you stay informed about the users in your organization and take appropriate actions when necessary.

## Table Usage Guide

The `googleworkspace_people_directory_people` table provides insights into user profiles within Google Workspace Directory People. As an IT administrator, explore user-specific details through this table, including email addresses, phone numbers, and other profile details. Utilize it to uncover information about users, such as their roles, the groups they belong to, and their profile's metadata.

## Examples

### Basic info
This query is useful for obtaining key information about individuals in a Google Workspace directory, such as their display name, primary email address, job title, and primary contact number. It's a practical tool for HR teams or managers who need to quickly access or compile this information for communication or organizational purposes.

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