package authentication

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/kpdowns/todoist-cli/authentication/types"
)

// Server handles the callback from Todoist.com during the authentication flow
type Server interface {
	StartTemporaryServerToListenForResponse(csrfGUID string) (*types.AuthenticationResponse, error)
}

type server struct {
}

// NewAuthenticationServer creates a new instance of the server
func NewAuthenticationServer() Server {
	return &server{}
}

// Starts a server listening on the port defined in configuration to handle the callback from Todoist.com
func (s *server) StartTemporaryServerToListenForResponse(guid string) (*types.AuthenticationResponse, error) {
	response, err := s.startServer(guid)
	if err != nil {
		return nil, err
	}

	if response.Code == "" {
		return nil, errors.New(errorNoAuthCodeReceived)
	}

	return response, err
}

func (s *server) startServer(guid string) (*types.AuthenticationResponse, error) {
	var waitGroup = &sync.WaitGroup{}
	waitGroup.Add(1)
	var server = &http.Server{Addr: ":8123"}

	var response *types.AuthenticationResponse
	var err error
	http.HandleFunc("/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		response, err = s.handleOauthResponse(w, r, guid)
		waitGroup.Done()
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("OauthFlow error: %v", err)
		}
	}()

	waitGroup.Wait()

	return response, err
}

func (s *server) handleOauthResponse(w http.ResponseWriter, r *http.Request, csrfGUID string) (*types.AuthenticationResponse, error) {
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
