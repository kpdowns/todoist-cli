package authentication

import (
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/todoist/responses"
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

func TestIfCodeIsNotSetWhenAttemptingToSignInThenErrorOccurs(t *testing.T) {
	api := &mocks.MockAPI{}
	service := NewService(api)

	err := service.SignIn("")
	if err == nil {
		t.Errorf("Expected error to be received because code was set to empty string")
	}

	if err != nil && err.Error() != errorNoCodeAvailableToSignInWith {
		t.Errorf("Expected '%s', but got '%s'", errorNoCodeAvailableToSignInWith, err.Error())
	}
}

func TestIfCodeIsSetWhenAttemptingToSignInAndApiReturnsNoErrorsThenNoErrorsAreReturned(t *testing.T) {
	api := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: "access-token"}, nil
		},
	}
	service := NewService(api)

	err := service.SignIn("test")
	if err != nil {
		t.Errorf("Expected no error to be returned because the mock api returned no errors")
	}
}
