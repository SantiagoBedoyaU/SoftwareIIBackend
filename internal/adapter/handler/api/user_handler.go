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
		if errors.Is(err, domain.UserNotFoundErr) {
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
		if errors.Is(err, domain.UserAlreadyExistErr) || errors.Is(err, domain.AdminRoleNotAllowedErr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (h *UserHandler) LoadUserByCSV(ctx *gin.Context) {
	// var data []domain.User
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
