package characters

import (
	"github.com/pcittadini/battlego/lib/rest"
	"os"
)

var Realm = "magtheridon"

func GetCharacterDetails(member string) (details string, err error) {

	r := new(rest.Endpoint)
	r.Host = "https://eu.api.battle.net"
	r.Path = "/wow/character/" + Realm + "/" + member + "?fields=items&locale=en_GB&apikey=" + os.Getenv("BATTLE_API_KEY")
	r.Method = "GET"

	res, err := r.Do()

	if err != nil {
		return "", err
	}

	return res.Result, nil
}

func PushCharacterToEs(characterDetails string) (res string, status string, err error) {

	r := new(rest.Endpoint)
	r.Host = "http://localhost:9200"
	r.Path = "/characters/items/?pretty"
	r.Method = "POST"

	r.Body = characterDetails

	response, err := r.Do()
	if err != nil {
		return "", response.StatusCode, err
	}

	return response.Result, response.StatusCode, nil
}
