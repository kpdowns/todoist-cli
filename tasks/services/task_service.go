package services

import (
	"errors"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/tasks/repositories"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/requests/commands"
)

const (
	errorNotCurrentlyAuthenticated   = "Error, you are not currently logged in."
	errorOccurredDuringSyncOperation = "Error occurred while syncing with Todoist."
	errorNoContent                   = "Task content must be provided when adding a task."
	errorNoTaskToComplete            = "The requested task does not exist."
	errorFailedToCompleteTask        = "An error occurred while flagging the task as completed on Todoist, please try again."
)

// TaskService provides functionality to retrieve and update tasks on Todoist
type TaskService interface {
	GetAllTasks() (types.TaskList, error)
	AddTask(content string, due string, priority int) error
	CompleteTask(taskID uint32) error
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

	command := requests.NewCommand(accessToken.AccessToken, commands.ItemAdd, arguments)
	err = s.api.ExecuteSyncCommand(command)
	if err != nil {
		return errors.New(errorOccurredDuringSyncOperation)
	}

	return nil
}

func (s *taskService) CompleteTask(taskID uint32) error {
	isAuthenticated, err := s.authenticationService.IsAuthenticated()
	if err != nil || !isAuthenticated {
		return errors.New(errorNotCurrentlyAuthenticated)
	}

	accessToken, _ := s.authenticationService.GetAccessToken()

	taskToComplete, err := s.taskRepository.Get(taskID)
	if err != nil {
		return errors.New(errorNoTaskToComplete)
	}

	arguments := make(map[string]interface{})
	arguments["id"] = taskToComplete.TodoistID
	command := requests.NewCommand(accessToken.AccessToken, commands.ItemClose, arguments)
	err = s.api.ExecuteSyncCommand(command)
	if err != nil {
		return errors.New(errorFailedToCompleteTask)
	}

	return nil
}
