package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

// API provides functions for interacting with the Todoist API
type API interface {
	GetAccessToken(code string) (*responses.AccessToken, error)
	RevokeAccessToken() error
}

type api struct {
	config config.TodoistCliConfiguration
}

// NewAPI creates a new instance of the API to interact with Todoist
func NewAPI(config config.TodoistCliConfiguration) API {
	return &api{
		config: config,
	}
}

// GetAccessToken returns the bearer token provided by the Todoist API while authenticating
func (a *api) GetAccessToken(code string) (*responses.AccessToken, error) {
	accessTokenURL := fmt.Sprintf("%s/oauth/access_token?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		a.config.Client.TodoistURL,
		a.config.Client.ClientID,
		a.config.Client.ClientSecret,
		code,
		a.config.Client.OauthRedirectURL,
	)

	var buffer []byte
	response, err := http.Post(accessTokenURL, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		errorMessage := fmt.Sprintf("Error occurred while retrieving the access token, code was %d\n", response.StatusCode)
		return nil, errors.New(errorMessage)
	}

	defer response.Body.Close()

	var accessToken responses.AccessToken
	err = json.NewDecoder(response.Body).Decode(&accessToken)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

// RevokeAccessToken revokes the current access token effectively logging the user out
func (a *api) RevokeAccessToken() error {
	revokeAccessTokenURL := fmt.Sprintf("%s/sync/v8/access_tokens/revoke", a.config.Client.TodoistURL)

	requestBody := &requests.RevokeAccessToken{
		ClientID:     a.config.Client.ClientID,
		ClientSecret: a.config.Client.ClientSecret,
		AccessToken:  a.config.Authentication.AccessToken,
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	response, err := http.Post(revokeAccessTokenURL, "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != 204 {
		errorMessage := fmt.Sprintf("Failed to revoke access token, response status was %d\n", response.StatusCode)
		return errors.New(errorMessage)
	}

	return nil
}
