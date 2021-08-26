# Table: googleworkspace_drive

List the user's shared drives.

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
  and use_domain_admin_access = true;
```
