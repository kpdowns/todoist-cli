package tasks

import (
	"errors"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

func TestGivenGettingAllTasksWhenAuthenticationServiceReturnsErrorThenErrorIsReturned(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		IsAuthenticatedErrorToReturn: errors.New("test error"),
	}
	mockAPI := &mocks.MockAPI{}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	_, err := taskService.GetAllTasks()
	if err == nil {
		t.Error("Expected error to be returned but didn't get any")
	}

	if err.Error() != errorNotCurrentlyAuthenticated {
		t.Errorf("Expected '%s', got '%s'", errorNotCurrentlyAuthenticated, err.Error())
	}
}

func TestGivenGettingAllTasksWhenApiReturnsErrorThenErrorIsReturned(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncQueryFunction: func(syncQuery requests.SyncQuery) (*responses.SyncQueryResponse, error) {
			return nil, errors.New("Error")
		},
	}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	_, err := taskService.GetAllTasks()
	if err == nil {
		t.Error("Expected error to be returned but didn't get any")
	}

	if err.Error() != errorOccurredDuringSyncOperation {
		t.Errorf("Expected '%s', got '%s'", errorOccurredDuringSyncOperation, err.Error())
	}
}
