package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"softwareIIbackend/internal/adapter/config"
	"softwareIIbackend/internal/adapter/handler/api"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalln(err)
	}

	healthcheckHandler := api.NewHealtcheckHandler()
	router.GET("/health", healthcheckHandler.HealthCheck)

	srv := http.Server{
		Addr:         config.Server.Addr(),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
		log.Printf("Server is running on port %v", srv.Addr)
	}()

	<-ctx.Done()
	log.Println("shutting down...")
	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Fatalln(err)
	}
}
