package types

// AuthenticationResponse contains the code received as part of the oauth process, the code is used to retrieve an access token
type AuthenticationResponse struct {
	Code string
}
