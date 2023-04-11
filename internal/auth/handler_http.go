package auth

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ultimathul3/notes-backend/internal/domain"
)

var ErrInvalidOrExpiredAccessToken = errors.New("invalid or expired access token")

type jwtManager interface {
	GenerateTokens(userID int64) (string, uuid.UUID, error)
	ParseAccessToken(token string) (int64, error)
}

type HandlerHTTP struct {
	uuc                  domain.UserUsecase
	suc                  domain.SessionUsecase
	jwt                  jwtManager
	refreshTokenTTL      time.Duration
	maxUserSessionsCount int64
}

func NewHandlerHTTP(
	router *gin.Engine,
	uuc domain.UserUsecase,
	suc domain.SessionUsecase,
	jwt jwtManager,
	refreshTokenTTL time.Duration,
	maxUserSessionsCount int64,
) {
	handler := &HandlerHTTP{
		uuc:                  uuc,
		suc:                  suc,
		jwt:                  jwt,
		refreshTokenTTL:      refreshTokenTTL,
		maxUserSessionsCount: maxUserSessionsCount,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", handler.signIn)
		auth.POST("/sign-up", handler.signUp)
	}
}
