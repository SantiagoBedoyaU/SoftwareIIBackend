package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
)

type UserService interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	LoadUserByCSV(ctx context.Context, users []*domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, newPassword string) error
}

type UserRepository interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUserPassword(ctx context.Context, user *domain.User) error
}
