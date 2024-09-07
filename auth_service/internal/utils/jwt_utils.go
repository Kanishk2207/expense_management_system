package utils

import (
	"user-service/internal/configs"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user *models.User) (string, error) {
	config := configs.LoadConfig()

	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"exp":     int(GetCurrentUnixTime()) + config.JWTExpiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}
