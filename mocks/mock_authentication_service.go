package mocks

import "github.com/kpdowns/todoist-cli/authentication/types"

// MockAuthenticationService stores the access token in-memory and defines operations that act on it
type MockAuthenticationService struct {
	AuthenticatedStateToReturn   bool
	AccessTokenToReturn          string
	IsAuthenticatedErrorToReturn error
	GetAccessTokenErrorToReturn  error
	SignInErrorToReturn          error
	SignOutErrorToReturn         error
	OathURL                      string
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not by examining the in-memory access token
func (s *MockAuthenticationService) IsAuthenticated() (bool, error) {
	return s.AuthenticatedStateToReturn, s.IsAuthenticatedErrorToReturn
}

// GetAccessToken returns the access token stored in-memory
func (s *MockAuthenticationService) GetAccessToken() (*types.AccessToken, error) {
	return &types.AccessToken{AccessToken: s.AccessTokenToReturn}, s.GetAccessTokenErrorToReturn
}

// SignIn updates the in-memory access token
func (s *MockAuthenticationService) SignIn(string) error {
	return s.SignInErrorToReturn
}

// SignOut resets the access token stored in-memory
func (s *MockAuthenticationService) SignOut() error {
	return s.SignOutErrorToReturn
}

// GetOauthURL returns the configured oauth url
func (s *MockAuthenticationService) GetOauthURL(string) string {
	return s.OathURL
}
