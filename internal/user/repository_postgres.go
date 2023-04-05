package user

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

func (r *RepositoryPostgres) Create(ctx context.Context, user *domain.User) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO users (login, name, password_hash)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		user.Login, user.Name, user.PasswordHash,
	).Scan(&user.ID); err != nil {
		return 0, domain.ErrUserAlreadyExists
	}

	return user.ID, nil
}

func (r *RepositoryPostgres) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return nil, nil
}
