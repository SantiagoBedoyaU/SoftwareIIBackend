package main

import (
	"log"
	"softwareIIbackend/cmd/api/middleware"
	"softwareIIbackend/docs"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func corsConfig() cors.Config {
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	return config
}

func (app *application) setupRoutes() *gin.Engine {
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalln(err)
	}
	router.Use(cors.New(corsConfig()))

	// routes
	router.GET("/health", app.HealthCheckHandler)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/sign-in", app.SignInHandler)
		v1.POST("/recover-password", app.RecoverPasswordHandler)
		v1.POST("/reset-password", app.ResetPasswordHandler)

		protected := v1.Group("", middleware.AuthMiddleware(app.services.authService))
		{
			user := protected.Group("/users")
			{
				user.POST("/", app.CreateUserHandler)
				user.GET("/:dni", app.GetUserByDNIHandler)
				user.GET("/me", app.GetMyInformationHandler)
				user.PATCH("/me", app.UpdateMyInformationHandler)
				user.POST("/load-by-csv", app.LoadUserByCSVHandler)
				user.POST("/reset-password", app.UpdateUserPasswordHandler)
				user.PATCH("/assign-role", app.UpdateUserRoleHandler)
			}
			appointment := protected.Group("/appointments")
			{
				appointment.GET("/", app.GetAppointmentsHandler)
				appointment.POST("/add-appointment", app.GetAppointmentsHandler)
			}
		}

	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
