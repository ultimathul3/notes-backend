package session

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repository domain.SessionRepository
}

func NewUsecase(repository domain.SessionRepository) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (u *Usecase) Create(ctx context.Context, input *domain.CreateSessionDTO) (int64, error) {
	session := domain.Session{
		UserID:       *input.UserID,
		RefreshToken: *input.RefreshToken,
		Fingerprint:  *input.Fingerprint,
		ExpiresIn:    *input.ExpiresIn,
	}

	return u.repository.Create(ctx, &session)
}
