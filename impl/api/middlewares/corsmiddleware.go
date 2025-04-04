package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

// Middleware to allow CORS to handle requests from UI
func CORSMiddleware() gin.HandlerFunc {

	config, _ := conf.GetConfig()
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.Api.CorsAllowedOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			log.Debug("CORS preflight received")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
