package service

import (
	"context"
	"fmt"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"
)

type AppointmentService struct {
	appointmentRepository port.AppointmentRepository
	UserService           port.UserService
	emailService          port.EmailService
}

func NewAppointmentService(appointmentRepository port.AppointmentRepository, UserService port.UserService, emailService port.EmailService) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
		UserService:           UserService,
		emailService:          emailService,
	}
}

func (s *AppointmentService) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	return s.appointmentRepository.GetByDateRange(ctx, startDate, endDate, doctorID, patientID)
}

func (s *AppointmentService) CreateAppointment(ctx context.Context, appointment *domain.Appointment) error {
	endDate := appointment.DateTime.Add(15 * time.Minute)
	// we can't create two appointments with the same date
	if _, err := s.appointmentRepository.GetByDateRange(ctx, appointment.DateTime, endDate); err == nil {
		return domain.ErrAlreadyHaveAnAppointment
	}
	user := s.UserService.GetUser(ctx, dni)
	_ = s.emailService.SendAppointmentEmail(ctx, fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, password)

	return s.appointmentRepository.CreateAppointment(ctx, appointment)
}
