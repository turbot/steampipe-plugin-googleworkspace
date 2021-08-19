package googleworkspace

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type googleworkspaceConfig struct {
	CredentialFile        *string `cty:"credential_file"`
	ImpersonatedUserEmail *string `cty:"impersonated_user_email"`
	ClientSecretFile      *string `cty:"client_secret_file"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"credential_file": {
		Type: schema.TypeString,
	},
	"impersonated_user_email": {
		Type: schema.TypeString,
	},
	"client_secret_file": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &googleworkspaceConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) googleworkspaceConfig {
	if connection == nil || connection.Config == nil {
		return googleworkspaceConfig{}
	}
	config, _ := connection.Config.(googleworkspaceConfig)
	return config
}
