package requests

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/beevik/guid"
)

// Command contains commands to execute against Todoist
type Command struct {
	Token    string
	Commands []CommandDetail
}

// CommandDetail is an individual command to be executed
type CommandDetail struct {
	Type        CommandType `json:"type"`
	TemporaryID string      `json:"temp_id"`
	UUID        string      `json:"uuid"`
	Arguments   interface{} `json:"args"`
}

// NewCommand creates a new instance of a Todoist Sync Command
func NewCommand(token string, commandType CommandType, arguments interface{}) Command {
	return Command{
		Token: token,
		Commands: []CommandDetail{
			CommandDetail{
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
	commandStringAsQueryString := url.QueryEscape(string(commandStringAsJSON))
	return fmt.Sprintf("token=%s&commands=%s", c.Token, commandStringAsQueryString)
}