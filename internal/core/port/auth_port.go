package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
)

type AuthService interface {
	GetAuthToken(ctx context.Context, dni string, role domain.UserRole) (string, error)
	RecoverPassword(ctx context.Context, fullname, email string) error
}
