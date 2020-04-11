package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/actions/login"
	"github.com/kpdowns/todoist-cli/actions/logout"
	"github.com/kpdowns/todoist-cli/actions/tasks"
	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/storage"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/spf13/cobra"
)

// Initialize creates an instance of the root command and registers all other commands of the Todoist-CLI
func Initialize() error {
	var rootCommand = &cobra.Command{
		Use:   "todoist-cli",
		Short: "A CLI tool that provides functionality that integrates with Todoist.com",
		Long:  "Todoist-CLI is a tool that allows you to interact with Todoist.com directly from the command line without using a browser.",
	}

	config, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	currentExecutablePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Cannot determine the location todoist-cli is running from")
	}

	outputStream := os.Stdout
	api := todoist.NewAPI(*config)

	authenticationFilePath := fmt.Sprintf("%s/authentication.data", currentExecutablePath)
	authenticationServer := authentication.NewAuthenticationServer()
	authenticationRepository := authentication.NewAuthenticationRepository(storage.NewFile(authenticationFilePath))
	authenticationService := authentication.NewAuthenticationService(api, authenticationRepository, *config, authenticationServer)

	rootCommand.AddCommand(login.NewLoginCommand(outputStream, authenticationService, guid.NewString()))
	rootCommand.AddCommand(logout.NewLogoutCommand(outputStream, authenticationService))
	rootCommand.AddCommand(tasks.NewTasksCommand(api, outputStream, authenticationService))

	return rootCommand.Execute()
}
