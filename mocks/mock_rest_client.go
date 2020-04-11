package mocks

import "net/http"

// MockHTTPClient is the mock http client
type MockHTTPClient struct {
	DoFunction func(r *http.Request) (*http.Response, error)
}

// Do executes the configured Do function
func (m *MockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	if m.DoFunction != nil {
		return m.DoFunction(r)
	}
	panic("Method call Do used but not configured")
}
