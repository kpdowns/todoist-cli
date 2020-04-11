package list

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/tasks"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenReceiveNotAuthenticatedErrorMessage(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, nil)
	listTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := errorNotCurrentlyAuthenticated
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}

func TestGivenAnAuthenticatedClientWhenThereAreNoTasksThenTextSayingThereAreNoTasksIsWrittenToTheOutputStream(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncQueryFunction: func(syncQuery requests.SyncQuery) (*responses.SyncQueryResponse, error) {
			syncResponse := &responses.SyncQueryResponse{
				Items: []responses.Item{},
			}

			return syncResponse, nil
		},
	}
	mockOutputStream := &bytes.Buffer{}

	taskService := tasks.NewTaskService(mockAPI, mockAuthenticationService)

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)
	listTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := noTasksMessage
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}

func TestGivenAnAuthenticatedClientWhenThereAreTasksThenTheTasksAreWrittenToTheOutputStream(t *testing.T) {
	itemReturned := responses.Item{
		TodoistID: 1,
		Priority:  1,
		Content:   "test",
		Due: &responses.Due{
			DateString: "",
		},
	}
	taskToBeWritten := itemReturned.ToTask()

	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockAPI := &mocks.MockAPI{
		ExecuteSyncQueryFunction: func(syncQuery requests.SyncQuery) (*responses.SyncQueryResponse, error) {
			syncResponse := &responses.SyncQueryResponse{
				Items: []responses.Item{
					itemReturned,
				},
			}

			return syncResponse, nil
		},
	}
	mockOutputStream := &bytes.Buffer{}

	taskService := tasks.NewTaskService(mockAPI, mockAuthenticationService)

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)
	listTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := taskToBeWritten.AsString() + "\n"
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}
