// +build ignore

package main

import "os"

var secretsFileContents = `
package secrets

const (
	// ClientID is the ID of the client communicating with Todoist
	ClientID string = ""

	// ClientSecret is the secret used when authenticating the todoist-cli with Todoist
	ClientSecret string = ""
)

`

func main() {
	const secretsFileName = "secrets.go"
	if _, err := os.Stat(secretsFileName); err == nil {
		panic("The secrets file already exists, continuing may overwrite your secrets. Halting...")
	}

	settingsFile, err := os.Create("secrets.go")
	if err != nil {
		panic("Failed to create 'secrets.go' file")
	}

	defer settingsFile.Close()
	settingsFile.WriteString(secretsFileContents)
}
