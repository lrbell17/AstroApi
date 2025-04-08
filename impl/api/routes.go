package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/middlewares"
)

func SetupRouter(authHandler *handlers.AuthHandler, exoplanetHandler *handlers.ExoplanetHandler, starHandler *handlers.StarHandler) *gin.Engine {
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
	return r
}
