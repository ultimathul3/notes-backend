package session

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

type jwtManager interface {
	GenerateTokens(userID int64) (string, uuid.UUID, error)
	ParseAccessToken(token string) (int64, error)
}

type Usecase struct {
	repository           domain.SessionRepository
	jwt                  jwtManager
	refreshTokenTTL      time.Duration
	maxUserSessionsCount int64
}

func NewUsecase(
	repository domain.SessionRepository,
	jwt jwtManager,
	refreshTokenTTL time.Duration,
	maxUserSessionsCount int64,
) *Usecase {
	return &Usecase{
		repository:           repository,
		jwt:                  jwt,
		refreshTokenTTL:      refreshTokenTTL,
		maxUserSessionsCount: maxUserSessionsCount,
	}
}

func (u *Usecase) GenerateTokens(userID int64) (string, uuid.UUID, error) {
	return u.jwt.GenerateTokens(userID)
}

func (u *Usecase) GetMaxUserSessionsCount() int64 {
	return u.maxUserSessionsCount
}

func (u *Usecase) Create(ctx context.Context, input *domain.CreateSessionDTO) (int64, error) {
	session := domain.Session{
		UserID:       input.UserID,
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
		ExpiresIn:    time.Now().Add(u.refreshTokenTTL),
	}

	return u.repository.Create(ctx, &session)
}

func (u *Usecase) GetUserSessionsCount(ctx context.Context, userID int64) int64 {
	return u.repository.GetUserSessionsCount(ctx, userID)
}

func (u *Usecase) DeleteAllUserSessions(ctx context.Context, userID int64) {
	u.repository.DeleteAllUserSessions(ctx, userID)
}

func (u *Usecase) RefreshUserSession(ctx context.Context, input *domain.RefreshSessionDTO) error {
	return nil
}
