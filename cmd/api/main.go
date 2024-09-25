package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"softwareIIbackend/internal/adapter/config"
	"softwareIIbackend/internal/adapter/handler/api"
	"softwareIIbackend/internal/adapter/middleware"
	"softwareIIbackend/internal/adapter/repository/mongodb"
	"softwareIIbackend/internal/adapter/service/mailersend"
	"softwareIIbackend/internal/core/service"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.New()

	ctx := context.Background()
	dbconn, err := mongodb.NewMongodbConnection(ctx, config.Database)
	if err != nil {
		log.Fatalln(err)
	}
	defer dbconn.Disconnect(ctx)

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalln(err)
	}

	// email service with mailersend
	emailService := mailersend.NewEmailService(&config.Notification)

	// health
	healthcheckHandler := api.NewHealtcheckHandler()

	// user
	userRepo := mongodb.NewUserRepository("users", dbconn)
	userService := service.NewUserService(userRepo, emailService)
	userHandler := api.NewUserHandler(userService)

	// auth
	authService := service.NewAuthService(&config.Auth, emailService)
	authHandler := api.NewAuthHandler(authService, userService)

	// routes
	router.GET("/health", healthcheckHandler.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/sign-in", authHandler.SignIn)
		v1.POST("/recover-password", authHandler.RecoverPassword)
		v1.POST("/reset-password", authHandler.ResetPassword)

		user := v1.Group("/users", middleware.AuthMiddleware(authService))
		{
			user.GET("/:dni", userHandler.GetUserByDNI)
			user.POST("/", userHandler.CreateUser)
			user.POST("/load-by-csv", userHandler.LoadUserByCSV)
			user.POST("/reset-password", userHandler.ResetPassword)
		}

	}

	srv := http.Server{
		Addr:         config.Server.Addr(),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		log.Printf("Server is running on %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
