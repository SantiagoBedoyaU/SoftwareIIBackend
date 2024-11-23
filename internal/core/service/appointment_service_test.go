package service

import (
	"context"
	"softwareIIbackend/internal/core/domain"
	mock_port "softwareIIbackend/internal/core/port/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetByDateRage(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockAppointmentRepository := mock_port.NewMockAppointmentRepository(controller)
	mockUserService := mock_port.NewMockUserService(controller)
	mockEmailService := mock_port.NewMockEmailService(controller)
	appointmentService := NewAppointmentService(mockAppointmentRepository, mockUserService, mockEmailService)

	t.Run("should return appointments by date range", func(t *testing.T) {
		ctx := context.Background()
		startDate := time.Now()
		endDate := startDate.Add(24 * 5)
		mockAppointmentRepository.
			EXPECT().
			GetByDateRange(ctx, startDate, endDate, "", "").
			Return([]domain.Appointment{}, nil)
		results, err := appointmentService.GetByDateRange(ctx, startDate, endDate, "", "")
		assert.NoError(t, err)
		assert.Len(t, results, 0)
	})
}
