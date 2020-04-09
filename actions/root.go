package actions

import (
	"fmt"
	"os"

	"github.com/kpdowns/todoist-cli/actions/authenticate"
	"github.com/kpdowns/todoist-cli/config"
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

	outputStream := os.Stdout
	api := todoist.NewAPI(*config)

	rootCommand.AddCommand(authenticate.NewAuthenticateCommand(config, outputStream, api))

	return rootCommand.Execute()
}
