package service

import (
	"os"
	"softwareIIbackend/internal/core/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) GetAuthToken(dni string, role domain.UserRole) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"sub":  dni,
		"role": role,
	})

	return claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
