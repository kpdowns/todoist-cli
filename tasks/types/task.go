package types

import (
	"fmt"
	"time"
)

// Task is an item to do
type Task struct {
	TodoistID int64
	DayOrder  int32
	Checked   int16
	Content   string
	DueDate   time.Time
	Priority  int16
}

// AsString returns a tab delimited string representing the task
func (i *Task) AsString() string {
	priorityString := ""
	switch priority := i.Priority; priority {
	case 4:
		priorityString = "High"
	case 3:
		priorityString = "Medium"
	case 2:
		priorityString = "Normal"
	case 1:
		priorityString = "Low"
	}

	return fmt.Sprintf("%d\t%s\t%s",
		i.TodoistID,
		priorityString,
		i.Content,
	)
}
