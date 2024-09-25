package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"math/big"
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

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	// we can't have another user with the same DNI
	if _, err := s.repo.GetUser(ctx, user.DNI); err == nil {
		return domain.UserAlreadyExistErr
	}

	// we can't have another user with the same email
	if _, err := s.repo.GetUserByEmail(ctx, user.Email); err == nil {
		return domain.UserAlreadyExistErr
	}

	// we can't allow create user with role admin
	if user.Role == domain.AdminRole {
		return domain.AdminRoleNotAllowedErr
	}

	password := s.generatePassword(ctx)
	// TODO: sent an email with the user password
	log.Println("Password", password)

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) LoadUserByCSV(ctx context.Context, users []*domain.User) error {
	errs := make([]error, 0)
	for _, user := range users {
		if err := s.CreateUser(ctx, user); err != nil {
			errs = append(errs, err)
		}
	}

	err := errors.Join(errs...)
	return err
}

func (s *UserService) generatePassword(ctx context.Context) string {
	// Generate a cryptographically secure random number
	n, err := rand.Int(rand.Reader, big.NewInt(36))
	if err != nil {
		// Handle error, e.g. log and return an error
		log.Println(err)
		return ""
	}

	// Convert the number to a base64-encoded string
	password := base64.StdEncoding.EncodeToString([]byte(n.String()))

	// Trim the password to a fixed length (e.g. 12 characters)
	password = password[:12]

	return password
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
