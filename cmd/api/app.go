package main

import (
	"softwareIIbackend/internal/adapter/repository/mongodb"
	"softwareIIbackend/internal/adapter/service/mailgun"
	"softwareIIbackend/internal/config"
	"softwareIIbackend/internal/core/port"
	"softwareIIbackend/internal/core/service"

	"github.com/go-co-op/gocron/v2"
)

type services struct {
	emailService           port.EmailService
	userService            port.UserService
	authService            port.AuthService
	appointmentService     port.AppoitmentService
	reportService	         port.ReportService
	unavailableTimeService port.UnavailableTimeService
}

type Application struct {
	config    *config.Config
	services  services
	scheduler gocron.Scheduler
}

func NewApplication(config *config.Config, dbconn *mongodb.MongoDBConnection) *Application {
	// email service with mailgun
	emailService := mailgun.NewEmailService(&config.Notification)
	// user
	userRepo := mongodb.NewUserRepository("users", dbconn)
	userService := service.NewUserService(userRepo, emailService)
	// auth
	authService := service.NewAuthService(&config.Auth, emailService)
	// appointment
	appointmentRepo := mongodb.NewAppointmentRepository("appointments", dbconn)
	appointmentService := service.NewAppointmentService(appointmentRepo, userService, emailService)
	// report
	reportService := service.NewReportService(appointmentRepo, userRepo)
	// unavailable time
	unavailableTimeRepo := mongodb.NewUnavailableTimeRepository("unavailable_time", dbconn)
	unavailableTimeService := service.NewUnavailableTimeService(unavailableTimeRepo, userService)

	// scheduler
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	app := &Application{
		config: config,
		services: services{
			emailService:           emailService,
			userService:            userService,
			authService:            authService,
			appointmentService:     appointmentService,
			unavailableTimeService: unavailableTimeService,
     		reportService: 		    reportService,
		},
		scheduler: scheduler,
	}
	return app
}
