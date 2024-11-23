package port

import (
	"context"
	"time"
)

//go:generate mockgen -source=email_port.go -destination=mocks/mock_email_port.go -typed

type EmailService interface {
	SendPasswordEmail(ctx context.Context, fullname, email, password string) error
	SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error
	SendAppointmentEmail(ctx context.Context, fullname, email string, date time.Time) error
}
