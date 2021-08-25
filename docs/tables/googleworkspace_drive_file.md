# Table: googleworkspace_drive_file

Query information about metadata or content of a file shared in user's domain, or shared in the specified shared drive.

**Note:**

A specific `corpora` must be defined in all queries to this table.

## Examples

### Basic info

```sql
select
  name,
  id,
  mime_type,
  created_time
from
  googleworkspace_drive_file
where
  corpora = 'domain';
```

### List files in a specific shared drive

```sql
select
  name,
  id,
  drive_id,
  mime_type,
  created_time
from
  googleworkspace_drive_file
where
  corpora = 'drive'
  and drive_id = '0AOO1xYnxwu_EUk9PVA';
```

### List files shared in domain

```sql
select
  name,
  id,
  drive_id,
  mime_type,
  created_time
from
  googleworkspace_drive_file
where
  corpora = 'domain';
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
  googleworkspace_drive_file
where
  corpora = 'domain'
  and query = 'modifiedTime > ''2021-08-15T00:00:00'' and (mimeType contains ''image/'' or mimeType contains ''video/'')';
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
  googleworkspace_drive_file
where
  corpora = 'domain'
  and query = 'name contains ''steampipe''';
```
