package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/kpdowns/todoist-cli/config"
	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/rest"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
	"github.com/stretchr/testify/assert"
)

func TestRetrievingAccessTokens(t *testing.T) {
	config := config.TodoistCliConfiguration{}

	t.Run("When retrieving an access token and the Todoist API is unavailable, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("Rest client error")
			},
		}

		api := NewAPI(config)

		accessToken, err := api.GetAccessToken("code")
		assert.Nil(t, accessToken)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorCommunicatingWithTodoistAPI)
		}
	})

	t.Run("When retrieving an access token and a response that does not indicate success is received, then an error is returned", func(t *testing.T) {
		expectedStatusCode := 400

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: expectedStatusCode,
				}, nil
			},
		}

		api := NewAPI(config)

		accessToken, err := api.GetAccessToken("code")
		assert.Nil(t, accessToken)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorRetrievingAccessToken)
		}
	})

	t.Run("When retrieving an access token and the token cannot be decoded from JSON, then an error is returned", func(t *testing.T) {
		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				body := ""
				responseBody := ioutil.NopCloser(bytes.NewReader([]byte(body)))
				return &http.Response{
					StatusCode: 200,
					Body:       responseBody,
				}, nil
			},
		}

		api := NewAPI(config)

		accessToken, err := api.GetAccessToken("code")
		assert.Nil(t, accessToken)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorMalformedResponse)
		}
	})

	t.Run("When retrieving an access token and the response is valid, an access token is returned", func(t *testing.T) {
		expectedTokenValue := "test-token"

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				body := fmt.Sprintf(`{"access_token": "%s", "type": "bearer"}`, expectedTokenValue)
				responseBody := ioutil.NopCloser(bytes.NewReader([]byte(body)))
				return &http.Response{
					StatusCode: 200,
					Body:       responseBody,
				}, nil
			},
		}

		api := NewAPI(config)

		accessToken, err := api.GetAccessToken("code")
		assert.Nil(t, err)
		if assert.NotNil(t, accessToken) {
			assert.Equal(t, accessToken.AccessToken, expectedTokenValue)
		}
	})

}

func TestRevokingAccessToken(t *testing.T) {
	config := config.TodoistCliConfiguration{}

	t.Run("When revoking an access token and no token is provided, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{}

		api := NewAPI(config)

		err := api.RevokeAccessToken("")
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorRevokingAccessToken)
		}
	})

	t.Run("When revoking an access token and the Todoist API is unavailable, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("Rest client error")
			},
		}

		api := NewAPI(config)

		err := api.RevokeAccessToken("access-token")
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorCommunicatingWithTodoistAPI)
		}
	})

	t.Run("When revoking an access token and the response status code is not '204 - No Content', then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		err := api.RevokeAccessToken("access-token")
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorRevokingAccessToken)
		}
	})

	t.Run("When revoking an access token and the response status code is '204 - No Content', then no error is returned and the token can be considered revoked", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 204,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		err := api.RevokeAccessToken("access-token")
		assert.Nil(t, err)
	})
}

func TestExecutingSyncQueries(t *testing.T) {
	config := config.TodoistCliConfiguration{}

	t.Run("When executing a sync query and the Todoist API is unavailable, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("Rest client error")
			},
		}

		api := NewAPI(config)

		query := requests.NewQuery("token", "*", []requests.ResourceType{"all"})
		response, err := api.ExecuteSyncQuery(query)
		assert.Nil(t, response)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorCommunicatingWithTodoistAPI)
		}
	})

	t.Run("When executing a sync query and the response does not indicate success, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		query := requests.NewQuery("token", "*", []requests.ResourceType{"all"})
		response, err := api.ExecuteSyncQuery(query)
		assert.Nil(t, response)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorExecutingQuery)
		}
	})

	t.Run("When executing a sync query and the response indicate success but the body of the response can't be decoded, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		query := requests.NewQuery("token", "*", []requests.ResourceType{"all"})
		response, err := api.ExecuteSyncQuery(query)
		assert.Nil(t, response)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorMalformedResponse)
		}
	})

	t.Run("When executing a sync query and the response indicate success and the body can be decoded, the sync response is returned", func(t *testing.T) {

		expectedResponseObject := &responses.Query{
			IsFullSync: true,
			SyncToken:  "test-sync-token",
		}
		expectedReponseString, _ := json.Marshal(expectedResponseObject)

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(expectedReponseString))),
				}, nil
			},
		}

		api := NewAPI(config)

		query := requests.NewQuery("token", "*", []requests.ResourceType{"all"})
		response, err := api.ExecuteSyncQuery(query)

		assert.Nil(t, err)
		if assert.NotNil(t, response) {
			assert.Equal(t, response, expectedResponseObject)
		}
	})
}

func TestExecutingSyncCommands(t *testing.T) {
	config := config.TodoistCliConfiguration{}

	t.Run("When executing a sync command and the Todoist API is unavailable, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return nil, errors.New("Rest client error")
			},
		}

		api := NewAPI(config)

		arguments := make(map[string]interface{})
		arguments["content"] = "test-content"

		command := requests.NewCommand("test-token", "item_add", arguments)
		err := api.ExecuteSyncCommand(command)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorCommunicatingWithTodoistAPI)
		}
	})

	t.Run("When executing a sync command and the status code does not indicate success, then an error is returned", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		arguments := make(map[string]interface{})
		arguments["content"] = "test-content"

		command := requests.NewCommand("test-token", "item_add", arguments)
		err := api.ExecuteSyncCommand(command)
		if assert.NotNil(t, err) {
			assert.Equal(t, err.Error(), errorExecutingCommand)
		}
	})

	t.Run("When executing a sync command and the status code indicates success, then it is assumed the command executed successfully", func(t *testing.T) {

		rest.Client = &mocks.MockHTTPClient{
			DoFunction: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(""))),
				}, nil
			},
		}

		api := NewAPI(config)

		arguments := make(map[string]interface{})
		arguments["content"] = "test-content"

		command := requests.NewCommand("test-token", "item_add", arguments)
		err := api.ExecuteSyncCommand(command)
		assert.Nil(t, err)
	})
}
