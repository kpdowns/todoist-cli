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
The todoist-cli loads configuration files from the configuration files located in `./config.yml`. A sample configuration is provided in `./config.sample.yml`. Please update the sample configuration file and rename it to `config.yml` in order for the todoist-cli to read it.

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
In order to run the tests for the todoist-cli, you can run the following command in the root directory.

```
go test ./...
```