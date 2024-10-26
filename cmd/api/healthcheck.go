package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (app *application) HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
