package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/middlewares"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

const tlsPort = ":8443"

func SetupRouter(authHandler *handlers.AuthHandler, exoplanetHandler *handlers.ExoplanetHandler, starHandler *handlers.StarHandler) {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	auth := r.Group("/api")
	{
		auth.POST("/login", authHandler.Login)
		auth.GET("/session", middlewares.JwtAuthMiddleware(), authHandler.Session)
	}

	api := r.Group("/api").Use(middlewares.JwtAuthMiddleware())
	{
		api.GET("/exoplanets/:id", exoplanetHandler.GetById)
		api.GET("/exoplanets/habitable", exoplanetHandler.GetHabitablePlanets)
		api.POST("/exoplanets", exoplanetHandler.Post)

		api.GET("/stars/:id", starHandler.GetById)
		api.GET("/stars", starHandler.SearchByName)
		api.POST("/stars", starHandler.Post)
	}

	config, _ := conf.GetConfig()
	log.Debugf("Loading SSL cert from %v and SSL key from %v", config.Api.SSLCertPath, config.Api.SSLKeyPath)

	log.Infof("Starting Astro API on port %v", tlsPort)
	if err := r.RunTLS(tlsPort, config.Api.SSLCertPath, config.Api.SSLKeyPath); err != nil {
		log.Fatalf("Unable to start API: %v", err)
	}

}
