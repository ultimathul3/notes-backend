package user

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

func (r *RepositoryPostgres) Create(ctx context.Context, user domain.User) (int64, error) {
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

func (r *RepositoryPostgres) Get(ctx context.Context, login, passwordHash string) (domain.User, error) {
	var user domain.User

	if err := r.conn.QueryRow(
		ctx,
		`SELECT id, login, name, password_hash FROM users
		 WHERE login=$1 AND password_hash=$2`,
		login, passwordHash,
	).Scan(&user.ID, &user.Login, &user.Name, &user.PasswordHash); err != nil {
		return domain.User{}, domain.ErrInvalidLoginOrPassword
	}

	return user, nil
}

func (r *RepositoryPostgres) GetUserIdByLogin(ctx context.Context, login string) (int64, error) {
	var id int64

	if err := r.conn.QueryRow(
		ctx,
		`SELECT id FROM users
		 WHERE login=$1`,
		login,
	).Scan(&id); err != nil {
		return 0, domain.ErrUserNotFound
	}

	return id, nil
}
