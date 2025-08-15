package handlers

import (
	"final-project/models"
	"final-project/repositories"
	"final-project/schemas"
	"final-project/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context){
	var request schemas.ListUserRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	users, err := repositories.ListUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []schemas.ListUserResponse
	for _, user := range users{
		_, status := utils.CekStatusUser(user.Status)

		if request.Search != nil {
			if !strings.Contains(user.Nama, *request.Search) {
				continue
			}
		}

		if request.Status != nil {
			if user.Status != *request.Status {
				continue
			}
		}

		response = append(response, schemas.ListUserResponse{
			ID:     user.ID,
			Nama:   user.Nama,
			Email:  user.Email,
			NoHP:   user.NoHP,
			Status: status,
		})
	}
	c.JSON(200, response)
}

func CreateUser(c *gin.Context){
	var req schemas.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{
		Nama:  req.Nama,
		Email: req.Email,
		NoHP:  req.NoHP,
	}

	err := repositories.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := repositories.GetUserByID(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var request schemas.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := repositories.GetUserByID(uint(idInt))
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	user.Nama = request.Nama
	user.Email = request.Email
	user.NoHP = request.NoHP
	user.Status = *request.Status
	

	err = repositories.UpdateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = repositories.DeleteUser(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}