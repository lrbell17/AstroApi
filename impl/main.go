package main

import (
	"flag"

	"github.com/lrbell17/astroapi/impl/api"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/repos"
	"github.com/lrbell17/astroapi/impl/api/services"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/database"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting Exoplanet API")
	exoplanetRepo := repos.NewExoplanetRepo(database.DB)
	exoplanetService := services.NewExoplanetService(exoplanetRepo)
	exoplanetHandler := handlers.NewExoplanetHandler(exoplanetService)

	r := api.SetupRouter(exoplanetHandler)
	r.Run(":8080")

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
	err := conf.InitConfig(configFile)
	if err != nil {
		log.Fatal("Failed to build configuration, exiting.")
	}

	// Initialize DB
	database.Connect()
	err = database.InitDb()
	if err != nil {
		log.Fatalf("DB initialization failed: %v", err)
	}
}
