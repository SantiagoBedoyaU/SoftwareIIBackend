package main

import (
	"errors"
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
)

// GenerateAttendanceReportHandler
// @Router			/reports/attendance-report [get]
// @Summary			Generate a report about the amount of patients that not assist to their appointments
// @Description		Generate a report about the amount of patients that not assist to their appointments
// @Tags Report
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) GenerateAttendanceReportHandler(ctx *gin.Context) {
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
	reports, err := app.services.reportService.GenerateAttendanceReport(ctx, startTime, endTime)
	if err != nil {
		if errors.Is(err, domain.ErrNotAnAdminRole) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidDates) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidEndDate) {
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
	ctx.JSON(http.StatusOK, reports)
}

// GenerateWaitingTimeReportHandler
// @Router			/reports/waiting-time-report [get]
// @Summary			Generate a report about the waiting time of the patients for their appointments
// @Description		Generate a report about the waiting time of the patients for their appointments
// @Tags Report
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) GenerateWaitingTimeReportHandler(ctx *gin.Context) {
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
	reports, err := app.services.reportService.GenerateWaitingTimeReport(ctx, startTime, endTime)
	if err != nil {
		if errors.Is(err, domain.ErrNotAnAdminRole) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidDates) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidEndDate) {
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
	ctx.JSON(http.StatusOK, reports)
}

// GenerateUsersDNIReportHandler
// @Router			/reports/users-dni-report [get]
// @Summary			Generate a report about the percentage of users with different type of DNI
// @Description		Generate a report about the percentage of users with different type of DNI
// @Tags Report
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) GenerateUsersDNIReportHandler(ctx *gin.Context) {
	reports, err := app.services.reportService.GenerateUsersDNIReport(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrNotAnAdminRole) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, reports)
}

// GenerateMostConsultedDoctorsReportHandler
// @Router			/reports/most-consulted-doctors [get]
// @Summary			Generate a report with the doctors that have more realized appointments.
// @Description		Generate a report with the doctors that have more realized appointments.
// @Tags Report
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) GenerateMostConsultedDoctorsReportHandler(ctx *gin.Context) {
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
	reports, err := app.services.reportService.GenerateMostConsultedDoctorsReport(ctx, startTime, endTime)
	if err != nil {
		if errors.Is(err, domain.ErrNotAnAdminRole) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidDates) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrNotValidEndDate) {
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
	ctx.JSON(http.StatusOK, reports)
}
