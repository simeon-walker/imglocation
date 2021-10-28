package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xiam/exif"
)

// processDir walks a directory to find images to process
func processDir(dir string) {
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".JPG" {
			tags := processImage(path)
			fmt.Println(path, tags)
			return nil
		}
		return err
	})
	if err != nil {
		log.Fatalf("Unable to process directory: %s\n%s\n", dir, err)
	}
}

// processImage obtains address data for the GPS location
func processImage(path string) []string {
	data, err := exif.Read(path)
	if err != nil {
		log.Panic(err)
	}

	coords := parseCoord(
		data.Tags["Latitude"],
		data.Tags["North or South Latitude"],
		data.Tags["Longitude"],
		data.Tags["East or West Longitude"],
	)
	loc := decodeLocation(apikey, coords.Latitude, coords.Longitude)

	var tags []string
	for _, item := range loc.Items {
		if item.Address.District != item.Address.City {
			tags = append(tags, item.Address.District)
		}
		tags = append(tags, item.Address.City)
		tags = append(tags, item.Address.County)
		tags = append(tags, item.Address.State)
		tags = append(tags, item.Address.CountryName)
	}

	return tags
}
