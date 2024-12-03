package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Command struct {
	name string
	args []string
	fn   Handler
}

type Handler func(context.Context, ...string)

var ErrCommandNotFound = errors.New("command not found")

func ParseCommand(s string) Command {
	words := strings.Split(s, " ")
	if len(words) == 0 {
		return Command{}
	}

	fn := GetHandler(words[0])

	return Command{
		name: words[0],
		args: words[1:],
		fn:   fn,
	}
}

func Exist(cmd string) bool {
	fn := GetHandler(cmd)
	return fn != nil
}

func GetHandler(cmd string) Handler {
	switch cmd {
	case "exit":
		return Exit
	case "echo":
		return Echo
	case "type":
		return Type
	case "pwd":
		return Pwd
	default:
		return nil
	}
}

func (c Command) Execute(ctx context.Context) {
	if c.fn != nil {
		c.fn(ctx, c.args...)
	} else {
		cmd := exec.CommandContext(ctx, c.name, c.args...)
		output, err := cmd.Output()
		if err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Printf("%s: command not found\n", c.name)
			} else {
				fmt.Println(err)
			}
			return
		}
		fmt.Print(string(output))
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
	path, err := exec.LookPath(args[0])
	if err == nil {
		fmt.Printf("%s is %s\n", args[0], path)
		return
	}

	fmt.Printf("%s: not found\n", args[0])
}

func Pwd(ctx context.Context, args ...string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pwd)
}
