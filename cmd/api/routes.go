package main

import (
	"log"
	"softwareIIbackend/cmd/api/middleware"
	"softwareIIbackend/docs"
	"time"

	helmet "github.com/danielkov/gin-helmet"
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

func (app *Application) setupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(helmet.Default())
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
				user.POST("", app.CreateUserHandler)
				user.GET("", app.GetUsersByRoleHandler)
				user.GET("/:dni", app.GetUserByDNIHandler)
				user.GET("/me", app.GetMyInformationHandler)
				user.PATCH("/me", app.UpdateMyInformationHandler)
				user.POST("/load-by-csv", app.LoadUserByCSVHandler)
				user.POST("/reset-password", app.UpdateUserPasswordHandler)
				user.PATCH("/assign-role", app.UpdateUserRoleHandler)
			}
			appointment := protected.Group("/appointments")
			{
				appointment.GET("", app.GetAppointmentsHandler)
				appointment.POST("", app.CreateAppointmentHandler)
				appointment.GET("/my-history", app.GetAppointmentsHistoryHandler)
				appointment.PATCH("/:id", app.CancelAppointmentHandler)
				appointment.PATCH("/:id/add-procedure", app.AddAppointmentProcedureHandler)
			}
			reports := protected.Group("/reports")
			{
				reports.GET("/attendance-report", app.GenerateAttendanceReportHandler)
				reports.GET("/waiting-time-report", app.GenerateWaitingTimeReportHandler)
				reports.GET("/users-dni-report", app.GenerateUsersDNIReportHandler)
				reports.GET("/most-consulted-doctors", app.GenerateMostConsultedDoctorsReportHandler)
      }
			at := protected.Group("/unavailable-times")
			{
				at.GET("", app.GetUnavailableTimeHandler)
				at.POST("", app.CreateUnavailableTimeHandler)
				at.PATCH("/:id", app.UpdateUnavailableTimeHandler)
				at.DELETE("/:id", app.DeleteUnavailableTimeHandler)
			}
		}

	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
