package tasks

import "fmt"

// Task is an item to do
type Task struct {
	TodoistID int64
	DayOrder  int32
	Checked   int16
	Content   string
	DueDate   string
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

// TaskList is a list of unordered tasks
type TaskList []Task

func (a TaskList) Len() int { return len(a) }

func (a TaskList) Less(i, j int) bool {
	isHigherInDayOrderList := a[i].DayOrder > a[j].DayOrder
	return isHigherInDayOrderList
}

func (a TaskList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
