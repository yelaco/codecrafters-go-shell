package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/internal/commands"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	rd := bufio.NewReader(os.Stdin)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			panic(err)
		}

		command, err := commands.ParseCommand(line)
		if err != nil {
			fmt.Println(err.Error())
		}
		command.Execute()

		fmt.Println()
	}
}
