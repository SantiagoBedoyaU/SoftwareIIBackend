package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// generate jwt auth middleware for gin
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtToken := ctx.GetHeader("authorization")
		if jwtToken == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		// validate jwt token
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(jwtToken, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			log.Println(err)
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		ctx.Set("userDNI", claims["sub"])
		ctx.Set("userRole", claims["role"])
	}
}
