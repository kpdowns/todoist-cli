package add

import (
	"errors"
	"fmt"
	"io"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/services"
	"github.com/spf13/cobra"
)

const (
	successfullyAddedTask = "Task has been added"

	errorNotCurrentlyAuthenticated = "Error, you are not currently logged in"
	errorContentNotProvided        = "Error, content must be provided when creating a task"
	errorTaskNotAdded              = "Error, the task could not be added, please try again later"
)

type dependencies struct {
	outputStream          io.Writer
	authenticationService authentication.Service
	taskService           services.TaskService
}

// NewAddTaskCommand creates an instance of the command that adds a task on Todoist
func NewAddTaskCommand(o io.Writer, a authentication.Service, t services.TaskService) *cobra.Command {
	var dependencies = &dependencies{
		outputStream:          o,
		authenticationService: a,
		taskService:           t,
	}

	content := ""
	due := ""
	priority := 1

	var addTaskCommand = &cobra.Command{
		Use:   "add",
		Short: "Add task",
		Long:  "Adds a task",
		Args:  cobra.OnlyValidArgs,
		Run: func(command *cobra.Command, args []string) {
			err := execute(dependencies, content, due, priority)
			if err != nil {
				fmt.Fprint(dependencies.outputStream, err.Error())
			}
		},
	}

	addTaskCommand.Flags().StringVarP(&content, "content", "c", "", "the content of the task")
	addTaskCommand.Flags().StringVarP(&due, "due", "d", "today", "the due date of the task (either in plain-text 'today', 'tomorrow', etc, or in long format)")
	addTaskCommand.Flags().IntVarP(&priority, "priority", "p", 1, "the priority of the task, options are 1 - 4 with 4 being the highest")

	return addTaskCommand
}

func execute(d *dependencies, content string, due string, priority int) error {
	if content == "" {
		return errors.New(errorContentNotProvided)
	}

	isAuthenticated, _ := d.authenticationService.IsAuthenticated()
	if !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	err := d.taskService.AddTask(content, due, priority)
	if err != nil {
		return errors.New(errorTaskNotAdded)
	}

	fmt.Fprint(d.outputStream, successfullyAddedTask)
	return nil
}
