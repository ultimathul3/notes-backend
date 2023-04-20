package session

import (
	"context"

	"github.com/google/uuid"
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

func (r *RepositoryPostgres) Create(ctx context.Context, session domain.Session) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO sessions (user_id, refresh_token, fingerprint, expires_in)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		session.UserID, session.RefreshToken, session.Fingerprint, session.ExpiresIn,
	).Scan(&session.ID); err != nil {
		return 0, err
	}

	return session.ID, nil
}

func (r *RepositoryPostgres) GetCountByUserID(ctx context.Context, userID int64) int64 {
	var count int64

	if err := r.conn.QueryRow(
		ctx,
		`SELECT COUNT(*) FROM sessions
		 WHERE user_id=$1`,
		userID,
	).Scan(&count); err != nil {
		return 0
	}

	return count
}

func (r *RepositoryPostgres) DeleteAllByUserID(ctx context.Context, userID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM sessions
		 WHERE user_id=$1
		 RETURNING id`,
		userID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) GetByRefreshToken(ctx context.Context, refreshToken uuid.UUID) (domain.Session, error) {
	var session domain.Session

	if err := r.conn.QueryRow(
		ctx,
		`SELECT id, user_id, refresh_token, fingerprint, expires_in
		 FROM sessions
		 WHERE refresh_token=$1`,
		refreshToken,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.RefreshToken,
		&session.Fingerprint,
		&session.ExpiresIn,
	); err != nil {
		return domain.Session{}, domain.ErrInvalidOrExpiredRefreshToken
	}

	return session, nil
}

func (r *RepositoryPostgres) Update(ctx context.Context, input domain.UpdateSessionDTO) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE sessions
		 SET refresh_token=$1, expires_in=$2
		 WHERE id=$3
		 RETURNING id`,
		input.RefreshToken, input.ExpiresIn, input.ID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) DeleteByID(ctx context.Context, id int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM sessions
		 WHERE id=$1
		 RETURNING id`,
		id,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) DeleteByRefreshToken(ctx context.Context, userID int64, refreshToken uuid.UUID) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM sessions
		 WHERE user_id=$1 AND refresh_token=$2
		 RETURNING id`,
		userID, refreshToken,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
