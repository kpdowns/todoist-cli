package requests

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/beevik/guid"
)

func TestSyncCommandSerialization(t *testing.T) {

	t.Run("Given a command, when converting the command to query string, the token is a parameter, and the commands are a URL escaped JSON object", func(t *testing.T) {
		arguments := make(map[string]interface{})
		arguments["color"] = 1
		arguments["name"] = "project1"

		tempID := guid.NewString()
		uuid := guid.NewString()

		command := NewCommand("token", "project_add", arguments)
		command.Commands[0].TemporaryID = tempID
		command.Commands[0].UUID = uuid

		commandString := fmt.Sprintf(`[{"type":"project_add","temp_id":"%s","uuid":"%s","args":{"color":1,"name":"project1"}}]`, tempID, uuid)
		urlEscaptedCommandString := url.QueryEscape(commandString)

		expected := fmt.Sprintf(`token=token&commands=%s`, urlEscaptedCommandString)
		actual := command.ToQueryString()

		assert.Equal(t, expected, actual)
	})
}
