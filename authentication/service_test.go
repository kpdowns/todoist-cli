package authentication

import (
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

func TestIfAccessTokenIsSetThenIsAlreadyAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		Repository: mocks.MockAuthenticationRepository{
			AccessToken: "test",
		},
	}

	isAuthenticated, _ := mockAuthenticationService.IsAuthenticated()
	if !isAuthenticated {
		t.Error("If the access token is set, then the todoist-cli should be authenticated")
	}
}

func TestIfAccessTokenIsNotSetThenTheTodoistCliIsNotAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		Repository: mocks.MockAuthenticationRepository{
			AccessToken: "",
		},
	}

	isAuthenticated, _ := mockAuthenticationService.IsAuthenticated()
	if isAuthenticated {
		t.Error("If the access token is not set, then the todoist-cli should not be authenticated")
	}
}

func TestIfCodeIsNotSetWhenAttemptingToSignInThenErrorOccurs(t *testing.T) {
	api := &mocks.MockAPI{}
	repository := &mocks.MockAuthenticationRepository{
		AccessToken: "",
	}
	service := NewService(api, repository)

	err := service.SignIn("")
	if err == nil {
		t.Errorf("Expected error to be received because code was set to empty string")
	}

	if err != nil && err.Error() != errorNoCodeAvailableToSignInWith {
		t.Errorf("Expected '%s', but got '%s'", errorNoCodeAvailableToSignInWith, err.Error())
	}
}

func TestIfCodeIsSetWhenAttemptingToSignInAndApiReturnsNoErrorsThenNoErrorsAreReturnedAndAccessTokenIsSaved(t *testing.T) {
	api := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: "access-token"}, nil
		},
	}
	repository := &mocks.MockAuthenticationRepository{
		AccessToken: "",
	}
	service := NewService(api, repository)

	err := service.SignIn("test")
	if err != nil {
		t.Errorf("Expected no error to be returned because the mock api returned no errors")
	}

	isAuthenticated, err := service.IsAuthenticated()
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isAuthenticated {
		t.Errorf("Expected to be authenticated")
	}
}

func TestIfSigningOutAndNoErrorsAreRecievedFromTheApiThenTheAccessTokenIsRemoved(t *testing.T) {
	api := &mocks.MockAPI{
		RevokeAccessTokenFunction: func(accessToken string) error {
			return nil
		},
	}
	repository := &mocks.MockAuthenticationRepository{
		AccessToken: "access-token",
	}
	service := NewService(api, repository)

	isAuthenticated, err := service.IsAuthenticated()
	if err != nil || !isAuthenticated {
		t.Errorf("Expected to be authenticated because we haven't yet signed out")
	}

	err = service.SignOut()
	if err != nil {
		t.Errorf("Expected no errors, but received '%s'", err.Error())
	}

	isAuthenticated, err = service.IsAuthenticated()
	if err != nil || isAuthenticated {
		t.Errorf("Expected to not be authenticated because we have just signed out")
	}
}
