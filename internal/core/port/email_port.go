package port

import "context"

type EmailService interface {
	SendPasswordEmail(ctx context.Context, fullname, email, password string) error
	SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error
}
