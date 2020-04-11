package mocks

import "github.com/kpdowns/todoist-cli/authentication/types"

// MockAuthenticationRepository handles persistence of the access token to be used with the Todoist API in-memory
type MockAuthenticationRepository struct {
	AccessToken string
}

// GetAccessToken returns the access token from memory
func (r *MockAuthenticationRepository) GetAccessToken() (*types.AccessToken, error) {
	return &types.AccessToken{AccessToken: r.AccessToken}, nil
}

// DeleteAccessToken deletes the access token in memory
func (r *MockAuthenticationRepository) DeleteAccessToken() error {
	r.UpdateAccessToken("")
	return nil
}

// UpdateAccessToken updates the access token in memory
func (r *MockAuthenticationRepository) UpdateAccessToken(token string) error {
	r.AccessToken = token
	return nil
}
