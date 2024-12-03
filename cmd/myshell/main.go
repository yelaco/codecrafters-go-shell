package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Wait for user input
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		line, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command, err := commands.ParseCommand(strings.TrimSpace(line))
		if err != nil {
			fmt.Println(err.Error())
		}
		command.Execute(context.Background())
	}
}
