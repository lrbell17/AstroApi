package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lrbell17/astroapi/impl/api/auth"
	"github.com/lrbell17/astroapi/impl/api/services"
	log "github.com/sirupsen/logrus"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get JWT from header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": services.ErrNoJwt})
			c.Abort()
			return
		}
		jwtStr := strings.TrimPrefix(authHeader, "Bearer ")
		if jwtStr == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": services.ErrInvalidJwt})
			c.Abort()
			return
		}

		// Validate JWT
		token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				log.Error(jwt.ErrSignatureInvalid)
				return nil, jwt.ErrSignatureInvalid
			}
			return auth.GetJwkPublicKey(), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": services.ErrInvalidJwt})
			c.Abort()
			return
		}

		c.Next()
	}
}
