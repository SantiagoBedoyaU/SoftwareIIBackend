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
