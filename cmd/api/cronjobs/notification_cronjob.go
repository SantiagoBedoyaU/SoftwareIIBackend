package cronjobs

import (
	"context"
	"fmt"
	"log"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func worker(ctx context.Context, userService port.UserService, emailService port.EmailService, jobs <-chan domain.Appointment) {
	for job := range jobs {
		user, err := userService.GetUser(ctx, job.PatientID)
		if err != nil {
			log.Println("NotificationCronJob: ", err)
		}
		err = emailService.SendAppointmentEmail(ctx, fmt.Sprintf("%s %s", user.FirstName, user.LastName), user.Email, job.StartDate)
		if err != nil {
			log.Println("NotificationCronJob: ", err)
		}
	}
}

func NotificationCronJob(
	ctx context.Context,
	scheduler gocron.Scheduler,
	appointmentsService port.AppoitmentService,
	emailService port.EmailService,
	userService port.UserService,
) error {
	jobs := make(chan domain.Appointment, 10)
	log.Println("NotificationCronJob: Initialized")
	_, err := scheduler.NewJob(
		gocron.DailyJob(1, gocron.NewAtTimes(
			gocron.NewAtTime(00, 00, 00),
		)),
		gocron.NewTask(func() {
			now := time.Now()
			tomorrow := now.Add(24 * time.Hour)
			appointments, err := appointmentsService.GetByDateRange(ctx, now, tomorrow, "", "")
			if err != nil {
				log.Println("NotificationCronJob: ", err)
			}

			for range len(appointments) {
				go worker(ctx, userService, emailService, jobs)
			}

			for _, appointment := range appointments {
				jobs <- appointment
			}
		}),
	)
	return err
}
