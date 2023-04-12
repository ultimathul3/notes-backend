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

	UpdateSessionDTO struct {
		ID           int64     `json:"id"`
		RefreshToken uuid.UUID `json:"refresh_token"`
		ExpiresIn    time.Time `json:"expires_in"`
	}
)

type SessionUsecase interface {
	GenerateTokens(userID int64) (string, uuid.UUID, error)
	GetMaxUserSessionsCount() int64
	Create(ctx context.Context, input CreateSessionDTO) (int64, error)
	GetCountByUserID(ctx context.Context, userID int64) int64
	DeleteAllByUserID(ctx context.Context, userID int64)
	Refresh(ctx context.Context, input RefreshSessionDTO) (string, uuid.UUID, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session Session) (int64, error)
	GetCountByUserID(ctx context.Context, userID int64) int64
	DeleteAllByUserID(ctx context.Context, userID int64) error
	GetByRefreshToken(ctx context.Context, refreshToken uuid.UUID) (Session, error)
	Update(ctx context.Context, input UpdateSessionDTO) error
	DeleteByID(ctx context.Context, id int64) error
}

func (r *RefreshSessionDTO) Validate() error {
	if r.RefreshToken == nil {
		return errors.New("empty refresh token")
	}
	return nil
}

var (
	ErrInvalidOrExpiredAccessToken  = errors.New("invalid or expired access token")
	ErrInvalidOrExpiredRefreshToken = errors.New("invalid or expired refresh token")
	ErrInvalidFingerPrint           = errors.New("invalid fingerprint")
)
