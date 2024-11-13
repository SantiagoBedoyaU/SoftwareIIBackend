package main

import (
	"net/http"
	"softwareIIbackend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

func (app *application) AddAppointmentProcedureHandler(ctx *gin.Context) {
	appointmentID := ctx.Param("id")
	var req domain.Procedure
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	app.services.appointmentService.AddAppointmentProcedure(ctx, appointmentID, req)
}
