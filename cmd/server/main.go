package main

import (
	"api_go/internal/database"
	"api_go/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inicializa a conexão com o banco
	database.InitDB()
	defer database.Pool.Close()

	router := gin.Default()

	// 3. Definição das Rotas
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", handlers.GetApiStatusHandler)

		// Rotas de Usuário agrupadas
		users := v1.Group("/users")
		{
			users.GET("", handlers.CreateUserHandler)     // Listar todos
			users.GET("/:uuid", handlers.GetUsersHandler) // Buscar um
			users.POST("", handlers.CreateUserHandler)    // Criar
			//users.PUT("/:uuid", handlers.)    // Atualizar (Dynamic/Squirrel)
			users.DELETE("/:uuid", handlers.DeleteUser) // Soft Delete
		}
	}

	log.Printf("Server running on http://localhost:8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
