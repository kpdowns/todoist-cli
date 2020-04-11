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
		ExecuteSyncQueryFunction: func(syncQuery requests.Query) (*responses.Query, error) {
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

func TestGivenAddingATaskWhenNotAuthenticatedThenError(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		IsAuthenticatedErrorToReturn: errors.New("test error"),
	}
	mockAPI := &mocks.MockAPI{}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	err := taskService.AddTask("content", "today", 1)
	if err == nil {
		t.Error("Expected error to be returned but didn't get any")
	}

	if err.Error() != errorNotCurrentlyAuthenticated {
		t.Errorf("Expected '%s', got '%s'", errorNotCurrentlyAuthenticated, err.Error())
	}
}

func TestGivenAddingATaskWhenNoContentIsProvidedForTheTaskThenError(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncCommandFunction: func(command requests.Command) error {
			return nil
		},
	}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	err := taskService.AddTask("", "today", 1)
	if err == nil {
		t.Error("Expected error to be returned but didn't get any")
	}

	if err.Error() != errorNoContent {
		t.Errorf("Expected '%s', got '%s'", errorNoContent, err.Error())
	}
}

func TestGivenAddingATaskWhenTheApiReturnsAnErrorThenError(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncCommandFunction: func(command requests.Command) error {
			return errors.New("test error")
		},
	}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	err := taskService.AddTask("content", "today", 1)
	if err == nil {
		t.Error("Expected error to be returned but didn't get any")
	}

	if err.Error() != errorOccurredDuringSyncOperation {
		t.Errorf("Expected '%s', got '%s'", errorOccurredDuringSyncOperation, err.Error())
	}
}

func TestGivenAddingATaskWhenTheApiDoesNotReturnAnErrorThenTheTaskWasAdded(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncCommandFunction: func(command requests.Command) error {
			return nil
		},
	}

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	err := taskService.AddTask("content", "today", 1)
	if err != nil {
		t.Errorf("Expected no error, but got '%s'", err.Error())
	}
}
