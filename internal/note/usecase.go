package note

import (
	"context"
	"time"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.NoteRepository
}

func NewUsecase(repo domain.NoteRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, userID, notebookID int64, input domain.CreateUpdateNoteDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	return u.repo.Create(ctx, domain.Note{
		Title:      *input.Title,
		Body:       *input.Body,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		UserID:     userID,
		NotebookID: notebookID,
	})
}

func (u *Usecase) GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]domain.Note, error) {
	return u.repo.GetAllByNotebookID(ctx, userID, notebookID)
}

func (u *Usecase) Update(ctx context.Context, noteID, userID, notebookID int64, input domain.CreateUpdateNoteDTO) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return u.repo.Update(ctx, domain.Note{
		ID:         noteID,
		Title:      *input.Title,
		Body:       *input.Body,
		UpdatedAt:  time.Now(),
		UserID:     userID,
		NotebookID: notebookID,
	})
}
