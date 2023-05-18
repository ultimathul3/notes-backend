package domain

import (
	"context"
	"errors"
	"time"
)

type (
	SharedTodoList struct {
		ID         int64 `json:"id"`
		WhoseID    int64 `json:"whose_id"`
		WhomID     int64 `json:"whom_id"`
		TodoListID int64 `json:"note_id"`
		Accepted   bool  `json:"accepted"`
	}

	CreateSharedTodoListDTO struct {
		Login      *string `json:"login"`
		TodoListID *int64  `json:"todo_list_id"`
	}

	SharedTodoListInfo struct {
		ID         int64  `json:"id"`
		OwnerLogin string `json:"owner_login"`
		OwnerName  string `json:"owner_name"`
		Title      string `json:"title"`
		Accepted   bool   `json:"accepted"`
	}

	SharedTodoListData struct {
		Items     []TodoItem `json:"items,omitempty"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
	}

	OutgoingSharedTodoListInfo struct {
		ID             int64  `json:"id"`
		RecipientLogin string `json:"recipient_login"`
		RecipientName  string `json:"recipient_name"`
		Accepted       bool   `json:"accepted"`
	}

	GetSharedTodoListsInfoResponse struct {
		SharedTodoListsInfo []SharedTodoListInfo `json:"shared_todo_lists,omitempty"`
		Count               int                  `json:"count"`
	}

	GetOutgoingSharedTodoListsInfoResponse struct {
		OutgoingSharedTodoListsInfo []OutgoingSharedTodoListInfo `json:"shared_todo_lists,omitempty"`
		Count                       int                          `json:"count"`
	}
)

type SharedTodoListUsecase interface {
	Create(ctx context.Context, whoseID, whomID, noteID int64) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetAllInfo(ctx context.Context, whomID int64) ([]SharedTodoListInfo, error)
	Accept(ctx context.Context, id, whomID int64) error
	GetDataByID(ctx context.Context, id, whomID int64) (SharedTodoListData, error)
	GetOutgoingInfoByTodoListID(ctx context.Context, noteID, whoseID int64) ([]OutgoingSharedTodoListInfo, error)
}

type SharedTodoListRepository interface {
	Create(ctx context.Context, sharedTodoList SharedTodoList) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetAllInfo(ctx context.Context, whomID int64) ([]SharedTodoListInfo, error)
	Accept(ctx context.Context, id, whomID int64) error
	GetTimestampsByID(ctx context.Context, id, whomID int64) (time.Time, time.Time, error)
	GetItemsByID(ctx context.Context, id, whomID int64) ([]TodoItem, error)
	GetOutgoingInfoByTodoListID(ctx context.Context, noteID, whoseID int64) ([]OutgoingSharedTodoListInfo, error)
}

func (cs *CreateSharedTodoListDTO) Validate() error {
	if cs.Login == nil {
		return errors.New("empty whom id")
	}
	if cs.TodoListID == nil {
		return errors.New("empty todo list id")
	}
	return nil
}

var (
	ErrImpossibleToShareTodoListWithYourself = errors.New("impossible to share todo list with yourself")
	ErrSharedTodoListNotFound                = errors.New("shared todo list not found")
	ErrSharedTodoListsNotFound               = errors.New("shared todo lists not found")
	ErrTodoListHasAlreadyBeenShared          = errors.New("todo list has already been shared")
)
