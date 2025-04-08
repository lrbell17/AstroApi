package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/api/auth"
	"github.com/lrbell17/astroapi/impl/conf"
)

// Handler for authentication API
type AuthHandler struct {
}

// Constructor for auth handler
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Handler to authenticate dummy user and return cookie with JWT
func (h *AuthHandler) Login(c *gin.Context) {
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

// Handler to check if the user is already logged in
func (h *AuthHandler) Session(c *gin.Context) {
	if user, exists := c.Get("user"); exists {
		c.JSON(200, gin.H{"message": "Authenticated", "user": user})
	} else {
		c.AbortWithStatus(401)
	}
}
