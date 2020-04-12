package tasks

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCommandCreation(t *testing.T) {

	t.Run("Sub command to list tasks is added", func(t *testing.T) {

		mockOutputStream := &bytes.Buffer{}
		mockAuthenticationService := &mocks.MockAuthenticationService{}
		mockTaskService := &mocks.MockTaskService{}

		taskCommand := NewTasksCommand(mockOutputStream, mockAuthenticationService, mockTaskService)

		registeredCommands := taskCommand.Commands()

		found := false
		for _, registeredCommand := range registeredCommands {
			if registeredCommand.Use == "list" {
				found = true
				break
			}
		}

		assert.True(t, found)

	})

	t.Run("Sub command to create task is added", func(t *testing.T) {

		mockOutputStream := &bytes.Buffer{}
		mockAuthenticationService := &mocks.MockAuthenticationService{}
		mockTaskService := &mocks.MockTaskService{}

		taskCommand := NewTasksCommand(mockOutputStream, mockAuthenticationService, mockTaskService)

		registeredCommands := taskCommand.Commands()

		found := false
		for _, registeredCommand := range registeredCommands {
			if registeredCommand.Use == "add" {
				found = true
				break
			}
		}

		assert.True(t, found)

	})

}
