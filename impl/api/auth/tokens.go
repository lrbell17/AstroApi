package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

// Generate JWT for user
func GenerateJWTForUser(userID string, expiry int) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * time.Duration(expiry)).Unix(),
		"admin": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwt, err := token.SignedString(GetPrivateKey())
	if err != nil {
		log.Errorf("Failed to generate JWT for user %v: %v", userID, err)
	}
	return jwt, err
}
