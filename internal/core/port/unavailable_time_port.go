package port

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	"time"
)

//go:generate mockgen -source=unavailable_time_port.go -destination=mocks/mock_unavailable_time_port.go -typed

type UnavailableTimeService interface {
	GetUnavailableTime(ctx context.Context, startDate, endDate time.Time, doctorID string) ([]domain.UnavailableTime, error)
	CreateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error
	UpdateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error
	DeleteUnavailableTime(ctx context.Context, id string) error
}

type UnavailableTimeRepository interface {
	GetUnavailableTime(ctx context.Context, startDate, endDate time.Time, doctorID string) ([]domain.UnavailableTime, error)
	CreateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error
	UpdateUnavailableTime(ctx context.Context, at *domain.UnavailableTime) error
	DeleteUnavailableTime(ctx context.Context, id string) error
}
