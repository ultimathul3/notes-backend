package note

import (
	"context"
	"fmt"
	"strings"
	"time"

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
		return 0, domain.ErrNoteNotFound
	}

	return note.ID, nil
}

func (r *RepositoryPostgres) GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]domain.Note, error) {
	var notes []domain.Note

	rows, err := r.conn.Query(
		ctx,
		`SELECT id, title, body, created_at, updated_at
		 FROM notes
		 WHERE user_id=$1 AND notebook_id=$2`,
		userID, notebookID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := domain.Note{}
		err := rows.Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *RepositoryPostgres) Update(ctx context.Context, note domain.Note) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE notes
		 SET title=$1, body=$2, updated_at=$3
		 WHERE user_id=$4 AND notebook_id=$5 AND id=$6
		 RETURNING id`,
		note.Title, note.Body, note.UpdatedAt, note.UserID, note.NotebookID, note.ID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, noteID, userID, notebookID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM notes
		 WHERE user_id=$1 AND notebook_id=$2 AND id=$3
		 RETURNING id`,
		userID, notebookID, noteID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) Patch(ctx context.Context, noteID, userID, notebookID int64, input domain.PatchNoteDTO) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argID))
		args = append(args, *input.Body)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	if err := r.conn.QueryRow(
		ctx,
		fmt.Sprintf(
			`UPDATE notes
			 SET %s, updated_at=$%d
			 WHERE user_id=$%d AND notebook_id=$%d AND id=$%d
			 RETURNING id`,
			setQuery, argID, argID+1, argID+2, argID+3,
		),
		append(args, time.Now(), userID, notebookID, noteID)...,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
