package port

import (
	"context"
	"time"
	"softwareIIbackend/internal/core/domain"
)

type ReportService interface {
	GenerateAttendanceReport(ctx context.Context, startDate, endDate time.Time) (*domain.AttendanceReport, error)
	GenerateWaitingTimeReport(ctx context.Context, startDate, endDate time.Time) (*domain.WaitingTimeReport, error)
}