package tasks

import (
	"errors"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/kpdowns/todoist-cli/todoist/requests"
)

const (
	errorNotCurrentlyAuthenticated   = "Error, you are not currently logged in"
	errorOccurredDuringSyncOperation = "Error occurred while syncing with Todoist"
	errorNoContent                   = "Task content must be provided when adding a task"
)

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	GetAllTasks() ([]types.Task, error)
	AddTask(content string) error
}

type service struct {
	api                   todoist.API
	authenticationService authentication.Service
}

// NewTaskService creates a new instance of the task service
func NewTaskService(api todoist.API, authenticationService authentication.Service) Service {
	return &service{
		api:                   api,
		authenticationService: authenticationService,
	}
}

// GetAllTasks returns a list of tasks to do, sorted by day order
func (s *service) GetAllTasks() ([]types.Task, error) {
	isAuthenticated, err := s.authenticationService.IsAuthenticated()
	if err != nil || !isAuthenticated {
		return nil, errors.New(errorNotCurrentlyAuthenticated)
	}

	accessToken, _ := s.authenticationService.GetAccessToken()
	resourceTypes := []requests.ResourceType{"items"}
	syncQuery := requests.NewQuery(accessToken.AccessToken, "*", resourceTypes)

	syncResponse, err := s.api.ExecuteSyncQuery(syncQuery)
	if err != nil {
		return nil, errors.New(errorOccurredDuringSyncOperation)
	}

	var tasks types.TaskList
	for _, item := range syncResponse.Items {
		newTask := item.ToTask()
		tasks = append(tasks, newTask)
	}

	sortedTasks := tasks.SortByDueDateThenSortByPriority()
	return sortedTasks, nil
}

func (s *service) AddTask(content string) error {
	if content == "" {
		return errors.New(errorNoContent)
	}

	isAuthenticated, err := s.authenticationService.IsAuthenticated()
	if err != nil || !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	accessToken, _ := s.authenticationService.GetAccessToken()

	arguments := struct {
		Content string `json:"content"`
	}{
		Content: content,
	}

	command := requests.NewCommand(accessToken.AccessToken, "item_add", arguments)
	err = s.api.ExecuteSyncCommand(command)
	if err != nil {
		return errors.New(errorOccurredDuringSyncOperation)
	}

	return nil
}
