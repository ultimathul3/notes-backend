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
	repo                 domain.SessionRepository
	jwt                  jwtManager
	refreshTokenTTL      time.Duration
	maxUserSessionsCount int64
}

func NewUsecase(
	repo domain.SessionRepository,
	jwt jwtManager,
	refreshTokenTTL time.Duration,
	maxUserSessionsCount int64,
) *Usecase {
	return &Usecase{
		repo:                 repo,
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

func (u *Usecase) Create(ctx context.Context, input domain.CreateSessionDTO) (int64, error) {
	session := domain.Session{
		UserID:       input.UserID,
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
		ExpiresIn:    time.Now().Add(u.refreshTokenTTL),
	}

	return u.repo.Create(ctx, session)
}

func (u *Usecase) GetCountByUserID(ctx context.Context, userID int64) int64 {
	return u.repo.GetCountByUserID(ctx, userID)
}

func (u *Usecase) DeleteAllByUserID(ctx context.Context, userID int64) {
	u.repo.DeleteAllByUserID(ctx, userID)
}

func (u *Usecase) Refresh(ctx context.Context, input domain.RefreshSessionDTO) (string, uuid.UUID, error) {
	session, err := u.repo.GetByRefreshToken(ctx, *input.RefreshToken)
	if err != nil {
		return "", uuid.Nil, err
	}

	if session.Fingerprint != input.Fingerprint {
		u.repo.DeleteByID(ctx, session.ID)
		return "", uuid.Nil, domain.ErrInvalidFingerPrint
	}

	if time.Now().After(session.ExpiresIn) {
		u.repo.DeleteByID(ctx, session.ID)
		return "", uuid.Nil, domain.ErrInvalidOrExpiredRefreshToken
	}

	accessToken, refreshToken, err := u.jwt.GenerateTokens(session.ID)
	if err != nil {
		return "", uuid.Nil, err
	}

	if err := u.repo.Update(ctx, domain.UpdateSessionDTO{
		ID:           session.ID,
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().Add(u.refreshTokenTTL),
	}); err != nil {
		return "", uuid.Nil, err
	}

	return accessToken, refreshToken, nil
}
