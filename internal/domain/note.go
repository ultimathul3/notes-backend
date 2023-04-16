package domain

import (
	"context"
	"errors"
	"time"
)

type (
	Note struct {
		ID         int64     `json:"id"`
		Title      string    `json:"title"`
		Body       string    `json:"body"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		UserID     int64     `json:"user_id,omitempty"`
		NotebookID int64     `json:"notebook_id,omitempty"`
	}

	CreateNoteDTO struct {
		Title *string `json:"title"`
		Body  *string `json:"body"`
	}

	GetAllNotesResponse struct {
		Notes []Note `json:"notes,omitempty"`
		Count int    `json:"count"`
	}
)

type NoteUsecase interface {
	Create(ctx context.Context, userID, notebookID int64, input CreateNoteDTO) (int64, error)
	GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]Note, error)
}

type NoteRepository interface {
	Create(ctx context.Context, note Note) (int64, error)
	GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]Note, error)
}

func (cn *CreateNoteDTO) Validate() error {
	if cn.Title == nil {
		return errors.New("empty title")
	}
	if cn.Body == nil {
		return errors.New("empty body")
	}
	return nil
}

var ErrNotebookNotFound = errors.New("notebook not found")
