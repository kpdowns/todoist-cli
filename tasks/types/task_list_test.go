package types

import (
	"testing"
	"time"
)

func TestGivenListOfTasksWhenSortingTasksThenListOfTasksIsOrderedByDueDateThenPriority(t *testing.T) {
	tasks := TaskList{
		Task{
			TodoistID: 2,
			DueDate:   getDateDisregardingError("2020-04-10"),
			Priority:  3,
		},
		Task{
			TodoistID: 1,
			DueDate:   getDateDisregardingError("2020-04-10"),
			Priority:  4,
		},
		Task{
			TodoistID: 5,
			DueDate:   getDateDisregardingError("2020-04-12"),
			Priority:  2,
		},
		Task{
			TodoistID: 4,
			DueDate:   getDateDisregardingError("2020-04-12"),
			Priority:  3,
		},
		Task{
			TodoistID: 3,
			DueDate:   getDateDisregardingError("2020-04-11"),
			Priority:  3,
		},
	}

	sortedTasks := tasks.SortByDueDateThenSortByPriority()
	sortedTasks[0].assertTaskHasID(t, 1)
	sortedTasks[1].assertTaskHasID(t, 2)
	sortedTasks[2].assertTaskHasID(t, 3)
	sortedTasks[3].assertTaskHasID(t, 4)
	sortedTasks[4].assertTaskHasID(t, 5)
}

func (task *Task) assertTaskHasID(t *testing.T, id int64) {
	if task.TodoistID != id {
		t.Errorf("Expected task to have id of %d, instead had id of %d", id, task.TodoistID)
	}
}

func getDateDisregardingError(dateString string) time.Time {
	expectedDateFormat := "2006-01-02"
	date, _ := time.Parse(expectedDateFormat, dateString)
	return date
}
