package tasks

import (
	"io"

	"github.com/kpdowns/todoist-cli/actions/tasks/add"
	"github.com/kpdowns/todoist-cli/actions/tasks/complete"
	"github.com/kpdowns/todoist-cli/actions/tasks/list"
	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/services"
	"github.com/spf13/cobra"
)

// NewTasksCommand creates a new instance of the authentication command
func NewTasksCommand(o io.Writer, authenticationService authentication.Service, taskService services.TaskService) *cobra.Command {
	var tasksCommand = &cobra.Command{
		Use:   "tasks",
		Short: "Manage tasks",
		Long:  "Manage tasks on Todoist.com",
	}

	tasksCommand.AddCommand(list.NewListTasksCommand(o, authenticationService, taskService))
	tasksCommand.AddCommand(add.NewAddTaskCommand(o, authenticationService, taskService))
	tasksCommand.AddCommand(complete.NewCompleteTaskCommand(o, authenticationService, taskService))

	return tasksCommand
}
