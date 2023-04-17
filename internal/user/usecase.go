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
	repo           domain.UserRepository
	passwordHasher Hasher
}

func NewUsecase(repo domain.UserRepository, passwordHasher Hasher) *Usecase {
	return &Usecase{
		repo:           repo,
		passwordHasher: passwordHasher,
	}
}

func (u *Usecase) Create(ctx context.Context, input domain.CreateUserDTO) (int64, error) {
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

	return u.repo.Create(ctx, user)
}

func (u *Usecase) GetID(ctx context.Context, input domain.GetUserIdDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	passwordHash, err := u.passwordHasher.Hash([]byte(*input.Password))
	if err != nil {
		return 0, err
	}

	return u.repo.GetID(ctx, *input.Login, fmt.Sprintf("%x", passwordHash))
}

func (u *Usecase) GetUserIdByLogin(ctx context.Context, input domain.GetUserIdByLoginDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	return u.repo.GetUserIdByLogin(ctx, *input.Login)
}
