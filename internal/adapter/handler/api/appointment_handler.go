package api

import (
	"net/http"
	"softwareIIbackend/internal/core/port"
	"time"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentService port.AppoitmentService
}

func NewAppointmentHandler(appointmentService port.AppoitmentService) *AppointmentHandler {
	return &AppointmentHandler{appointmentService: appointmentService}
}

func (h *AppointmentHandler) GetAppointments(ctx *gin.Context) {
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
	appointments, err := h.appointmentService.GetByDateRange(ctx, startTime, endTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, appointments)
}
