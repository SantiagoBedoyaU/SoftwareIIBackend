package api

import (
	"errors"
	"net/http"
	"softwareIIbackend/internal/adapter/repository"
	"softwareIIbackend/internal/core/port"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetUserByDNI(ctx *gin.Context) {
	dni := ctx.Param("dni")
	user, err := h.svc.GetUser(ctx, dni)
	if err != nil {
		if errors.Is(err, repository.UserNotFoundErr) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
