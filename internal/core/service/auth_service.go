package service

import (
	"context"
	"softwareIIbackend/internal/config"
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
		cfg:          cfg,
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
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(20 * time.Minute).Unix(),
		"sub": email,
	})

	token, err := claims.SignedString([]byte(s.cfg.JwtSecret))
	if err != nil {
		return err
	}
	return s.emailService.SendRecoverPasswordEmail(ctx, fullname, email, token)
}

func (s *AuthService) VerifyAccessToken(_ context.Context, accessToken string, claims *jwt.MapClaims) error {
	_, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JwtSecret), nil
	})
	return err
}
