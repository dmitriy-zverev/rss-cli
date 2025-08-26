package command

import (
	"fmt"
	"os"
)

type Commands struct {
	handlerFunctions map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	if err := c.handlerFunctions[cmd.Name](s, cmd); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if len(c.handlerFunctions) < 1 {
		c.handlerFunctions = make(map[string]func(*State, Command) error)
	}

	c.handlerFunctions[name] = f
}
