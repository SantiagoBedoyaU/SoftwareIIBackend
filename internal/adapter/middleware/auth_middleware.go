package middleware

import (
	"softwareIIbackend/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// generate jwt auth middleware for gin
func AuthMiddleware(authService port.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken := c.GetHeader("authorization")
		if jwtToken == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// validate jwt token
		claims := jwt.MapClaims{}
		if err := authService.VerifyAccessToken(c, jwtToken, &claims); err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		c.Set("userDNI", claims["sub"])
		c.Set("userRole", claims["role"])
	}
}
