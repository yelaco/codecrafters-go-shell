package commands

import (
	"errors"
	"strings"
)

type Command struct {
	cmd  string
	args []string
}

var ErrCommandNotFound = errors.New("nonexistent: not found")

func ParseCommand(s string) (Command, error) {
	words := strings.Split(s, "")
	if len(words) == 0 {
		return Command{}, nil
	}

	// Check if command exists
	if !Exist(words[0]) {
		return Command{}, ErrCommandNotFound
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
