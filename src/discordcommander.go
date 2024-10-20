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

func main() {
	arguments := os.Args[1:] // Program call arguements, ignoring directory first entry

	taskMaps := map[string]int{
		"-at":   1,
		"-list": 2,
	}

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

	for _, task := range tasks {
		fmt.Println(task)
	}

}

/*func main() {
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
}*/
