package mocks

import (
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

// MockAPI implements the TodoistAPI interface and wraps functions that can be set mocked
type MockAPI struct {
	RevokeAccessTokenFunction func(accessToken string) error
	GetAccessTokenFunction    func(code string) (*responses.AccessToken, error)
	ExecuteSyncQueryFunction  func(syncQuery requests.SyncQuery) (*responses.SyncQueryResponse, error)
}

// RevokeAccessToken executes the function configured for revoking the TodoistAPI access token
func (a *MockAPI) RevokeAccessToken(accessToken string) error {
	if a.RevokeAccessTokenFunction != nil {
		return a.RevokeAccessTokenFunction(accessToken)
	}
	panic("Method call RevokeAccessToken used but not configured")
}

// GetAccessToken executes the function configured for retrieving a TodoistAPI access token
func (a *MockAPI) GetAccessToken(code string) (*responses.AccessToken, error) {
	if a.GetAccessTokenFunction != nil {
		return a.GetAccessTokenFunction(code)
	}
	panic("Method call GetAccessToken used but not configured")
}

// ExecuteSyncQuery executes the function configured for executing sync queries against the Todoist API
func (a *MockAPI) ExecuteSyncQuery(syncQuery requests.SyncQuery) (*responses.SyncQueryResponse, error) {
	if a.ExecuteSyncQueryFunction != nil {
		return a.ExecuteSyncQueryFunction(syncQuery)
	}
	panic("Method call GetAccessToken used but not configured")
}
