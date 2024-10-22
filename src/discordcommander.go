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
	AUTHENTICATE         = iota
	LIST                 = iota
	ADD                  = iota
	ADD_SUBCOMMAND_GROUP = iota
	ADD_SUBCOMMAND       = iota
	DELETE               = iota
)

var taskMaps = map[string]int{
	"-at":      AUTHENTICATE,
	"-list":    LIST,
	"-add":     ADD,
	"-add-scg": ADD_SUBCOMMAND_GROUP,
	"-add-sc":  ADD_SUBCOMMAND,
	"-del":     DELETE,
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

	// Sorting and ordering engine (SYNCHRONOUS)
	tasks := []task{}
	subtasks := []int{}
	listIdx := -1 // For if list command is present, makes sending to end of slice easier
	for _, argtask := range argtasks {
		switch argtask.taskID {
		case LIST:
			listIdx = len(tasks)
			tasks = append(tasks, ListValidation(argtask, &subtasks))

		//TODO: Bundle together addition operations on the same task
		case ADD:
			tasks = append(tasks, ComboAddSCGVerification(argtask, &subtasks, ADD))

		case ADD_SUBCOMMAND_GROUP:
			tasks = append(tasks, ComboAddSCGVerification(argtask, &subtasks, ADD_SUBCOMMAND_GROUP))

		case ADD_SUBCOMMAND:
			tasks = append(tasks, AddSubcommandVerification(argtask, &subtasks))

		case DELETE:
			tasks = slices.Insert(tasks, 0, DeleteValidation(argtask, &subtasks))
			listIdx++ // To make sure listIdx is still accurate
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
			fmt.Print("Authentication status...		")

		case LIST:
			fmt.Print("List grab status...			")

		case ADD:
			fmt.Print("Command addition status...		")

		case ADD_SUBCOMMAND_GROUP:
			fmt.Print("Command subgroup addition status...	")

		case ADD_SUBCOMMAND:
			fmt.Print("Subcommand addition status...		")

		case DELETE:
			if len(t.args) > 0 {
				fmt.Print("Command \"" + string(t.args[len(t.args)-1]) + "\" deletion status...	")
			} else {
				fmt.Print("Command \"" + "\" deletion status...	")
			}

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
