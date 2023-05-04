package todolist

import (
	"context"
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

func (r *RepositoryPostgres) Create(ctx context.Context, list domain.TodoList) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO todo_lists (title, created_at, updated_at, user_id, notebook_id)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id`,
		list.Title, list.CreatedAt, list.UpdatedAt, list.UserID, list.NotebookID,
	).Scan(&list.ID); err != nil {
		return 0, domain.ErrNotebookNotFound
	}

	return list.ID, nil
}

func (r *RepositoryPostgres) GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]domain.TodoList, error) {
	var lists []domain.TodoList

	rows, err := r.conn.Query(
		ctx,
		`SELECT id, title, created_at, updated_at
		 FROM todo_lists
		 WHERE user_id=$1 AND notebook_id=$2`,
		userID, notebookID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		list := domain.TodoList{}
		err := rows.Scan(&list.ID, &list.Title, &list.CreatedAt, &list.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *RepositoryPostgres) Update(ctx context.Context, todoListID, userID, notebookID int64, input domain.UpdateTodoListDTO) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE todo_lists
		 SET title=$1, updated_at=$2
		 WHERE user_id=$3 AND notebook_id=$4 AND id=$5
		 RETURNING id`,
		input.Title, time.Now(), userID, notebookID, todoListID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, todoListID, userID, notebookID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM todo_lists
		 WHERE user_id=$1 AND notebook_id=$2 AND id=$3
		 RETURNING id`,
		userID, notebookID, todoListID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
