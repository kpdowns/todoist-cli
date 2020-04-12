package authentication

import (
	"fmt"
	"testing"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/authentication/types"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/todoist/responses"
	"github.com/stretchr/testify/assert"
)

func TestGeneratingOauthURL(t *testing.T) {

	t.Run("When generating an Oauth URL, the URL has the expected format", func(t *testing.T) {

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

		assert.Equal(t, expectedURL, generatedOathURL)

	})

}

func TestAccessTokenBeingSetDeterminesWhetherClientIsAuthenticated(t *testing.T) {

	t.Run("When the authentication repository returns an access token, then the client is authenticated", func(t *testing.T) {

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
		assert.True(t, isAuthenticated)

	})

	t.Run("When the authentication repository does not return an access token, then the client is not authenticated", func(t *testing.T) {

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
		assert.False(t, isAuthenticated)

	})

}

func TestSigningIn(t *testing.T) {

	t.Run("When code is set to an empty string, then an error is returned", func(t *testing.T) {

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
		assert.NotNil(t, err)
		assert.Equal(t, errorNoCodeAvailableToSignInWith, err.Error())

	})

	t.Run("When code is not set to an empty string and no errors are received, then the access token is persisted and the client is authenticated", func(t *testing.T) {

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
		assert.Nil(t, err)
		assert.Equal(t, "access-token", mockRepository.AccessToken)

	})

	t.Run("When no valid response is returned from Todoist, then an error is returned", func(t *testing.T) {

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
		assert.NotNil(t, err)

	})

}

func TestSigningOut(t *testing.T) {

	t.Run("When signing out and no errors are received, then no errors are returned and the client is no longer authenticated", func(t *testing.T) {

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

		err := service.SignOut()
		assert.Nil(t, err)
		assert.Equal(t, "", mockRepository.AccessToken)

	})

}
