package main

import (
	client "./lib"
	"encoding/json"
	"log"
)

func main() {

	log.Println("[ battleGO ]")

	r, err := client.ClientRest()
	if err != nil {
		log.Println("[ ERROR ] " + err.Error())
	}

	re := r.Body()
	character := new(client.Character)

	err = json.Unmarshal(re, character)
	if err != nil {
		log.Println("[ ERROR ] " + err.Error())
	}

	log.Println(character.Items.Feet.DisplayInfoID)
}
