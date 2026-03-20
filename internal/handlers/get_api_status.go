package handlers

import (
	"api_go/internal/database"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApiStatusHandler(c *gin.Context) {
	// Tenta um ping no banco para garantir que a conexão com o Supabase está ativa
	err := database.Pool.Ping(context.Background())

	dbStatus := "online"
	if err != nil {
		dbStatus = "offline"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "running",
		"version":  "v1",
		"database": dbStatus,
	})
}
