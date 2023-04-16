package note

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

func (r *RepositoryPostgres) Create(ctx context.Context, note domain.Note) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO notes (title, body, created_at, updated_at, user_id, notebook_id)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		note.Title, note.Body, note.CreatedAt, note.UpdatedAt, note.UserID, note.NotebookID,
	).Scan(&note.ID); err != nil {
		return 0, domain.ErrNotebookNotFound
	}

	return note.ID, nil
}
