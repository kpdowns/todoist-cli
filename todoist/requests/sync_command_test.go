package requests

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/beevik/guid"
)

func TestWhenGettingTheCommandsQueryStringThenTheQueryStringIsAsExpected(t *testing.T) {
	arguments := struct {
		Name  string `json:"name"`
		Color int    `json:"color"`
	}{
		Name:  "project1",
		Color: 1,
	}

	tempID := guid.NewString()
	uuid := guid.NewString()

	command := NewCommand("token", "project_add", arguments)
	command.Commands[0].TemporaryID = tempID
	command.Commands[0].UUID = uuid

	expectedCommandString := fmt.Sprintf(`[{"type":"project_add","temp_id":"%s","uuid":"%s","args":{"name":"project1","color":1}}]`, tempID, uuid)
	expectedEscapedCommandString := url.QueryEscape(expectedCommandString)
	expectedQueryString := fmt.Sprintf(`token=token&commands=%s`, expectedEscapedCommandString)
	actualQueryString := command.ToQueryString()
	if expectedQueryString != actualQueryString {
		t.Errorf("Expected '%s', got '%s'", expectedQueryString, actualQueryString)
	}
}
