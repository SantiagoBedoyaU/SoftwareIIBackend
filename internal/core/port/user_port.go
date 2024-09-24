package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
)

type UserService interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, newPassword string) error
}

type UserRepository interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, user *domain.User) error
}
