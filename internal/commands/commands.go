package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	cmd  string
	args []string
	fn   func(...string)
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

	fn, exist := GetHandlerIfExist(words[0])
	if !exist {
		return Command{}, fmt.Errorf("%s: %w", words[0], ErrCommandNotFound)
	}

	return Command{
		cmd:  words[0],
		args: words[1:],
		fn:   fn,
	}, nil
}

func GetHandlerIfExist(cmd string) (func(...string), bool) {
	switch cmd {
	case "exit":
		return Exit, true
	default:
		return nil, false
	}
}

func (c *Command) Execute() {
	if c.fn != nil {
		c.fn(c.args...)
	}
}

func Exit(args ...string) {
	if len(args) == 0 {
		// TODO: expect arguments
		return
	}
	code, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	os.Exit(code)
}
