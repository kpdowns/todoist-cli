package responses

// Item is a task on Todoist
type Item struct {
	TodoistID int64  `json:"id"`
	DayOrder  int32  `json:"day_order"`
	Checked   int16  `json:"checked"`
	Content   string `json:"content"`
	Due       *Due   `json:"due"`
	Priority  int16  `json:"priority"`
}
