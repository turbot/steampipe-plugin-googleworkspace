---
title: "Steampipe Table: googleworkspace_drive - Query Google Workspace Drives using SQL"
description: "Allows users to query Google Workspace Drives, providing detailed information about each drive, including the drive's type, name, theme, and other metadata."
---

# Table: googleworkspace_drive - Query Google Workspace Drives using SQL

Google Workspace Drive is a cloud-based storage solution provided by Google. It allows users to store, share, and synchronize files across devices. Google Workspace Drive supports various file formats and offers features like real-time collaboration, powerful search, and seamless integration with Google Workspace apps.

## Table Usage Guide

The `googleworkspace_drive` table provides insights into Drives within Google Workspace. As a system administrator or a DevOps engineer, explore specific details about each drive through this table, including the drive's type, name, theme, and other metadata. Utilize it to understand the distribution and organization of files, monitor the usage of storage space, and manage access control for sensitive files.

**Important Notes**
- To filter the resource using `name`, or `created_time` you must set `use_domain_admin_access` setting as true** in the where clause, and for that you must have admin access in the domain. See [Shared drive-specific query terms](https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties) for information on `use_domain_admin_access` setting.

## Examples

### Basic info
Discover the segments that are hidden within your Google Workspace Drive. This allows you to assess elements within your Drive, such as identifying instances where files or folders have been hidden, to better manage your resources.

```sql+postgres
select
  name,
  id,
  created_time,
  hidden
from
  googleworkspace_drive;
```

```sql+sqlite
select
  name,
  id,
  created_time,
  hidden
from
  googleworkspace_drive;
```

### List hidden drives
Discover the segments that contain hidden drives within the Google Workspace. This can be beneficial in identifying any potentially unauthorized or suspicious activity.

```sql+postgres
select
  name,
  id,
  created_time,
  hidden
from
  googleworkspace_drive
where
  hidden;
```

```sql+sqlite
select
  name,
  id,
  created_time,
  hidden
from
  googleworkspace_drive
where
  hidden = 1;
```

### List drives that allows access to users outside the domain
Explore which Google Workspace drives permit access to users outside of the domain. This is useful for assessing potential security risks and ensuring data is shared appropriately.

```sql+postgres
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  not domain_users_only;
```

```sql+sqlite
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  domain_users_only = 0;
```

### List drives older than 90 days
Determine the areas in which Google Workspace drives have been in use for more than 90 days. This can be useful for identifying potentially outdated or unused resources, contributing to more efficient resource management.

```sql+postgres
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  created_time <= current_date - interval '90 days'
  and use_domain_admin_access;
```

```sql+sqlite
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  created_time <= date('now','-90 day')
  and use_domain_admin_access;
```

### List drives using the [query filter](https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties)
Explore which Google Workspace drives were created after August 1, 2021, and contain 'steampipe' in their names. This is useful for administrators who want to monitor the creation and naming of drives within their domain.

```sql+postgres
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  query = 'createdTime > ''2021-08-01T07:00:00'' and name contains ''steampipe'''
  and use_domain_admin_access;
```

```sql+sqlite
select
  name,
  id,
  created_time,
  domain_users_only
from
  googleworkspace_drive
where
  query = 'createdTime > ''2021-08-01T07:00:00'' and name contains ''steampipe'''
  and use_domain_admin_access;
```