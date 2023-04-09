package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	Session struct {
		ID           int64     `json:"id"`
		UserID       int64     `json:"user_id"`
		RefreshToken uuid.UUID `json:"refresh_token"`
		Fingerprint  string    `json:"fingerprint"`
		ExpiresIn    time.Time `json:"expires_in"`
	}

	CreateSessionDTO struct {
		UserID       *int64     `json:"user_id"`
		RefreshToken *uuid.UUID `json:"refresh_token"`
		Fingerprint  *string    `json:"fingerprint"`
		ExpiresIn    *time.Time `json:"expires_in"`
	}
)

type SessionUsecase interface {
	Create(ctx context.Context, session *CreateSessionDTO) (int64, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (int64, error)
}
