package main

import (
	"net/http"
	"softwareIIbackend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// AddAppointmentProcedureHandler
// @Router			/appointments/{id}/add-procedure [patch]
// @Summary			Add appointment procedure
// @Description		Add appointment procedure
// @Tags Appointment
// @Param			body body domain.AppointmentPatch true		"Procedure Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) AddAppointmentProcedureHandler(ctx *gin.Context) {
	appointmentID := ctx.Param("id")
	var req domain.AppointmentPatch
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := app.services.appointmentService.AddAppointmentProcedure(ctx, appointmentID, req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusAccepted)
}
