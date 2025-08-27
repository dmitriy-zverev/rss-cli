package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmitriy-zverev/rss-cli/internal/rssfeed"
)

func HandlerAggregate(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	fmt.Println("Getting RSS Feed...")

	rssFeed, err := rssfeed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errors.New("cannot get rss feed")
	}

	fmt.Println("Got it!")
	fmt.Println(rssFeed)

	return nil
}
