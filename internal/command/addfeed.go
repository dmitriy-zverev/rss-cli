package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 3 {
		return errors.New("cannot add empty feed")
	}

	userName := s.Cfg.CurrentUserName
	userID, err := s.Db.GetUserID(context.Background(), userName)
	if err != nil {
		return errors.New("user with provided login was not found")
	}

	feedName := cmd.Args[1]
	feedURL := cmd.Args[2]

	feedRow := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    userID,
	}
	if _, err := s.Db.CreateFeed(context.Background(), feedRow); err != nil {
		return errors.New("cannot insert row to database")
	}

	fmt.Printf("Feed %s with address (%s) was successfully added!\n", feedName, feedURL)

	newFeed, err := s.Db.GetFeed(context.Background(), feedName)
	if err != nil {
		return errors.New("feed has beed added unsuccessfully")
	}

	fmt.Println(newFeed)

	return nil
}
