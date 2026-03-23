package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func GetUser(ctx context.Context, uuid string) (*models.User, error) {

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	table_name := "users_db"

	var user = new(models.User)

	query, args, err := psql.
		Select("uuid", "name", "age", "email", "created_at", "updated_at").
		From(table_name).
		Where(squirrel.Eq{"uuid": uuid}).
		ToSql()

	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return nil, err
	}

	err = database.Pool.QueryRow(ctx, query, args...).
		Scan(
			&user.UUID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)

		return nil, err
	}

	return user, nil
}
