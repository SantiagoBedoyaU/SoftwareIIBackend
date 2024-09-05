package main

import (
	"log"
	"softwareIIbackend/internal/adapter/handler/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	healthcheckHandler := api.NewHealtcheckHandler()
	router.GET("/health", healthcheckHandler.HealthCheck)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
	log.Println("Server is running on port 8080")
}
