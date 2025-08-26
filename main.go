package main

import (
	"fmt"
	"os"

	"github.com/dmitriy-zverev/rss-cli/internal/command"
	"github.com/dmitriy-zverev/rss-cli/internal/config"
)

func main() {
	userConfig, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		os.Exit(1)
	}

	userState := command.State{Cfg: &userConfig}
	userCommands := command.Commands{}

	userCommands.Register(command.LOGIN_CMD, command.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("error: you did not specify user login")
		os.Exit(1)
	}

	userCommand := command.Command{Name: os.Args[1], Args: os.Args[1:]}
	userCommands.Run(&userState, userCommand)
}
