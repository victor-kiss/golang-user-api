package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"
)

func CreateUser(user *models.User) error {

	// create the raw query

	query := `INSERT INTO users_db (name,age,email,password) VALUES ($1,$2,$3,$4) RETURNING id,uuid, created_at, updated_at`

	// execute the query and map the returning values to the user struct
	err := database.Pool.QueryRow(context.Background(), query, user.Name, user.Age, user.Email, user.Password).Scan(&user.ID, &user.UUID, &user.CreatedAt, &user.UpdatedAt)

	return err
}
