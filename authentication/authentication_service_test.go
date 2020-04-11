package authentication

import (
	"errors"
	"fmt"
	"testing"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/authentication/types"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

func TestGeneratedOauthURLHasTheExpectedFormatFromValuesInConfiguration(t *testing.T) {
	guid := guid.NewString()
	mockAPI := &mocks.MockAPI{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "test",
	}
	mockServer := &mocks.MockAuthenticationServer{}
	configuration := &config.TodoistCliConfiguration{
		Client: config.ClientConfiguration{
			TodoistURL:          "url",
			ClientID:            "clientId",
			RequiredPermissions: "permissions",
		},
	}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	expectedURL := fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		configuration.Client.TodoistURL,
		configuration.Client.ClientID,
		configuration.Client.RequiredPermissions,
		guid)
	generatedOathURL := service.GetOauthURL(guid)

	if expectedURL != generatedOathURL {
		t.Errorf("The generated Oauth URL was not in the correct format")
	}
}

func TestIfAccessTokenIsSetThenIsAlreadyAuthenticated(t *testing.T) {
	mockAPI := &mocks.MockAPI{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "test",
	}
	mockServer := &mocks.MockAuthenticationServer{}
	configuration := &config.TodoistCliConfiguration{
		Client: config.ClientConfiguration{
			TodoistURL:          "url",
			ClientID:            "clientId",
			RequiredPermissions: "permissions",
		},
	}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	isAuthenticated, _ := service.IsAuthenticated()
	if !isAuthenticated {
		t.Error("If the access token is set, then the todoist-cli should be authenticated")
	}
}

func TestIfAccessTokenIsNotSetThenTheTodoistCliIsNotAuthenticated(t *testing.T) {
	mockAPI := &mocks.MockAPI{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "",
	}
	mockServer := &mocks.MockAuthenticationServer{}
	configuration := &config.TodoistCliConfiguration{
		Client: config.ClientConfiguration{
			TodoistURL:          "url",
			ClientID:            "clientId",
			RequiredPermissions: "permissions",
		},
	}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	isAuthenticated, _ := service.IsAuthenticated()
	if isAuthenticated {
		t.Error("If the access token is not set, then the todoist-cli should not be authenticated")
	}
}

func TestGivenCodeIsNotSetWhenAttemptingToSignInThenErrorOccurs(t *testing.T) {
	mockAPI := &mocks.MockAPI{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "",
	}
	mockServer := &mocks.MockAuthenticationServer{}
	configuration := &config.TodoistCliConfiguration{
		Client: config.ClientConfiguration{
			TodoistURL:          "url",
			ClientID:            "clientId",
			RequiredPermissions: "permissions",
		},
	}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	err := service.SignIn("")
	if err == nil {
		t.Errorf("Expected error to be received because code was set to empty string")
	}

	if err != nil && err.Error() != errorNoCodeAvailableToSignInWith {
		t.Errorf("Expected '%s', but got '%s'", errorNoCodeAvailableToSignInWith, err.Error())
	}
}

func TestIfCodeIsSetWhenAttemptingToSignInAndApiReturnsNoErrorsThenNoErrorsAreReturnedAndAccessTokenIsSaved(t *testing.T) {
	mockAPI := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: "access-token"}, nil
		},
	}
	mockServer := &mocks.MockAuthenticationServer{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "",
	}
	configuration := &config.TodoistCliConfiguration{}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	err := service.SignIn("test")
	if err != nil {
		t.Errorf("Expected no error to be returned because the mock mockAPI returned no errors")
	}

	isAuthenticated, err := service.IsAuthenticated()
	if err != nil {
		t.Errorf(err.Error())
	}

	if !isAuthenticated {
		t.Errorf("Expected to be authenticated")
	}
}

func TestWhenSigningOutAndNoErrorsAreRecievedFromTheApiThenTheAccessTokenIsRemoved(t *testing.T) {
	mockAPI := &mocks.MockAPI{
		RevokeAccessTokenFunction: func(accessToken string) error {
			return nil
		},
	}
	mockServer := &mocks.MockAuthenticationServer{}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "access-token",
	}
	configuration := &config.TodoistCliConfiguration{}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

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

func TestWhenErrorOccursAfterTheCallbackFromTodoistThenAnErrorIsReturned(t *testing.T) {
	mockAPI := &mocks.MockAPI{
		RevokeAccessTokenFunction: func(accessToken string) error {
			return nil
		},
	}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "access-token",
	}
	mockServer := &mocks.MockAuthenticationServer{
		AuthenticationResponseErrorToReturn: errors.New("test"),
	}
	configuration := &config.TodoistCliConfiguration{}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	err := service.SignIn("code")
	if err.Error() != errorNoAuthCodeReceived {
		t.Errorf("Expected '%s', but received '%s'", errorNoAuthCodeReceived, err.Error())
	}
}

func TestWhenSigningInAndCodeWasNotReturnedFromTodoistThenAnErrorShouldBeReturned(t *testing.T) {
	mockAPI := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: ""}, nil
		},
	}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "access-token",
	}
	mockServer := &mocks.MockAuthenticationServer{
		AuthenticationResponseToReturn: types.AuthenticationResponse{},
	}
	configuration := &config.TodoistCliConfiguration{}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	err := service.SignIn("code")
	if err == nil {
		t.Errorf("Expected '%s', but received nothing", errorNoAccessTokenReceived)
	}
}

func TestWhenSigningInAndNoErrorsOccurThenSuccessfullySignedIn(t *testing.T) {
	mockAPI := &mocks.MockAPI{
		GetAccessTokenFunction: func(code string) (*responses.AccessToken, error) {
			return &responses.AccessToken{AccessToken: "access-token"}, nil
		},
	}
	mockRepository := &mocks.MockAuthenticationRepository{
		AccessToken: "access-token",
	}
	mockServer := &mocks.MockAuthenticationServer{
		AuthenticationResponseToReturn: types.AuthenticationResponse{},
	}
	configuration := &config.TodoistCliConfiguration{}

	service := NewAuthenticationService(mockAPI, mockRepository, *configuration, mockServer)

	err := service.SignIn("code")
	if err != nil {
		t.Errorf("Expected '%s', but received '%s'", errorNoAccessTokenReceived, err.Error())
	}
}
