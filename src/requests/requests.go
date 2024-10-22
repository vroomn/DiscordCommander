package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
)

var ApplicationID string
var AuthenticationToken string

const APPLICATION_ENDPOINT string = "https://discord.com/api/v10/applications/"

func InitVars(appID string, authToken string) {
	ApplicationID = appID
	AuthenticationToken = authToken
}

func GenericRequest(method string, url string, payload io.Reader) *http.Response {
	request, err := http.NewRequest(
		method,
		url,
		payload,
	)
	if err != nil {
		log.Fatal("Failed to create new HTTP request!")
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", AuthenticationToken)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Do command failed!")
	}

	return response
}

// Subtask (request) types
const (
	GET_GLOBAL          = iota
	GET_SERVER_PRESENCE = iota
	GET_SERVER_COMMANDS = iota
)

func AddSubtasks(currentSubtasks *[]int, newSubtasks ...int) {
	for _, newSubtask := range newSubtasks {
		if newSubtask > GET_SERVER_COMMANDS { // NOTE: Could be removed at release time
			fmt.Println("Invalid subtask type!")
			os.Exit(1)
		}

		if !slices.Contains(*currentSubtasks, newSubtask) {
			*currentSubtasks = append(*currentSubtasks, newSubtask)
		}
	}
}
