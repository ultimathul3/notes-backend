package sharedtodolist

import (
	"context"
	"errors"
	"time"

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

func (r *RepositoryPostgres) Create(ctx context.Context, sharedTodoList domain.SharedTodoList) (int64, error) {
	if err := r.conn.QueryRow(
		ctx,
		`INSERT INTO shared_todo_lists (whose_id, whom_id, todo_list_id, accepted)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id`,
		sharedTodoList.WhoseID, sharedTodoList.WhomID, sharedTodoList.TodoListID, sharedTodoList.Accepted,
	).Scan(&sharedTodoList.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == domain.UniqueViolation {
				return 0, domain.ErrTodoListHasAlreadyBeenShared
			}
		}
		return 0, err
	}

	return sharedTodoList.ID, nil
}

func (r *RepositoryPostgres) Delete(ctx context.Context, id, whomID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`DELETE FROM shared_todo_lists
		 WHERE id=$1 AND (whom_id=$2 OR whose_id=$2)
		 RETURNING id`,
		id, whomID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) GetAllInfo(ctx context.Context, whomID int64) ([]domain.SharedTodoListInfo, error) {
	var lists []domain.SharedTodoListInfo

	rows, err := r.conn.Query(
		ctx,
		`SELECT s.id, u.login, u.name, t.title, s.accepted
		 FROM shared_todo_lists s
		 LEFT JOIN users u ON u.id=s.whose_id
		 LEFT JOIN todo_lists t ON t.id=s.todo_list_id
		 WHERE s.whom_id=$1`,
		whomID,
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

func (r *RepositoryPostgres) Accept(ctx context.Context, id, whomID int64) error {
	if err := r.conn.QueryRow(
		ctx,
		`UPDATE shared_todo_lists
		 SET accepted=true
		 WHERE id=$1 AND whom_id=$2
		 RETURNING id`,
		id, whomID,
	).Scan(nil); err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) GetTimestampsByID(ctx context.Context, id, whomID int64) (time.Time, time.Time, error) {
	var createdAt, updatedAt time.Time

	if err := r.conn.QueryRow(
		ctx,
		`SELECT t.created_at, t.updated_at
		 FROM shared_todo_lists s
		 LEFT JOIN todo_lists t ON t.id=s.todo_list_id
		 WHERE s.id=$1 AND s.whom_id=$2 AND s.accepted=true`,
		id, whomID,
	).Scan(&createdAt, &updatedAt); err != nil {
		return time.Time{}, time.Time{}, err
	}

	return createdAt, updatedAt, nil
}

func (r *RepositoryPostgres) GetItemsByID(ctx context.Context, id, whomID int64) ([]domain.TodoItem, error) {
	var items []domain.TodoItem

	rows, err := r.conn.Query(
		ctx,
		`SELECT t.id, t.body, t.done
		 FROM shared_todo_lists s
		 JOIN todo_items t ON t.todo_list_id=s.todo_list_id
		 WHERE s.id=$1 AND s.whom_id=$2 AND s.accepted=true
		 ORDER BY t.id`,
		id, whomID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := domain.TodoItem{}
		err := rows.Scan(&item.ID, &item.Body, &item.Done)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *RepositoryPostgres) GetOutgoingInfoByTodoListID(ctx context.Context, noteID, whoseID int64) ([]domain.OutgoingSharedTodoListInfo, error) {
	var lists []domain.OutgoingSharedTodoListInfo

	rows, err := r.conn.Query(
		ctx,
		`SELECT s.id, u.login, u.name, s.accepted
		 FROM shared_todo_lists s
		 LEFT JOIN users u ON u.id=s.whom_id
		 LEFT JOIN todo_lists t ON t.id=s.todo_list_id
		 WHERE s.whose_id=$1 AND s.todo_list_id=$2 AND accepted=true`,
		whoseID, noteID,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		list := domain.OutgoingSharedTodoListInfo{}
		err := rows.Scan(&list.ID, &list.RecipientLogin, &list.RecipientName, &list.Accepted)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}
