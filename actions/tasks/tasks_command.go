package tasks

import (
	"io"

	"github.com/kpdowns/todoist-cli/actions/tasks/add"
	"github.com/kpdowns/todoist-cli/actions/tasks/list"
	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/spf13/cobra"
)

// NewTasksCommand creates a new instance of the authentication command
func NewTasksCommand(api todoist.API, o io.Writer, auth authentication.Service) *cobra.Command {
	var tasksCommand = &cobra.Command{
		Use:   "tasks",
		Short: "Manage tasks",
		Long:  "Manage tasks on Todoist.com",
	}

	taskService := tasks.NewTaskService(api, auth)
	tasksCommand.AddCommand(list.NewListTasksCommand(o, auth, taskService))
	tasksCommand.AddCommand(add.NewAddTaskCommand(o, auth, taskService))

	return tasksCommand
}
