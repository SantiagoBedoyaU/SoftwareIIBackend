package middleware

import (
	"net/http"
	"softwareIIbackend/internal/core/port"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// generate jwt auth middleware for gin
func AuthMiddleware(authService port.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("authorization")
		parts := strings.Split(bearerToken, "Bearer ")
		if len(parts) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}
		jwtToken := parts[1]
		if jwtToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		// validate jwt token
		claims := jwt.MapClaims{}
		if err := authService.VerifyAccessToken(ctx, jwtToken, &claims); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		ctx.Set("userDNI", claims["sub"])
		ctx.Set("userRole", claims["role"])
	}
}
