package authentication

import (
	"bufio"
	"os"
)

const authenticationFilePath = "./authentication.data"

// Repository handles persistence of the access token to be used with the Todoist API
type Repository interface {
	GetAccessToken() (string, error)
	DeleteAccessToken() error
	UpdateAccessToken(token string) error
}

type repository struct{}

// NewAuthenticationRepository creates a new instance of the repository that writes to file
func NewAuthenticationRepository() Repository {
	return &repository{}
}

func (r *repository) GetAccessToken() (string, error) {
	authenticationFile, err := r.getAuthenticationFile()
	if err != nil {
		return "", err
	}

	defer authenticationFile.Close()

	scanner := bufio.NewScanner(authenticationFile)
	scanner.Scan()

	accessToken := scanner.Text()
	return accessToken, nil
}

func (r *repository) DeleteAccessToken() error {
	authenticationFile, err := r.getAuthenticationFile()
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

func (r *repository) UpdateAccessToken(token string) error {
	r.DeleteAccessToken()

	authenticationFile, err := r.getAuthenticationFile()
	if err != nil {
		return err
	}
	defer authenticationFile.Close()

	_, err = authenticationFile.WriteString(token)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) getAuthenticationFile() (*os.File, error) {
	authenticationFile, err := os.OpenFile(authenticationFilePath, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return nil, err
	}
	return authenticationFile, nil
}
