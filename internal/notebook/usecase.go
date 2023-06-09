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

func (u *Usecase) GetAllByUserID(ctx context.Context, userID int64) ([]domain.Notebook, error) {
	return u.repo.GetAllByUserID(ctx, userID)
}

func (u *Usecase) Update(ctx context.Context, notebook domain.Notebook) error {
	if err := notebook.Validate(); err != nil {
		return err
	}

	return u.repo.Update(ctx, notebook)
}

func (u *Usecase) Delete(ctx context.Context, id, userID int64) error {
	return u.repo.Delete(ctx, id, userID)
}
