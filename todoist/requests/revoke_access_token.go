package requests

// RevokeAccessToken is the request used to logout of Todoist
type RevokeAccessToken struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}
