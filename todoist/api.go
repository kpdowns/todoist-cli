package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

// API provides functions for interacting with the Todoist API
type API interface {
	GetAccessToken(code string) (*responses.AccessToken, error)
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
