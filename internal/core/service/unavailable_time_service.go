package service

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"
)

type UnavailableTimeService struct {
	unavailableTimeRepository port.UnavailableTimeRepository
	userService               port.UserService
}

func NewUnavailableTimeService(unavailableTimeRepository port.UnavailableTimeRepository, userService port.UserService) *UnavailableTimeService {
	return &UnavailableTimeService{
		unavailableTimeRepository: unavailableTimeRepository,
		userService:               userService,
	}
}

func (s *UnavailableTimeService) GetUnavailableTime(ctx context.Context, startDate, endDate time.Time, doctorID string) ([]domain.UnavailableTime, error) {
	return s.unavailableTimeRepository.GetUnavailableTime(ctx, startDate, endDate, doctorID)
}

func (s *UnavailableTimeService) CreateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error {
	user, err := s.userService.GetUser(ctx, at.DoctorID)
	if err != nil {
		return err
	}
	if user.Role != domain.MedicRole {
		return domain.ErrNotAMedicRole
	}
	return s.unavailableTimeRepository.CreateUnavailableTime(ctx, at)
}

func (s *UnavailableTimeService) UpdateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error {
	return s.unavailableTimeRepository.UpdateUnavailableTime(ctx, at)
}

func (s *UnavailableTimeService) DeleteUnavailableTime(ctx context.Context, id string) error {
	return s.unavailableTimeRepository.DeleteUnavailableTime(ctx, id)
}
