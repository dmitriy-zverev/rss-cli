package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	if err := s.Db.ResetUsers(context.Background()); err != nil {
		return errors.New("cannot reset database")
	}

	fmt.Printf("The database has been reset\n")

	return nil
}
