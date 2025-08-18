package handlers

import (
	"final-project/models"
	"final-project/repositories"
	"final-project/schemas"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var request schemas.LoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin *models.Admin
    if strings.Contains(request.Login, "@") {
        admin, err = repositories.FindByEmail(request.Login)
    } else {
        admin, err = repositories.FindByPhone(request.Login)
    }

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
    }

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	claims := jwt.MapClaims{
		"id":   admin.ID,
		"nama": admin.Nama,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSceret := "rahasia"

	tokenString, err := token.SignedString([]byte(jwtSceret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
