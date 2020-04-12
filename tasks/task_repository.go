package tasks

import (
	"encoding/json"
	"errors"

	"github.com/kpdowns/todoist-cli/storage"
	"github.com/kpdowns/todoist-cli/tasks/types"
)

const (
	errorRepositoryNotAbleToGetTask     = "An error occurred while retrieving the persisted tasks"
	errorRepositoryTaskNotFound         = "The requested task does not exist"
	errorRepositoryErrorPersistingTasks = "An error occurred while persisting the list of tasks to disk"
	errorRepositoryErrorDeletingTasks   = "An error occurred deleting the persisted tasks"
)

// Repository handles persisting the task with the cli's own internal identifier
type Repository interface {
	GetAll() (types.TaskList, error)
	Get(types.TaskID) (*types.Task, error)
	CreateAll(types.TaskList) (types.TaskList, error)
	DeleteAll() error
}

type repository struct {
	file storage.File
}

// NewTaskRepository creates a new instance of a repository that handles persistence of tasks
func NewTaskRepository(file storage.File) Repository {
	return &repository{
		file: file,
	}
}

// GetAll retrieves all tasks, error if an error occurs while retrieving the tasks
func (r *repository) GetAll() (types.TaskList, error) {
	contents, err := r.file.ReadContents()
	if err != nil {
		return nil, errors.New(errorRepositoryNotAbleToGetTask)
	}

	var tasks types.TaskList
	err = json.Unmarshal([]byte(contents), &tasks)
	if err != nil {
		return nil, errors.New(errorRepositoryNotAbleToGetTask)
	}

	return tasks, nil
}

// Get retrieves a single task with the provided id, error if the task does not exist
func (r *repository) Get(taskID types.TaskID) (*types.Task, error) {
	tasks, err := r.GetAll()
	if err != nil {
		return nil, errors.New(errorRepositoryNotAbleToGetTask)
	}

	for _, task := range tasks {
		if task.ID == taskID {
			return &task, nil
		}
	}

	return nil, errors.New(errorRepositoryTaskNotFound)
}

// CreateAll persists all tasks with a generated id for later retrieval, returns a new list of tasks with the generated ids populated if there is no error
func (r *repository) CreateAll(tasks types.TaskList) (types.TaskList, error) {
	var tasksToPersist types.TaskList

	id := 1
	for _, task := range tasks {
		taskToPersist := types.Task{
			ID:        types.TaskID(id),
			Checked:   task.Checked,
			Content:   task.Content,
			DayOrder:  task.DayOrder,
			DueDate:   task.DueDate,
			Priority:  task.Priority,
			TodoistID: task.TodoistID,
		}
		tasksToPersist = append(tasksToPersist, taskToPersist)
		id++
	}

	taskString, _ := json.Marshal(tasksToPersist)
	err := r.file.OverwriteContents(string(taskString))
	if err != nil {
		return nil, errors.New(errorRepositoryErrorPersistingTasks)
	}

	return tasksToPersist, nil
}

// DeleteAll deletes all tasks that have been persisted, returns error if an error occurs
func (r *repository) DeleteAll() error {
	err := r.file.OverwriteContents("")
	if err != nil {
		return errors.New(errorRepositoryErrorDeletingTasks)
	}

	return nil
}
