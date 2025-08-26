package command

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("error: cannot login with empty username")
	}

	if err := s.Cfg.SetUser(cmd.Args[1]); err != nil {
		return err
	}

	fmt.Printf("The user %s has been set\n", cmd.Args[1])

	return nil
}
