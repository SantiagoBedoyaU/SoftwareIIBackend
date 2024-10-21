package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo         port.UserRepository
	emailService port.EmailService
}

func NewUserService(repo port.UserRepository, emailService port.EmailService) *UserService {
	return &UserService{repo: repo, emailService: emailService}
}

func (s *UserService) GetUser(ctx context.Context, DNI string) (*domain.User, error) {
	return s.repo.GetUser(ctx, DNI)
}
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) GetUserInformation(ctx context.Context) (*domain.User, error) {
	dni := ctx.Value("userDNI").(string)
	user, err := s.repo.GetUser(ctx, dni)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	var Authorized float64 = ctx.Value("userRole").(float64)
	// we can't have another user with the same DNI
	if _, err := s.repo.GetUser(ctx, user.DNI); err == nil {
		return domain.ErrUserAlreadyExist
	}

	// we can't have another user with the same email
	if _, err := s.repo.GetUserByEmail(ctx, user.Email); err == nil {
		return domain.ErrUserAlreadyExist
	}

	// we can't allow create user with role admin, unless the creator is an admin
	if user.Role == domain.AdminRole && Authorized != float64(domain.AdminRole) {
		return domain.ErrAdminRoleNotAllowed
	}

	password := strings.ToLower(fmt.Sprintf("%s-%d", strings.Split(user.Email, "@")[0], rand.Intn(1000)))
	_ = s.emailService.SendPasswordEmail(ctx, fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
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

func (s *UserService) UpdateUserInformation(ctx context.Context, firstName, lastName, email string) error {
	dni := ctx.Value("userDNI").(string)
	currentUser, err := s.repo.GetUser(ctx, dni)
	if err != nil {
		return err
	}
	emailUser, err := s.GetUserByEmail(ctx, email)
	if err == nil && currentUser.ID != emailUser.ID {
		return domain.ErrUserEmailAlreadyInUse
	}

	user := domain.User{DNI: dni, FirstName: firstName, LastName: lastName, Email: email}
	if err := s.repo.UpdateUserInformation(ctx, &user); err != nil {
		return err
	}
	return nil
}
