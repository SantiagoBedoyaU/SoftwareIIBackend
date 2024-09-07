package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"softwareIIbackend/internal/adapter/config"
	"softwareIIbackend/internal/adapter/handler/api"
	"softwareIIbackend/internal/adapter/repository/mongodb"
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

	// health
	healthcheckHandler := api.NewHealtcheckHandler()

	// user
	userRepo := mongodb.NewUserRepository("users", dbconn)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	// routes
	router.GET("/health", healthcheckHandler.HealthCheck)
	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("/:dni", userHandler.GetUserByDNI)
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
