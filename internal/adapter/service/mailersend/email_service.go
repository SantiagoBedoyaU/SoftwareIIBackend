package mailersend

import (
	"context"
	"fmt"
	"softwareIIbackend/internal/adapter/config"
	"time"

	"github.com/mailersend/mailersend-go"
)

type EmailService struct {
	cfg *config.NotificationConfig
	ms  *mailersend.Mailersend
}

func NewEmailService(cfg *config.NotificationConfig) *EmailService {
	ms := mailersend.NewMailersend(cfg.MailerSendAPIToken)
	return &EmailService{
		cfg: cfg,
		ms:  ms,
	}
}

func (s *EmailService) sendEmail(ctx context.Context, fullname, email, subject, text string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "Salud Y Vida",
		Email: s.cfg.MailerSendFromEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  fullname,
			Email: email,
		},
	}

	message := s.ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetText(text)

	_, err := s.ms.Email.Send(ctx, message)
	return err
}

func (s *EmailService) SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error {
	subject := "Password Recovery"

	url := fmt.Sprintf("%s/security/password-reset?at=%s", s.cfg.FrontendURL, token)
	text := fmt.Sprintf("Hey, %s. Your password recovery link is: %s", fullname, url)

	return s.sendEmail(ctx, fullname, email, subject, text)
}

func (s *EmailService) SendPasswordEmail(ctx context.Context, fullname, email, password string) error {
	subject := "Password Assigned"
	text := fmt.Sprintf("Hey, %s. Your password to log in in the system is: %s", fullname, password)

	return s.sendEmail(ctx, fullname, email, subject, text)
}
