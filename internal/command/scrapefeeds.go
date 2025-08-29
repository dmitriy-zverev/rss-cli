package command

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
	"github.com/dmitriy-zverev/rss-cli/internal/rssfeed"
	"github.com/google/uuid"
)

func scrapeFeeds(s *State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	markedFeedParams := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now().UTC(),
		ID:        feed.ID,
	}
	if err := s.Db.MarkFeedFetched(context.Background(), markedFeedParams); err != nil {
		return err
	}

	feedData, err := rssfeed.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range feedData.Channel.Item {
		parsedTime, err := time.Parse(TIME_LAYOUT_FORMAT_PUB_DATE, item.PubDate)
		if err != nil {
			return err
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parsedTime.UTC(),
			FeedID:      feed.ID,
		}
		if _, err := s.Db.CreatePost(context.Background(), postParams); err != nil {
			if strings.Contains(fmt.Sprintf("%v", err), "duplicate key value violates unique constraint") {
				continue
			}
			return err
		}
	}

	return nil
}
