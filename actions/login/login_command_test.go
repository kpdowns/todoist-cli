package login

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/todoist/responses"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/mocks"

	"github.com/beevik/guid"
)

func TestIfNotAlreadyAuthenticatedThenTheOauthUrlIsWrittenToTheConsoleSoThatTheUserCanFollowTheLink(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	guid := guid.NewString()
	authenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}

	loginCommand := NewLoginCommand(mockOutputStream, authenticationService, guid)
	loginCommand.Run(loginCommand, []string{})

	expectedURL := authenticationService.GetOauthURL(guid)
	expectedText := fmt.Sprintf(oauthInitiationText, expectedURL)
	actualText := mockOutputStream.String()
	actualPromptThatShouldHaveBeenReceived := strings.Split(actualText, "\n")[0]
	if actualPromptThatShouldHaveBeenReceived != expectedText {
		t.Errorf("Expected '%s', but got '%s'", expectedText, actualPromptThatShouldHaveBeenReceived)
	}
}

func TestIfAlreadyAuthenticatedThenErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	guid := guid.NewString()
	authenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}

	loginCommand := NewLoginCommand(mockOutputStream, authenticationService, guid)
	loginCommand.Run(loginCommand, []string{})

	actualText := mockOutputStream.String()
	expectedText := errorAlreadyAuthenticatedText
	if actualText != expectedText {
		t.Errorf("Expected '%s', but got '%s'", errorAlreadyAuthenticatedText, actualText)
	}
}

func TestIfTodoistCliIsNotAlreadyAuthenticatedThenNoAuthenticationErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	guid := guid.NewString()
	authenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}

	loginCommand := NewLoginCommand(mockOutputStream, authenticationService, guid)
	loginCommand.Run(loginCommand, []string{})

	actualText := mockOutputStream.String()
	if actualText == errorAlreadyAuthenticatedText {
		t.Errorf("Expected '%s', but got '%s'", errorAlreadyAuthenticatedText, actualText)
	}
}

func TestIfTodoistCliIsNotAlreadyAuthenticatedAndNoErrorsOccurThenTheUserIsAuthenticated(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}

	mockAPI := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: "access-token"}, nil
		},
	}
	mockAuthenticationRepository := &mocks.MockAuthenticationRepository{}
	mockAuthenticationServer := &mocks.MockAuthenticationServer{}
	config := &config.TodoistCliConfiguration{}

	authenticationService := authentication.NewAuthenticationService(mockAPI, mockAuthenticationRepository, *config, mockAuthenticationServer)

	guid := guid.NewString()

	loginCommand := NewLoginCommand(mockOutputStream, authenticationService, guid)
	loginCommand.Run(loginCommand, []string{})

	isAuthenticated, err := authenticationService.IsAuthenticated()
	if err != nil || !isAuthenticated {
		t.Error("Expected to be logged in")
	}
}
