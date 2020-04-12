package mocks

import (
	"github.com/kpdowns/todoist-cli/tasks/types"
)

// MockTaskRepository provides overrides for the functions of the repository for testing purposes
type MockTaskRepository struct {
	GetAllFunc    func() (types.TaskList, error)
	GetFunc       func(types.TaskID) (*types.Task, error)
	CreateAllFunc func(types.TaskList) (types.TaskList, error)
	DeleteAllFunc func() error
}

// GetAll retrieves all tasks, error if an error occurs while retrieving the tasks
func (r *MockTaskRepository) GetAll() (types.TaskList, error) {
	return r.GetAllFunc()
}

// Get retrieves a single task with the provided id, error if the task does not exist
func (r *MockTaskRepository) Get(taskID types.TaskID) (*types.Task, error) {
	return r.GetFunc(taskID)
}

// CreateAll persists all tasks with a generated id for later retrieval, returns a new list of tasks with the generated ids populated if there is no error
func (r *MockTaskRepository) CreateAll(tasks types.TaskList) (types.TaskList, error) {
	return r.CreateAllFunc(tasks)
}

// DeleteAll deletes all tasks that have been persisted, returns error if an error occurs
func (r *MockTaskRepository) DeleteAll() error {
	return r.DeleteAllFunc()
}
