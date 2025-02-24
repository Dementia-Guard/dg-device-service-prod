package controllers

import (
	"api/dto"
	"api/services"
	"api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
func GetUsers(c *gin.Context) {
	users, err := services.GetUsersService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
		return
	}
	utils.SuccessResponse(c,http.StatusOK,"Server Online","SUCCESS",users)
}

func CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := services.CreateUserService(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}