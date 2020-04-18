package services

import (
	"errors"
	"testing"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/tasks/repositories"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
	"github.com/stretchr/testify/assert"
)

func TestGettingAllTasks(t *testing.T) {

	t.Run("When getting all tasks and the client is not authenticated, then an error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			IsAuthenticatedErrorToReturn: errors.New("test error"),
		}
		mockAPI := &mocks.MockAPI{}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		_, err := taskService.GetAllTasks()
		assert.NotNil(t, err)
		assert.Equal(t, errorNotCurrentlyAuthenticated, err.Error())

	})

	t.Run("When getting all tasks and an error is returned from the API, then an error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(syncQuery requests.Query) (*responses.Query, error) {
				return nil, errors.New("Error")
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		_, err := taskService.GetAllTasks()

		assert.NotNil(t, err)
		assert.Equal(t, errorOccurredDuringSyncOperation, err.Error())

	})

	t.Run("When getting all tasks and and no error occurs, then the tasks are saved into the task repository", func(t *testing.T) {

		wasCreateAllCalled := false

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(query requests.Query) (*responses.Query, error) {
				return &responses.Query{}, nil
			},
		}
		mockRepository := &mocks.MockTaskRepository{
			CreateAllFunc: func(types.TaskList) (types.TaskList, error) {
				wasCreateAllCalled = true
				return nil, nil
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		_, err := taskService.GetAllTasks()

		assert.Nil(t, err)
		assert.True(t, wasCreateAllCalled)

	})

	t.Run("When getting all tasks and and no error occurs, then the tasks are sorted before being saved in the repository", func(t *testing.T) {

		item1 := responses.Item{
			TodoistID: 2,
			DayOrder:  1,
			Priority:  1,
			Due: &responses.Due{
				DateString: "2020-04-13",
			},
		}

		item2 := responses.Item{
			TodoistID: 1,
			DayOrder:  1,
			Priority:  4,
			Due: &responses.Due{
				DateString: "2020-04-13",
			},
		}

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(query requests.Query) (*responses.Query, error) {
				return &responses.Query{
					Items: []responses.Item{
						item1,
						item2,
					},
				}, nil
			},
		}
		repository := repositories.NewTaskRepository(&mocks.MockFile{})
		taskService := NewTaskService(mockAPI, mockAuthenticationService, repository)

		returnedTasks, err := taskService.GetAllTasks()

		assert.Nil(t, err)

		repositoryTasks, _ := repository.GetAll()
		assert.NotNil(t, repositoryTasks)
		assert.Equal(t, repositoryTasks, returnedTasks)

		assert.Equal(t, item2.TodoistID, repositoryTasks[0].TodoistID)
		assert.Equal(t, item1.TodoistID, repositoryTasks[1].TodoistID)

	})

	t.Run("When getting all tasks and and no error occurs, and an error occurs while persisting the tasks, then an error is returned", func(t *testing.T) {

		expectedError := errors.New("Error")

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(query requests.Query) (*responses.Query, error) {
				return &responses.Query{}, nil
			},
		}
		mockRepository := &mocks.MockTaskRepository{
			CreateAllFunc: func(types.TaskList) (types.TaskList, error) {
				return nil, expectedError
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		_, err := taskService.GetAllTasks()

		assert.NotNil(t, err)
		assert.Equal(t, expectedError.Error(), err.Error())

	})

}

func TestAddingANewTask(t *testing.T) {

	t.Run("When adding a task and the client is not authenticated, then error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			IsAuthenticatedErrorToReturn: errors.New("test error"),
		}
		mockAPI := &mocks.MockAPI{}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		err := taskService.AddTask("content", "today", 1)
		assert.NotNil(t, err)
		assert.Equal(t, errorNotCurrentlyAuthenticated, err.Error())

	})

	t.Run("When adding a task and no content is provided, then error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncCommandFunction: func(command requests.Command) error {
				return nil
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		err := taskService.AddTask("", "today", 1)
		assert.NotNil(t, err)
		assert.Equal(t, errorNoContent, err.Error())

	})

	t.Run("When adding a task and the api returns an error, then an error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncCommandFunction: func(command requests.Command) error {
				return errors.New("test error")
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		err := taskService.AddTask("content", "today", 1)
		assert.NotNil(t, err)
		assert.Equal(t, errorOccurredDuringSyncOperation, err.Error())

	})

	t.Run("When adding a task and no error occurs, then no error is returned", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncCommandFunction: func(command requests.Command) error {
				return nil
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, nil)

		err := taskService.AddTask("content", "today", 1)

		assert.Nil(t, err)

	})

}

func TestCompletingATask(t *testing.T) {

	t.Run("When completing a task, and the client is not authenticated, then an error is returned", func(t *testing.T) {

		mockAPI := &mocks.MockAPI{}
		mockRepository := &mocks.MockTaskRepository{}
		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: false,
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		err := taskService.CompleteTask(1)
		assert.NotNil(t, err)
		assert.Equal(t, errorNotCurrentlyAuthenticated, err.Error())

	})

	t.Run("When completing a task, and the task does not exist, then an error is returned", func(t *testing.T) {

		mockAPI := &mocks.MockAPI{}
		mockRepository := &mocks.MockTaskRepository{
			GetFunc: func(uint32) (*types.Task, error) {
				return nil, errors.New("Test error")
			},
		}
		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		err := taskService.CompleteTask(1)
		assert.NotNil(t, err)
		assert.Equal(t, errorNoTaskToComplete, err.Error())

	})

	t.Run("When completing a task, and the task exists, if when calling the API to complete the task an error occurs then an error is returned", func(t *testing.T) {

		mockRepository := &mocks.MockTaskRepository{
			GetFunc: func(uint32) (*types.Task, error) {
				return &types.Task{
					ID: 1,
				}, nil
			},
		}
		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncCommandFunction: func(command requests.Command) error {
				return errors.New("Test error")
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		err := taskService.CompleteTask(1)
		assert.NotNil(t, err)
		assert.Equal(t, errorFailedToCompleteTask, err.Error())

	})

	t.Run("When completing a task, and the task exists, no error occurs while calling the Todoist API, then no error is returned", func(t *testing.T) {

		mockRepository := &mocks.MockTaskRepository{
			GetFunc: func(uint32) (*types.Task, error) {
				return &types.Task{
					ID: 1,
				}, nil
			},
		}
		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncCommandFunction: func(command requests.Command) error {
				return nil
			},
		}

		taskService := NewTaskService(mockAPI, mockAuthenticationService, mockRepository)

		err := taskService.CompleteTask(1)
		assert.Nil(t, err)

	})

}
