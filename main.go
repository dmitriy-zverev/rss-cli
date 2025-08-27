package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/dmitriy-zverev/rss-cli/internal/command"
	"github.com/dmitriy-zverev/rss-cli/internal/config"
	"github.com/dmitriy-zverev/rss-cli/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	userConfig, err := config.Read()
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", userConfig.DBUrl)
	if err != nil {
		fmt.Printf("error connecting to database: %v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	userState := command.State{Cfg: &userConfig, Db: dbQueries}
	userCommands := command.Commands{}

	userCommands.Register(command.LOGIN_CMD, command.HandlerLogin)
	userCommands.Register(command.REGISTER_CMD, command.HandlerRegister)
	userCommands.Register(command.RESET_CMD, command.HandlerReset)

	if len(os.Args) < 2 {
		fmt.Println("error: you did not specify user login")
		os.Exit(1)
	}

	userCommand := command.Command{Name: os.Args[1], Args: os.Args[1:]}
	userCommands.Run(&userState, userCommand)
}
