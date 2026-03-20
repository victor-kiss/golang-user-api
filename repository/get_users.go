package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"
)

func GetUsers() ([]models.User, error) {

	// get all the active users
	query := `SELECT name,age,email,create_at FROM users_db WHERE deleted_at IS NULL`

	// run the query and get the rows
	rows, err := database.Pool.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {

		var user models.User

		// get the rows value and and map in the user struct

		err := rows.Scan(user.Name, user.Age, user.Email, user.CreatedAt)

		if err != nil {
			return nil, err
		}

		// insert in the user's array the user
		users = append(users, user)
	}
	return users, nil
}
