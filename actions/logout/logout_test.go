package logout

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenLoggingOutThrowsAnError(t *testing.T) {
	var (
		mockOutputStream = &bytes.Buffer{}
		dependencies     = &dependencies{
			outputStream: mockOutputStream,
			authenticationService: &mocks.MockAuthenticationService{
				AccessToken: "",
			},
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
		dependencies     = &dependencies{
			outputStream: mockOutputStream,
			api: &mocks.MockAPI{
				RevokeAccessTokenFunction: func(accessToken string) error { return nil },
			},
			authenticationService: &mocks.MockAuthenticationService{
				AccessToken: "access-token",
			},
		}
	)

	err := execute(dependencies)
	if err != nil {
		t.Errorf("When the todoist-cli is authenticated and no errors are returned when revoking the access tokens, then no errors should be returned")
	}
}
