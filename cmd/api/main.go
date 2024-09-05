package main

import (
	"log"
	"softwareIIbackend/internal/adapter/handler/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalln(err)
	}

	healthcheckHandler := api.NewHealtcheckHandler()
	router.GET("/health", healthcheckHandler.HealthCheck)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running on port 8080")
}
