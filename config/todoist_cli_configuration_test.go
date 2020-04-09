package config

import "testing"

func TestIfAccessTokenIsSetThenIsAlreadyAuthenticated(t *testing.T) {
	todoistCliConfiguration := &TodoistCliConfiguration{
		Authentication: AuthenticationConfiguration{
			AccessToken: "access-token",
		},
	}

	if !todoistCliConfiguration.IsAuthenticated() {
		t.Error("If the access token is set on the authentication configuration, then the todoist-cli should be authenticated")
	}
}

func TestIfAccessTokenIsNotSetThenTheTodoistCliIsNotAuthenticated(t *testing.T) {
	todoistCliConfiguration := &TodoistCliConfiguration{
		Authentication: AuthenticationConfiguration{
			AccessToken: "",
		},
	}

	if todoistCliConfiguration.IsAuthenticated() {
		t.Error("If the access token is not set on the authentication configuration, then the todoist-cli should not be authenticated")
	}
}
