package handlers

import (
	"net/http"
	"strings"

	"api_go/internal/repository"
	"api_go/internal/utils"

	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	if err := utils.ValidateUUID(uuid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.DeleteUser(ctx.Request.Context(), uuid)

	if err != nil {

		if strings.Contains(err.Error(), "no result rows in set") {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error failed to delete user"})
		return
	}

	//Status 204 No Content
	// REST pattern for deletion that no return content

	ctx.Status(http.StatusNoContent)
}
