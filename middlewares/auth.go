package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")

		if tokenString == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if len(tokenString) < 7 || strings.ToUpper(tokenString[:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		jwtSecret := []byte("rahasia")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}