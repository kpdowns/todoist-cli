package mocks

import "github.com/kpdowns/todoist-cli/todoist"

// MockAuthenticationService stores the access token in-memory and defines operations that act on it
type MockAuthenticationService struct {
	AccessToken string
	API         todoist.API
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not by examining the in-memory access token
func (service *MockAuthenticationService) IsAuthenticated() (bool, error) {
	isAccessTokenEmptyString := service.AccessToken == ""
	if isAccessTokenEmptyString {
		return false, nil
	}

	return true, nil
}

// GetAccessToken returns the access token stored in-memory
func (service *MockAuthenticationService) GetAccessToken() (string, error) {
	return service.AccessToken, nil
}

// SignIn updates the in-memory access token
func (service *MockAuthenticationService) SignIn(code string) error {

	accessToken, err := service.API.GetAccessToken(code)
	if err != nil {
		return err
	}

	service.AccessToken = accessToken.AccessToken
	return nil
}

// SignOut resets the access token stored in-memory
func (service *MockAuthenticationService) SignOut() error {
	service.AccessToken = ""
	return nil
}
