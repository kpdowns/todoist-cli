package tasks

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestWhenCreatingTaskCommandThenSubcommandToListTasksIsCreated(t *testing.T) {
	mockAPI := &mocks.MockAPI{}
	mockOutputStream := &bytes.Buffer{}
	mockAuthenticationService := &mocks.MockAuthenticationService{}

	taskCommand := NewTasksCommand(mockAPI, mockOutputStream, mockAuthenticationService)

	registeredCommands := taskCommand.Commands()
	if len(registeredCommands) == 0 {
		t.Errorf("Expected subcommands to be registered")
	}

	found := false
	for _, registeredCommand := range registeredCommands {
		if registeredCommand.Use == "list" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected that a subtask of 'list' was registered")
	}
}
