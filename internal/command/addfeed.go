package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 3 {
		return errors.New("cannot add empty feed")
	}

	feedName := cmd.Args[1]
	feedURL := cmd.Args[2]

	feedRow := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}
	if _, err := s.Db.CreateFeed(context.Background(), feedRow); err != nil {
		return errors.New("cannot insert row to database")
	}

	if err := HandlerFollow(s, Command{
		Name: FOLLOW_CMD,
		Args: []string{FOLLOW_CMD, feedURL},
	}, user); err != nil {
		return err
	}

	fmt.Printf("Feed '%s' (%s) was successfully added!\n", feedName, feedURL)

	if _, err := s.Db.GetFeed(context.Background(), feedName); err != nil {
		return errors.New("feed has beed added unsuccessfully")
	}

	return nil
}
