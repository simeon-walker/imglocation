package main

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

// parseCoord converts coords to a float based on value and reference (i.e. direction N/S, E/W)
func parseCoord(latVal, latRef, lngVal, lngRef string) Coordinate {
	lat := parseCoordString(latVal)
	lng := parseCoordString(lngVal)

	if latRef == "S" { // N is "+", S is "-"
		lat *= -1
	}
	if lngRef == "W" { // E is "+", W is "-"
		lng *= -1
	}
	return Coordinate{lat, lng}
}

// parseCoordString parses a comma separated string containing a co-ordinate in degrees, minutes, seconds.
func parseCoordString(val string) float64 {
	parts := strings.Split(val, ",")
	degrees, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	minutes, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	seconds, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)

	return degrees + (minutes / 60) + (seconds / 3600)
}
