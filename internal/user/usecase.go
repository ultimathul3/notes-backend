package user

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repository domain.UserRepository
}

func NewUsecase(repository domain.UserRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (u *Usecase) Create(ctx context.Context, user *domain.User) (int64, error) {
	return u.repository.Create(ctx, user)
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.repository.GetByID(ctx, id)
}
