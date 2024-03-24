package util

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ClearTable(ctx context.Context, name string, con *pgxpool.Pool) error {
	_, err := con.Exec(ctx, "TRUNCATE TABLE " + name + ";")
	return err
}
