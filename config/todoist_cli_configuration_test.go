package config

import "testing"

func TestIfAccessTokenIsSetThenIsAlreadyAuthenticated(t *testing.T) {
	authenticationConfiguration := &AuthenticationConfiguration{
		AccessToken: "access-token",
	}

	if !authenticationConfiguration.IsAuthenticated() {
		t.Error("If the access token is set on the authentication configuration, then the todoist-cli should be authenticated")
	}
}

func TestIfAccessTokenIsNotSetThenTheTodoistCliIsNotAuthenticated(t *testing.T) {
	authenticationConfiguration := &AuthenticationConfiguration{
		AccessToken: "",
	}

	if authenticationConfiguration.IsAuthenticated() {
		t.Error("If the access token is not set on the authentication configuration, then the todoist-cli should not be authenticated")
	}
}
