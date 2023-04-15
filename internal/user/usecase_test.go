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
	expectedID := int64(1)
	repo.On(
		"Create",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.User")).
		Return(expectedID, nil)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	resultID, err := usecase.Create(context.Background(), domain.CreateUserDTO{
		Login:    toPtr("login"),
		Name:     toPtr("name"),
		Password: toPtr("password"),
	})

	assert.Equal(t, expectedID, resultID)
	assert.Nil(t, err)
}

func TestUserCreateError(t *testing.T) {
	repo := &mocks.UserRepository{}
	expectedID := int64(0)
	repo.On(
		"Create",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.User")).
		Return(expectedID, domain.ErrUserAlreadyExists)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	resultID, err := usecase.Create(context.Background(), domain.CreateUserDTO{
		Login:    toPtr("login"),
		Name:     toPtr("name"),
		Password: toPtr("password"),
	})

	assert.Equal(t, expectedID, resultID)
	assert.ErrorIs(t, err, domain.ErrUserAlreadyExists)
}

func TestGetUserID(t *testing.T) {
	repo := &mocks.UserRepository{}
	expectedID := int64(1)
	repo.On(
		"GetID",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).
		Return(expectedID, nil)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	resultID, err := usecase.GetID(context.Background(), domain.GetUserIDDTO{
		Login:    toPtr("login"),
		Password: toPtr("password"),
	})

	assert.Equal(t, expectedID, resultID)
	assert.Nil(t, err)
}
