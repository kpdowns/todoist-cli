package logout

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenLoggingOutThrowsAnError(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}

	logoutCommand := NewLogoutCommand(mockOutputStream, mockAuthenticationService)
	logoutCommand.Run(logoutCommand, []string{})

	expectedPrompt := errorNotCurrentlyAuthenticated
	actualPrompt := mockOutputStream.String()
	if expectedPrompt != actualPrompt {
		t.Errorf("Received '%s', expected '%s'", actualPrompt, expectedPrompt)
	}
}

func TestIfAuthenticatedAndRevokingAccessTokensReturnsNoErrorsThenNoErrorsAreReturned(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}

	logoutCommand := NewLogoutCommand(mockOutputStream, mockAuthenticationService)
	logoutCommand.Run(logoutCommand, []string{})

	expectedPrompt := successfullyLoggedOut
	actualPrompt := mockOutputStream.String()
	if expectedPrompt != actualPrompt {
		t.Errorf("Received '%s', expected '%s'", actualPrompt, expectedPrompt)
	}
}

func TestIfAuthenticatedAndLoggingOutThenSuccessfullyLoggedOut(t *testing.T) {
	mockOutputStream := &bytes.Buffer{}
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}

	logoutCommand := NewLogoutCommand(mockOutputStream, mockAuthenticationService)
	logoutCommand.Run(logoutCommand, []string{})

	isLoggedOut, _ := mockAuthenticationService.IsAuthenticated()
	if !isLoggedOut {
		t.Error("Expected to have been logged out")
	}
}
