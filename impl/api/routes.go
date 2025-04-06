package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/auth"
	"github.com/lrbell17/astroapi/impl/api/handlers"
	"github.com/lrbell17/astroapi/impl/api/middlewares"
	"github.com/lrbell17/astroapi/impl/conf"
)

func SetupRouter(exoplanetHandler *handlers.ExoplanetHandler, starHandler *handlers.StarHandler) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	auth := r.Group("/api")
	{
		auth.POST("/login", Login)
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

// Handler to authenticate dummy user and return cookie with JWT
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != "admin" || password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	config, _ := conf.GetConfig()
	token, err := auth.GenerateJWTForUser(username, config.Api.JwtExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT for user"})
	}

	c.SetCookie("jwt", token, config.Api.JwtExpiry, "/", config.Api.JwtDomain, false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})

}
