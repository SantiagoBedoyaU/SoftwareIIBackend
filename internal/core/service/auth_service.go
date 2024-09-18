package service

import (
	"context"
	"softwareIIbackend/internal/adapter/config"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	cfg          *config.AuthConfig
	emailService port.EmailService
}

func NewAuthService(cfg *config.AuthConfig, emailService port.EmailService) *AuthService {
	return &AuthService{
		emailService: emailService,
	}
}

func (s *AuthService) GetAuthToken(ctx context.Context, dni string, role domain.UserRole) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"sub":  dni,
		"role": role,
	})

	return claims.SignedString([]byte(s.cfg.JwtSecret))
}

func (s *AuthService) RecoverPassword(ctx context.Context, fullname, email string) error {
	return s.emailService.SendRecoverPasswordEmail(ctx, fullname, email)
}
