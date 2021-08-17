![image](https://hub.steampipe.io/images/plugins/turbot/googleworkspace-social-graphic.png)

# Google Workspace Plugin for Steampipe

Use SQL to query users, groups, org units and more from your Google Workspace.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/googleworkspace)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/googleworkspace/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-googleworkspace/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install googleworkspace
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/googleworkspace#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/googleworkspace#configuration).

Run a query:

```sql
select
  summary,
  hangout_link,
  start_time,
  end_time
from
  googleworkspace_calendar_my_event
where
  start_time > now()::timestamp
  and end_time < ('now'::timestamp + interval '1 day');
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-googleworkspace
cd steampipe-plugin-googleworkspace
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/googleworkspace.spc
```

Try it!

```
steampipe query
> .inspect googleworkspace
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-googleworkspace/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Google Workspace Plugin](https://github.com/turbot/steampipe-plugin-googleworkspace/labels/help%20wanted)
