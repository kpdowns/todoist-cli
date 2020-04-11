package mocks

import (
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

// MockAPI implements the TodoistAPI interface and wraps functions that can be set mocked
type MockAPI struct {
	RevokeAccessTokenFunction  func(accessToken string) error
	GetAccessTokenFunction     func(code string) (*responses.AccessToken, error)
	ExecuteSyncQueryFunction   func(query requests.Query) (*responses.Query, error)
	ExecuteSyncCommandFunction func(command requests.Command) error
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
func (a *MockAPI) ExecuteSyncQuery(query requests.Query) (*responses.Query, error) {
	if a.ExecuteSyncQueryFunction != nil {
		return a.ExecuteSyncQueryFunction(query)
	}
	panic("Method call ExecuteSyncQuery used but not configured")
}

// ExecuteSyncCommand executes the function configured for executing sync commands against Todoist
func (a *MockAPI) ExecuteSyncCommand(command requests.Command) error {
	if a.ExecuteSyncCommandFunction != nil {
		return a.ExecuteSyncCommandFunction(command)
	}
	panic("Method call ExecuteSyncCommand used but not configured")
}
