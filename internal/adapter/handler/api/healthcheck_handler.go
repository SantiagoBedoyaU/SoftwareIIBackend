package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
}

func NewHealtcheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
