package service

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, DNI string) (*domain.User, error) {
	return s.repo.GetUser(ctx, DNI)
}
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) UpdateUserPassword(ctx context.Context, newPassword string) error {
	dni := ctx.Value("userDNI")
	user, err := s.repo.GetUser(ctx, dni.(string))
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	err = s.repo.UpdateUserPassword(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
