package repository

import (
	"api_go/internal/database"
	"api_go/internal/models"
	"context"
	"time"

	"github.com/Masterminds/squirrel"
)

func UpdateUser(ctx context.Context, uuid string, updates map[string]any) (*models.User, error) {

	table_name := "users_db"

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	//  make the query

	// select all values to return the new object

	queryBuilder := psql.Update(table_name).
		SetMap(updates).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"uuid": uuid, "deleted_at": nil}).
		Suffix("RETURNING id, uuid, created_at, updated_at, name, age, email")

	sql, args, err := queryBuilder.ToSql()

	if err != nil {
		return nil, err
	}

	// create the return struct
	updatedUser := &models.User{}

	// execute the query and map the result
	err = database.Pool.QueryRow(ctx, sql, args...).Scan(
		&updatedUser.ID,
		&updatedUser.UUID,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
		&updatedUser.Name,
		&updatedUser.Age,
		&updatedUser.Email,
	)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
