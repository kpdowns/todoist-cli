package list

import (
	"bytes"
	"testing"

	"github.com/kpdowns/todoist-cli/tasks/services"
	"github.com/kpdowns/todoist-cli/tasks/types"
	"github.com/kpdowns/todoist-cli/todoist/requests"
	"github.com/kpdowns/todoist-cli/todoist/responses"
	"github.com/stretchr/testify/assert"

	"github.com/kpdowns/todoist-cli/mocks"
)

func TestNotAuthenticated(t *testing.T) {
	mockAuthenticationService := &mocks.MockAuthenticationService{
		AuthenticatedStateToReturn: false,
	}
	mockOutputStream := &bytes.Buffer{}

	listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, nil)
	listTaskCommand.Execute()

	assert.Equal(t, errorNotCurrentlyAuthenticated, mockOutputStream.String())
}

func TestWrittingToOutputStream(t *testing.T) {

	t.Run("When authenticated and there are no tasks, then message is written to output stream", func(t *testing.T) {

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(syncQuery requests.Query) (*responses.Query, error) {
				syncResponse := &responses.Query{
					Items: []responses.Item{},
				}

				return syncResponse, nil
			},
		}
		mockOutputStream := &bytes.Buffer{}

		mockTaskRepository := &mocks.MockTaskRepository{
			CreateAllFunc: func(types.TaskList) (types.TaskList, error) { return nil, nil },
		}

		taskService := services.NewTaskService(mockAPI, mockAuthenticationService, mockTaskRepository)

		listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)
		listTaskCommand.Execute()

		assert.Equal(t, noTasksMessage, mockOutputStream.String())

	})

	t.Run("When authenticated and there are tasks, those tasks are written to output stream", func(t *testing.T) {
		itemReturned := responses.Item{
			TodoistID: 1,
			Priority:  1,
			Content:   "test",
			Due: &responses.Due{
				DateString: "",
			},
		}
		taskToBeWritten := itemReturned.ToTask()

		mockAuthenticationService := &mocks.MockAuthenticationService{
			AuthenticatedStateToReturn: true,
		}
		mockAPI := &mocks.MockAPI{
			ExecuteSyncQueryFunction: func(syncQuery requests.Query) (*responses.Query, error) {
				syncResponse := &responses.Query{
					Items: []responses.Item{
						itemReturned,
					},
				}

				return syncResponse, nil
			},
		}
		mockOutputStream := &bytes.Buffer{}

		mockTaskRepository := &mocks.MockTaskRepository{
			CreateAllFunc: func(types.TaskList) (types.TaskList, error) {
				return types.TaskList{taskToBeWritten}, nil
			},
		}

		taskService := services.NewTaskService(mockAPI, mockAuthenticationService, mockTaskRepository)

		listTaskCommand := NewListTasksCommand(mockOutputStream, mockAuthenticationService, taskService)
		listTaskCommand.Execute()

		assert.Equal(t, taskToBeWritten.AsString()+"\n", mockOutputStream.String())

	})

}
