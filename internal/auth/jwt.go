package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

func GenerateJWT(userID uint, tenantID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"tenant_id": tenantID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
