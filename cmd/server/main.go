package main

import (
	"api_go/internal/database"
	"api_go/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. Inicializa a conexão com o Supabase (PGX)
	database.InitDB()

	defer database.Pool.Close()

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		// Agora as URLs serão /api/status e /api/users
		api.GET("/status", handlers.GetApiStatusHandler)
		api.POST("/create_user", handlers.CreateUserHandler)
	}

	port := ":8000"

	if err := router.Run(port); err != nil {
		log.Fatalf("server error %v", err)
	}

}
