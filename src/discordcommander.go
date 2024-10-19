package main

import (
	"log"
	"os"
)

type CLICommands struct {
	args         []string
	len          int
	iterationIdx int
}

func main() {
	cmdArgs := CLICommands{os.Args[1:], len(os.Args[1:]), 0}

	//cmdArgs := os.Args[1:]
	if len(cmdArgs.args) <= 0 {
		log.Fatal("Expected input!")
	}

	// Check if the program has access to the keys
	cmdArgs.AuthenticationInit()

	for cmdArgs.iterationIdx = 0; cmdArgs.iterationIdx < cmdArgs.len; cmdArgs.iterationIdx++ {
		switch cmdArgs.args[cmdArgs.iterationIdx] {
		case "-list":
			cmdArgs.ListCommandsHandler()

		case "-del":
			cmdArgs.DeleteHandler()
		}
	}
}
