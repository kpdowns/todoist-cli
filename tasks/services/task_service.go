package services

import (
	"errors"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/repositories"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/kpdowns/todoist-cli/todoist/requests"
)

const (
	errorNotCurrentlyAuthenticated   = "Error, you are not currently logged in"
	errorOccurredDuringSyncOperation = "Error occurred while syncing with Todoist"
	errorNoContent                   = "Task content must be provided when adding a task"
)

// TaskService provides functionality to handle the access token used by the Todoist API
type TaskService interface {
	GetAllTasks() (types.TaskList, error)
	AddTask(content string, due string, priority int) error
}

type taskService struct {
	api                   todoist.API
	authenticationService authentication.Service
	taskRepository        repositories.TaskRepository
}

// NewTaskService creates a new instance of the task service
func NewTaskService(api todoist.API, authenticationService authentication.Service, taskRepository repositories.TaskRepository) TaskService {
	return &taskService{
		api:                   api,
		authenticationService: authenticationService,
		taskRepository:        taskRepository,
	}
}

// GetAllTasks returns a list of tasks to do, sorted by day order
func (s *taskService) GetAllTasks() (types.TaskList, error) {
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

	persistedTasks, err := s.taskRepository.CreateAll(sortedTasks)
	if err != nil {
		return nil, err
	}

	return persistedTasks, nil
}

func (s *taskService) AddTask(content string, due string, priority int) error {
	if content == "" {
		return errors.New(errorNoContent)
	}

	isAuthenticated, err := s.authenticationService.IsAuthenticated()
	if err != nil || !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	accessToken, _ := s.authenticationService.GetAccessToken()

	arguments := make(map[string]interface{})
	arguments["content"] = content

	if due != "" {
		arguments["due"] = &requests.Due{
			Value: due,
		}
	}
	if priority != 0 {
		arguments["priority"] = priority
	}

	command := requests.NewCommand(accessToken.AccessToken, "item_add", arguments)
	err = s.api.ExecuteSyncCommand(command)
	if err != nil {
		return errors.New(errorOccurredDuringSyncOperation)
	}

	return nil
}
