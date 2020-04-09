package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/spf13/viper"
)

// ClientConfiguration contains configuration for communicating with the Todoist API
type ClientConfiguration struct {
	TodoistURL          string `mapstructure:"todoist_url" yaml:"todoist_url"`
	ClientID            string `mapstructure:"client_id" yaml:"client_id"`
	ClientSecret        string `mapstructure:"client_secret" yaml:"client_secret"`
	RequiredPermissions string `mapstructure:"required_permissions" yaml:"required_permissions"`
	AppServiceURL       string `mapstructure:"app_service_url" yaml:"app_service_url"`
	OauthRedirectURL    string `mapstructure:"oauth_redirect_url" yaml:"oauth_redirect_url"`
}

// AuthenticationConfiguration contains the configured authentication related properties
type AuthenticationConfiguration struct {
	AccessToken string `mapstructure:"access_token" yaml:"access_token"`
}

// TodoistCliConfiguration contains the configuration required for the TodoistCli to function
type TodoistCliConfiguration struct {
	Client         ClientConfiguration         `mapstructure:"client" yaml:"client"`
	Authentication AuthenticationConfiguration `mapstructure:"auth" yaml:"auth"`
}

// LoadConfiguration loads the configuration file located in ./config.yml, emits an error if the configuration file is not valid
func LoadConfiguration() (*TodoistCliConfiguration, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.New("Error occurred while reading configuration file, no configuration file could be found")
	}

	var config TodoistCliConfiguration
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.New("Error occurred while reading configuration file, the configuration file is invalid")
	}

	return &config, nil
}

// IsAuthenticated checks whether the Todoist-cli is authenticated or not
func (todoistCliConfiguration *TodoistCliConfiguration) IsAuthenticated() bool {
	return todoistCliConfiguration.Authentication.AccessToken != ""
}

// StoreAccessToken stores the access token to be used for API requests in configuration
func (todoistCliConfiguration *TodoistCliConfiguration) StoreAccessToken(accessToken string) error {
	todoistCliConfiguration.Authentication.AccessToken = accessToken
	configToWrite, err := yaml.Marshal(todoistCliConfiguration)
	if err != nil {
		return err
	}

	fmt.Println(string(configToWrite))

	configFile, err := os.OpenFile("./config.yml", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		return err
	}

	defer configFile.Close()
	configFile.Truncate(0)
	configFile.Seek(0, 0)

	_, err = configFile.WriteString(string(configToWrite))
	if err != nil {
		return err
	}

	return nil
}
