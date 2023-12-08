---
title: "Steampipe Table: googleworkspace_drive_my_file - Query Google Workspace Drive Files using SQL"
description: "Allows users to query Google Workspace Drive Files, providing insights into file details, ownership, sharing settings, and more."
---

# Table: googleworkspace_drive_my_file - Query Google Workspace Drive Files using SQL

Google Workspace Drive is a cloud storage service within Google Workspace that allows users to store, sync, and share files. It provides a centralized way to manage files, including documents, spreadsheets, presentations, and more. Google Workspace Drive helps users collaborate on files in real time and access them from any device.

## Table Usage Guide

The `googleworkspace_drive_my_file` table provides insights into files within Google Workspace Drive. As a Google Workspace administrator, explore file-specific details through this table, including ownership, sharing settings, and associated metadata. Utilize it to uncover information about files, such as those shared externally, the permissions associated with each file, and the verification of sharing policies.

## Examples

### Basic info
Discover the segments that have recently been created within your Google Workspace Drive. This can help you keep track of new additions and manage your files more efficiently.

```sql+postgres
select
  name,
  id,
  mime_type,
  created_time
from
  googleworkspace_drive_my_file;
```

```sql+sqlite
select
  name,
  id,
  mime_type,
  created_time
from
  googleworkspace_drive_my_file;
```

### List files shared by other users
Gain insights into files shared with you by other users in Google Workspace. This is particularly useful for understanding the scope of shared resources and identifying who has shared files with you.

```sql+postgres
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

```sql+sqlite
select
  name,
  id,
  mime_type,
  created_time,
  owned_by_me,
  shared,
  json_extract(sharing_user, '$.displayName') as sharing_user_name
from
  googleworkspace_drive_my_file
where
  not owned_by_me
  and sharing_user is not null;
```

### List image or video files modified after a specific date
Analyze your Google Workspace Drive to pinpoint specific image or video files that have been modified after a certain date. This can be useful to track recent changes or updates to media files in your drive.

```sql+postgres
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  query = 'modifiedTime > "2021-08-15T00:00:00" and (mimeType contains "image/" or mimeType contains "video/")';
```

```sql+sqlite
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  strftime('%Y-%m-%dT%H:%M:%S', created_time) > "2021-08-15T00:00:00" and (mime_type like '%image/%' or mime_type like '%video/%');
```

### List files using the [query filter](https://developers.google.com/drive/api/v3/search-files)
Explore which files in your Google Workspace Drive contain the term "Steampipe". This can be particularly useful for quickly locating specific documents or resources related to Steampipe within your workspace.

```sql+postgres
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  query = 'name contains "steampipe"';
```

```sql+sqlite
select
  name,
  id,
  mime_type,
  created_time,
  web_view_link
from
  googleworkspace_drive_my_file
where
  query = 'name contains "steampipe"';
```