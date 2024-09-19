package port

import "context"

type EmailService interface {
	SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error
}
