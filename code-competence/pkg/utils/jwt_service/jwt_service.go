package jwt_service

import (
	"github.com/DewaBiara/Code-Competence/pkg/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateToken(user *entity.User) (string, error)
	GetClaims(c *echo.Context) jwt.MapClaims
}
