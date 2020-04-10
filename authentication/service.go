package authentication

import (
	"bufio"
	"errors"
	"os"

	"github.com/kpdowns/todoist-cli/todoist"
)

const authenticationFilePath = "./authentication.data"
const errorNoCodeAvailableToSignInWith = "No code was provided while attempting to sign-in to the Todoist API"

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	IsAuthenticated() (bool, error)
	GetAccessToken() (string, error)
	SignIn(code string) error
	SignOut() error
}

type service struct {
	api todoist.API
}

// NewService creates a new instance of the Authentication service
func NewService(api todoist.API) Service {
	return &service{
		api: api,
	}
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not
func (service *service) IsAuthenticated() (bool, error) {
	accessToken, err := service.GetAccessToken()
	if err != nil {
		return false, err
	}

	isAccessTokenEmptyString := accessToken == ""
	if isAccessTokenEmptyString {
		return false, nil
	}

	return true, nil
}

// GetAccessToken retrieves the current access token
func (service *service) GetAccessToken() (string, error) {
	authenticationFile, err := service.getAuthenticationFile()
	if err != nil {
		return "", err
	}

	defer authenticationFile.Close()

	scanner := bufio.NewScanner(authenticationFile)
	scanner.Scan()

	accessToken := scanner.Text()
	return accessToken, nil
}

// SignIn signs into Todoist.com and stores the access token to be used for API requests
func (service *service) SignIn(code string) error {
	if code == "" {
		return errors.New(errorNoCodeAvailableToSignInWith)
	}

	err := service.deleteAccessToken()
	if err != nil {
		return err
	}

	authenticationFile, err := service.getAuthenticationFile()
	if err != nil {
		return err
	}
	defer authenticationFile.Close()

	token, err := service.api.GetAccessToken(code)
	if err != nil {
		return err
	}

	_, err = authenticationFile.WriteString(token.AccessToken)
	if err != nil {
		return err
	}

	return nil
}

// SignOut signs out of Todoist.com and deletes the stored access token
func (service *service) SignOut() error {
	authenticationFile, err := service.getAuthenticationFile()
	if err != nil {
		return err
	}

	defer authenticationFile.Close()

	err = authenticationFile.Truncate(0)
	if err != nil {
		return err
	}

	_, err = authenticationFile.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (service *service) deleteAccessToken() error {
	authenticationFile, err := service.getAuthenticationFile()
	if err != nil {
		return err
	}

	defer authenticationFile.Close()

	err = authenticationFile.Truncate(0)
	if err != nil {
		return err
	}

	_, err = authenticationFile.Seek(0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (service *service) getAuthenticationFile() (*os.File, error) {
	authenticationFile, err := os.OpenFile(authenticationFilePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	return authenticationFile, nil
}
