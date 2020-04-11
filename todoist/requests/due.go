package requests

// Due is the representation of a due date (only supports simple plain-text dates)
type Due struct {
	Value string `json:"string"`
}
