package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/middlewares"
)

func SetupRouter(exoplanetHandler *handlers.ExoplanetHandler, starHandler *handlers.StarHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api").Use(middlewares.JwtAuthMiddleware())
	{
		api.GET("/exoplanets/:id", exoplanetHandler.GetById)
		api.POST("/exoplanets", exoplanetHandler.Post)

		api.GET("/stars/:id", starHandler.GetById)
		api.POST("/stars", starHandler.Post)
	}
	return r
}
