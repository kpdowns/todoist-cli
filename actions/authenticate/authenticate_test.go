package authenticate

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/actions/authenticate/types"
	"github.com/kpdowns/todoist-cli/config"
)

var authenticationFunctionReturningEmptyObject = func(csrfGUID string) (*types.AuthenticationResponse, error) {
	return &types.AuthenticationResponse{}, nil
}

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

	execute(dependencies, authenticationFunctionReturningEmptyObject)

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

	err := execute(dependencies, authenticationFunctionReturningEmptyObject)
	if err == nil {
		t.Errorf("If the todoist-cli is already authenticated, the authentication should not be allowed again until after the client has logged out")
	}
}

func TestIfTodoistCliIsNotAlreadyAuthenticatedThenNoAuthenticationErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	var (
		configuration = &config.TodoistCliConfiguration{
			Authentication: config.AuthenticationConfiguration{},
			Client:         config.ClientConfiguration{},
		}
		dependencies = &dependencies{
			config:       configuration,
			outputStream: os.Stdout,
		}
	)

	err := execute(dependencies, authenticationFunctionReturningEmptyObject)
	if err != nil && err.Error() == errorAlreadyAuthenticatedText {
		t.Errorf("If the todoist-cli is not authenticated, then there should be no already authenticated error thrown when executing the command")
	}
}

func TestIfErrorOccursAfterTheCallbackFromTodoistThenThatErrorIsReturned(t *testing.T) {
	var (
		configuration = &config.TodoistCliConfiguration{
			Authentication: config.AuthenticationConfiguration{},
			Client:         config.ClientConfiguration{},
		}
		dependencies = &dependencies{
			config:       configuration,
			outputStream: os.Stdout,
		}
	)

	expectedError := errors.New("Error occurring during callback")
	authenticationFunctionReturningAnError := func(csrfGUID string) (*types.AuthenticationResponse, error) {
		return nil, expectedError
	}

	actualError := execute(dependencies, authenticationFunctionReturningAnError)

	if actualError.Error() != expectedError.Error() {
		t.Errorf("The error that we receive when listening for the callback from Todoist should be returned if it occurs")
	}
}
func TestIfCodeWasNotReturnedFromTodoistThenAnErrorShouldBeReturned(t *testing.T) {
	var (
		configuration = &config.TodoistCliConfiguration{
			Authentication: config.AuthenticationConfiguration{},
			Client:         config.ClientConfiguration{},
		}
		dependencies = &dependencies{
			config:       configuration,
			outputStream: os.Stdout,
		}
	)

	err := execute(dependencies, authenticationFunctionReturningEmptyObject)
	if err.Error() != errorNoAuthCodeReceived {
		t.Errorf("An error stating that no authentication code was received back from Todoist should be returned if we did not receive an authentication code")
	}
}
