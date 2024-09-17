package api

import (
	"errors"
	"net/http"
	"softwareIIbackend/internal/adapter/repository"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authService port.AuthService
	userService port.UserService
}

func NewAuthHandler(authService port.AuthService, userService port.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var req domain.Auth
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := h.userService.GetUser(ctx, req.DNI)
	if err != nil {
		if errors.Is(err, repository.UserNotFoundErr) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "DNI or Password incorrect",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "DNI or Password incorrect",
		})
		return
	}

	accessToken, err := h.authService.GetAuthToken(user.DNI, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
