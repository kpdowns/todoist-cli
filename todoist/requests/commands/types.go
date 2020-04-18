package commands

// CommandType is the type of sync command being executed
type CommandType string

const (
	// ItemClose is a command that marks a task as completed, it is a simplified version of item_complete
	ItemClose CommandType = CommandType("item_close")

	// ItemAdd is a command that adds an item based on the arguments provided
	ItemAdd CommandType = CommandType("item_add")
)
