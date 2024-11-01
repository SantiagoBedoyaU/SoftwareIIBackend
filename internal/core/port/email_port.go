package port

import (
	"context"
	"time"
)

type EmailService interface {
	SendPasswordEmail(ctx context.Context, fullname, email, password string) error
	SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error
	SendAppointmentEmail(ctx context.Context, fullname, email string, date time.Time) error
}
