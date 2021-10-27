package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/config"
	"github.com/xiam/exif"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

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

var (
	filePath = ""
	coords   Coordinate
)

// parseCoordString parses a comma separated string containing a co-ordinate in degrees, minutes, seconds.
func parseCoordString(val string) float64 {
	parts := strings.Split(val, ",")
	degrees, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	minutes, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	seconds, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)

	return degrees + (minutes / 60) + (seconds / 3600)
}

// parseCoord converts coords to a float based on value and reference (i.e. direction N/S, E/W)
func parseCoord(latVal, latRef, lngVal, lngRef string) *Coordinate {
	lat := parseCoordString(latVal)
	lng := parseCoordString(lngVal)

	if latRef == "S" { // N is "+", S is "-"
		lat *= -1
	}

	if lngRef == "W" { // E is "+", W is "-"
		lng *= -1
	}

	return &Coordinate{lat, lng}
}

// readexif extracts the GPS location from im age EXIF tags.
func main() {
	conf, cfgerr := config.NewConfig("ini", "app.conf")
	if cfgerr != nil {
		log.Fatal(cfgerr)
	}
	apikey, cfgerr := conf.String("heremaps::apikey")
	if cfgerr != nil {
		log.Fatal(cfgerr)
	}

	flag.StringVar(&filePath, "f", "", "Image file.")
	flag.Parse()

	if filePath == "" {
		fmt.Printf("Please use [-f] option to specify an image file.\n")
		os.Exit(1)
	}

	data, err := exif.Read(filePath)
	if err != nil {
		log.Panic(err)
	}

	coords = *parseCoord(
		data.Tags["Latitude"],
		data.Tags["North or South Latitude"],
		data.Tags["Longitude"],
		data.Tags["East or West Longitude"],
	)
	fmt.Printf("Lat: %f, Lon: %f\n", coords.Latitude, coords.Longitude)

	loc := decodeLocation(apikey, coords.Latitude, coords.Longitude)

	for _, item := range loc.Items {
		if item.Address.District != item.Address.City {
			fmt.Println(item.Address.District)
		}
		fmt.Println(item.Address.City)
		fmt.Println(item.Address.County)
		fmt.Println(item.Address.State)
		fmt.Println(item.Address.CountryName)
	}

}
