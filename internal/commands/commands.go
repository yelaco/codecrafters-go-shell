package commands

import (
	"errors"
	"fmt"
	"strings"
)

type Command struct {
	cmd  string
	args []string
}

var ErrCommandNotFound = errors.New("command not found")

func ParseCommand(s string) (Command, error) {
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return Command{}, nil
	}
	for i := range len(words) {
		words[i] = strings.TrimSpace(words[i])
	}

	// Check if command exists
	if !Exist(words[0]) {
		return Command{}, fmt.Errorf("%s: %w", words[0], ErrCommandNotFound)
	}

	return Command{
		cmd:  words[0],
		args: words[1:],
	}, nil
}

func Exist(cmd string) bool {
	switch cmd {
	default:
		return false
	}
}

func (c *Command) Execute() {
}
