package guild

import (
	"../rest"
	"os"
)

var Realm = "magtheridon"

func GetGuildDetails(name string) (details string, err error) {

	r := new(rest.Endpoint)
	r.Host = "https://eu.api.battle.net"
	r.Path = "/wow/guild/" + Realm + "/" + name + "?fields=members&locale=fr_FR&apikey=" + os.Getenv("BATTLE_API_KEY")
	r.Method = "GET"

	res, err := r.Do()

	if err != nil {
		return err
	}

	return res.Result, nil
}

func GetCharacterDetails() {
	return
}

func PushGuildToEs(guildDetails string) (err error) {

	r := new(rest.Endpoint)
	r.Host = "http://localhost:9200"
	r.Path = "/guilds/"
	r.Method = "PUT"

	r.Body = guildDetails

	res, err := r.Do()
	if err != nil {
		return err
	}

	return res.Result, nil
}
