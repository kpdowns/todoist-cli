package tasks

import (
	"errors"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/spf13/cobra"
)

type dependencies struct {
	outputStream          io.Writer
	authenticationService authentication.Service
	taskService           Service
}

// NewListTasksCommand creates an instance of the command that prints all tasks to the console
func NewListTasksCommand(o io.Writer, a authentication.Service, t Service) *cobra.Command {
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
				fmt.Fprintln(o, err.Error())
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

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', 0)
	for _, task := range tasks {
		fmt.Fprintln(writer, task.AsString())
	}
	writer.Flush()

	return nil
}
