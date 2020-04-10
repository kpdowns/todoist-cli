package authentication

import (
	"bufio"
	"os"
)

const authenticationFilePath = "./authentication.data"

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	IsAuthenticated() (bool, error)
	GetAccessToken() (string, error)
	SaveAccessToken(accessToken string) error
	DeleteAccessToken() error
}

type service struct{}

// NewService creates a new instance of the Authentication service
func NewService() Service {
	return &service{}
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

// UpdateAccessToken stores the access token to be used for API requests
func (service *service) SaveAccessToken(accessToken string) error {
	err := service.DeleteAccessToken()
	if err != nil {
		return err
	}

	authenticationFile, err := service.getAuthenticationFile()
	if err != nil {
		return err
	}
	defer authenticationFile.Close()

	_, err = authenticationFile.WriteString(accessToken)
	if err != nil {
		return err
	}

	return nil
}

// DeleteAccessToken deletes the currently stored access token
func (service *service) DeleteAccessToken() error {
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
