# todoist-cli
CLI tool for interacting with Todoist directly from the command line

## Getting started
To get started developing the Todoist-cli please make sure that you have:

- At least Go version 1.14 installed.
- gcc installed and on path

### Configuration
The Todoist-cli loads configuration files from the configuration files located in `./config.yml`. A sample configuration is provided in `./config.sample.yml`. Please update the sample configuration file and rename it to `config.yml` in order for the Todoist-cli to read it.

The values used in the configuration file can be found after registering an application on Todoist at https://developer.todoist.com/.

> Please do not commit your version of the `config.yml` file to source control. Doing so will leak sensitive configuration details of your Todoist application.

### Building the cli
For convenience, a launch configuration for Visual Studio Code is provided that will allow you to get started debugging immediately.

To build an executable version that can be shipped you can run the following command from the root directory.

```
go build
```

This will create an executable that can be shipped with a version of the `config.yml` file so that your application can be run.

### Running tests
In order to run the tests for the Todoist-cli, you can run the following command in the root directory.

```
go test
```