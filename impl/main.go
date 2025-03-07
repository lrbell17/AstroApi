package main

import (
	"flag"
	"time"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/db"
	log "github.com/sirupsen/logrus"
)

var (
	config *conf.Config
)

func main() {

	log.Info("Starting Astro API")

	log.Infof("Database host: %v", config.Database.Host)
	log.Infof("Database port: %v", config.Database.Port)
	log.Infof("Database user: %v", config.Database.User)
	log.Infof("Database password: %v", config.Database.Pass)
	log.Infof("Database name: %v", config.Database.Name)

	// keep container alive for debugging
	time.Sleep(10 * time.Minute)
}

func init() {

	// Get config file path from flags
	var configFile string
	flag.StringVar(&configFile, "c", configFile, "config file path")
	flag.Parse()
	if configFile == "" {
		log.Fatalf("Config file path flag '-c' is required")
	}

	// Initialize config
	err := conf.InitConfig("/app/conf/config.yaml")
	if err != nil {
		log.Fatal("Failed to build configuration, exiting.")
	}
	config, _ = conf.GetConfig()

	db.InitDb()
}
