package api

import (
	"errors"
	"fmt"
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

// SignIn
// @Router			/sign-in [post]
// @Summary			Authenticate user by DNI and Password
// @Description		Authenticate user by DNI and Password
// @Param			body body domain.Auth true	"User credentials"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			401	{object}	interface{}
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
		if errors.Is(err, domain.ErrUserNotFound) {
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

	accessToken, err := h.authService.GetAuthToken(ctx, user.DNI, user.Role)
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

func (h *AuthHandler) RecoverPassword(ctx *gin.Context) {
	var req domain.RecoverPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.userService.GetUser(ctx, req.DNI)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fullname := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	if err := h.authService.RecoverPassword(ctx, fullname, user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password recovery email sent",
	})
}

func (h *AuthHandler) ResetPassword(ctx *gin.Context) {
	var req domain.ResetPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims := jwt.MapClaims{}
	err := h.authService.VerifyAccessToken(ctx, req.AccessToken, &claims)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid access token",
		})
		return
	}

	// call user.service to update password
	email := claims["sub"].(string)
	user, err := h.userService.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Set("userDNI", user.DNI)
	if err := h.userService.UpdateUserPassword(ctx, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
