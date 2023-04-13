package utils

import (
	"time"

	"1_JWT_Auth/config"
	"1_JWT_Auth/models"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWT_SECRET))
}
