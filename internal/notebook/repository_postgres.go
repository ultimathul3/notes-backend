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

func (r *RepositoryPostgres) GetAllByUserID(ctx context.Context, userID int64) ([]domain.Notebook, error) {
	var notebooks []domain.Notebook

	rows, err := r.conn.Query(
		ctx,
		`SELECT id, description
		 FROM notebooks
		 WHERE user_id=$1`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		notebook := domain.Notebook{}
		err := rows.Scan(&notebook.ID, &notebook.Description)
		if err != nil {
			return nil, err
		}
		notebooks = append(notebooks, notebook)
	}

	return notebooks, nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, id, userID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM notebooks
		 WHERE user_id=$1 AND id=$2
		 RETURNING id`,
		userID, id,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) Update(ctx context.Context, notebook domain.Notebook) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE notebooks
		 SET description=$1
		 WHERE user_id=$2 AND id=$3
		 RETURNING id`,
		notebook.Description, notebook.UserID, notebook.ID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
