package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"

	"github.com/Masterminds/squirrel"
)

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	table_name := "users_db"

	query_builder := psql.
		Insert(table_name).
		Columns("name", "age", "email", "password").
		Values(user.Name, user.Age, user.Email, user.Password).
		Suffix("Returning uuid, created_at, updated_at")

	query, args, err := query_builder.ToSql()

	// create the raw query

	// execute the query and map the returning values to the user struct
	err = database.Pool.QueryRow(ctx, query, args...).Scan(&user.ID, &user.UUID, &user.CreatedAt, &user.UpdatedAt)

	return err
}
