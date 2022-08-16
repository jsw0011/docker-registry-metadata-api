package main

import (
	"docker_image_metadata_api/restapi"
	"docker_image_metadata_api/server/configuration"
	"docker_image_metadata_api/server/image_api"
	"docker_image_metadata_api/server/image_tag_api"
	"docker_image_metadata_api/server/log"
	"docker_image_metadata_api/server/metadata_api"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var cfg configuration.AppConfiguration

func init() {
	// parse flags
	var confFile string

	flag.StringVar(&confFile, "c", "", "Configuration file \"./config.yml\"")

	flag.Parse()

	if confFile == "" {
		fmt.Fprintf(os.Stderr, "Missing configuration file\n")
		os.Exit(1)
	}

	// parse configuration
	configuration.ParseConfiguration(&confFile, &cfg)

	// init logger
	if cfg.General.LogToConsole || cfg.General.LogFile == "" {
		log.InitLogger(nil, cfg.General.LogLevel)
	} else {
		log.InitLogger(&cfg.General.LogFile, cfg.General.LogLevel)
	}
	log.Logger.Info("Log library initialized")
}

func main() {
	log.Logger.Info("Preparing REST API...")

	rc := restapi.Config{
		Logger:      log.Logger.Infof,
		ImageAPI:    image_api.GetAPI(&cfg),
		ImageTagAPI: image_tag_api.GetAPI(&cfg),
		MetadataAPI: metadata_api.GetAPI(&cfg),
	}

	h, e := restapi.Handler(rc)

	if e != nil {

		log.Logger.Criticalf("Error creating REST handler: %s ...Exiting\n", e)

		os.Exit(4)

	}

	// note that this runs on all interfaces right now

	log.Logger.Infof("Starting to serve REST API on http://localhost:%d", cfg.General.Port)

	// Run the standard http server
	log.Logger.Criticalf(http.ListenAndServe(":"+strconv.FormatUint(uint64(cfg.General.Port), 10), h).Error())

}
