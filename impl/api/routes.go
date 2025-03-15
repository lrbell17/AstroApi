package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/handlers"
)

func SetupRouter(exoplanetHandler *handlers.ExoplanetHandler, starHandler *handlers.StarHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/exoplanets/:id", exoplanetHandler.GetById)

		api.GET("/stars/:id", starHandler.GetById)
		api.POST("/stars", starHandler.PostStar)
	}
	return r
}
