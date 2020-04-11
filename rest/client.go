package rest

import (
	"bytes"
	"net/http"
	"time"
)

var (
	// Client is the HTTP client to be used
	Client HTTPClient
)

// HTTPClient is an HTTP client that performs requests
type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{
		Timeout: time.Second * 10,
	}
}

// Post sends a post request to the url with a body
func Post(url string, contentType string, body *bytes.Buffer) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", contentType)
	return Client.Do(request)
}

// Get sends a get request to the url
func Get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	return Client.Do(request)
}
