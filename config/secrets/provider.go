package secrets

//go:generate go run generate_secrets_file.go

// Secrets are the secrets required for the Todoist CLI
type Secrets struct {
	// ClientID is the ID of the client communicating with Todoist
	ClientID string

	// ClientSecret is the secret used when authenticating the todoist-cli with Todoist
	ClientSecret string
}

// GetSecrets returns the secrets required for the todoist-cli
func GetSecrets() Secrets {
	return Secrets{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
	}
}
