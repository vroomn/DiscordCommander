package main

import (
	"DiscordCommander/requests"
	"log"
	"net/http"
	"os"
	"strings"
)

func authVerify() bool {
	response := requests.GenericRequest(
		http.MethodGet,
		strings.Join([]string{requests.APPLICATION_ENDPOINT, requests.ApplicationID}, ""),
		nil,
	)

	if response.StatusCode != 200 {
		return false
	}
	defer response.Body.Close()

	return true
}

func (c CLICommands) AuthenticationInit() {
	_, valid := os.Stat("tokens.env")
	if valid != nil {
		passedAuthentication := false

		// Check that the cli options don't contain it
		for i := 0; i < len(c.args); i++ {
			if c.args[i] == "-at" {
				if len(c.args) < i+3 {
					log.Fatalln("Insufficient input for \"-at\" command!")
				}

				requests.InitVars(c.args[i+1], c.args[i+2])
				if !authVerify() {
					log.Fatalln("Malformed \"-at\" command!")
				} else {
					passedAuthentication = true
					log.Println("Successfuly checked authentication")
				}

				// Spit out the new tokens.env file
				data := []byte("APPLICATION_ID=\"" + c.args[i+1] + "\"\n" + "AUTHENTICATION_TOKEN=\"" + c.args[i+2] + "\"")
				err := os.WriteFile("tokens.env", data, 0777)
				if err != nil {
					log.Fatalln("Failed to write authentication info to token file!")
				}
			}
		}

		if !passedAuthentication {
			log.Fatalln("You haven't added your authentication information (Application ID & Authentication token), use \"-at {appID} {authToken}\"")
		}
	} else {
		tokenData, err := os.ReadFile("tokens.env")
		if err != nil {
			log.Fatalln("Failed to open token file!")
		}

		appIDString, authTokenString, _ := strings.Cut(string(tokenData), "\n")
		_, appID, _ := strings.Cut(appIDString, "=\"")
		_, authToken, _ := strings.Cut(authTokenString, "=\"")

		requests.InitVars(strings.Replace(appID, "\"", "", -1), strings.Replace(authToken, "\"", "", -1))
	}

	if authVerify() {
		log.Println("Successfuly checked authentication")
	}
}
