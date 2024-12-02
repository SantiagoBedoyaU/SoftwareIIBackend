package main

import (
	"encoding/csv"
	"errors"
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetUserByDNIHandler
// @Router			/users/{dni} [get]
// @Summary			Get user by DNI
// @Description		Get user by DNI
// @Tags User
// @Param			dni path string true	"User DNI"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}  	domain.User
// @Failure			404	{object}	interface{}
func (app *Application) GetUserByDNIHandler(ctx *gin.Context) {
	dni := ctx.Param("dni")
	user, err := app.services.userService.GetUser(ctx, dni)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// CreateUserHandler
// @Router			/users [post]
// @Summary			Create an regular or admin user
// @Description		Create an regular or admin user
// @Tags User
// @Param			body body domain.User true	"User Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	domain.User
// @Failure			404	{object}	interface{}
func (app *Application) CreateUserHandler(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := app.services.userService.CreateUser(ctx, &user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExist) || errors.Is(err, domain.ErrAdminRoleNotAllowed) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// LoadUserByCSVHandler
// @Router			/users/load-by-csv [post]
// @Summary			Load user information by CSV
// @Description		Load user information by CSV
// @Tags User
// @Param			file formData file true	"CSV file with user information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			multipart/form-data
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) LoadUserByCSVHandler(ctx *gin.Context) {
	multipartFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	csvFile, err := multipartFile.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var users []*domain.User
	for i := 0; i < len(records); i++ {
		typeDNI, err := strconv.Atoi(records[i][0])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of DNI, it should be int"})
			return
		}
		role, err := strconv.Atoi(strings.TrimSpace(records[i][5]))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid role, it should be int"})
			return
		}

		user := domain.User{
			TypeDNI:   domain.UserTypeDNI(typeDNI),
			DNI:       records[i][1],
			FirstName: records[i][2],
			LastName:  records[i][3],
			Email:     records[i][4],
			Role:      domain.UserRole(role),
		}
		users = append(users, &user)
	}

	if err := app.services.userService.LoadUserByCSV(ctx, users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

// UpdateUserPasswordHandler
// @Router			/users/reset-password [post]
// @Summary			Reset the password of an user by DNI
// @Description		Reset the password of an user by DNI
// @Tags User
// @Param			body body domain.UpdatePassword true	"User password"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			401	{object}	interface{}
func (app *Application) UpdateUserPasswordHandler(ctx *gin.Context) {
	var req domain.UpdatePassword
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = app.services.userService.UpdateUserPassword(ctx, req.NewPassword)
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

// GetMyInformationHandler
// @Router			/users/me [get]
// @Summary			Get authenticated user information
// @Description		Get authenticated user information
// @Tags User
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	domain.User
// @Failure			404	{object}	interface{}
func (app *Application) GetMyInformationHandler(ctx *gin.Context) {
	user, err := app.services.userService.GetUserInformation(ctx)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// UpdateMyInformationHandler
// @Router			/users/me [patch]
// @Summary			Update authenticated user information
// @Description		Update authenticated user information
// @Tags User
// @Param			body body domain.UpdateUser true	"User information to update"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) UpdateMyInformationHandler(ctx *gin.Context) {
	var req domain.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := app.services.userService.UpdateUserInformation(ctx, &req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// UpdateUserRoleHandler
// @Router			/users/assign-role [patch]
// @Summary			Assign user role by an admin
// @Description		Assign user role by an admin
// @Tags User
// @Param			body body domain.UpdateRole true	"Role to update"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (app *Application) UpdateUserRoleHandler(ctx *gin.Context) {
	var req domain.UpdateRole
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := app.services.userService.UpdateUserRole(ctx, req.DNI, req.NewRole); err != nil {
		if err == domain.ErrNotAnAdminRole {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err == domain.ErrUserNotFound {
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
	ctx.Status(http.StatusNoContent)
}

// GetUsersByRoleHandler
// @Router			/users/ [get]
// @Summary			Get users by role
// @Description		Get appointments by role
// @Tags User
// @Param			role query domain.UserRole true	"Role ID"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	[]domain.User
// @Failure			404	{object}	interface{}
func (app *Application) GetUsersByRoleHandler(ctx *gin.Context) {
	role := ctx.Query("role")
	if role == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "role query param must be provided"})
		return
	}
	roleInt, err := strconv.Atoi(role)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "role query param must be a number"})
		return
	}
	users, err := app.services.userService.GetUsersByRole(ctx, domain.UserRole(roleInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
