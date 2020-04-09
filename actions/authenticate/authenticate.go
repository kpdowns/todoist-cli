package authenticate

import (
	"fmt"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

type dependencies struct {
	config *config.TodoistCliConfiguration
}

// NewAuthenticateCommand creates a new instance of the authentication command
func NewAuthenticateCommand(config *config.TodoistCliConfiguration) *cobra.Command {
	var dependencies = &dependencies{
		config: config,
	}

	var authenticateCommand = &cobra.Command{
		Use:   "authenticate",
		Short: "Start the authentication process against Todoist.com",
		Long:  "Starts the Oauth login flow on Todoist.com which will allow Todoist-cli to access your tasks and projects on Todoist.com",
		Run: func(command *cobra.Command, args []string) {
			execute(command, args, dependencies)
		},
	}

	return authenticateCommand
}

func execute(command *cobra.Command, args []string, dependencies *dependencies) {

}

func generateOauthURL(config *config.TodoistCliConfiguration, guid string) string {
	return fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		config.Client.TodoistURL,
		config.Client.ClientID,
		config.Client.RequiredPermissions,
		guid)
}
