package authentication

import (
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfAccessTokenIsSetThenIsAlreadyAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AccessToken: "access-token",
	}

	isAuthenticated, _ := mockAuthenticationService.IsAuthenticated()
	if !isAuthenticated {
		t.Error("If the access token is set, then the todoist-cli should be authenticated")
	}
}

func TestIfAccessTokenIsNotSetThenTheTodoistCliIsNotAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AccessToken: "",
	}

	isAuthenticated, _ := mockAuthenticationService.IsAuthenticated()
	if isAuthenticated {
		t.Error("If the access token is not set, then the todoist-cli should not be authenticated")
	}
}
