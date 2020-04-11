package mocks

import "github.com/kpdowns/todoist-cli/authentication/types"

// MockAuthenticationServer provides mocked functionality to handle Oauth responses from Todoist
type MockAuthenticationServer struct {
	AuthenticationResponseErrorToReturn error
	AuthenticationResponseToReturn      types.AuthenticationResponse
}

// StartTemporaryServerToListenForResponse returns the configured state and response
func (s *MockAuthenticationServer) StartTemporaryServerToListenForResponse(guid string) (*types.AuthenticationResponse, error) {
	return &s.AuthenticationResponseToReturn, s.AuthenticationResponseErrorToReturn
}
