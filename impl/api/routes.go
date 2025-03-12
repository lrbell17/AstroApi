package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/handlers"
)

func SetupRouter(exoplanetHandler *handlers.ExoplanetHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/exoplanets/:id", exoplanetHandler.GetByID)
	}
	return r
}
