package session

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type RepositoryPostgres struct {
	conn *pgx.Conn
}

func NewRepositoryPostgres(conn *pgx.Conn) *RepositoryPostgres {
	return &RepositoryPostgres{
		conn: conn,
	}
}

func (r *RepositoryPostgres) Create(ctx context.Context, session *domain.Session) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO sessions (user_id, refresh_token, fingerprint, expires_in)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		session.UserID, session.RefreshToken, session.Fingerprint, session.ExpiresIn,
	).Scan(&session.ID); err != nil {
		return 0, err
	}

	return session.ID, nil
}
