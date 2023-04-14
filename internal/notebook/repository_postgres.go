package notebook

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type RepositoryPostgres struct {
	conn *pgxpool.Pool
}

func NewRepositoryPostgres(conn *pgxpool.Pool) *RepositoryPostgres {
	return &RepositoryPostgres{
		conn: conn,
	}
}

func (r *RepositoryPostgres) Create(ctx context.Context, notebook domain.Notebook) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO notebooks (description, user_id)
		 VALUES ($1, $2)
		 RETURNING id`,
		notebook.Description, notebook.UserID,
	).Scan(&notebook.ID); err != nil {
		return 0, err
	}

	return notebook.ID, nil
}
