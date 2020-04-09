package main

import (
	"fmt"
	"os"

	"github.com/kpdowns/todoist-cli/actions"
	"github.com/kpdowns/todoist-cli/config"
)

func main() {
	configuration, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	actions.Initialize(configuration)
}
