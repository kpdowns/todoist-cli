package complete

import (
	"errors"
	"fmt"
	"io"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/services"
	"github.com/spf13/cobra"
)

const (
	errorFailedToCompleteTask      = "An error occurred while completing the task"
	errorNotCurrentlyAuthenticated = "Error, you are not currently logged in"
	successTaskFlaggedAsCompleted  = "The task has successfully been completed"
)

type dependencies struct {
	outputStream          io.Writer
	authenticationService authentication.Service
	taskService           services.TaskService
}

// NewCompleteTaskCommand creates an instance of the command that completes tasks
func NewCompleteTaskCommand(o io.Writer, a authentication.Service, t services.TaskService) *cobra.Command {
	var dependencies = &dependencies{
		outputStream:          o,
		authenticationService: a,
		taskService:           t,
	}

	taskID := 0

	var completeTaskCommand = &cobra.Command{
		Use:   "complete",
		Short: "Complete task",
		Long:  "Flag a task as completed given a task id",
		Args:  cobra.OnlyValidArgs,
		Run: func(command *cobra.Command, args []string) {
			err := execute(dependencies, uint32(taskID))
			if err != nil {
				fmt.Fprint(o, err.Error())
			}
		},
	}

	completeTaskCommand.Flags().IntVarP(&taskID, "id", "i", 0, "the id of the task to flag as completed")

	return completeTaskCommand
}

func execute(d *dependencies, taskID uint32) error {
	isAuthenticated, _ := d.authenticationService.IsAuthenticated()
	if !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	err := d.taskService.CompleteTask(taskID)
	if err != nil {
		return errors.New(errorFailedToCompleteTask)
	}

	fmt.Fprintf(d.outputStream, successTaskFlaggedAsCompleted)
	return nil
}
