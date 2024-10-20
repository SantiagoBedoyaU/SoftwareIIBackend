package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GetAuthToken(ctx context.Context, dni string, role domain.UserRole) (string, error)
	RecoverPassword(ctx context.Context, fullname, email string) error
	VerifyAccessToken(ctx context.Context, accessToken string, claims *jwt.MapClaims) error
}
