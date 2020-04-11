package mocks

import "github.com/kpdowns/todoist-cli/tasks/types"

// MockTaskService implements the TaskService interface and allows functions to be mocked
type MockTaskService struct {
	GetAllTasksFunctionToExecute func() ([]types.Task, error)
	AddTaskFunctionToExecute     func(content string) error
}

// AddTask executes the function configured in GetAllTasksFunctionToExecute
func (s *MockTaskService) AddTask(content string) error {
	if s.AddTaskFunctionToExecute != nil {
		return s.AddTaskFunctionToExecute(content)
	}
	panic("Method call AddTaskFunctionToExecute used but not configured")
}

// GetAllTasks executes the function configured in AddTaskFunctionToExecute
func (s *MockTaskService) GetAllTasks() ([]types.Task, error) {
	if s.GetAllTasksFunctionToExecute != nil {
		return s.GetAllTasksFunctionToExecute()
	}
	panic("Method call GetAllTasksFunctionToExecute used but not configured")
}
