package jwtauth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	accessTokenTTL time.Duration
	secretKey      string
}

func NewJWT(accessTokenTTL time.Duration, secretKey string) *JWT {
	return &JWT{
		accessTokenTTL: accessTokenTTL,
		secretKey:      secretKey,
	}
}

func (j *JWT) GenerateTokens(userID int64) (string, uuid.UUID, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenTTL)),
		Subject:   fmt.Sprintf("%d", userID),
	})

	signedAccesToken, err := accessToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", uuid.Nil, err
	}

	refreshToken := uuid.New()

	return signedAccesToken, refreshToken, nil
}

func (j *JWT) ParseAccessToken(accessToken string) (int64, error) {
	var sub string

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, ok = claims["sub"].(string)
		if !ok {
			return 0, errors.New("missing sub field")
		}
	}

	return strconv.ParseInt(sub, 10, 64)
}
