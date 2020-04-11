package authenticate

import (
	"errors"
	"fmt"
	"io"

	"github.com/kpdowns/todoist-cli/authentication"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

const (
	oauthInitiationText       = "To authenticate todoist-cli, please navigate to %s"
	successfullyAuthenticated = "Successfully authenticated"

	errorAlreadyAuthenticatedText = "The todoist-cli is already authenticated"
	errorDuringAuthentication     = "An error occurred while authenticating against Todoist.com, please try again"
)

type dependencies struct {
	config                *config.TodoistCliConfiguration
	outputStream          io.Writer
	guid                  string
	authenticationService authentication.Service
}

// NewAuthenticateCommand creates a new instance of the authentication command
func NewAuthenticateCommand(outputStream io.Writer, authenticationService authentication.Service, guid string) *cobra.Command {
	var dependencies = &dependencies{
		outputStream:          outputStream,
		guid:                  guid,
		authenticationService: authenticationService,
	}

	var authenticateCommand = &cobra.Command{
		Use:   "authenticate",
		Short: "Start the authentication process against Todoist.com",
		Long:  "Starts the Oauth login flow on Todoist.com which will allow Todoist-cli to access your tasks and projects on Todoist.com",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			err := execute(dependencies)
			if err != nil {
				fmt.Fprint(outputStream, err.Error())
			}
		},
	}

	return authenticateCommand
}

func execute(d *dependencies) error {
	isAuthenticated, err := d.authenticationService.IsAuthenticated()
	if isAuthenticated {
		return errors.New(errorAlreadyAuthenticatedText)
	}

	if err != nil {
		return err
	}

	oauthInitiationURL := d.authenticationService.GetOauthURL(d.guid)
	promptText := fmt.Sprintf(oauthInitiationText, oauthInitiationURL)
	fmt.Fprintln(d.outputStream, promptText)

	err = d.authenticationService.SignIn(d.guid)
	if err != nil {
		return errors.New(errorDuringAuthentication)
	}

	fmt.Fprint(d.outputStream, successfullyAuthenticated)

	return nil
}
