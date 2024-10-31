package service

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"
)

type AppointmentService struct {
	appointmentRepository port.AppointmentRepository
}

func NewAppointmentService(appointmentRepository port.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		appointmentRepository: appointmentRepository,
	}
}

func (s *AppointmentService) GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error) {
	return s.appointmentRepository.GetByDateRange(ctx, startDate, endDate, doctorID, patientID)
}
