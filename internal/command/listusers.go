package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerListUsers(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return errors.New("cannot get users")
	}

	for _, user := range users {
		if user == s.Cfg.CurrentUserName {
			fmt.Printf("%s (current)\n", user)
			continue
		}
		fmt.Printf("%s\n", user)
	}

	return nil
}
