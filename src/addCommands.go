package main

import (
	"DiscordCommander/requests"
	"sync"
)

// Same basic function for both, just different type identifier
func ComboAddSCGVerification(argtask ArgumentTask, subtasks *[]int, taskType int) Task {
	argtaskLen := len(argtask.options)
	if argtaskLen < 1 {
		return Task{taskType, []string{}, "Must have a command type!"}
	}

	switch argtask.options[0] {
	case "global":
		if argtaskLen < 4 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return Task{taskType, argtask.options[0:4], ""}

	case "server":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_PRESENCE, requests.GET_SERVER_COMMANDS)
		return Task{taskType, argtask.options[0:5], ""}

	default:
		return Task{taskType, []string{}, "Invalid selection argument"}
	}

	return Task{taskType, []string{}, "Not enough command arguments"}
}

func AddSubcommandVerification(argtask ArgumentTask, subtasks *[]int) Task {
	argtaskLen := len(argtask.options)
	if argtaskLen < 1 {
		return Task{ADD_SUBCOMMAND, []string{}, "Must have a command type!"}
	}

	switch argtask.options[0] {
	case "global":
		if argtaskLen < 4 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return Task{ADD_SUBCOMMAND, argtask.options[0:4], ""}

	case "global-subgroup":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return Task{ADD_SUBCOMMAND, argtask.options[0:5], ""}

	case "server":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_COMMANDS, requests.GET_SERVER_PRESENCE)
		return Task{ADD_SUBCOMMAND, argtask.options[0:5], ""}

	case "server-subgroup":
		if argtaskLen < 6 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_COMMANDS, requests.GET_SERVER_PRESENCE)
		return Task{ADD_SUBCOMMAND, argtask.options[0:6], "Invalid selection argument"}

	default:
		return Task{ADD_SUBCOMMAND, []string{}, "Invalid selection argument"}
	}

	return Task{ADD_SUBCOMMAND, []string{}, "Not enough command arguments"}
}

type AsyncJointAddArr struct {
	mu        sync.Mutex
	jointAdds []JointAdd
}

type JointAdd struct {
	name             string
	primaryAdd       *Task
	subcommandGroups []*Task
	subcommands      []*Task
}

// Add to joint tasks
func (arr *AsyncJointAddArr) TaskAppend(task *Task, wg *sync.WaitGroup) {
	// Get the name of the command
	var name string
	switch task.taskType {
	case ADD:
		name = task.args[len(task.args)-2]

	case ADD_SUBCOMMAND_GROUP:
		name = task.args[len(task.args)-4]

	case ADD_SUBCOMMAND:
		if task.args[0] == "server" || task.args[0] == "server-subgroup" {
			name = task.args[2]
		} else {
			name = task.args[1]
		}
	}

	var sibling *JointAdd = nil
	arr.mu.Lock()
	defer arr.mu.Unlock()
	for i, v := range arr.jointAdds {
		if v.name == name {
			sibling = &arr.jointAdds[i]
		}
	}

	if sibling == nil {
		newAdd := JointAdd{name, nil, []*Task{}, []*Task{}}
		switch task.taskType {
		case ADD:
			newAdd.primaryAdd = task

		case ADD_SUBCOMMAND_GROUP:
			newAdd.subcommandGroups = append(newAdd.subcommandGroups, task)

		case ADD_SUBCOMMAND:
			newAdd.subcommands = append(newAdd.subcommands, task)

		}
		arr.jointAdds = append(arr.jointAdds, newAdd)

	} else {
		switch task.taskType {
		case ADD:
			(*sibling).primaryAdd = task

		case ADD_SUBCOMMAND_GROUP:
			(*sibling).subcommandGroups = append((*sibling).subcommandGroups, task)

		case ADD_SUBCOMMAND:
			(*sibling).subcommands = append((*sibling).subcommands, task)
		}
	}

	wg.Done()
}

func (arr *AsyncJointAddArr) cull(wg *sync.WaitGroup) {

	wg.Done()
}
