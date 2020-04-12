package types

import (
	"fmt"
	"time"
)

// Task is an item to do
type Task struct {
	ID        TaskID
	TodoistID int64
	DayOrder  int32
	Checked   int16
	Content   string
	DueDate   time.Time
	Priority  int
}

// AsString returns a tab delimited string representing the task
func (i *Task) AsString() string {
	priorityString := ""
	switch priority := i.Priority; priority {
	case 4:
		priorityString = "Very Urgent"
	case 3:
		priorityString = "Urgent"
	case 2:
		priorityString = "Normal"
	case 1:
		priorityString = "Low"
	}

	return fmt.Sprintf("%d\t%s\t%s",
		i.ID,
		priorityString,
		i.Content,
	)
}
