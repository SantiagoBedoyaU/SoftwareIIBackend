package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
)

type UserService interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
}
