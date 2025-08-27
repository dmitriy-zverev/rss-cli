package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("cannot follow empty feed URL")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedURL := cmd.Args[1]
	feedID, err := s.Db.GetFeedIDByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feedID,
	}

	createdFeed, err := s.Db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("%s - %s\n", createdFeed.UserName, createdFeed.FeedName)

	return nil
}
