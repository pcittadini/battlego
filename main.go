package main

import (
	"./lib/guild"
	"encoding/json"
	"log"
	"os"
)

func main() {

	type FileGuilds struct {
		Items []struct {
			Score     int    `json:"score"`
			WorldRank int    `json:"world_rank"`
			AreaRank  int    `json:"area_rank"`
			RealmRank int    `json:"realm_rank"`
			Name      string `json:"name"`
			URL       string `json:"url"`
		}
	}

	log.Println("[ battleGO ]")

	guilds := new(FileGuilds)
	fd, _ := os.Open("./resources/eu_magtheridon_tier18.json")

	dec := json.NewDecoder(fd)

	err := dec.Decode(&guilds.Items)
	if err != nil {
		panic(err)
	}

	for _, g := range guilds.Items {
		log.Println("[indexing] " + g.Name + " , " + g.URL)
		r, err := guild.GetGuildDetails(g.Name)

		if err != nil {
			log.Println("[ ERROR ] " + err.Error())
		}

		guild := new(guild.Guild)
		err = json.Unmarshal([]byte(r), &guild)
		if err != nil {
			log.Println("[ ERROR ] " + err.Error())
		}

		log.Print(guild.Members)
	}

}
