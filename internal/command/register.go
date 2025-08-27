package command

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dmitriy-zverev/rss-cli/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return errors.New("cannot register with empty username")
	}

	userName := cmd.Args[1]
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	if _, err := s.Db.CreateUser(context.Background(), userParams); err != nil {
		return errors.New("user with provided login already exists")
	}

	if err := s.Cfg.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("Nice choice! The username %s was reserved for you\n", cmd.Args[1])

	return nil
}
