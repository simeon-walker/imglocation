package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/v2/core/config"
)

var (
	apikey   string
	imageDir string
)

// readexif extracts the GPS location from im age EXIF tags.
func main() {
	conf, cfgerr := config.NewConfig("ini", "app.conf")
	if cfgerr != nil {
		log.Fatal(cfgerr)
	}
	apikey, cfgerr = conf.String("heremaps::apikey")
	if cfgerr != nil {
		log.Fatal(cfgerr)
	}

	flag.StringVar(&imageDir, "d", "", "Image directory")
	flag.Parse()

	if imageDir == "" {
		fmt.Printf("Please specify an directory.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	processDir(imageDir)
}
