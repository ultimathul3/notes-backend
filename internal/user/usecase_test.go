package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultimathul3/notes-backend/internal/domain"
	"github.com/ultimathul3/notes-backend/internal/domain/mocks"
	"github.com/ultimathul3/notes-backend/pkg/hash"
)

func toPtr[T any](t T) *T {
	return &t
}

func TestUserCreate(t *testing.T) {
	repo := &mocks.UserRepository{}
	repo.On(
		"Create",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.User")).
		Return(int64(1), nil)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	_, err := usecase.Create(context.Background(), domain.CreateUserDTO{
		Login:    toPtr("login"),
		Name:     toPtr("name"),
		Password: toPtr("password"),
	})

	assert.Nil(t, err)
}

func TestUserCreateError(t *testing.T) {
	repo := &mocks.UserRepository{}
	repo.On(
		"Create",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.User")).
		Return(int64(0), domain.ErrUserAlreadyExists)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	_, err := usecase.Create(context.Background(), domain.CreateUserDTO{
		Login:    toPtr("login"),
		Name:     toPtr("name"),
		Password: toPtr("password"),
	})

	assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
}
