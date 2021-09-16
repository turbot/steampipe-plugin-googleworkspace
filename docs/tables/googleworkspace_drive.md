# Table: googleworkspace_drive

List the user's shared drives.

**NOTE:**

- To filter the resource using `name`, or `created_time`; **you must set `use_domain_admin_access` setting as true** in the where clause, and for that **you must have admin access** in the domain. See [Shared drive-specific query terms](https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties) for information on `use_domain_admin_access` setting.

## Examples

### Basic info

```sql
select
  name,
  id,
  created_time,
  hidden
from
  googleworkspace_drive;
```

### List hidden drives

```sql
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

### List drives that allows access to users outside the domain

```sql
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

### List drives older than 90 days

```sql
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

### List drives using the [query filter](https://developers.google.com/drive/api/v3/ref-search-terms#drive_properties)

```sql
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
