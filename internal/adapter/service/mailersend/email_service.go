package mailersend

import (
	"context"
	"fmt"
	"softwareIIbackend/internal/adapter/config"
	"time"

	"github.com/mailersend/mailersend-go"
)

type EmailService struct {
	ms *mailersend.Mailersend
}

func NewEmailService(cfg *config.NotificationConfig) *EmailService {
	ms := mailersend.NewMailersend(cfg.MailerSendAPIToken)
	return &EmailService{
		ms: ms,
	}
}

func (s *EmailService) sendEmail(ctx context.Context, fullname, email, subject, text string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "Salud Y Vida",
		Email: "saludyvida@trial-pxkjn41ykz64z781.mlsender.net",
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

	url := fmt.Sprintf("http://localhost:8000/password-reset?at=%s", token)
	text := fmt.Sprintf("Hey, %s. Your password recovery link is: %s", fullname, url)

	return s.sendEmail(ctx, fullname, email, subject, text)
}
