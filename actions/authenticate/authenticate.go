package authenticate

import (
	"errors"
	"fmt"
	"io"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

const (
	oauthInitiationText = "To authenticate todoist-cli, please navigate to %s"

	errorAlreadyAuthenticatedText = "The todoist-cli is already authenticated"
)

type dependencies struct {
	config       *config.TodoistCliConfiguration
	outputStream io.Writer
	guid         string
}

// NewAuthenticateCommand creates a new instance of the authentication command
func NewAuthenticateCommand(config *config.TodoistCliConfiguration, outputStream io.Writer) *cobra.Command {
	var dependencies = &dependencies{
		config:       config,
		outputStream: outputStream,
		guid:         guid.NewString(),
	}

	var authenticateCommand = &cobra.Command{
		Use:   "authenticate",
		Short: "Start the authentication process against Todoist.com",
		Long:  "Starts the Oauth login flow on Todoist.com which will allow Todoist-cli to access your tasks and projects on Todoist.com",
		Run: func(command *cobra.Command, args []string) {
			err := execute(command, args, dependencies)
			if err != nil {
				fmt.Fprintln(outputStream, err.Error())
			}
		},
	}

	return authenticateCommand
}

func execute(command *cobra.Command, args []string, dependencies *dependencies) error {
	if dependencies.config.Authentication.IsAuthenticated() {
		return errors.New(errorAlreadyAuthenticatedText)
	}

	oauthInitiationURL := generateOauthURL(dependencies.config, dependencies.guid)
	promptText := fmt.Sprintf(oauthInitiationText, oauthInitiationURL)
	fmt.Fprintln(dependencies.outputStream, promptText)

	return nil
}

func generateOauthURL(config *config.TodoistCliConfiguration, guid string) string {
	return fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		config.Client.TodoistURL,
		config.Client.ClientID,
		config.Client.RequiredPermissions,
		guid)
}
