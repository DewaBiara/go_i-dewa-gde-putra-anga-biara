package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/suryaadi44/go-training-restful/config"
	"github.com/suryaadi44/go-training-restful/models"
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
