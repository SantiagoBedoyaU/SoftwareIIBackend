package main

import (
	"net/http"
	"time"
	"errors"
	"softwareIIbackend/internal/core/domain"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHandler
// @Router			/appointments [get]
// @Summary			Get appointments by date range
// @Description		Get appointmentes by date range
// @Tags Appointment
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			doctor_id query string false	"Doctor ID"
// @Param			patient_id query string false	"Patient ID"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]domain.Appointment
// @Failure			404	{object}	interface{}
func (app *application) GetAppointmentsHandler(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	doctorID := ctx.Query("doctor_id")
	patientID := ctx.Query("patient_id")

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

	if patientID == "" && doctorID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "patient_id or doctor_id should be provided",
		})
		return
	}

	appointments, err := app.services.appointmentService.GetByDateRange(ctx, startTime, endTime, doctorID, patientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}

// CreateAppointment
// @Router			/appointments/add-appointment [post]
// @Summary			Create an appointment
// @Description		Create an appointment
// @Tags User
// @Param			body body domain.User true	"Appointment Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	domain.Appointment
// @Failure			404	{object}	interface{}
func (app *application) CreateAppointmentHandler(ctx *gin.Context) {
	var appointment domain.Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := app.services.appointmentService.CreateAppointment(ctx, &appointment); err != nil {
		if errors.Is(err, domain.ErrAlreadyHaveAnAppointment) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, appointment)
}
