package domain

import (
	"context"
	"errors"
)

type (
	SharedNote struct {
		ID       int64 `json:"id"`
		WhoseID  int64 `json:"whose_id"`
		WhomID   int64 `json:"whom_id"`
		NoteID   int64 `json:"note_id"`
		Accepted bool  `json:"accepted"`
	}

	CreateSharedNoteDTO struct {
		Login  *string `json:"login"`
		NoteID *int64  `json:"note_id"`
	}

	IncomingSharedNote struct {
		ID         int64  `json:"id"`
		OwnerLogin string `json:"owner_login"`
		OwnerName  string `json:"owner_name"`
		Title      string `json:"title"`
	}

	GetAllIncomingSharedNotesResponse struct {
		IncomingSharedNotes []IncomingSharedNote `json:"incoming_shared_notes,omitempty"`
		Count               int                  `json:"count"`
	}
)

type SharedNoteUsecase interface {
	Create(ctx context.Context, whoseID, whomID, noteID int64) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetIncomingSharedNotes(ctx context.Context, whomID int64) ([]IncomingSharedNote, error)
	Accept(ctx context.Context, id, whomID int64) error
}

type SharedNoteRepository interface {
	Create(ctx context.Context, sharedNote SharedNote) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetIncomingSharedNotes(ctx context.Context, whomID int64) ([]IncomingSharedNote, error)
	Accept(ctx context.Context, id, whomID int64) error
}

func (cs *CreateSharedNoteDTO) Validate() error {
	if cs.Login == nil {
		return errors.New("empty whom id")
	}
	if cs.NoteID == nil {
		return errors.New("empty note id")
	}
	return nil
}

var (
	ErrImpossibleToShareNoteWithYourself = errors.New("impossible to share note with yourself")
	ErrSharedNoteNotFound                = errors.New("shared note not found")
	ErrAlreadyShared                     = errors.New("already shared")
)
