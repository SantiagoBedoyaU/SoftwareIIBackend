package main

import (
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUnavailableTimeHandler
// @Router			/unavailable-times [get]
// @Summary			Get unavailable-times by date range
// @Description		Get unavailable-times by date range
// @Tags UnavailableTimes
// @Param			start_date query string true	"Start Date with format YYYY-MM-DD"
// @Param			end_date query string true	"End Date with format YYYY-MM-DD"
// @Param			doctor_id query string false	"Doctor ID"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]domain.UnavailableTime
// @Failure			404	{object}	interface{}
func (app *Application) GetUnavailableTimeHandler(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	doctorID := ctx.Query("doctor_id")
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
	if doctorID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "doctor_id query param is required",
		})
		return
	}
	ats, err := app.services.unavailableTimeService.GetUnavailableTime(ctx, startTime, endTime, doctorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, ats)
}

// CreateUnavailableTimeHandler
// @Router			/unavailable-times [post]
// @Summary			Create unavailable-time
// @Description		Create unavailable-time
// @Tags UnavailableTimes
// @Param			body body domain.UnavailableTime true	"Unavailable Time Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) CreateUnavailableTimeHandler(ctx *gin.Context) {
	var req domain.UnavailableTime
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := app.services.unavailableTimeService.CreateUnavailableTime(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusCreated)
}

// UpdateUnavailableTimeHandler
// @Router			/unavailable-times/{id} [patch]
// @Summary			Update unavailable-time
// @Description		Update unavailable-time
// @Tags UnavailableTimes
// @Param			id path string true	"Unavailable time ID"
// @Param			body body domain.UnavailableTime true	"Unavailable Time Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) UpdateUnavailableTimeHandler(ctx *gin.Context) {
	atID := ctx.Param("id")
	var req domain.UnavailableTime
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	req.ID = atID
	err := app.services.unavailableTimeService.UpdateUnavailableTime(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// DeleteUnavailableTimeHandler
// @Router			/unavailable-times/{id} [delete]
// @Summary			Delete unavailable-time
// @Description		Delete unavailable-time
// @Tags UnavailableTimes
// @Param			id path string true	"Unavailable time ID"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) DeleteUnavailableTimeHandler(ctx *gin.Context) {
	atID := ctx.Param("id")
	err := app.services.unavailableTimeService.DeleteUnavailableTime(ctx, atID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
