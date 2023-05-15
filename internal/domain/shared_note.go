package domain

import (
	"context"
	"errors"
	"time"
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

	SharedNoteInfo struct {
		ID         int64  `json:"id"`
		OwnerLogin string `json:"owner_login"`
		OwnerName  string `json:"owner_name"`
		Title      string `json:"title"`
		Accepted   bool   `json:"accepted"`
	}

	SharedNoteData struct {
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	OutgoingSharedNoteInfo struct {
		ID             int64  `json:"id"`
		RecipientLogin string `json:"recipient_login"`
		RecipientName  string `json:"recipient_name"`
		Accepted       bool   `json:"accepted"`
	}

	GetSharedNotesInfoResponse struct {
		SharedNotesInfo []SharedNoteInfo `json:"shared_notes,omitempty"`
		Count           int              `json:"count"`
	}

	GetOutgoingSharedNotesInfoResponse struct {
		OutgoingSharedNotesInfo []OutgoingSharedNoteInfo `json:"shared_notes,omitempty"`
		Count                   int                      `json:"count"`
	}
)

type SharedNoteUsecase interface {
	Create(ctx context.Context, whoseID, whomID, noteID int64) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetAllInfo(ctx context.Context, whomID int64) ([]SharedNoteInfo, error)
	Accept(ctx context.Context, id, whomID int64) error
	GetDataByID(ctx context.Context, id, whomID int64) (SharedNoteData, error)
	GetOutgoingInfoByNoteID(ctx context.Context, noteID, whoseID int64) ([]OutgoingSharedNoteInfo, error)
}

type SharedNoteRepository interface {
	Create(ctx context.Context, sharedNote SharedNote) (int64, error)
	Delete(ctx context.Context, id, whomID int64) error
	GetAllInfo(ctx context.Context, whomID int64) ([]SharedNoteInfo, error)
	Accept(ctx context.Context, id, whomID int64) error
	GetDataByID(ctx context.Context, id, whomID int64) (SharedNoteData, error)
	GetOutgoingInfoByNoteID(ctx context.Context, noteID, whoseID int64) ([]OutgoingSharedNoteInfo, error)
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
