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

func (u *Usecase) Create(ctx context.Context, input *domain.CreateUserDTO) (int64, error) {
	user := domain.User{
		Login:        *input.Login,
		Name:         *input.Name,
		PasswordHash: *input.Password,
	}

	return u.repository.Create(ctx, &user)
}

func (u *Usecase) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.repository.GetByID(ctx, id)
}
