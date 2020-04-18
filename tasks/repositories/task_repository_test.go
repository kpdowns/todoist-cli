package repositories

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/kpdowns/todoist-cli/mocks"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/stretchr/testify/assert"
)

func TestGettingAllTasks(t *testing.T) {

	t.Run("When getting all tasks, if no error occurs, then the tasks are returned", func(t *testing.T) {

		expectedTasks := types.TaskList{
			{
				ID:        1,
				Checked:   0,
				Content:   "test-tasks",
				DayOrder:  1,
				DueDate:   time.Now().UTC(),
				Priority:  1,
				TodoistID: 1,
			},
		}
		expectedBytes, _ := json.Marshal(expectedTasks)
		expectedContents := string(expectedBytes)

		inMemoryFile := &mocks.MockFile{
			Contents: string(expectedContents),
		}

		repository := NewTaskRepository(inMemoryFile)

		actualTasks, err := repository.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, actualTasks)
		assert.Equal(t, expectedTasks, actualTasks)

	})

	t.Run("When getting all tasks, if an error occurs while deserializing the tasks the contents on disk, then the error is returned", func(t *testing.T) {

		inMemoryFile := &mocks.MockFile{
			Contents: string("not valid json"),
		}

		repository := NewTaskRepository(inMemoryFile)

		tasks, err := repository.GetAll()
		assert.Equal(t, errorRepositoryNotAbleToGetTask, err.Error())
		assert.NotNil(t, err)
		assert.Nil(t, tasks)

	})

	t.Run("When getting all tasks, if an error occurs, then the error is returned", func(t *testing.T) {

		inMemoryFile := &mocks.MockFile{
			ReadError: errors.New("test error"),
		}

		repository := NewTaskRepository(inMemoryFile)

		tasks, err := repository.GetAll()
		assert.Equal(t, errorRepositoryNotAbleToGetTask, err.Error())
		assert.NotNil(t, err)
		assert.Nil(t, tasks)

	})

}

func TestGettingAnIndividualTask(t *testing.T) {

	t.Run("When retrieving a single task, if an error occurs, then an error is returned", func(t *testing.T) {

		inMemoryFile := &mocks.MockFile{
			ReadError: errors.New("test error"),
		}

		repository := NewTaskRepository(inMemoryFile)

		tasks, err := repository.Get(1)
		assert.Equal(t, errorRepositoryNotAbleToGetTask, err.Error())
		assert.NotNil(t, err)
		assert.Nil(t, tasks)

	})

	t.Run("When retrieving a single task, if the task exists, then the task is returned", func(t *testing.T) {

		taskToBeRetrieved := &types.Task{
			ID: 1,
		}

		expectedTasks := types.TaskList{*taskToBeRetrieved}
		expectedBytes, _ := json.Marshal(expectedTasks)
		expectedContents := string(expectedBytes)

		inMemoryFile := &mocks.MockFile{
			Contents: string(expectedContents),
		}

		repository := NewTaskRepository(inMemoryFile)

		task, err := repository.Get(taskToBeRetrieved.ID)
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, taskToBeRetrieved, task)

	})

	t.Run("When retrieving a single task, if the task does not exist, then an error is returned", func(t *testing.T) {

		expectedTasks := types.TaskList{
			types.Task{},
		}
		expectedBytes, _ := json.Marshal(expectedTasks)
		expectedContents := string(expectedBytes)

		inMemoryFile := &mocks.MockFile{
			Contents: string(expectedContents),
		}

		repository := NewTaskRepository(inMemoryFile)

		task, err := repository.Get(1)
		assert.NotNil(t, err)
		assert.Equal(t, errorRepositoryTaskNotFound, err.Error())
		assert.Nil(t, task)

	})

}

func TestPersistingAllTasks(t *testing.T) {

	t.Run("Given a list of tasks, when persisting the tasks, the tasks are assigned an id before being written to storage", func(t *testing.T) {
		tasksToWrite := types.TaskList{
			types.Task{
				// expected to have id of 1 when read back from storage
				TodoistID: 100,
			},
			types.Task{
				// expected to have id of 2 when read back from storage
				TodoistID: 200,
			},
		}

		inMemoryFile := &mocks.MockFile{}
		repository := NewTaskRepository(inMemoryFile)

		_, err := repository.CreateAll(tasksToWrite)

		assert.Nil(t, err)

		var tasksAfterBeingWritten types.TaskList
		json.Unmarshal([]byte(inMemoryFile.Contents), &tasksAfterBeingWritten)
		assert.Equal(t, uint32(1), tasksAfterBeingWritten[0].ID)
		assert.Equal(t, int64(100), tasksAfterBeingWritten[0].TodoistID)
		assert.Equal(t, uint32(2), tasksAfterBeingWritten[1].ID)
		assert.Equal(t, int64(200), tasksAfterBeingWritten[1].TodoistID)

	})

	t.Run("Given a list of tasks, when persisting the tasks, the returned tasks contain the same tasks as those persisted to storage", func(t *testing.T) {
		tasksToWrite := types.TaskList{
			types.Task{
				// expected to have id of 1 when read back from storage
				TodoistID: 100,
			},
			types.Task{
				// expected to have id of 2 when read back from storage
				TodoistID: 200,
			},
		}

		inMemoryFile := &mocks.MockFile{}
		repository := NewTaskRepository(inMemoryFile)

		tasksAfterBeingWritten, err := repository.CreateAll(tasksToWrite)

		assert.Nil(t, err)
		var storedTasks types.TaskList
		json.Unmarshal([]byte(inMemoryFile.Contents), &storedTasks)
		assert.Equal(t, storedTasks[0].ID, tasksAfterBeingWritten[0].ID)
		assert.Equal(t, storedTasks[0].TodoistID, tasksAfterBeingWritten[0].TodoistID)
		assert.Equal(t, storedTasks[1].ID, tasksAfterBeingWritten[1].ID)
		assert.Equal(t, storedTasks[1].TodoistID, tasksAfterBeingWritten[1].TodoistID)

	})

	t.Run("Given a list of tasks, when persisting the tasks and an error occurs while persisting the tasks to disk, an error is returned", func(t *testing.T) {
		tasksToWrite := types.TaskList{}
		inMemoryFile := &mocks.MockFile{
			OverwriteError: errors.New("test error"),
		}
		repository := NewTaskRepository(inMemoryFile)

		tasksAfterBeingWritten, err := repository.CreateAll(tasksToWrite)
		assert.NotNil(t, err)
		assert.Nil(t, tasksAfterBeingWritten)
		assert.Equal(t, errorRepositoryErrorPersistingTasks, err.Error())

	})

}

func TestDeletingTasks(t *testing.T) {

	t.Run("When deleting all tasks and an error occurs, then an error is returned", func(t *testing.T) {

		inMemoryFile := &mocks.MockFile{
			OverwriteError: errors.New("test error"),
		}

		repository := NewTaskRepository(inMemoryFile)

		err := repository.DeleteAll()
		assert.NotNil(t, err)
		assert.Equal(t, errorRepositoryErrorDeletingTasks, err.Error())

	})

	t.Run("When deleting all tasks and and no error occurs, the tasks are deleted from the disk", func(t *testing.T) {

		inMemoryFile := &mocks.MockFile{
			Contents: "test contents of file",
		}

		repository := NewTaskRepository(inMemoryFile)

		err := repository.DeleteAll()
		assert.Nil(t, err)
		assert.Equal(t, "", inMemoryFile.Contents)

	})

}
