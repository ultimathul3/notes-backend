package search

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

func getWhereQuery(search domain.Search) (string, []any, int) {
	whereValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if search.Title != "" {
		whereValues = append(whereValues, fmt.Sprintf("LOWER(title) LIKE '%%' || LOWER($%d) || '%%'", argID))
		args = append(args, search.Title)
		argID++
	}

	whereValues = append(whereValues, "")
	whereQuery := strings.Join(whereValues, " AND ")

	return whereQuery, args, argID
}

func (r *RepositoryPostgres) GetAllNotes(ctx context.Context, userID int64, search domain.Search) ([]domain.Note, error) {
	var notes []domain.Note

	whereQuery, args, argID := getWhereQuery(search)

	rows, err := r.conn.Query(
		ctx,
		fmt.Sprintf(
			`SELECT id, title, body, created_at, updated_at, notebook_id
			 FROM notes
			 WHERE %s user_id=$%d`,
			whereQuery, argID,
		),
		append(args, userID)...,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := domain.Note{}
		err := rows.Scan(&note.ID, &note.Title, &note.Body, &note.CreatedAt, &note.UpdatedAt, &note.NotebookID)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (r *RepositoryPostgres) GetAllTodoLists(ctx context.Context, userID int64, search domain.Search) ([]domain.TodoList, error) {
	var lists []domain.TodoList

	whereQuery, args, argID := getWhereQuery(search)

	rows, err := r.conn.Query(
		ctx,
		fmt.Sprintf(
			`SELECT id, title, created_at, updated_at, notebook_id
			 FROM todo_lists
			 WHERE %s user_id=$%d`,
			whereQuery, argID,
		),
		append(args, userID)...,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		list := domain.TodoList{}
		err := rows.Scan(&list.ID, &list.Title, &list.CreatedAt, &list.UpdatedAt, &list.NotebookID)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (r *RepositoryPostgres) GetAllSharedNotes(ctx context.Context, userID int64, search domain.Search) ([]domain.Note, error) {
	return nil, nil
}
