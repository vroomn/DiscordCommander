package requests

import (
	"io"
	"log"
	"net/http"
)

var ApplicationID string
var AuthenticationToken string

const APPLICATION_ENDPOINT string = "https://discord.com/api/v10/applications/"

func InitVars(appID string, authToken string) {
	ApplicationID = appID
	AuthenticationToken = authToken

	//fmt.Println(ApplicationID)
	//fmt.Println(AuthenticationToken)
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
