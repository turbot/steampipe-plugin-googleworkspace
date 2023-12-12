package googleworkspace

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type googleworkspaceConfig struct {
	CredentialFile        *string `hcl:"credential_file"`
	Credentials           *string `hcl:"credentials"`
	ImpersonatedUserEmail *string `hcl:"impersonated_user_email"`
	TokenPath             *string `hcl:"token_path"`
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
