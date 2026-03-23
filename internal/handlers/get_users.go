package handlers

import (
	"api_go/internal/models"
	"api_go/internal/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(c *gin.Context) {

	// c.Request.Context() is important if the client cancel the request

	users, err := repository.GetUsers(c.Request.Context())

	if err != nil {

		// debug
		fmt.Printf("Error fetching users: %v\n", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if the list come null return a empty array maintaining the JSON contract strong for the front-end
	if users == nil {
		users = []models.User{}
	}

	c.JSON(http.StatusOK, users)
}
