package authentication

import (
	"errors"
	"fmt"

	"github.com/kpdowns/todoist-cli/authentication/types"
	"github.com/kpdowns/todoist-cli/config"

	"github.com/kpdowns/todoist-cli/todoist"
)

const (
	errorNoCodeAvailableToSignInWith     = "No code was provided while attempting to sign-in to the Todoist API"
	errorInvalidTokenReturnedFromTodoist = "Authentication failed, no access token was returned"
	errorAuthenticationRejected          = "The authentication request was rejected"
	errorPotentialCsrfAttack             = "Potential CSRF, the state provided to Todoist did not match what was returned"
	errorNoAuthCodeReceived              = "No authorization code was received"
	errorNoAccessTokenReceived           = "Error during authentication, no access token could be retrieved"
)

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	IsAuthenticated() (bool, error)
	GetAccessToken() (*types.AccessToken, error)
	SignIn(code string) error
	SignOut() error
	GetOauthURL(guid string) string
}

type service struct {
	api        todoist.API
	repository Repository
	config     config.TodoistCliConfiguration
	server     Server
}

// NewAuthenticationService creates a new instance of the Authentication service
func NewAuthenticationService(api todoist.API, repository Repository, config config.TodoistCliConfiguration, server Server) Service {
	return &service{
		api:        api,
		repository: repository,
		config:     config,
		server:     server,
	}
}

// IsAuthenticated checks whether the todoist-cli is authenticated or not
func (s *service) IsAuthenticated() (bool, error) {
	accessToken, err := s.repository.GetAccessToken()
	if err != nil {
		return false, err
	}

	isAccessTokenEmptyString := accessToken.AccessToken == ""
	if isAccessTokenEmptyString {
		return false, nil
	}

	return true, nil
}

// GetAccessToken retrieves the current access token
func (s *service) GetAccessToken() (*types.AccessToken, error) {
	return s.repository.GetAccessToken()
}

// SignIn signs into Todoist using the provided guid as a CSRF token
func (s *service) SignIn(guid string) error {
	if guid == "" {
		return errors.New(errorNoCodeAvailableToSignInWith)
	}

	response, err := s.server.StartTemporaryServerToListenForResponse(guid)
	if err != nil {
		return errors.New(errorNoAuthCodeReceived)
	}

	token, err := s.api.GetAccessToken(response.Code)
	if err != nil {
		return err
	}

	if token.AccessToken == "" {
		return errors.New(errorNoAccessTokenReceived)
	}

	err = s.repository.UpdateAccessToken(token.AccessToken)
	if err != nil {
		return err
	}

	return nil
}

// SignOut signs out of Todoist.com and deletes the stored access token
func (s *service) SignOut() error {

	accessToken, err := s.repository.GetAccessToken()
	if err != nil {
		return err
	}

	err = s.api.RevokeAccessToken(accessToken.AccessToken)
	if err != nil {
		return err
	}

	err = s.repository.DeleteAccessToken()
	if err != nil {
		return err
	}

	return nil
}

// GetOauthURL generates a URL with provided guid as a CSRF protection token
func (s *service) GetOauthURL(guid string) string {
	return fmt.Sprintf("%s/oauth/authorize?client_id=%s&scope=%s&state=%s",
		s.config.Client.TodoistURL,
		s.config.Client.ClientID,
		s.config.Client.RequiredPermissions,
		guid)
}
