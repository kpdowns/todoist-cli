package main

import (
	"fmt"

	"github.com/kpdowns/todoist-cli/config"
)

func main() {

	configuration, err := config.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configuration)
}
