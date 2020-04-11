package list

import (
	"errors"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks"
	"github.com/spf13/cobra"
)

const (
	noTasksMessage                 = "No tasks to complete across any of your projects"
	errorNotCurrentlyAuthenticated = "Error, you are not currently logged in"
)

type dependencies struct {
	outputStream          io.Writer
	authenticationService authentication.Service
	taskService           tasks.Service
}

// NewListTasksCommand creates an instance of the command that prints all tasks to the console
func NewListTasksCommand(o io.Writer, a authentication.Service, t tasks.Service) *cobra.Command {
	var dependencies = &dependencies{
		outputStream:          o,
		authenticationService: a,
		taskService:           t,
	}

	var listTasksCommand = &cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Long:  "List tasks for all or only a specific project",
		Run: func(command *cobra.Command, args []string) {
			err := execute(dependencies)
			if err != nil {
				fmt.Fprint(o, err.Error())
			}
		},
	}

	return listTasksCommand
}

func execute(d *dependencies) error {
	isAuthenticated, _ := d.authenticationService.IsAuthenticated()
	if !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	tasks, err := d.taskService.GetAllTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Fprint(d.outputStream, noTasksMessage)
	}

	writer := tabwriter.NewWriter(d.outputStream, 0, 8, 1, '\t', 0)
	for _, task := range tasks {
		fmt.Fprintln(writer, task.AsString())
	}
	writer.Flush()

	return nil
}
