package complete

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestNotAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	completeTasksCommand := NewCompleteTaskCommand(mockOutputStream, mockAuthenticationService, nil)
	completeTasksCommand.Execute()

	assert.Equal(t, errorNotCurrentlyAuthenticated, mockOutputStream.String())
}

func TestWrittingToOutputStream(t *testing.T) {

	t.Run("When authenticated and an error occurs while completing the task, then message is written to output stream", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}

		mockOutputStream := &bytes.Buffer{}

		mockTaskService := &mocks.MockTaskService{
			CompleteTaskFunc: func(uint32) error {
				return errors.New("Test error")
			},
		}

		completeTasksCommand := NewCompleteTaskCommand(mockOutputStream, mockAuthenticationService, mockTaskService)
		completeTasksCommand.Execute()

		assert.Equal(t, errorFailedToCompleteTask, mockOutputStream.String())

	})

	t.Run("When authenticated and no error occurs while completing the task, then message is written to output stream", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}

		mockOutputStream := &bytes.Buffer{}

		mockTaskService := &mocks.MockTaskService{
			CompleteTaskFunc: func(uint32) error {
				return nil
			},
		}

		completeTasksCommand := NewCompleteTaskCommand(mockOutputStream, mockAuthenticationService, mockTaskService)
		completeTasksCommand.Execute()

		assert.Equal(t, successTaskFlaggedAsCompleted, mockOutputStream.String())

	})

}
