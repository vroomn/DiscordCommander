package main

import (
	"DiscordCommander/requests"
	"log"
	"net/http"
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
			switch cmdArgs.args[cmdArgs.iterationIdx+1] {
			case "global":
				if cmdArgs.len < cmdArgs.iterationIdx+3 {
					log.Fatalln("Insufficinet data to run \"-del global\" command!")
				}

				// Get the command ID
				var commandID string
				commandPresent := false
				commands := getGlobalCommands()
				for _, command := range commands {
					if command.Name == cmdArgs.args[cmdArgs.iterationIdx+2] {
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
					log.Println("Successfully deleted command " + cmdArgs.args[cmdArgs.iterationIdx+2])
				} else {
					log.Fatalln("Failed to delete command " + cmdArgs.args[cmdArgs.iterationIdx+2])
				}

			default:
				log.Fatalln("Invalid organization tag")
			}
		}
	}
}
