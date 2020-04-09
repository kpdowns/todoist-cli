package actions

import (
	"fmt"
	"os"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

var (
	rootCommand = &cobra.Command{
		Use:   "todoist-cli",
		Short: "A CLI tool that provides functionality that integrates with Todoist.com",
		Long:  "Todoist-CLI is a tool that allows you to interact with Todoist.com directly from the command line without using a browser.",
	}
)

// ExecuteRootCommand executes the root command of the Todoist-CLI
func ExecuteRootCommand() error {
	return rootCommand.Execute()
}

func init() {
	cobra.OnInitialize(initializeConfiguration)
}

func initializeConfiguration() {
	_, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
