package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
)

//go:generate mockgen -source=user_port.go -destination=mocks/mock_user_port.go -typed

type UserService interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	GetUsersByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error)
	GetUserInformation(ctx context.Context) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	LoadUserByCSV(ctx context.Context, users []*domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUserPassword(ctx context.Context, newPassword string) error
	UpdateUserInformation(ctx context.Context, user *domain.UpdateUser) error
	UpdateUserRole(ctx context.Context, dni string, role domain.UserRole) error
}

type UserRepository interface {
	GetUser(ctx context.Context, DNI string) (*domain.User, error)
	GetUsersByRole(ctx context.Context, role domain.UserRole) ([]domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUserPassword(ctx context.Context, user *domain.User) error
	UpdateUserInformation(ctx context.Context, user *domain.User) error
	UpdateUserRole(ctx context.Context, updateRole *domain.UpdateRole) error
	GenerateUsersDNIReport(ctx context.Context) (int64, int64, int64, error) 
}
