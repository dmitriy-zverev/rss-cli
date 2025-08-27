package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("cannot unfollow empty feed URL")
	}

	feedURL := cmd.Args[1]

	feedUnfollowParams := database.DeleteFeedFollowsForUserParams{
		Url:    feedURL,
		UserID: user.ID,
	}

	if err := s.Db.DeleteFeedFollowsForUser(context.Background(), feedUnfollowParams); err != nil {
		return err
	}

	fmt.Printf("%s has unfollowed %s\n", user.Name, feedURL)

	return nil
}
