package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"softwareIIbackend/internal/adapter/repository/mongodb"
	"softwareIIbackend/internal/config"
	"syscall"
	"time"
)

// @title					Software2Backend
// @version					1.0
// @description				API para el backend de Software2
// @schemes         		https http
// @host 					useless-ayn-santiagobedoya-423a6091.koyeb.app
// @BasePath				/api/v1
func main() {
	config := config.New()

	ctx := context.Background()
	dbconn, err := mongodb.NewMongodbConnection(ctx, config.Database)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := dbconn.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	app := NewApplication(config, dbconn)

	srv := http.Server{
		Addr:         config.Server.Addr(),
		Handler:      app.setupRoutes(),
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
