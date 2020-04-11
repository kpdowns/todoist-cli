package tasks

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenReceiveNotAuthenticatedErrorMessage(t *testing.T) {
	dependencies := &dependencies{
		authenticationService: &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: false,
		},
	}

	err := execute(dependencies)
	if err == nil {
		t.Errorf("Expected that the client would recieve an error but none was received")
	}

	if err != nil && err.Error() != errorNotCurrentlyAuthenticated {
		t.Errorf("Received an error of '%s', but expected '%s'", err.Error(), errorNotCurrentlyAuthenticated)
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

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)

	listTaskCommand.Run(listTaskCommand, []string{})

	textExpectedToBeWrittenToConsole := noTasksMessage + "\n"
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

	taskService := NewTaskService(mockAPI, mockAuthenticationService)

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)

	listTaskCommand.Run(listTaskCommand, []string{})

	textExpectedToBeWrittenToConsole := taskToBeWritten.AsString() + "\n"
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}
