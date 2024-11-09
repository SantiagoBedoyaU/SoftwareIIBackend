package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"time"
)

//go:generate mockgen -source=appointment_port.go -destination=mocks/mock_appointment_port.go -typed

type AppoitmentService interface {
	GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error)
	CreateAppointment(ctx context.Context, appointment *domain.Appointment) error
	CancelAppointment(ctx context.Context, id string) error
}

type AppointmentRepository interface {
	GetByDateRange(ctx context.Context, startDate, endDate time.Time, doctorID, patientID string) ([]domain.Appointment, error)
	CreateAppointment(ctx context.Context, appointment *domain.Appointment) error
	CancelAppointment(ctx context.Context, id string) error
}
