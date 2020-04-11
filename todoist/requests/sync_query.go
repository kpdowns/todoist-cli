package requests

import (
	"fmt"
)

// Query is a type of query that can be made against the Todoist API
type Query interface {
	ToQueryString() string
}

type syncQuery struct {
	Token         string
	SyncToken     string
	ResourceTypes ResourceTypes
}

// NewQuery creates a new instance of a Query
func NewQuery(token string, syncToken string, resourceTypes ResourceTypes) Query {
	return &syncQuery{
		Token:         token,
		SyncToken:     syncToken,
		ResourceTypes: resourceTypes,
	}
}

// ToQueryString converts the Query into a url query string to be provided on requests to Todoist
func (q *syncQuery) ToQueryString() string {
	return fmt.Sprintf(`token=%s&sync_token=%s&resource_types=%s`, q.Token, q.SyncToken, q.ResourceTypes.ToString())
}
