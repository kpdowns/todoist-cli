package types

import (
	"fmt"
	"time"

	"github.com/fatih/color"
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
		priorityString = color.RedString("Very Urgent")
	case 3:
		priorityString = color.YellowString("Urgent")
	case 2:
		priorityString = color.BlueString("Normal")
	case 1:
		priorityString = color.WhiteString("Low")
	}

	return fmt.Sprintf("[%d] %s\t| %s",
		i.ID,
		priorityString,
		i.Content,
	)
}
