package service

import (
	"context"
	"softwareIIbackend/internal/config"
	"softwareIIbackend/internal/core/domain"
	mock_port "softwareIIbackend/internal/core/port/mocks"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAuthToken(t *testing.T) {
	cfg := config.New()
	controller := gomock.NewController(t)
	defer controller.Finish()
	emailService := mock_port.NewMockEmailService(controller)
	authService := NewAuthService(&cfg.Auth, emailService)

	ctx := context.Background()
	token, err := authService.GetAuthToken(ctx, "testing", domain.PatientRole)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestRecoverPassword(t *testing.T) {
	cfg := config.New()
	controller := gomock.NewController(t)
	defer controller.Finish()
	emailService := mock_port.NewMockEmailService(controller)
	authService := NewAuthService(&cfg.Auth, emailService)

	ctx := context.Background()
	emailService.
		EXPECT().
		SendRecoverPasswordEmail(ctx, "Testing Fullname", "testing@email.com", gomock.Any())

	err := authService.RecoverPassword(ctx, "Testing Fullname", "testing@email.com")
	assert.NoError(t, err)
}

func TestVerifyAccessToken(t *testing.T) {
	cfg := config.New()
	controller := gomock.NewController(t)
	defer controller.Finish()
	emailService := mock_port.NewMockEmailService(controller)
	authService := NewAuthService(&cfg.Auth, emailService)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(29 * time.Minute).Unix(),
		"sub": "testing",
	})
	token, err := claims.SignedString([]byte(""))
	assert.NoError(t, err)

	ctx := context.Background()
	cl := jwt.MapClaims{}
	err = authService.VerifyAccessToken(ctx, token, &cl)
	assert.NoError(t, err)
}
