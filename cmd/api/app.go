package main

import (
	"softwareIIbackend/internal/adapter/repository/mongodb"
	"softwareIIbackend/internal/adapter/service/mailgun"
	"softwareIIbackend/internal/config"
	"softwareIIbackend/internal/core/port"
	"softwareIIbackend/internal/core/service"
)

type services struct {
	emailService       port.EmailService
	userService        port.UserService
	authService        port.AuthService
	appointmentService port.AppoitmentService
}

type application struct {
	config   *config.Config
	services services
}

func NewApplication(config *config.Config, dbconn *mongodb.MongoDBConnection) *application {
	// email service with mailgun
	emailService := mailgun.NewEmailService(&config.Notification)
	// user
	userRepo := mongodb.NewUserRepository("users", dbconn)
	userService := service.NewUserService(userRepo, emailService)
	// auth
	authService := service.NewAuthService(&config.Auth, emailService)
	// appointment
	appointmentRepo := mongodb.NewAppointmentRepository("appointments", dbconn)
	appointmentService := service.NewAppointmentService(appointmentRepo)

	app := &application{
		config: config,
		services: services{
			emailService:       emailService,
			userService:        userService,
			authService:        authService,
			appointmentService: appointmentService,
		},
	}
	return app
}
