package main

import (
	"fmt"
	"os"
	"slices"
)

type ArgumentTask struct {
	taskID  int
	options []string
}

const (
	AUTHENTICATE = iota
	LIST         = iota
	ADD          = iota
	DELETE       = iota
)

var taskMaps = map[string]int{
	"-at":   AUTHENTICATE,
	"-list": LIST,
	"-add":  ADD,
	"-del":  DELETE,
}

type task struct {
	taskType int
	args     []string
	errMsg   string
}

func main() {
	arguments := os.Args[1:] // Program call arguements, ignoring directory first entry

	// Break given arguements into organizable commands
	argtasks := []ArgumentTask{}
	for _, argument := range arguments {
		taskID, exists := taskMaps[argument]
		if exists {
			argtasks = append(argtasks, ArgumentTask{taskID: taskID})
		} else {
			if len(argtasks) == 0 {
				fmt.Println("Must have a command present first!")
				os.Exit(1)
			}
			argtasks[len(argtasks)-1].options = append(argtasks[len(argtasks)-1].options, argument)
		}
	}

	// Perform async authentication check'

	authErrChan := make(chan string)
	go Authenticate(argtasks, authErrChan)

	// Sorting and ordering engine
	tasks := []task{}
	subtasks := []int{}
	listIdx := -1 // For if list command is present, makes sending to back easier
	for _, argtask := range argtasks {
		switch argtask.taskID {
		case LIST:
			listIdx = len(tasks)
			tasks = append(tasks, ListValidation(argtask, &subtasks))

		case ADD:
			/* Possible Subtasks
			- Get servers present in
			*/

			fmt.Println(argtask)

		case DELETE:
			/* Possible Subtasks
			- Get global commands
			- Get servers present in
			- Get specific server commands
			*/

			tasks = append(tasks, task{DELETE, []string{}, ""}) // Testing hack
		}
	}

	// Sort tasks into appropriate order
	// Place list command last
	if listIdx != -1 && listIdx != len(tasks)-1 {
		tasks = append(tasks, tasks[listIdx])
		tasks = append(tasks[:listIdx], tasks[listIdx+1:]...) // Remove old entry
	}

	// Print out action statuses
	tasks = slices.Insert(tasks, 0, task{AUTHENTICATE, []string{}, <-authErrChan}) // Workaround to get auth to show up
	for _, t := range tasks {
		fmt.Print("\033[1m")

		switch t.taskType {
		case AUTHENTICATE:
			fmt.Print("Authentication status...	")

		case LIST:
			fmt.Print("List grab status...		")

		case DELETE:
			fmt.Print("Command deletion status...	")

		default: // Only for development, to delete later
			fmt.Println("Unknown command, skipping")
			continue
		}

		// Print formatted status
		if t.errMsg == "" {
			fmt.Println("\033[38;5;34mSuccess\033[0m")
		} else {
			fmt.Println("\033[38;5;196mFailure\033[38;5;m\n  ↳ " + t.errMsg + "\033[0m")
		}
	}
}
