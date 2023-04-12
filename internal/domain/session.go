package domain

import (
	"context"
	"errors"
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
		UserID       int64     `json:"user_id"`
		RefreshToken uuid.UUID `json:"refresh_token"`
		Fingerprint  string    `json:"fingerprint"`
	}

	RefreshSessionDTO struct {
		UserID       int64      `json:"user_id"`
		RefreshToken *uuid.UUID `json:"refresh_token"`
		Fingerprint  string     `json:"fingerprint"`
	}
)

type SessionUsecase interface {
	GenerateTokens(userID int64) (string, uuid.UUID, error)
	GetMaxUserSessionsCount() int64
	Create(ctx context.Context, input *CreateSessionDTO) (int64, error)
	GetCountByUserID(ctx context.Context, userID int64) int64
	DeleteAllByUserID(ctx context.Context, userID int64)
	Refresh(ctx context.Context, input *RefreshSessionDTO) error
}

type SessionRepository interface {
	Create(ctx context.Context, session *Session) (int64, error)
	GetCountByUserID(ctx context.Context, userID int64) int64
	DeleteAllByUserID(ctx context.Context, userID int64)
}

var ErrInvalidOrExpiredAccessToken = errors.New("invalid or expired access token")
