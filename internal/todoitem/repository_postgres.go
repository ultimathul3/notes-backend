package todoitem

import (
	"context"
	"fmt"
	"strings"

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

func (r *RepositoryPostgres) Create(ctx context.Context, userID, notebookID int64, item domain.TodoItem) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO todo_items (body, done, todo_list_id)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		item.Body, item.Done, item.TodoListID,
	).Scan(&item.ID); err != nil {
		return 0, err
	}

	return item.ID, nil
}

func (r *RepositoryPostgres) GetAllByListID(ctx context.Context, userID, notebookID, listID int64) ([]domain.TodoItem, error) {
	var items []domain.TodoItem

	rows, err := r.conn.Query(
		ctx,
		`SELECT ti.id, ti.body, ti.done, ti.todo_list_id
		 FROM todo_items ti
		 LEFT JOIN todo_lists tl ON tl.id=ti.todo_list_id
		 WHERE ti.todo_list_id=$1 AND tl.user_id=$2 AND tl.notebook_id=$3
		 ORDER BY ti.id`,
		listID, userID, notebookID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := domain.TodoItem{}
		err := rows.Scan(&item.ID, &item.Body, &item.Done, &item.TodoListID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *RepositoryPostgres) Patch(ctx context.Context, itemID, userID, notebookID, listID int64, input domain.PatchTodoItemDTO) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Body != nil {
		setValues = append(setValues, fmt.Sprintf("body=$%d", argID))
		args = append(args, *input.Body)
		argID++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	if err := r.conn.QueryRow(
		ctx,
		fmt.Sprintf(
			`UPDATE todo_items
			 SET %s
			 FROM todo_lists
		 	 WHERE todo_items.id=$%d AND todo_items.todo_list_id=$%d AND todo_lists.user_id=$%d AND todo_lists.notebook_id=$%d
			 RETURNING todo_items.id`,
			setQuery, argID, argID+1, argID+2, argID+3,
		),
		append(args, itemID, listID, userID, notebookID)...,
	).Scan(nil); err != nil {
		fmt.Println(err.Error())
		return domain.ErrTodoListNotFound
	}

	return nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, itemID, userID, notebookID, listID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM todo_items ti
		 USING todo_lists tl
		 WHERE ti.id=$1 AND ti.todo_list_id=$2 AND tl.user_id=$3 AND tl.notebook_id=$4
		 RETURNING ti.id`,
		itemID, listID, userID, notebookID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}
