package mailgun

import (
	"context"
	"fmt"
	"log"
	"softwareIIbackend/internal/config"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type EmailService struct {
	cfg    *config.NotificationConfig
	client *mailgun.MailgunImpl
}

func NewEmailService(cfg *config.NotificationConfig) *EmailService {
	client := mailgun.NewMailgun(cfg.MailgunDomain, cfg.MailgunAPIKey)
	return &EmailService{
		cfg:    cfg,
		client: client,
	}
}
func (s *EmailService) SendPasswordEmail(ctx context.Context, fullname, email, password string) error {
	sender := "accounts@saludyvida.com"
	subject := "SignIn Password"
	body := fmt.Sprintf(`
		Hey %s!
		Your password to Sign In on Salud y Vida is: %s
	`, fullname, password)
	recipient := email

	message := s.client.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	resp, id, err := s.client.Send(ctx, message)

	if err != nil {
		return err
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}

func (s *EmailService) SendRecoverPasswordEmail(ctx context.Context, fullname, email, token string) error {
	sender := "accounts@saludyvida.com"
	subject := "Password Reset"
	url := fmt.Sprintf("%s/security/password-reset?at=%s", s.cfg.FrontendURL, token)
	body := fmt.Sprintf(`
		Hey %s!
		Your password your password reset URL is here: %s
	`, fullname, url)
	recipient := email

	message := s.client.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	resp, id, err := s.client.Send(ctx, message)

	if err != nil {
		return err
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}

func (s *EmailService) SendAppointmentEmail(ctx context.Context, fullname, email string, date time.Time) error {
	sender := "accounts@saludyvida.com"
	subject := "Appointment assigned"
	body := fmt.Sprintf(`
		Hey %s!
		Your appointment in our medical center was assigned
		The designated date : %s
	`, fullname, date)
	recipient := email

	message := s.client.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	resp, id, err := s.client.Send(ctx, message)

	if err != nil {
		return err
	}

	log.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
