package main

import "DiscordCommander/requests"

func DeleteValidation(argtask ArgumentTask, subtasks *[]int) Task {
	if len(argtask.options) < 1 {
		return Task{DELETE, []string{}, "Must have a command type!"}
	}

	switch argtask.options[0] {
	case "global":
		if len(argtask.options) < 2 { // Not enough options
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_GLOBAL)
		return Task{DELETE, []string{argtask.options[1]}, ""}

	case "server":
		if len(argtask.options) < 3 { // Not enough options
			break
		}

		requests.AddSubtasks(subtasks, requests.GET_SERVER_PRESENCE, requests.GET_SERVER_COMMANDS)
		return Task{DELETE, []string{argtask.options[1], argtask.options[2]}, ""}

	default:
		return Task{DELETE, []string{argtask.options[len(argtask.options)-1]}, "Invalid selection argument"}
	}

	return Task{DELETE, []string{}, "Not enough command arguments"}
}

/*
func (c CLICommands) DeleteHandler() {
	switch c.args[c.iterationIdx+1] {
	case "global":
		if c.len < c.iterationIdx+3 {
			log.Fatalln("Insufficinet data to run \"-del global\" command!")
		}

		// Get the command ID
		var commandID string
		commandPresent := false
		commands := getGlobalCommands()
		for _, command := range commands {
			if command.Name == c.args[c.iterationIdx+2] {
				commandID = command.ID
				commandPresent = true
			}
		}
		if !commandPresent {
			log.Fatalln("Unable to find command to delete!")
		}

		response := requests.GenericRequest(
			http.MethodDelete,
			requests.APPLICATION_ENDPOINT+requests.ApplicationID+"/commands/"+commandID,
			nil,
		)

		if response.StatusCode == 204 {
			log.Println("Successfully deleted command " + c.args[c.iterationIdx+2])
		} else {
			log.Fatalln("Failed to delete command " + c.args[c.iterationIdx+2])
		}

	case "clan":
		if c.len < c.iterationIdx+4 {
			log.Fatalln("Insufficinet data to run \"-del clan\" command!")
		}

		servers := getServers()
		serverFound := false
		var serverID string
		for _, server := range servers {
			if server.Name == c.args[c.iterationIdx+2] {
				serverID = server.ID
				serverFound = true
			}
		}
		if !serverFound {
			log.Fatalln("Cannot find guild to delete from!")
		}

		serverCmds := getServerCommands(serverID)
		cmdFound := false
		var commandID string
		for _, command := range serverCmds {
			if command.Name == c.args[c.iterationIdx+3] {
				cmdFound = true
				commandID = command.ID
			}
		}
		if !cmdFound {
			log.Fatalln("Cannot find command in the server " + c.args[c.iterationIdx+2])
		}

		response := requests.GenericRequest(
			http.MethodDelete,
			requests.APPLICATION_ENDPOINT+requests.ApplicationID+"/guilds/"+serverID+"/commands/"+commandID,
			nil,
		)

		if response.StatusCode == 204 {
			log.Println("Successfully deleted command " + c.args[c.iterationIdx+3])
		} else {
			log.Fatalln("Failed to delete command " + c.args[c.iterationIdx+3])
		}

	default:
		log.Fatalln("Invalid organization tag")
	}
}
*/
