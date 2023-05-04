package domain

import (
	"context"
	"errors"
	"time"
	"unicode/utf8"
)

type (
	TodoList struct {
		ID         int64     `json:"id"`
		Title      string    `json:"title"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		UserID     int64     `json:"user_id,omitempty"`
		NotebookID int64     `json:"notebook_id,omitempty"`
	}

	CreateTodoListDTO struct {
		Title *string `json:"title"`
	}

	GetAllTodoListsResponse struct {
		TodoLists []TodoList `json:"todo_lists,omitempty"`
		Count     int        `json:"count"`
	}

	UpdateTodoListDTO struct {
		Title *string `json:"title"`
	}
)

type TodoListUsecase interface {
	Create(ctx context.Context, userID, notebookID int64, input CreateTodoListDTO) (int64, error)
	GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]TodoList, error)
	Update(ctx context.Context, todoListID, userID, notebookID int64, input UpdateTodoListDTO) error
	Delete(ctx context.Context, todoListID, userID, notebookID int64) error
}

type TodoListRepository interface {
	Create(ctx context.Context, list TodoList) (int64, error)
	GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]TodoList, error)
	Update(ctx context.Context, todoListID, userID, notebookID int64, input UpdateTodoListDTO) error
	Delete(ctx context.Context, todoListID, userID, notebookID int64) error
}

func (ct *CreateTodoListDTO) Validate() error {
	if ct.Title == nil {
		return errors.New("empty title")
	}
	if utf8.RuneCountInString(*ct.Title) == 0 || utf8.RuneCountInString(*ct.Title) > 64 {
		return errors.New("title length must be from 1 to 64 characters")
	}
	return nil
}

func (ut *UpdateTodoListDTO) Validate() error {
	if ut.Title == nil {
		return errors.New("empty title")
	}
	if utf8.RuneCountInString(*ut.Title) == 0 || utf8.RuneCountInString(*ut.Title) > 64 {
		return errors.New("title length must be from 1 to 64 characters")
	}
	return nil
}

var ErrTodoListNotFound = errors.New("todo list not found")
