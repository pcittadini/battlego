package lib

import (
	"github.com/go-resty/resty"
	"os"
)

func ClientRest() (res *resty.Response, err error) {

	api_key := os.Getenv("BATTLE_API_KEY")

	// Sample of using Request.SetQueryString method
	resp, err := resty.SetHostURL("https://eu.api.battle.net").
		R().
		SetQueryString("locale=en_GB&apikey=" + api_key).
		Get("/wow/character/magtheridon/nikotine")

	if err != nil {
		panic(err)
	}

	return resp, nil
}
