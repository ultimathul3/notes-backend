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
	expectedUser := domain.User{
		ID:    1,
		Login: "login",
		Name:  "name",
	}
	repo.On(
		"Get",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).
		Return(expectedUser, nil)

	usecase := NewUsecase(repo, hash.NewSHA256Hasher([]byte("salt")))
	resultUser, err := usecase.Get(context.Background(), domain.GetUserDTO{
		Login:    toPtr("login"),
		Password: toPtr("password"),
	})

	assert.Equal(t, expectedUser.ID, resultUser.ID)
	assert.Equal(t, expectedUser.Name, resultUser.Name)
	assert.Equal(t, expectedUser.Login, resultUser.Login)
	assert.Nil(t, err)
}
