package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func decodeLocation(apikey string, lat, lon float64) location {

	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + apikey + "&at=" + fmt.Sprint(lat) + "," + fmt.Sprint(lon)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data location
	json.Unmarshal(body, &data)
	return data
}
