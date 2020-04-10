package mocks

import "github.com/kpdowns/todoist-cli/todoist/responses"

// MockAPI implements the TodoistAPI interface and wraps functions that can be set mocked
type MockAPI struct {
	RevokeAccessTokenFunction func(accessToken string) error
	GetAccessTokenFunction    func(code string) (*responses.AccessToken, error)
}

// RevokeAccessToken executes the function configured for revoking the TodoistAPI access token
func (a *MockAPI) RevokeAccessToken(accessToken string) error {
	if a.RevokeAccessTokenFunction != nil {
		return a.RevokeAccessTokenFunction(accessToken)
	}
	panic("Method call used but not configured")
}

// GetAccessToken executes the function configured for retrieving a TodoistAPI access token
func (a *MockAPI) GetAccessToken(code string) (*responses.AccessToken, error) {
	if a.GetAccessTokenFunction != nil {
		return a.GetAccessTokenFunction(code)
	}
	panic("Method call used but not configured")
}
