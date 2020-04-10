package mocks

import "github.com/kpdowns/todoist-cli/todoist"

// MockAuthenticationService stores the access token in-memory and defines operations that act on it
type MockAuthenticationService struct {
	Repository MockAuthenticationRepository
	API        todoist.API
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not by examining the in-memory access token
func (s *MockAuthenticationService) IsAuthenticated() (bool, error) {
	accessToken, err := s.Repository.GetAccessToken()
	if err != nil {
		return false, err
	}

	isAccessTokenEmptyString := accessToken == ""
	if isAccessTokenEmptyString {
		return false, nil
	}

	return true, nil
}

// GetAccessToken returns the access token stored in-memory
func (s *MockAuthenticationService) GetAccessToken() (string, error) {
	accessToken, err := s.Repository.GetAccessToken()
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// SignIn updates the in-memory access token
func (s *MockAuthenticationService) SignIn(code string) error {

	accessToken, err := s.API.GetAccessToken(code)
	if err != nil {
		return err
	}

	s.Repository.UpdateAccessToken(accessToken.AccessToken)
	return nil
}

// SignOut resets the access token stored in-memory
func (s *MockAuthenticationService) SignOut() error {
	s.Repository.DeleteAccessToken()
	return nil
}
