package main

import (
	"fmt"
	"os"
)

type CLICommands struct {
	args         []string
	len          int
	iterationIdx int
}

type ArguementTask struct {
	taskID  int
	options []string
}

const (
	AUTHENTICATE = iota
	LIST         = iota
)

var taskMaps = map[string]int{
	"-at":   AUTHENTICATE,
	"-list": LIST,
}

func main() {
	arguments := os.Args[1:] // Program call arguements, ignoring directory first entry

	// Break given arguements into organizable commands
	tasks := []ArguementTask{}
	for _, argument := range arguments {
		taskID, exists := taskMaps[argument]
		if exists {
			tasks = append(tasks, ArguementTask{taskID: taskID})
		} else {
			if len(tasks) == 0 {
				fmt.Println("Must have a command present first!")
				os.Exit(1)
			}
			tasks[len(tasks)-1].options = append(tasks[len(tasks)-1].options, argument)
		}
	}

	// Perform async authentication check
	authErrChan := make(chan string)
	go Authenticate(tasks, authErrChan)

	//for _, task := range tasks {
	//	fmt.Println(task)
	//}

	// Assume the authentication is going to take the least time

	fmt.Print("\033[1mAuthentication status ... ")
	authErr := <-authErrChan
	if authErr == "" {
		fmt.Println("\033[38;5;34mSuccess\033[0m")
	} else {
		fmt.Println("\033[38;5;196mFailure\033[38;5;m\n  ↳ " + authErr + "\033[0m")
	}
}
