package main

import (
	"github.com/turbot/steampipe-plugin-googleworkspace/googleworkspace"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: googleworkspace.Plugin})
}
