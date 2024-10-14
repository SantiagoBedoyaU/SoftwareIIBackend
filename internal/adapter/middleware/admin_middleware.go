package middleware

import (
	"softwareIIbackend/internal/core/port"

	"github.com/gin-gonic/gin"
)

// Validate admin rol of the request
func AdminMiddleware(authService port.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := authService.ValidateAdminRol(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		ctx.Set("adminPermission", true)
	}
}