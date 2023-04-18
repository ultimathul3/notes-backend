package sharednote

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

func (r *RepositoryPostgres) Create(ctx context.Context, sharedNote domain.SharedNote) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO shared_notes (whose_id, whom_id, note_id, accepted)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		sharedNote.WhoseID, sharedNote.WhomID, sharedNote.NoteID, sharedNote.Accepted,
	).Scan(&sharedNote.ID); err != nil {
		return 0, err
	}

	return sharedNote.ID, nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, id, whomID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM shared_notes
		 WHERE id=$1 AND whom_id=$2
		 RETURNING id`,
		id, whomID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
