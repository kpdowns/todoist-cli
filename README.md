# todoist-cli

[![Build Status](https://travis-ci.com/kpdowns/todoist-cli.svg?branch=master)](https://travis-ci.com/kpdowns/todoist-cli) [![codecov](https://codecov.io/gh/kpdowns/todoist-cli/branch/master/graph/badge.svg)](https://codecov.io/gh/kpdowns/todoist-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/kpdowns/todoist-cli)](https://goreportcard.com/report/github.com/kpdowns/todoist-cli)

Todoist-CLI is an **unofficial** tool that allows you to interact with Todoist.com directly from the command line without using a browser. The motivation for it's existence is as a learning tool for myself.

Please familiarize yourself with the planned functionality for the Todoist-cli (items that are struck through are completed):


- ~~Authenticate againsts Todoist.com and save access token~~
- ~~Revoking of access tokens from Todoist.com~~
- Listing of tasks across all projects on Todoist.com
- Listing of all projects on Todoist.com
- Associating internal ids with Todoist tasks and projects to allow for easier management
- Allow creation of projects
- Allow deletion of projects
- Allow creation of basic tasks (with no priority or defined project)
- Allow creation of tasks with a user defined priority and project
- Allow reprioritization of tasks & changing due dates
- Allow flagging tasks as completed

 
## Getting started
To get started developing the Todoist-cli please make sure that you have:

- At least Go version 1.14 installed.

### Configuration
The Todoist-cli loads configuration files from the configuration files located in `./config.yml`. A sample configuration is provided in `./config.sample.yml`. Please update the sample configuration file and rename it to `config.yml` in order for the Todoist-cli to read it.

The values used in the configuration file can be found after registering an application on Todoist at https://developer.todoist.com/.

> Please do not commit your version of the `config.yml` file to source control. Doing so will leak sensitive configuration details of your Todoist application.

### Building the cli
For convenience, a launch configuration for Visual Studio Code is provided that will allow you to get started debugging immediately.

To build an executable version that can be shipped you can run the following command from the root directory.

```
go build -ldflags="-s -w"
```

This will create an executable that can be shipped with a version of the `config.yml` file so that your application can be run.

### Running tests
In order to run the tests for the Todoist-cli, you can run the following command in the root directory.

```
go test ./...
```