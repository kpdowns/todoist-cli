package mocks

// MockAuthenticationService stores the access token in-memory and defines operations that act on it
type MockAuthenticationService struct {
	AccessToken string
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

// SaveAccessToken updates the in-memory access token
func (service *MockAuthenticationService) SaveAccessToken(accessToken string) error {
	service.AccessToken = accessToken
	return nil
}

// DeleteAccessToken resets the access token stored in-memory
func (service *MockAuthenticationService) DeleteAccessToken() error {
	service.AccessToken = ""
	return nil
}
