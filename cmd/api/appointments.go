package main

import (
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAppointmentsHistoryHandler
// @Router			/appointments/my-history [get]
// @Summary			Get appointments user history
// @Description		Get appointments user history
// @Tags Appointment
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]domain.Appointment
// @Failure			404	{object}	interface{}
func (app *Application) GetAppointmentsHistoryHandler(ctx *gin.Context) {
	userDNI := ctx.Value("userDNI").(string)
	appointments, err := app.services.appointmentService.GetHistoryByUser(ctx, userDNI)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}

// GetAppointmentsHandler
// @Router			/appointments [get]
// @Summary			Get appointments by date range
// @Description		Get appointments by date range
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
func (app *Application) GetAppointmentsHandler(ctx *gin.Context) {
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

// CreateAppointmentHandler
// @Router			/appointments [post]
// @Summary			Create an appointment
// @Description		Create an appointment
// @Tags Appointment
// @Param			body body domain.Appointment true	"Appointment Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	domain.Appointment
// @Failure			404	{object}	interface{}
func (app *Application) CreateAppointmentHandler(ctx *gin.Context) {
	var appointment domain.Appointment
	if err := ctx.ShouldBindJSON(&appointment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := app.services.appointmentService.CreateAppointment(ctx, &appointment); err != nil {
		errorMap := map[error]int{
			domain.ErrAlreadyHaveAnAppointment: http.StatusBadRequest,
			domain.ErrUserNotFound:             http.StatusNotFound,
			domain.ErrNotAMedicRole:            http.StatusBadRequest,
			domain.ErrNotValidDates:            http.StatusBadRequest,
		}

		if statusCode, exists := errorMap[err]; exists {
			ctx.JSON(statusCode, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, appointment)
}

// CancelAppointmentHandler
// @Router			/appointments/{id} [patch]
// @Summary			Cancel an appointment by an id
// @Description		Cancel an appointment by an id
// @Tags Appointment
// @Param			id path string true	"Appointment id"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}  	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) CancelAppointmentHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := app.services.appointmentService.CancelAppointment(ctx, id); err != nil {
		if err == domain.ErrAppointmentNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err == domain.ErrInvalidIDFormat {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
