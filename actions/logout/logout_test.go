package logout

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"

	"github.com/kpdowns/todoist-cli/config"
)

func TestIfNotAuthenticatedThenLoggingOutThrowsAnError(t *testing.T) {
	var (
		mockOutputStream = &bytes.Buffer{}
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
		}
	)

	err := execute(dependencies)
	if err != nil && err.Error() != errorNotCurrentlyAuthenticated {
		t.Errorf("When the todoist-cli is not currently authenticated, an error must be thrown")
	}
}

func TestIfAuthenticatedAndRevokingAccessTokensReturnsNoErrorsThenNoErrorsAreReturned(t *testing.T) {
	var (
		mockOutputStream = &bytes.Buffer{}
		configuration    = &config.TodoistCliConfiguration{
			Client: config.ClientConfiguration{
				TodoistURL:          "url",
				ClientID:            "clientId",
				RequiredPermissions: "permissions",
			},
			Authentication: config.AuthenticationConfiguration{
				AccessToken: "access-token",
			},
		}
		dependencies = &dependencies{
			outputStream: mockOutputStream,
			config:       configuration,
			api: &mocks.MockAPI{
				RevokeAccessTokenFunction: func() error { return nil },
			},
		}
	)

	err := execute(dependencies)
	if err != nil {
		t.Errorf("When the todoist-cli is authenticated and no errors are returned when revoking the access tokens, then no errors should be returned")
	}
}
