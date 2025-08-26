package main

import (
	"fmt"
	"os"

	"github.com/dmitriy-zverev/rss-cli/internal/config"
)

func main() {
	userConfig, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		os.Exit(1)
	}

	userConfig.SetUser(config.ADMIN_USERNAME)

	userConfig, err = config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(userConfig)
}
