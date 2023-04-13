package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var (
	ErrMissingOrInvalidAuthHeader  = errors.New("missing or invalid authorization header")
	ErrInvalidOrExpiredAccessToken = errors.New("invalid or expired access token")
)

type jwtManager interface {
	GenerateTokens(userID int64) (string, uuid.UUID, error)
	ParseAccessToken(token string) (int64, error)
}

type TokenChecker struct {
	jwt jwtManager
}

func NewTokenChecker(jwt jwtManager) *TokenChecker {
	return &TokenChecker{
		jwt: jwt,
	}
}

func (t *TokenChecker) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerParts := strings.Split(header, "Bearer ")

		if len(headerParts) != 2 {
			log.Error("TokenChecker: ", ErrMissingOrInvalidAuthHeader)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": ErrMissingOrInvalidAuthHeader.Error(),
			})
			return
		}

		token := headerParts[1]

		userID, err := t.jwt.ParseAccessToken(token)
		if err != nil {
			log.Error("TokenChecker: ", ErrInvalidOrExpiredAccessToken)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": ErrInvalidOrExpiredAccessToken.Error(),
			})
		}

		c.Set("userID", userID)
		c.Next()
	}
}
