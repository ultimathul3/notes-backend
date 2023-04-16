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

func (u *Usecase) Create(ctx context.Context, userID, notebookID int64, input domain.CreateNoteDTO) (int64, error) {
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