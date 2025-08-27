package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("error: cannot login with empty username")
	}

	userName := cmd.Args[1]
	if _, err := s.Db.GetUser(context.Background(), userName); err != nil {
		return errors.New("user with provided login was not found")
	}

	if err := s.Cfg.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("Hello %s! Welcome back\n", userName)

	return nil
}
