package main

import (
	"log"
	"time"

	"github.com/lrbell17/astroapi/impl/conf"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.Info("Starting Astro API program")

	config, err := conf.InitConfig()
	if err != nil {
		log.Fatal("Failed to build configuration, exiting.")
	}

	logrus.Infof("Database host: %v", config.Database.Host)
	logrus.Infof("Database port: %v", config.Database.Port)
	logrus.Infof("Database user: %v", config.Database.User)
	logrus.Infof("Database password: %v", config.Database.Pass)

	// keep container alive for debugging
	time.Sleep(10 * time.Minute)
}
