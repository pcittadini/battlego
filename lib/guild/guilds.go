package guild

import (
	"github.com/pcittadini/battlego/lib/rest"
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
		return "", err
	}

	return res.Result, nil
}

func PushGuildToEs(guildDetails string) (res string, statuscode string, err error) {

	r := new(rest.Endpoint)
	r.Host = "http://localhost:9200"
	r.Path = "/guilds/memberlist/?pretty"
	r.Method = "POST"

	r.Body = guildDetails

	response, err := r.Do()
	if err != nil {
		return "", response.StatusCode, err
	}

	return response.Result, response.StatusCode, nil
}
