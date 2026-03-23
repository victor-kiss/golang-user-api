package repository

import (
	"api_go/internal/database"
	"context"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
)

func (r *userRepository) DeleteUser(ctx context.Context, uuid string) error {

	table_name := "users_db"

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	query, args, err := psql.Update(table_name).
		Set("deleted_at", time.Now()).
		Where(squirrel.Eq{"uuid": uuid}).
		ToSql()

	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return err
	}

	_, err = database.Pool.Exec(ctx, query, args...)

	return err
}
