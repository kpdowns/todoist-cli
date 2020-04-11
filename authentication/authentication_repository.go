package authentication

import (
	"errors"
	"strings"

	"github.com/kpdowns/todoist-cli/authentication/types"
	"github.com/kpdowns/todoist-cli/storage"
)

const (
	errorMalformedAuthenticationFile = "Error, the contents of the authentication file are malformed"
)

// Repository handles persistence of the access token to be used with the Todoist API
type Repository interface {
	GetAccessToken() (*types.AccessToken, error)
	DeleteAccessToken() error
	UpdateAccessToken(token string) error
}

type repository struct {
	file storage.File
}

// NewAuthenticationRepository creates a new instance of the repository that writes to storage
func NewAuthenticationRepository(file storage.File) Repository {
	return &repository{
		file: file,
	}
}

// GetAccessToken retrieves the access token from the contents of the storage
func (r *repository) GetAccessToken() (*types.AccessToken, error) {
	//todo - test for error if more than 1 line
	contents, err := r.file.ReadContents()
	if err != nil {
		return nil, nil
	}

	if len(strings.Split(contents, "\n")) >= 2 {
		return nil, errors.New(errorMalformedAuthenticationFile)
	}

	return &types.AccessToken{AccessToken: contents}, nil
}

// DeleteAccessToken removes the access token from storage
func (r *repository) DeleteAccessToken() error {
	err := r.file.OverwriteContents("")
	return err
}

// UpdateAccessToken overwrites the existing access token saved in storage
func (r *repository) UpdateAccessToken(token string) error {
	err := r.file.OverwriteContents(token)
	return err
}
