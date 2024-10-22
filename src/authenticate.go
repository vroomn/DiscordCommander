package main

import (
	"DiscordCommander/requests"
	"net/http"
	"os"
	"strings"
)

func Authenticate(tasks []ArgumentTask, errMsg chan string) {
	// Check the task list for auth info
	taskPresence := false
	for _, task := range tasks {
		if task.taskID != AUTHENTICATE {
			continue // Pass over the value
		}

		// The command is of-interest
		if len(task.options) == 2 {
			requests.InitVars(task.options[0], task.options[1])

			data := []byte("APPLICATION_ID=\"" + requests.ApplicationID + "\"\n" + "AUTHENTICATION_TOKEN=\"" + requests.AuthenticationToken + "\"")
			err := os.WriteFile("tokens.env", data, 0777)
			if err != nil {
				errMsg <- "Writing cached token info failed"
			}
		} else {
			errMsg <- "Invalid number of related args"
		}

		taskPresence = true
	}

	// Check the filesystem only if not in tasks
	_, filePresent := os.Stat("tokens.env")
	if !taskPresence && filePresent == nil {
		fileData, err := os.ReadFile("tokens.env")
		if err != nil {
			errMsg <- "Error when attempting to read cached tokens"
		}

		appIDString, authTokenString, _ := strings.Cut(string(fileData), "\n")
		_, appID, _ := strings.Cut(appIDString, "=\"")
		_, authToken, _ := strings.Cut(authTokenString, "=\"")

		requests.InitVars(strings.Replace(appID, "\"", "", -1), strings.Replace(authToken, "\"", "", -1))
	} else if filePresent != nil {
		errMsg <- "Token cache not found, did you \"-at'\"?"
	}

	// Do the final API check if there are tokens
	response := requests.GenericRequest(
		http.MethodGet,
		requests.APPLICATION_ENDPOINT+requests.ApplicationID,
		nil,
	)

	if response.StatusCode != 200 {
		errMsg <- "Invalid response code when authenticating, code: " + response.Status
	}
	defer response.Body.Close()

	errMsg <- ""
}
