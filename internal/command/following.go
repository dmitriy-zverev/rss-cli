package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
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
