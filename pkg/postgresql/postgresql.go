package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(ctx context.Context, username, password, host, port, db string) (*pgxpool.Pool, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		username, password, host, port, db,
	)

	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
