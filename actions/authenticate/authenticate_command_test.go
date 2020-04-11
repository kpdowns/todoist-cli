package authenticate

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"

	"github.com/beevik/guid"
)

func TestIfNotAlreadyAuthenticatedThenTheOauthUrlIsWrittenToTheConsoleSoThatTheUserCanFollowTheLink(t *testing.T) {
	var (
		mockOutputStream      = &bytes.Buffer{}
		guid                  = guid.NewString()
		outputStream          = mockOutputStream
		authenticationService = &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: false,
		}
	)

	authenticateCommand := NewAuthenticateCommand(outputStream, authenticationService, guid)
	authenticateCommand.Run(authenticateCommand, []string{})

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
	authenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}

	authenticateCommand := NewAuthenticateCommand(mockOutputStream, authenticationService, "")
	authenticateCommand.Run(authenticateCommand, []string{})

	actualText := mockOutputStream.String()
	expectedText := errorAlreadyAuthenticatedText
	if actualText != expectedText {
		t.Errorf("Expected '%s', but got '%s'", errorAlreadyAuthenticatedText, actualText)
	}
}

func TestIfTodoistCliIsNotAlreadyAuthenticatedThenNoAuthenticationErrorIsReturnedWhenExecutingCommand(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	authenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}

	authenticateCommand := NewAuthenticateCommand(mockOutputStream, authenticationService, "")
	authenticateCommand.Run(authenticateCommand, []string{})

	actualText := mockOutputStream.String()
	if actualText == errorAlreadyAuthenticatedText {
		t.Errorf("Expected '%s', but got '%s'", errorAlreadyAuthenticatedText, actualText)
	}
}
