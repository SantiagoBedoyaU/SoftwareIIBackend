package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHandler
// @Router			/appointments [get]
// @Summary			Get appointments by date range
// @Description		Get appointmentes by date range
// @Tags Appointment
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]domain.Appointment
// @Failure			404	{object}	interface{}
func (app *application) GetAppointmentsHandler(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	startTime, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid start_date: should be in this format YYYY-MM-DD",
		})
		return
	}
	endTime, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid end_date: should be in this format YYYY-MM-DD",
		})
		return
	}
	appointments, err := app.services.appointmentService.GetByDateRange(ctx, startTime, endTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}
