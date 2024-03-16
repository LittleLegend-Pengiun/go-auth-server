package controllers

import (
	"go-auth-server/initializers"
	"go-auth-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	tx := initializers.DB.Find(&users)

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot fetch user from database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users_list": users,
	})

}
