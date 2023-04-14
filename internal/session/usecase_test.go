package session

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultimathul3/notes-backend/internal/domain"
	"github.com/ultimathul3/notes-backend/internal/domain/mocks"
	"github.com/ultimathul3/notes-backend/pkg/jwtauth"
)

func toPtr[T any](t T) *T {
	return &t
}

func TestSessionCreate(t *testing.T) {
	repo := &mocks.SessionRepository{}
	repo.On(
		"Create",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.Session")).
		Return(int64(1), nil)

	usecase := NewUsecase(repo, jwtauth.NewJWT(10*time.Second, ""), 60*time.Second, 1)
	id, err := usecase.Create(context.Background(), domain.CreateSessionDTO{
		UserID:       1,
		RefreshToken: uuid.New(),
		Fingerprint:  "fingerprint",
	})

	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
}

func TestSessionRefresh(t *testing.T) {
	repo := &mocks.SessionRepository{}
	uuidTest := uuid.New()
	repo.On(
		"GetByRefreshToken",
		mock.AnythingOfType("*context.emptyCtx"),
		uuidTest).
		Return(
			domain.Session{
				ID:           1,
				UserID:       1,
				RefreshToken: uuidTest,
				Fingerprint:  "fingerprint",
				ExpiresIn:    time.Now().Add(-10 * time.Second),
			}, nil)
	repo.On(
		"GetByRefreshToken",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("uuid.UUID")).
		Return(
			domain.Session{
				ID:           1,
				UserID:       1,
				RefreshToken: uuid.New(),
				Fingerprint:  "fingerprint",
				ExpiresIn:    time.Now().Add(1 * time.Second),
			}, nil)
	repo.On(
		"Update",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("domain.UpdateSessionDTO")).
		Return(nil)
	repo.On(
		"DeleteByID",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.AnythingOfType("int64")).
		Return(nil)

	usecase := NewUsecase(repo, jwtauth.NewJWT(10*time.Second, ""), 60*time.Second, 1)

	_, _, err := usecase.Refresh(context.Background(), domain.RefreshSessionDTO{
		UserID:       1,
		RefreshToken: toPtr(uuid.New()),
		Fingerprint:  "other fingerprint",
	})
	assert.ErrorIs(t, err, domain.ErrInvalidFingerPrint)

	_, _, err = usecase.Refresh(context.Background(), domain.RefreshSessionDTO{
		UserID:       1,
		RefreshToken: toPtr(uuid.New()),
		Fingerprint:  "fingerprint",
	})
	assert.Nil(t, err)

	_, _, err = usecase.Refresh(context.Background(), domain.RefreshSessionDTO{
		UserID:       1,
		RefreshToken: toPtr(uuidTest),
		Fingerprint:  "fingerprint",
	})
	assert.ErrorIs(t, err, domain.ErrInvalidOrExpiredRefreshToken)
}
