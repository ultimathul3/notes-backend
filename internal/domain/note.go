package domain

import (
	"context"
	"errors"
	"time"
	"unicode/utf8"
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

	PatchNoteDTO struct {
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
	GetByID(ctx context.Context, userID, notebookID, noteID int64) (Note, error)
	Delete(ctx context.Context, noteID, userID, notebookID int64) error
	Patch(ctx context.Context, noteID, userID, notebookID int64, input PatchNoteDTO) error
}

type NoteRepository interface {
	Create(ctx context.Context, note Note) (int64, error)
	GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]Note, error)
	GetByID(ctx context.Context, userID, notebookID, noteID int64) (Note, error)
	Delete(ctx context.Context, noteID, userID, notebookID int64) error
	Patch(ctx context.Context, noteID, userID, notebookID int64, input PatchNoteDTO) error
}

func (cn *CreateNoteDTO) Validate() error {
	if cn.Title == nil {
		return errors.New("empty title")
	}
	if cn.Body == nil {
		return errors.New("empty body")
	}
	if utf8.RuneCountInString(*cn.Title) == 0 || utf8.RuneCountInString(*cn.Title) > 64 {
		return errors.New("title length must be from 1 to 64 characters")
	}
	return nil
}

func (pn *PatchNoteDTO) Validate() error {
	if pn.Title == nil && pn.Body == nil {
		return errors.New("empty title and body")
	}
	if pn.Title != nil && (utf8.RuneCountInString(*pn.Title) == 0 || utf8.RuneCountInString(*pn.Title) > 64) {
		return errors.New("title length must be from 1 to 64 characters")
	}
	return nil
}

var ErrNoteNotFound = errors.New("note not found")
