package requests

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/beevik/guid"
)

func TestWhenGettingTheCommandsQueryStringThenTheQueryStringIsAsExpected(t *testing.T) {
	arguments := make(map[string]interface{})
	arguments["color"] = 1
	arguments["name"] = "project1"

	tempID := guid.NewString()
	uuid := guid.NewString()

	command := NewCommand("token", "project_add", arguments)
	command.Commands[0].TemporaryID = tempID
	command.Commands[0].UUID = uuid

	expectedCommandString := fmt.Sprintf(`[{"type":"project_add","temp_id":"%s","uuid":"%s","args":{"color":1,"name":"project1"}}]`, tempID, uuid)
	expectedEscapedCommandString := url.QueryEscape(expectedCommandString)
	expectedQueryString := fmt.Sprintf(`token=token&commands=%s`, expectedEscapedCommandString)
	actualQueryString := command.ToQueryString()
	if expectedQueryString != actualQueryString {
		t.Errorf("Expected '%s', got '%s'", expectedQueryString, actualQueryString)
	}
}
