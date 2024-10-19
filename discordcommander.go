package main

import (
	"log"
	"os"
)

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) <= 0 {
		log.Fatal("Expected input!")
	}

	// Check if the program has access to the keys
	AuthenticationInit(cmdArgs)
}
