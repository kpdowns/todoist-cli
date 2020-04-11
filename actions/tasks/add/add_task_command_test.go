package add

import (
	"bytes"
	"errors"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestIfNotAuthenticatedThenReceiveNotAuthenticatedErrorMessage(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, nil)
	addTaskCommand.SetArgs([]string{
		`-c="test content"`,
	})

	addTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := errorNotCurrentlyAuthenticated
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}

func TestWhenContentIsNotProvidedThenErrorIsThrown(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, nil)
	addTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := errorContentNotProvided
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}

func TestWhenContentIsProvidedAndErrorOccursWhileAddingTaskThenTaskWasNotAdded(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockOutputStream := &bytes.Buffer{}
	mockTaskService := &mocks.MockTaskService{
		AddTaskFunctionToExecute: func(content string, due string, priority int) error {
			return errors.New("error while adding task")
		},
	}

	addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, mockTaskService)
	addTaskCommand.SetArgs([]string{
		`-c="test content"`,
	})

	addTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := errorTaskNotAdded
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}

func TestWhenContentIsProvidedAndNoErrorsOccurThenTaskWasAdded(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: true,
	}
	mockOutputStream := &bytes.Buffer{}
	mockTaskService := &mocks.MockTaskService{
		AddTaskFunctionToExecute: func(content string, due string, priority int) error {
			return nil
		},
	}

	addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, mockTaskService)
	addTaskCommand.SetArgs([]string{
		`-c="test content"`,
	})

	addTaskCommand.Execute()

	textExpectedToBeWrittenToConsole := successfullyAddedTask
	textThatWasWrittenToConsole := mockOutputStream.String()
	if textExpectedToBeWrittenToConsole != textThatWasWrittenToConsole {
		t.Errorf("Expected '%s' to be written to output stream, received '%s'", textExpectedToBeWrittenToConsole, textThatWasWrittenToConsole)
	}
}
