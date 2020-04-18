package requests

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/beevik/guid"
	"github.com/kpdowns/todoist-cli/todoist/requests/commands"
)

// Command contains commands to execute against Todoist
type Command struct {
	Token    string
	Commands []CommandDetail
}

// CommandDetail is an individual command to be executed
type CommandDetail struct {
	Type        commands.CommandType   `json:"type"`
	TemporaryID string                 `json:"temp_id"`
	UUID        string                 `json:"uuid"`
	Arguments   map[string]interface{} `json:"args"`
}

// NewCommand creates a new instance of a Todoist Sync Command
func NewCommand(token string, commandType commands.CommandType, arguments map[string]interface{}) Command {
	return Command{
		Token: token,
		Commands: []CommandDetail{
			{
				Type:        commandType,
				TemporaryID: guid.NewString(),
				UUID:        guid.NewString(),
				Arguments:   arguments,
			},
		},
	}
}

// ToQueryString generates a query string to be provided as part of the URL in a sync command
func (c *Command) ToQueryString() string {
	commandStringAsJSON, _ := json.Marshal(c.Commands)
	commandString := string(commandStringAsJSON)
	commandStringAsQueryString := url.QueryEscape(commandString)
	return fmt.Sprintf("token=%s&commands=%s", c.Token, commandStringAsQueryString)
}
