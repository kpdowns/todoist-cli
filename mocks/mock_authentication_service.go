package mocks

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
func (s *MockAuthenticationService) GetAccessToken() (string, error) {
	return s.AccessTokenToReturn, s.GetAccessTokenErrorToReturn
}

// SignIn updates the in-memory access token
func (s *MockAuthenticationService) SignIn(code string) error {
	return s.SignInErrorToReturn
}

// SignOut resets the access token stored in-memory
func (s *MockAuthenticationService) SignOut() error {
	return s.SignOutErrorToReturn
}

// GetOauthURL returns the configured oauth url
func (s *MockAuthenticationService) GetOauthURL(guid string) string {
	return s.OathURL
}
