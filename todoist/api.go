package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/rest"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
)

const (
	errorRetrievingAccessToken            = "An error occurred while retrieving your access token, please try again later"
	errorRevokingAccessToken              = "An error occurred while attempting to revoke your access token, please try again later"
	errorCommunicatingWithTodoistAPI      = "An error occurred while attempting to communicate with Todoist, please try again later"
	errorExecutingQuery                   = "An error occurred while executing your query, please try again later"
	errorExecutingQueryMalformedQuery     = "An error occurred while executing your query, the query was not valid"
	errorExecutingCommand                 = "An error occurred while executing your command, please try again later"
	errorExecutingCommandMalformedCommand = "An error occurred while executing your command, the command was not valid"
	errorMalformedResponse                = "An error occurred while trying to decode the response from Todoist, please try again later"
)

// API provides functions for interacting with the Todoist API
type API interface {
	GetAccessToken(code string) (*responses.AccessToken, error)
	RevokeAccessToken(accessToken string) error
	ExecuteSyncQuery(query requests.Query) (*responses.Query, error)
	ExecuteSyncCommand(command requests.Command) error
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
		a.config.TodoistURL,
		a.config.ClientID,
		a.config.ClientSecret,
		code,
		a.config.OauthRedirectURL,
	)

	var buffer []byte
	response, err := rest.Post(accessTokenURL, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return nil, errors.New(errorCommunicatingWithTodoistAPI)
	}

	if response.StatusCode != 200 {
		return nil, errors.New(errorRetrievingAccessToken)
	}

	defer response.Body.Close()

	var accessToken responses.AccessToken
	err = json.NewDecoder(response.Body).Decode(&accessToken)
	if err != nil {
		return nil, errors.New(errorMalformedResponse)
	}

	return &accessToken, nil
}

// RevokeAccessToken revokes the current access token effectively logging the user out
func (a *api) RevokeAccessToken(accessToken string) error {
	if accessToken == "" {
		return errors.New(errorRevokingAccessToken)
	}

	revokeAccessTokenURL := fmt.Sprintf("%s/sync/v8/access_tokens/revoke", a.config.TodoistURL)

	requestBody := &requests.RevokeAccessToken{
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		AccessToken:  accessToken,
	}

	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	response, err := rest.Post(revokeAccessTokenURL, "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return errors.New(errorCommunicatingWithTodoistAPI)
	}

	defer response.Body.Close()
	if response.StatusCode != 204 {
		return errors.New(errorRevokingAccessToken)
	}

	return nil
}

// ExecuteSyncQuery executes a query against Todoist and returns the response
func (a *api) ExecuteSyncQuery(query requests.Query) (*responses.Query, error) {
	url := fmt.Sprintf("%s/sync/v8/sync?%s", a.config.TodoistURL, query.ToQueryString())

	var buffer []byte
	response, err := rest.Post(url, "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return nil, errors.New(errorCommunicatingWithTodoistAPI)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New(errorExecutingQuery)
	}

	var queryResponse responses.Query
	err = json.NewDecoder(response.Body).Decode(&queryResponse)
	if err != nil {
		return nil, errors.New(errorMalformedResponse)
	}

	return &queryResponse, nil
}

// ExecuteSyncCommand executes a command against Todoist and returns the response
func (a *api) ExecuteSyncCommand(command requests.Command) error {
	url := fmt.Sprintf("%s/sync/v8/sync?%s", a.config.TodoistURL, command.ToQueryString())

	response, err := rest.Get(url)
	if err != nil {
		return errors.New(errorCommunicatingWithTodoistAPI)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New(errorExecutingCommand)
	}

	return nil
}
