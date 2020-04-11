package tasks

import (
	"io"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/spf13/cobra"
)

const (
	errorNotCurrentlyAuthenticated = "Error, you are not currently logged in"
)

// NewTasksCommand creates a new instance of the authentication command
func NewTasksCommand(api todoist.API, o io.Writer, auth authentication.Service) *cobra.Command {
	var tasksCommand = &cobra.Command{
		Use:   "tasks",
		Short: "Manage tasks",
		Long:  "Manage tasks on Todoist.com",
	}

	taskService := NewTaskService(api, auth)
	tasksCommand.AddCommand(NewListTasksCommand(o, auth, taskService))

	return tasksCommand
}
