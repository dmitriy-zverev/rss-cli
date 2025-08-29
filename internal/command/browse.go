package command

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("you did not specify command")
	}

	limitPosts := 2
	if len(cmd.Args) >= 2 {
		limitPosts, _ = strconv.Atoi(cmd.Args[1])
	}

	postsParams := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limitPosts),
	}
	posts, err := s.Db.GetPostsForUser(context.Background(), postsParams)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s: %s (%s)\n\n", post.Name, post.Title, post.Url)
		fmt.Printf("—————————————————-\n\n")
	}

	return nil
}
