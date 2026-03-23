package handlers

import (
	"api_go/internal/repository"
	"net/http"

	"api_go/internal/utils"

	"github.com/gin-gonic/gin"
)

// UpdateUserHandler processes partial updates for a user record identified by UUID.
func UpdateUserHandler(ctx *gin.Context) {
	// 1. Extract the UUID from the URL parameter (e.g., /api/v1/users/:uuid)
	uuid := ctx.Param("uuid")

	if err := utils.ValidateUUID(uuid); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Map the JSON request body to a map[string]any.
	// This captures only the fields sent by the client, allowing for partial updates (PATCH style).
	var updates map[string]any
	if err := ctx.ShouldBindJSON(&updates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload: " + err.Error()})
		return
	}

	// 3. Prevent unnecessary database calls if the update map is empty.
	if len(updates) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No fields provided for update"})
		return
	}

	// 4. Invoke the repository layer to perform the dynamic update.
	// We pass the request context to ensure the query can be canceled if the client disconnects.
	updatedUser, err := repository.UpdateUser(ctx.Request.Context(), uuid, updates)
	if err != nil {
		// If the record was not found or a database error occurred, return an error response.
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found or update failed"})
		return
	}

	// 5. Return the newly updated user object with a 200 OK status.
	ctx.JSON(http.StatusOK, updatedUser)
}
