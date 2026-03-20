package handlers

import (
	"api_go/internal/models"
	"api_go/repository"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {

	user := new(models.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.CreateUser(user)

	if err != nil {

		if strings.Contains(err.Error(), "unique_violation") || strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"error": "user already cadastred"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
