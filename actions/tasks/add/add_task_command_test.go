package add

import (
	"bytes"
	"errors"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNotAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, nil)
	addTaskCommand.SetArgs([]string{
		`-c="test content"`,
	})

	addTaskCommand.Execute()

	assert.Equal(t, errorNotCurrentlyAuthenticated, mockOutputStream.String())
}

func TestInputParameters(t *testing.T) {

	t.Run("If the priority is not valid, then an error stating so is written to the console", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: false,
		}
		mockOutputStream := &bytes.Buffer{}

		addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, nil)
		addTaskCommand.SetArgs([]string{
			`-content="test"`,
			`-p=5`,
		})

		addTaskCommand.Execute()

		assert.Equal(t, errorInvalidPriority, mockOutputStream.String())

	})

	t.Run("If no content is provided, then an error stating so is written to the console", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: false,
		}
		mockOutputStream := &bytes.Buffer{}

		addTaskCommand := NewAddTaskCommand(mockOutputStream, mockAuthenticationService, nil)
		addTaskCommand.Execute()

		assert.Equal(t, errorContentNotProvided, mockOutputStream.String())

	})
}

func TestAddingATask(t *testing.T) {

	t.Run("When creating a task and an error occurs, an error stating that the task wasn't added is written to console", func(t *testing.T) {
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

		assert.Equal(t, errorTaskNotAdded, mockOutputStream.String())

	})

	t.Run("When creating a task and no error occurs, then a message stating that the task was created is written to console", func(t *testing.T) {
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

		assert.Equal(t, successfullyAddedTask, mockOutputStream.String())

	})

}
