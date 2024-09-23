package api

import (
	"errors"
	"net/http"
	"softwareIIbackend/internal/adapter/repository"
	"softwareIIbackend/internal/core/domain"
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

func (h *UserHandler) ResetPassword(ctx *gin.Context) {
	var req domain.UpdatePassword
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = h.svc.UpdatePassword(ctx, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
