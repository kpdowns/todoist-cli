package authenticate

import (
	"fmt"
	"testing"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/config"
)

func TestGeneratedOauthURLHasTheExpectedFormatFromValuesInConfiguration(t *testing.T) {
	var (
		guid          = guid.NewString()
		configuration = &config.TodoistCliConfiguration{
			Client: config.ClientConfiguration{
				TodoistURL:          "url",
				ClientID:            "clientId",
				RequiredPermissions: "permissions",
			},
			Authentication: config.AuthenticationConfiguration{},
		}
	)

	expectedURL := fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		configuration.Client.TodoistURL,
		configuration.Client.ClientID,
		configuration.Client.RequiredPermissions,
		guid)
	generatedOathURL := generateOauthURL(configuration, guid)

	if expectedURL != generatedOathURL {
		t.Errorf("The generated Oauth URL was not in the correct format")
	}
}
