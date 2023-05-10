package domain

import (
	"context"
)

type (
	Search struct {
		Title       string
		ByNotes     bool
		ByTodoLists bool
	}

	SearchResult struct {
		Notes            []Note     `json:"notes,omitempty"`
		TodoLists        []TodoList `json:"todo_lists,omitempty"`
		SharedNotes      []Note     `json:"shared_notes,omitempty"`
		NotesCount       int        `json:"notes_count"`
		TodoListsCount   int        `json:"todo_lists_count"`
		SharedNotesCount int        `json:"shared_notes_count"`
	}
)

type SearchUsecase interface {
	GetAll(ctx context.Context, userID int64, search Search) (SearchResult, error)
}

type SearchRepository interface {
	GetAllNotes(ctx context.Context, userID int64, search Search) ([]Note, error)
	GetAllTodoLists(ctx context.Context, userID int64, search Search) ([]TodoList, error)
	GetAllSharedNotes(ctx context.Context, userID int64, search Search) ([]Note, error)
}
