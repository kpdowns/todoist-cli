package authenticate

import (
	"fmt"
	"testing"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/config"
)

func TestGeneratedOauthURLHasTheExpectedFormatFromValuesInConfiguration(t *testing.T) {
	var (
		guid                = guid.NewString()
		clientConfiguration = &config.ClientConfiguration{
			TodoistURL:          "url",
			ClientID:            "clientId",
			RequiredPermissions: "permissions",
		}
		authenticationConfiguration = &config.AuthenticationConfiguration{}
		configuration               = &config.TodoistCliConfiguration{
			Client:         *clientConfiguration,
			Authentication: *authenticationConfiguration,
		}
	)

	expectedURL := fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		clientConfiguration.TodoistURL,
		clientConfiguration.ClientID,
		clientConfiguration.RequiredPermissions,
		guid)
	generatedOathURL := generateOauthURL(configuration, guid)

	if expectedURL != generatedOathURL {
		t.Errorf("The generated Oauth URL was not in the correct format")
	}
}
