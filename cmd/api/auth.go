package main

import (
	"errors"
	"fmt"
	"net/http"
	"softwareIIbackend/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// SignInHandler
// @Router			/sign-in [post]
// @Summary			Authenticate user by DNI and Password
// @Description		Authenticate user by DNI and Password
// @Tags Auth
// @Param			body body domain.Auth true	"User credentials"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			401	{object}	interface{}
func (app *Application) SignInHandler(ctx *gin.Context) {
	var req domain.Auth
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, err := app.services.userService.GetUser(ctx, req.DNI)
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

	accessToken, err := app.services.authService.GetAuthToken(ctx, user.DNI, user.Role)
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

// RecoverPasswordHandler
// @Router			/recover-password [post]
// @Summary			Recover user password
// @Description		Recover user password
// @Tags Auth
// @Param			body body domain.RecoverPassword true	"Recover Passsword information"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) RecoverPasswordHandler(ctx *gin.Context) {
	var req domain.RecoverPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := app.services.userService.GetUser(ctx, req.DNI)
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
	if err := app.services.authService.RecoverPassword(ctx, fullname, user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password recovery email sent",
	})
}

// ResetPasswordHandler
// @Router			/reset-password [post]
// @Summary			Reset user password with verification token
// @Description		Reset user password with verification token
// @Tags Auth
// @Param			body body domain.ResetPassword true	"Reset User Password"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) ResetPasswordHandler(ctx *gin.Context) {
	var req domain.ResetPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims := jwt.MapClaims{}
	err := app.services.authService.VerifyAccessToken(ctx, req.AccessToken, &claims)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid access token",
		})
		return
	}

	// call user.service to update password
	email := claims["sub"].(string)
	user, err := app.services.userService.GetUserByEmail(ctx, email)
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
	if err := app.services.userService.UpdateUserPassword(ctx, req.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}
