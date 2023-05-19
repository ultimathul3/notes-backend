package domain

import (
	"context"
)

type (
	Search struct {
		Title             string
		ByNotes           bool
		ByTodoLists       bool
		BySharedNotes     bool
		BySharedTodoLists bool
	}

	SearchResult struct {
		Notes                []Note               `json:"notes,omitempty"`
		TodoLists            []TodoList           `json:"todo_lists,omitempty"`
		SharedNotes          []SharedNoteInfo     `json:"shared_notes,omitempty"`
		SharedTodoLists      []SharedTodoListInfo `json:"shared_todo_lists,omitempty"`
		NotesCount           int                  `json:"notes_count"`
		TodoListsCount       int                  `json:"todo_lists_count"`
		SharedNotesCount     int                  `json:"shared_notes_count"`
		SharedTodoListsCount int                  `json:"shared_todo_lists_count"`
	}
)

type SearchUsecase interface {
	GetAll(ctx context.Context, userID int64, search Search) (SearchResult, error)
}

type SearchRepository interface {
	GetAllNotes(ctx context.Context, userID int64, search Search) ([]Note, error)
	GetAllTodoLists(ctx context.Context, userID int64, search Search) ([]TodoList, error)
	GetAllAcceptedSharedNotes(ctx context.Context, userID int64, search Search) ([]SharedNoteInfo, error)
	GetAllAcceptedSharedTodoLists(ctx context.Context, userID int64, search Search) ([]SharedTodoListInfo, error)
}
