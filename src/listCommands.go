package main

import (
	"DiscordCommander/requests"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type GlobalCommands []struct {
	ID                       string      `json:"id"`
	ApplicationID            string      `json:"application_id"`
	Version                  string      `json:"version"`
	DefaultMemberPermissions interface{} `json:"default_member_permissions"`
	Type                     int         `json:"type"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	DmPermission             bool        `json:"dm_permission"`
	Contexts                 []int       `json:"contexts"`
	IntegrationTypes         []int       `json:"integration_types"`
	Nsfw                     bool        `json:"nsfw"`
	Handler                  int         `json:"handler,omitempty"`
}

func getGlobalCommands() GlobalCommands {
	response := requests.GenericRequest(
		http.MethodGet,
		requests.APPLICATION_ENDPOINT+requests.ApplicationID+"/commands",
		nil,
	)

	if response.StatusCode != 200 {
		log.Fatalln("Global list request failed!")
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Failed to read response body!")
	}
	response.Body.Close()

	var commands GlobalCommands
	json.Unmarshal(body, &commands)

	return commands
}

func (c CLICommands) ListCommandsHandler() {
	switch c.args[c.iterationIdx+1] {
	case "all":
		fmt.Println("All not yet implemented")

	case "global":
		commands := getGlobalCommands()

		for i := 0; i < len(commands); i++ {
			fmt.Println("Command " + strconv.Itoa(i+1) + ": " + commands[i].Name)
		}

	case "clan":
		if c.len < c.iterationIdx+3 {
			log.Fatalln("Insufficinet data to run list command!")
		}

		targetClan := c.args[c.iterationIdx+2]
		log.Println("Clan commands not yet recognized, got clan: " + targetClan)
	}
}
