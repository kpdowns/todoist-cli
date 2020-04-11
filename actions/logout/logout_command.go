package logout

import (
	"errors"
	"fmt"
	"io"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/spf13/cobra"
)

const (
	successfullyLoggedOut          = "Successfully logged out"
	errorNotCurrentlyAuthenticated = "You are not currently logged in, there are no access tokens to clear"
)

type dependencies struct {
	outputStream          io.Writer
	authenticationService authentication.Service
}

// NewLogoutCommand creates a new instance of the authentication command
func NewLogoutCommand(outputStream io.Writer, authenticationService authentication.Service) *cobra.Command {
	var dependencies = &dependencies{
		outputStream:          outputStream,
		authenticationService: authenticationService,
	}

	var authenticateCommand = &cobra.Command{
		Use:   "logout",
		Short: "Logout of Todoist.com",
		Long:  "Logout of Todoist.com by clearing saved access tokens and revoking access",
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

func execute(dependencies *dependencies) error {
	isAuthenticated, _ := dependencies.authenticationService.IsAuthenticated()
	if !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	err := dependencies.authenticationService.SignOut()
	if err != nil {
		return err
	}

	fmt.Fprint(dependencies.outputStream, successfullyLoggedOut)

	return nil
}
