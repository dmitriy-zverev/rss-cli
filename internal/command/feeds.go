package command

import (
	"context"
	"errors"
	"fmt"
)

func HandlerListFeeds(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return errors.New("cannot get users")
	}

	for _, feed := range feeds {
		userName, err := s.Db.GetUserNameFromID(context.Background(), feed.UserID)
		if err != nil {
			return errors.New("cannot get user by ID")
		}

		fmt.Printf("%s (%s) - %s\n", feed.Name, feed.Url, userName)
	}

	return nil
}
