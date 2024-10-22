package main

import "DiscordCommander/requests"

// Same basic function for both, just different type identifier
func ComboAddSCGVerification(argtask ArgumentTask, subtasks *[]int, taskType int) task {
	argtaskLen := len(argtask.options)
	if argtaskLen < 1 {
		return task{taskType, []string{}, "Must have a command type!"}
	}

	switch argtask.options[0] {
	case "global":
		if argtaskLen < 4 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return task{taskType, argtask.options[1:4], ""}

	case "server":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_PRESENCE, requests.GET_SERVER_COMMANDS)
		return task{taskType, argtask.options[1:5], ""}

	default:
		return task{taskType, []string{}, "Invalid selection argument"}
	}

	return task{taskType, []string{}, "Not enough command arguments"}
}

func AddSubcommandVerification(argtask ArgumentTask, subtasks *[]int) task {
	argtaskLen := len(argtask.options)
	if argtaskLen < 1 {
		return task{ADD_SUBCOMMAND, []string{}, "Must have a command type!"}
	}

	switch argtask.options[0] {
	case "global":
		if argtaskLen < 4 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return task{ADD_SUBCOMMAND, argtask.options[1:4], ""}

	case "global-subgroup":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return task{ADD_SUBCOMMAND, argtask.options[1:5], ""}

	case "server":
		if argtaskLen < 5 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_COMMANDS, requests.GET_SERVER_PRESENCE)
		return task{ADD_SUBCOMMAND, argtask.options[1:5], ""}

	case "server-subgroup":
		if argtaskLen < 6 {
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_COMMANDS, requests.GET_SERVER_PRESENCE)
		return task{ADD_SUBCOMMAND, argtask.options[1:6], "Invalid selection argument"}

	default:
		return task{ADD_SUBCOMMAND, []string{}, "Invalid selection argument"}
	}

	return task{ADD_SUBCOMMAND, []string{}, "Not enough command arguments"}
}
