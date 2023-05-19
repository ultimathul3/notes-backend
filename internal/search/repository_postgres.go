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

func (r *RepositoryPostgres) GetAllAcceptedSharedNotes(ctx context.Context, userID int64, search domain.Search) ([]domain.SharedNoteInfo, error) {
	var notes []domain.SharedNoteInfo

	whereQuery, args, argID := getWhereQuery(search)

	rows, err := r.conn.Query(
		ctx,
		fmt.Sprintf(
			`SELECT s.id, u.login, u.name, n.title, s.accepted
			 FROM shared_notes s
			 LEFT JOIN users u ON u.id=s.whose_id
			 LEFT JOIN notes n ON n.id=s.note_id
			 WHERE %s s.whom_id=$%d AND s.accepted=true`,
			whereQuery, argID,
		),
		append(args, userID)...,
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

func (r *RepositoryPostgres) GetAllAcceptedSharedTodoLists(ctx context.Context, userID int64, search domain.Search) ([]domain.SharedTodoListInfo, error) {
	var lists []domain.SharedTodoListInfo

	whereQuery, args, argID := getWhereQuery(search)

	rows, err := r.conn.Query(
		ctx,
		fmt.Sprintf(
			`SELECT s.id, u.login, u.name, t.title, s.accepted
			 FROM shared_todo_lists s
			 LEFT JOIN users u ON u.id=s.whose_id
			 LEFT JOIN todo_lists t ON t.id=s.todo_list_id
			 WHERE %s s.whom_id=$%d and s.accepted=true`,
			whereQuery, argID,
		),
		append(args, userID)...,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		list := domain.SharedTodoListInfo{}
		err := rows.Scan(&list.ID, &list.OwnerLogin, &list.OwnerName, &list.Title, &list.Accepted)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}
