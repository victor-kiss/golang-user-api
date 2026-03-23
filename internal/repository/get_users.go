package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (r *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {

	// Table name for UPDATE — must not shadow the database package import.
	table_name := "users_db"

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// 1. make the query

	sql, args, err := psql.Select("id", "uuid", "name", "age", "email", "created_at").
		From(table_name).
		Where(squirrel.Eq{"deleted_at": nil}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("error building query: %v", err)
	}

	// 2. query execution for multiples lines

	rows, err := database.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 3. map the results making the users struct
	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.UUID, &u.Name, &u.Age, &u.Email, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
