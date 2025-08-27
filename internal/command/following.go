package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%s's feeds:\n\n", user.Name)

	for _, feed := range feeds {
		fmt.Printf("    - %s\n", feed)
	}

	return nil
}
