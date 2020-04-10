package authentication

import (
	"errors"

	"github.com/kpdowns/todoist-cli/todoist"
)

const errorNoCodeAvailableToSignInWith = "No code was provided while attempting to sign-in to the Todoist API"

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	IsAuthenticated() (bool, error)
	GetAccessToken() (string, error)
	SignIn(code string) error
	SignOut() error
}

type service struct {
	api        todoist.API
	repository Repository
}

// NewService creates a new instance of the Authentication service
func NewService(api todoist.API, repository Repository) Service {
	return &service{
		api:        api,
		repository: repository,
	}
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not
func (s *service) IsAuthenticated() (bool, error) {
	accessToken, err := s.repository.GetAccessToken()
	if err != nil {
		return false, err
	}

	isAccessTokenEmptyString := accessToken != ""
	if isAccessTokenEmptyString {
		return false, nil
	}

	return true, nil
}

// GetAccessToken retrieves the current access token
func (s *service) GetAccessToken() (string, error) {
	return s.repository.GetAccessToken()
}

// SignIn signs into Todoist.com and stores the access token to be used for API requests
func (s *service) SignIn(code string) error {
	if code == "" {
		return errors.New(errorNoCodeAvailableToSignInWith)
	}

	token, err := s.api.GetAccessToken(code)
	if err != nil {
		return err
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

	err = s.api.RevokeAccessToken(accessToken)
	if err != nil {
		return err
	}

	err = s.repository.DeleteAccessToken()
	if err != nil {
		return err
	}

	return nil
}
