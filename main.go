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
	defer db.Close()

	dbQueries := database.New(db)

	userState := command.State{Cfg: &userConfig, Db: dbQueries}
	userCommands := command.Commands{}
	registerUserCommands(&userCommands)

	if len(os.Args) < 2 {
		fmt.Println("error: you did not specify command")
		os.Exit(1)
	}

	userCommand := command.Command{Name: os.Args[1], Args: os.Args[1:]}
	userCommands.Run(&userState, userCommand)
}

func registerUserCommands(cmds *command.Commands) {
	cmds.Register(command.LOGIN_CMD, command.HandlerLogin)
	cmds.Register(command.REGISTER_CMD, command.HandlerRegister)
	cmds.Register(command.RESET_CMD, command.HandlerReset)
	cmds.Register(command.LIST_USERS_CMD, command.HandlerListUsers)
	cmds.Register(command.AGG_CMD, command.HandlerAggregate)
	cmds.Register(command.ADD_FEED_CMD, command.HandlerAddFeed)
	cmds.Register(command.LIST_FEEDS_CMD, command.HandlerListFeeds)
	cmds.Register(command.FOLLOW_CMD, command.HandlerFollow)
	cmds.Register(command.FOLLOWING_CMD, command.HandlerFollowing)
}
