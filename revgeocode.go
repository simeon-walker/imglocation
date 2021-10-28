package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type location struct {
	Items []struct {
		Title      string `json:"title"`
		ID         string `json:"id"`
		ResultType string `json:"resultType"`
		Address    struct {
			Label       string `json:"label"`
			CountryCode string `json:"countryCode"`
			CountryName string `json:"countryName"`
			State       string `json:"state"`
			CountyCode  string `json:"countyCode"`
			County      string `json:"county"`
			City        string `json:"city"`
			District    string `json:"district"`
			Street      string `json:"street"`
			PostalCode  string `json:"postalCode"`
		} `json:"address"`
		Position struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
		Distance int `json:"distance"`
		MapView  struct {
			West  float64 `json:"west"`
			South float64 `json:"south"`
			East  float64 `json:"east"`
			North float64 `json:"north"`
		} `json:"mapView"`
	} `json:"items"`
}

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
