# todoist-cli

[![Build Status](https://travis-ci.com/kpdowns/todoist-cli.svg?branch=master)](https://travis-ci.com/kpdowns/todoist-cli) [![codecov](https://codecov.io/gh/kpdowns/todoist-cli/branch/master/graph/badge.svg)](https://codecov.io/gh/kpdowns/todoist-cli) [![Go Report Card](https://goreportcard.com/badge/github.com/kpdowns/todoist-cli)](https://goreportcard.com/report/github.com/kpdowns/todoist-cli)

todoist-cli is an **unofficial** tool that allows you to interact with Todoist.com directly from the command line without using a browser. The motivation for it's existence is as a learning tool for myself.

Please familiarize yourself with the planned functionality for the todoist-cli below - the list is in order of priority. Items that are struck through are completed:


- ~~Authenticate againsts Todoist.com and save access token~~
- ~~Revoking of access tokens from Todoist.com~~
- ~~Listing of tasks across all projects on Todoist.com~~
- ~~Associating internal ids with Todoist tasks and projects to allow for easier management~~
- ~~Allow creation of basic tasks~~
- ~~Allow creation of tasks with priority and due date~~
- Allow flagging tasks as completed
- Allow reprioritization of tasks & changing due dates
- Display the project associated with a task
- Listing of tasks in a specific project
- Listing of all projects on Todoist.com
- Allow creation of projects
- Allow deletion of projects
- Allow creation of a task associated with a project
 
## Getting started
To get started developing the todoist-cli please make sure that you have:

- At least Go version 1.14 installed.

### Configuration
The todoist-cli loads configuration from both `./config/todoist_cli_configuration.go` and `./config/secrets/secrets.go` (this file is not committed to source control for obvious reasons). Before starting development, please make a copy of the file located `./config/secrets/secrets.sample` and rename it to be `secrets.go`. You will then be able to set the values after registering an application on Todoist at https://developer.todoist.com/.

> Please do not commit your version of the `secrets.go` file to source control. Doing so will leak sensitive configuration details of your Todoist application. To make it simpler to not make this mistake, this file is explicitly excluded in the `.gitignore`.

### Building the cli
For convenience, a launch configuration for Visual Studio Code is provided that will allow you to get started debugging immediately.

To build an executable version that can be shipped you can run the following command from the root directory.

```
go build -ldflags="-s -w"
```

This will create an executable that can be released.

### Running tests
In order to run the tests for the todoist-cli, you can run the following command in the root directory.

```
go test ./...
```