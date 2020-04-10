package tasks

import (
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenReceiveNotAuthenticatedErrorMessage(t *testing.T) {
	dependencies := &dependencies{
		authenticationService: &mocks.MockAuthenticationService{
			Repository: mocks.MockAuthenticationRepository{
				AccessToken: "",
			},
		},
	}

	err := execute(dependencies)
	if err == nil {
		t.Errorf("Expected that the client would recieve an error but none was received")
	}

	if err != nil && err.Error() != errorNotCurrentlyAuthenticated {
		t.Errorf("Received an error of '%s', but expected '%s'", err.Error(), errorNotCurrentlyAuthenticated)
	}
}
