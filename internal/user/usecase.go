package user

import (
	"context"
	"fmt"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Hasher interface {
	Hash(data []byte) ([]byte, error)
}

type Usecase struct {
	repository     domain.UserRepository
	passwordHasher Hasher
}

func NewUsecase(repository domain.UserRepository, passwordHasher Hasher) *Usecase {
	return &Usecase{
		repository:     repository,
		passwordHasher: passwordHasher,
	}
}

func (u *Usecase) Create(ctx context.Context, input *domain.CreateUserDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	passwordHash, err := u.passwordHasher.Hash([]byte(*input.Password))
	if err != nil {
		return 0, err
	}

	user := domain.User{
		Login:        *input.Login,
		Name:         *input.Name,
		PasswordHash: fmt.Sprintf("%x", passwordHash),
	}

	return u.repository.Create(ctx, &user)
}

func (u *Usecase) GetID(ctx context.Context, input *domain.GetUserIDDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	passwordHash, err := u.passwordHasher.Hash([]byte(*input.Password))
	if err != nil {
		return 0, err
	}

	return u.repository.GetID(ctx, *input.Login, fmt.Sprintf("%x", passwordHash))
}
