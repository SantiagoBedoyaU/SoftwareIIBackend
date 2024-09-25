package middleware

import (
	"softwareIIbackend/internal/core/port"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// generate jwt auth middleware for gin
func AuthMiddleware(authService port.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("authorization")
		jwtToken := strings.Split(bearerToken, "Bearer ")[1]
		if jwtToken == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// validate jwt token
		claims := jwt.MapClaims{}
		if err := authService.VerifyAccessToken(ctx, jwtToken, &claims); err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set("userDNI", claims["sub"])
		ctx.Set("userRole", claims["role"])
	}
}
