package api

import (
	"encoding/csv"
	"errors"
	"net/http"
	"softwareIIbackend/internal/core/domain"
	"softwareIIbackend/internal/core/port"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc port.UserService
}

func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// GetUserByDNI
// @Router			/users/:dni [get]
// @Summary			Get user by DNI
// @Description		Get user by DNI
// @Param			query-id path string true	"User DNI"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) GetUserByDNI(ctx *gin.Context) {
	dni := ctx.Param("dni")
	user, err := h.svc.GetUser(ctx, dni)
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

// CreateUser
// @Router			/users [post]
// @Summary			Create an regular or admin user
// @Description		Create an regular or admin user
// @Param			body body domain.User true	"User Information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.CreateUser(ctx, &user); err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExist) || errors.Is(err, domain.ErrAdminRoleNotAllowed) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// CreateUser
// @Router			/users/load-by-csv [post]
// @Summary			Load user information by CSV
// @Description		Load user information by CSV
// @Param			file formData file true	"CSV file with user information"
// @Param			authorization header string true	"Authorization Token"
// @Accept			multipart/form-data
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) LoadUserByCSV(ctx *gin.Context) {
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
	for i := 1; i < len(records); i++ {
		typeDNI, err := strconv.Atoi(records[i][0])
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid type of DNI, it should be int"})
			return
		}
		role, err := strconv.Atoi(records[i][5])
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

	if err := h.svc.LoadUserByCSV(ctx, users); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

// ResetPassword
// @Router			/users/reset-password [post]
// @Summary			Reset the password of an user by DNI
// @Description		Reset the password of an user by DNI
// @Param			body body domain.UpdatePassword true	"User password"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			401	{object}	interface{}
func (h *UserHandler) ResetPassword(ctx *gin.Context) {
	var req domain.UpdatePassword
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = h.svc.UpdateUserPassword(ctx, req.NewPassword)
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

// GetMyInformation
// @Router			/users/me [get]
// @Summary			Get authenticated user information
// @Description		Get authenticated user information
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) GetMyInformation(ctx *gin.Context) {
	user, err := h.svc.GetUserInformation(ctx)
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

// GetMyInformation
// @Router			/users/me [patch]
// @Summary			Update authenticated user information
// @Description		Update authenticated user information
// @Param			body body domain.UpdateUser true	"User information to update"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) UpdateMyInformation(ctx *gin.Context) {
	var req domain.UpdateUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.svc.UpdateUserInformation(ctx, req.FirstName, req.LastName, req.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// UpdateUserRole
// @Router			/users/assign-role [patch]
// @Summary			Assign user role by an admin
// @Description		Assign user role by an admin
// @Param			body body domain.UpdateRole true	"Role to update"
// @Param			authorization header string true	"Authorization Token"
// @Accept			json
// @Produce			json
// @Success			200	{object}	interface{}
// @Failure			404	{object}	interface{}
func (h *UserHandler) UpdateUserRole(ctx *gin.Context) {
	var req domain.UpdateRole
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.svc.UpdateUserRole(ctx, req.DNI, req.NewRole); err != nil {
		if err == domain.ErrNotAnAdminRole{
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err == domain.ErrUserNotFound{
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