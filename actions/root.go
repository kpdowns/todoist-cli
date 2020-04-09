package actions

import (
	"os"

	"github.com/kpdowns/todoist-cli/actions/authenticate"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

// Initialize creates an instance of the root command and registers all other commands of the Todoist-CLI
func Initialize(config *config.TodoistCliConfiguration) error {
	var rootCommand = &cobra.Command{
		Use:   "todoist-cli",
		Short: "A CLI tool that provides functionality that integrates with Todoist.com",
		Long:  "Todoist-CLI is a tool that allows you to interact with Todoist.com directly from the command line without using a browser.",
	}

	var outputStream = os.Stdout

	rootCommand.AddCommand(authenticate.NewAuthenticateCommand(config, outputStream))

	return rootCommand.Execute()
}
