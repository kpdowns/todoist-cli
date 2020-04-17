package config

import "github.com/kpdowns/todoist-cli/config/secrets"

// TodoistCliConfiguration contains the configuration required for the TodoistCli to function
type TodoistCliConfiguration struct {
	TodoistURL          string
	ClientID            string
	ClientSecret        string
	RequiredPermissions string
	AppServiceURL       string
	OauthRedirectURL    string
}

// LoadConfiguration loads the configuration file located in ./config.yml, emits an error if the configuration file is not valid
func LoadConfiguration() *TodoistCliConfiguration {
	return &TodoistCliConfiguration{
		TodoistURL:          "https://todoist.com",
		ClientID:            secrets.ClientID,
		ClientSecret:        secrets.ClientSecret,
		RequiredPermissions: "data:read_write,data:delete,project:delete",
		AppServiceURL:       "http://127.0.0.1:8123",
		OauthRedirectURL:    "http://127.0.0.1:8123/oauth/access_token",
	}
}
