package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/guid"
	"github.com/fatih/color"
	"github.com/kpdowns/todoist-cli/actions/login"
	"github.com/kpdowns/todoist-cli/actions/logout"
	"github.com/kpdowns/todoist-cli/actions/tasks"
	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/storage"
	"github.com/kpdowns/todoist-cli/tasks/repositories"
	"github.com/kpdowns/todoist-cli/tasks/services"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/spf13/cobra"
)

// Initialize creates an instance of the root command and registers all other commands of the todoist-cli
func Initialize() error {
	var rootCommand = &cobra.Command{
		Use:   "todoist",
		Short: "A CLI tool that provides functionality that integrates with Todoist.com",
		Long:  "todoist-cli is a tool that allows you to interact with Todoist.com directly from the command line without using a browser.",
	}

	config := config.LoadConfiguration()

	currentExecutablePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Cannot determine the location todoist-cli is running from")
	}

	outputStream := color.Output
	api := todoist.NewAPI(*config)

	authenticationFilePath := fmt.Sprintf("%s/authentication.data", currentExecutablePath)
	authenticationServer := authentication.NewAuthenticationServer()
	authenticationRepository := authentication.NewAuthenticationRepository(storage.NewFile(authenticationFilePath))
	authenticationService := authentication.NewAuthenticationService(api, authenticationRepository, *config, authenticationServer)

	tasksFilePath := fmt.Sprintf("%s/tasks.data", currentExecutablePath)
	tasksFile := storage.NewFile(tasksFilePath)
	taskRepository := repositories.NewTaskRepository(tasksFile)
	taskService := services.NewTaskService(api, authenticationService, taskRepository)

	rootCommand.AddCommand(login.NewLoginCommand(outputStream, authenticationService, guid.NewString()))
	rootCommand.AddCommand(logout.NewLogoutCommand(outputStream, authenticationService))
	rootCommand.AddCommand(tasks.NewTasksCommand(outputStream, authenticationService, taskService))

	return rootCommand.Execute()
}
