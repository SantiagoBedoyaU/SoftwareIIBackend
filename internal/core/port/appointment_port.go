package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"time"
)

type AppoitmentService interface {
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]domain.Appointment, error)
}

type AppointmentRepository interface {
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]domain.Appointment, error)
}
