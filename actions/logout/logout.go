package logout

import (
	"errors"
	"fmt"
	"io"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/spf13/cobra"
)

const (
	successfullyLoggedOut          = "Successfully logged out"
	errorNotCurrentlyAuthenticated = "You are not currently logged in, there are no access tokens to clear"
)

type dependencies struct {
	config       *config.TodoistCliConfiguration
	outputStream io.Writer
	api          todoist.API
}

// NewLogoutCommand creates a new instance of the authentication command
func NewLogoutCommand(config *config.TodoistCliConfiguration, outputStream io.Writer, api todoist.API) *cobra.Command {
	var dependencies = &dependencies{
		config:       config,
		outputStream: outputStream,
		api:          api,
	}

	var authenticateCommand = &cobra.Command{
		Use:   "logout",
		Short: "Logout of Todoist.com",
		Long:  "Logout of Todoist.com by clearing saved access tokens and revoking access",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			err := execute(dependencies)
			if err != nil {
				fmt.Fprintln(outputStream, err.Error())
			}
		},
	}

	return authenticateCommand
}

func execute(dependencies *dependencies) error {
	if !dependencies.config.IsAuthenticated() {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	err := dependencies.api.RevokeAccessToken()
	if err != nil {
		return err
	}

	dependencies.config.StoreAccessToken("")
	fmt.Fprintln(dependencies.outputStream, successfullyLoggedOut)

	return nil
}
