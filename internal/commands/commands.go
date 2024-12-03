package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Command struct {
	cmd  string
	args []string
	fn   Handler
}

type Handler func(context.Context, ...string)

var ErrCommandNotFound = errors.New("command not found")

func ParseCommand(s string) (Command, error) {
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return Command{}, nil
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

func Exist(cmd string) bool {
	_, exist := GetHandlerIfExist(cmd)
	return exist
}

func GetHandlerIfExist(cmd string) (Handler, bool) {
	switch cmd {
	case "exit":
		return Exit, true
	case "echo":
		return Echo, true
	case "type":
		return Type, true
	default:
		return nil, false
	}
}

func (c *Command) Execute(ctx context.Context) {
	if c.fn != nil {
		c.fn(ctx, c.args...)
	}
}

func Exit(ctx context.Context, args ...string) {
	if len(args) == 0 {
		os.Exit(0)
	}
	code, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	os.Exit(code)
}

func Echo(ctx context.Context, args ...string) {
	fmt.Println(strings.Join(args, " "))
}

func Type(ctx context.Context, args ...string) {
	if len(args) == 0 {
		return
	}

	// check builtin
	if Exist(args[0]) {
		fmt.Printf("%s is a shell builtin\n", args[0])
		return
	}

	// check in $PATH
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, path := range paths {
		fullPath := filepath.Join(path, args[0])
		_, err := os.Stat(fullPath)
		if err == nil {
			fmt.Printf("%s is %s\n", args[0], fullPath)
			return
		}
	}

	fmt.Printf("%s: not found\n", args[0])
}
