package notebook

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.NotebookRepository
}

func NewUsecase(repo domain.NotebookUsecase) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, notebook domain.Notebook) (int64, error) {
	if err := notebook.Validate(); err != nil {
		return 0, err
	}

	return u.repo.Create(ctx, notebook)
}
