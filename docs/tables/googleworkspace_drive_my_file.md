# Table: googleworkspace_drive_my_file

Get metadata for files created by, opened by, or shared directly with the user.

## Examples

### Basic info

```sql
select
  name,
  id,
  mime_type,
  created_time
from
  googleworkspace_drive_my_file;
```

### List files shared by other users

```sql
select
  name,
  id,
  mime_type,
  created_time,
  owned_by_me,
  shared,
  sharing_user ->> 'displayName' as sharing_user_name
from
  googleworkspace_drive_my_file
where
  not owned_by_me
  and sharing_user is not null;
```

### List image or video files modified after a specific date

```sql
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  query = 'modifiedTime > ''2021-08-15T00:00:00'' and (mimeType contains ''image/'' or mimeType contains ''video/'')';
```

### List files using the [query filter](https://developers.google.com/drive/api/v3/search-files)

```sql
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  query = 'name contains ''steampipe''';
```
