package responses

// Query is the response received as a result of a sync query
type Query struct {
	IsFullSync bool   `json:"full_sync"`
	Items      []Item `json:"items"`
	SyncToken  string `json:"sync_token"`
}
