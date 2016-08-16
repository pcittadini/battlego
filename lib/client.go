package lib

import (
	"github.com/go-resty/resty"
)

func ClientRest() (res *resty.Response, err error) {

	//api_key := "xymxr7cgxdrgd8f79mtcejztsa8fwxzf"

	// Sample of using Request.SetQueryString method
	resp, err := resty.SetHostURL("https://eu.api.battle.net").
		R().
		SetQueryString("locale=en_GB&apikey=xymxr7cgxdrgd8f79mtcejztsa8fwxzf").
		Get("/wow/character/magtheridon/nikotine")

	if err != nil {
		panic(err)
	}

	return resp, nil
}
