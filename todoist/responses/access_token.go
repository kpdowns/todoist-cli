package responses

// AccessToken is the received access token from Todoist after signing in
type AccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
