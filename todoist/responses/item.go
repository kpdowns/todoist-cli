package responses

import (
	"time"

	"github.com/kpdowns/todoist-cli/actions/tasks/types"
)

// Item is a task on Todoist
type Item struct {
	TodoistID int64  `json:"id"`
	DayOrder  int32  `json:"day_order"`
	Checked   int16  `json:"checked"`
	Content   string `json:"content"`
	Due       *Due   `json:"due"`
	Priority  int16  `json:"priority"`
}

func (i *Item) ToTask() types.Task {
	dateFormat := "2006-01-02"
	dueDate, err := time.Parse(dateFormat, i.Due.DateString)
	if err != nil {
		dueDate = time.Now()
	}

	newTask := types.Task{
		Checked:   i.Checked,
		Content:   i.Content,
		DayOrder:  i.DayOrder,
		DueDate:   dueDate,
		Priority:  i.Priority,
		TodoistID: i.TodoistID,
	}

	return newTask
}
