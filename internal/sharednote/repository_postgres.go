package sharednote

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == domain.UniqueViolation {
				return 0, domain.ErrAlreadyShared
			}
		}
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

func (r *RepositoryPostgres) GetAllInfo(ctx context.Context, whomID int64) ([]domain.SharedNoteInfo, error) {
	var notes []domain.SharedNoteInfo

	rows, err := r.conn.Query(
		ctx,
		`SELECT s.id, u.login, u.name, n.title, s.accepted
		 FROM shared_notes s
		 LEFT JOIN users u ON u.id=s.whose_id
		 LEFT JOIN notes n ON n.id=s.note_id
		 WHERE s.whom_id=$1`,
		whomID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := domain.SharedNoteInfo{}
		err := rows.Scan(&note.ID, &note.OwnerLogin, &note.OwnerName, &note.Title, &note.Accepted)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *RepositoryPostgres) Accept(ctx context.Context, id, whomID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE shared_notes
		 SET accepted=true
		 WHERE id=$1 AND whom_id=$2
		 RETURNING id`,
		id, whomID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
