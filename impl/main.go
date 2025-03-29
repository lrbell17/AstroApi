package main

import (
	"flag"
	"fmt"

	"github.com/lrbell17/astroapi/impl/api"
	"github.com/lrbell17/astroapi/impl/api/auth"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/services"
	"github.com/lrbell17/astroapi/impl/cache"
	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/lrbell17/astroapi/impl/persistence/database"
	"github.com/lrbell17/astroapi/impl/persistence/repos"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Get config file path from flags and initialize config
	var configFile string
	flag.StringVar(&configFile, "c", configFile, "config file path")
	flag.Parse()
	if configFile == "" {
		panic("Config file path flag '-c' is required")
	}
	if err := conf.InitConfig(configFile); err != nil {
		panic(fmt.Sprintf("Failed to build configuration: %v", err))
	}
	log.Info("Configuration initilized successfully")

	// Initialize DB
	database.Connect()
	if err := database.InitDb(); err != nil {
		log.Fatalf("DB initialization failed: %v", err)
	}

	// Initialize Redis cache
	cache.Connect()

	// Load JWK
	if err := auth.LoadJwk(); err != nil {
		log.Fatal(err)
	}

	// Start API
	log.Info("Starting Astro API")

	starRepo := repos.NewStarRepo(database.DB)
	starService := services.NewStarService(starRepo)
	starHandler := handlers.NewStarHandler(*starService)

	exoplanetRepo := repos.NewExoplanetRepo(database.DB)
	exoplanetService := services.NewExoplanetService(exoplanetRepo, starRepo)
	exoplanetHandler := handlers.NewExoplanetHandler(exoplanetService)

	r := api.SetupRouter(exoplanetHandler, starHandler)
	r.Run(":8080")

}
