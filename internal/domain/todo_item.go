package domain

import (
	"context"
	"errors"
	"unicode/utf8"
)

type (
	TodoItem struct {
		ID         int64  `json:"id"`
		Body       string `json:"body"`
		Done       bool   `json:"done"`
		TodoListID int64  `json:"todo_list_id,omitempty"`
	}

	CreateTodoItemDTO struct {
		Body *string `json:"body"`
	}

	GetAllTodoItemsResponse struct {
		TodoItems []TodoItem `json:"todo_items,omitempty"`
		Count     int        `json:"count"`
	}

	PatchTodoItemDTO struct {
		Body *string `json:"body"`
		Done *bool   `json:"done"`
	}
)

type TodoItemUsecase interface {
	Create(ctx context.Context, userID, notebookID, listID int64, input CreateTodoItemDTO) (int64, error)
	GetAllByListID(ctx context.Context, userID, notebookID, listID int64) ([]TodoItem, error)
	Patch(ctx context.Context, itemID, userID, notebookID, listID int64, input PatchTodoItemDTO) error
	Delete(ctx context.Context, itemID, userID, notebookID, listID int64) error
}

type TodoItemRepository interface {
	Create(ctx context.Context, userID, notebookID int64, item TodoItem) (int64, error)
	GetAllByListID(ctx context.Context, userID, notebookID, listID int64) ([]TodoItem, error)
	Patch(ctx context.Context, itemID, userID, notebookID, listID int64, input PatchTodoItemDTO) error
	Delete(ctx context.Context, itemID, userID, notebookID, listID int64) error
}

func (ct *CreateTodoItemDTO) Validate() error {
	if ct.Body == nil {
		return errors.New("empty title")
	}
	if utf8.RuneCountInString(*ct.Body) == 0 || utf8.RuneCountInString(*ct.Body) > 128 {
		return errors.New("body length must be from 1 to 128 characters")
	}
	return nil
}

func (ut *PatchTodoItemDTO) Validate() error {
	if ut.Body == nil && ut.Done == nil {
		return errors.New("empty fields")
	}
	if ut.Body != nil && (utf8.RuneCountInString(*ut.Body) == 0 || utf8.RuneCountInString(*ut.Body) > 128) {
		return errors.New("body length must be from 1 to 128 characters")
	}
	return nil
}
