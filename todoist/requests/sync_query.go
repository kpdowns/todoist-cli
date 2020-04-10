package requests

import "fmt"

// SyncQuery is a type of query that can be made against the Todoist API
type SyncQuery interface {
	ToQueryString() string
}

type syncQuery struct {
	Token         string
	SyncToken     string
	ResourceTypes ResourceTypes
}

// NewSyncQuery creates a new instance of a SyncQuery
func NewSyncQuery(token string, syncToken string, resourceTypes ResourceTypes) SyncQuery {
	return &syncQuery{
		Token:         token,
		SyncToken:     syncToken,
		ResourceTypes: resourceTypes,
	}
}

// ToQueryString converts the SyncQuery into a url query string to be provided on requests to the Todoist API
func (q *syncQuery) ToQueryString() string {
	return fmt.Sprintf(`token=%s&sync_token=%s&resource_types=%s`, q.Token, q.SyncToken, q.ResourceTypes.ToString())
}
