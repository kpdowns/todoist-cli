package authenticate

import (
	"bytes"
	"fmt"
	"os"
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

func TestIfNotAlreadyAuthenticatedThenTheOauthUrlIsWrittenToTheConsoleSoThatTheUserCanFollowTheLink(t *testing.T) {
	var (
		mockOutputStream = &bytes.Buffer{}
		guid             = guid.NewString()
		configuration    = &config.TodoistCliConfiguration{
			Client: config.ClientConfiguration{
				TodoistURL:          "url",
				ClientID:            "clientId",
				RequiredPermissions: "permissions",
			},
			Authentication: config.AuthenticationConfiguration{},
		}
		dependencies = &dependencies{
			outputStream: mockOutputStream,
			config:       configuration,
			guid:         guid,
		}
	)

	expectedURL := generateOauthURL(configuration, guid)
	textExpectedToBeWrittenToConsole := fmt.Sprintf(oauthInitiationText, expectedURL) + "\n"

	execute(nil, nil, dependencies)

	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected that the url for the user to follow to start authentication is written to the console")
	}
}

func TestIfAlreadyAuthenticatedThenErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	var (
		configuration = &config.TodoistCliConfiguration{
			Authentication: config.AuthenticationConfiguration{
				AccessToken: "access-token",
			},
		}
		dependencies = &dependencies{
			config: configuration,
		}
	)

	err := execute(nil, nil, dependencies)
	if err == nil {
		t.Errorf("If the todoist-cli is already authenticated, the authentication should not be allowed again until after the client has logged out")
	}
}

func TestIfTodoistCliIsNotAlreadyAuthenticatedThenNoAuthenticationErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	var (
		configuration = &config.TodoistCliConfiguration{
			Authentication: config.AuthenticationConfiguration{
				AccessToken: "",
			},
			Client: config.ClientConfiguration{},
		}
		dependencies = &dependencies{
			config:       configuration,
			outputStream: os.Stdout,
		}
	)

	err := execute(nil, nil, dependencies)
	if err != nil && err.Error() != errorAlreadyAuthenticatedText {
		t.Errorf("If the todoist-cli is not authenticated, then there should be no already authenticated error thrown when executing the command")
	}
}
