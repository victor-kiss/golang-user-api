package handlers

import (
	"api_go/internal/repository"
	"net/http"

	"api_go/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	if err := utils.ValidateUUID(uuid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUser(ctx.Request.Context(), uuid)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
