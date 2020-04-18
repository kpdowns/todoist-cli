package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, sortedTasks[0].TodoistID, int64(1))
	assert.Equal(t, sortedTasks[1].TodoistID, int64(2))
	assert.Equal(t, sortedTasks[2].TodoistID, int64(3))
	assert.Equal(t, sortedTasks[3].TodoistID, int64(4))
	assert.Equal(t, sortedTasks[4].TodoistID, int64(5))
}

func getDateDisregardingError(dateString string) time.Time {
	expectedDateFormat := "2006-01-02"
	date, _ := time.Parse(expectedDateFormat, dateString)
	return date
}
