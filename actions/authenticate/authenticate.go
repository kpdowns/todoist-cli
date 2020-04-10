package authenticate

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/kpdowns/todoist-cli/authentication"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/actions/authenticate/types"
	"github.com/kpdowns/todoist-cli/config"
	"github.com/spf13/cobra"
)

const (
	oauthInitiationText       = "To authenticate todoist-cli, please navigate to %s"
	successfullyAuthenticated = "Successfully authenticated"

	errorAlreadyAuthenticatedText        = "The todoist-cli is already authenticated"
	errorInvalidTokenReturnedFromTodoist = "Authentication failed, no access token was returned"
	errorAuthenticationRejected          = "The authentication request was rejected"
	errorPotentialCsrfAttack             = "Potential CSRF, the state provided to Todoist did not match what was returned"
	errorNoAuthCodeReceived              = "No authorization code was received"
)

type dependencies struct {
	config                *config.TodoistCliConfiguration
	outputStream          io.Writer
	guid                  string
	authenticationService authentication.Service
}

// NewAuthenticateCommand creates a new instance of the authentication command
func NewAuthenticateCommand(config *config.TodoistCliConfiguration, outputStream io.Writer, authenticationService authentication.Service) *cobra.Command {
	var dependencies = &dependencies{
		config:                config,
		outputStream:          outputStream,
		guid:                  guid.NewString(),
		authenticationService: authenticationService,
	}

	var authenticateCommand = &cobra.Command{
		Use:   "authenticate",
		Short: "Start the authentication process against Todoist.com",
		Long:  "Starts the Oauth login flow on Todoist.com which will allow Todoist-cli to access your tasks and projects on Todoist.com",
		Args:  cobra.NoArgs,
		Run: func(command *cobra.Command, args []string) {
			authenticationFunction := startTemporaryServerToListenForResponse
			err := execute(dependencies, authenticationFunction)
			if err != nil {
				fmt.Fprintln(outputStream, err.Error())
			}
		},
	}

	return authenticateCommand
}

func execute(d *dependencies, authenticationFunction func(csrfGUID string) (*types.AuthenticationResponse, error)) error {
	isAuthenticated, err := d.authenticationService.IsAuthenticated()
	if isAuthenticated {
		return errors.New(errorAlreadyAuthenticatedText)
	}

	if err != nil {
		return err
	}

	oauthInitiationURL := generateOauthURL(d.config, d.guid)
	promptText := fmt.Sprintf(oauthInitiationText, oauthInitiationURL)
	fmt.Fprintln(d.outputStream, promptText)

	response, err := authenticationFunction(d.guid)
	if err != nil {
		return err
	}

	if response.Code == "" {
		return errors.New(errorNoAuthCodeReceived)
	}

	err = d.authenticationService.SignIn(response.Code)
	if err != nil {
		return err
	}
	fmt.Fprintln(d.outputStream, successfullyAuthenticated)

	return nil
}

func startTemporaryServerToListenForResponse(csrfGUID string) (*types.AuthenticationResponse, error) {
	var waitGroup = &sync.WaitGroup{}
	waitGroup.Add(1)
	var server = &http.Server{Addr: ":8123"}

	var authenticationResponse types.AuthenticationResponse
	var authenticationError error
	http.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		response, err := handleOauthResponse(w, r, csrfGUID)
		if err == nil {
			authenticationResponse = *response
		}

		authenticationError = err
		waitGroup.Done()
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("OauthFlow error: %v", err)
		}
	}()

	waitGroup.Wait()
	return &authenticationResponse, authenticationError
}

func handleOauthResponse(w http.ResponseWriter, r *http.Request, csrfGUID string) (*types.AuthenticationResponse, error) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Connection", "close")

	queryParameters := r.URL.Query()
	state := queryParameters.Get("state")
	if state == "" {

		return nil, errors.New(errorAuthenticationRejected)
	}

	if csrfGUID != state {
		return nil, errors.New(errorPotentialCsrfAttack)
	}

	code := queryParameters.Get("code")
	if code == "" {
		return nil, errors.New(errorNoAuthCodeReceived)
	}

	const successfulResponseHTML = `
		<html>
			<body>
				<p>
					<b>Authentication successful</b>, you can safely close this page.
				</p>
			</body>
		</html>
	`

	fmt.Fprint(w, successfulResponseHTML)

	return &types.AuthenticationResponse{
		Code: code,
	}, nil
}

func generateOauthURL(config *config.TodoistCliConfiguration, guid string) string {
	return fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		config.Client.TodoistURL,
		config.Client.ClientID,
		config.Client.RequiredPermissions,
		guid)
}
