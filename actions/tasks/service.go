package tasks

import (
	"errors"
	"sort"

	"github.com/kpdowns/todoist-cli/actions/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist/requests"

	"github.com/kpdowns/todoist-cli/authentication"
	"github.com/kpdowns/todoist-cli/todoist"
)

// Service provides functionality to handle the access token used by the Todoist API
type Service interface {
	GetAllTasks() ([]types.Task, error)
}

type service struct {
	api                   todoist.API
	authenticationService authentication.Service
}

// NewService creates a new instance of the Authentication service
func NewService(api todoist.API, authenticationService authentication.Service) Service {
	return &service{
		api:                   api,
		authenticationService: authenticationService,
	}
}

// GetAllTasks returns a list of tasks to do, sorted by day order
func (s *service) GetAllTasks() ([]types.Task, error) {
	isAuthenticated, err := s.authenticationService.IsAuthenticated()
	if err != nil {
		return nil, err
	}

	if !isAuthenticated {
		return nil, errors.New(errorNotCurrentlyAuthenticated)
	}

	accessToken, _ := s.authenticationService.GetAccessToken()
	resourceTypes := []requests.ResourceType{"items"}
	syncQuery := requests.NewSyncQuery(accessToken, "*", resourceTypes)

	syncResponse, err := s.api.ExecuteSyncQuery(syncQuery)
	if err != nil {
		return nil, err
	}

	var tasks types.TaskList
	for _, item := range syncResponse.Items {
		newTask := types.Task{
			Checked:   item.Checked,
			Content:   item.Content,
			DayOrder:  item.DayOrder,
			DueDate:   item.Due.DateString,
			Priority:  item.Priority,
			TodoistID: item.TodoistID,
		}
		tasks = append(tasks, newTask)
	}

	sort.Sort(tasks)

	return tasks, nil
}
