package jwtauth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParseValidToken(t *testing.T) {
	jwt := NewJWT(10*time.Second, "secret")

	expectedID := int64(123)
	accessToken, refreshToken, err := jwt.GenerateTokens(expectedID)
	if err != nil {
		t.Error(err)
	}

	if refreshToken == uuid.Nil {
		t.Error("not expected nil uuid")
	}

	resultID, err := jwt.ParseAccessToken(accessToken)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expectedID, resultID)
}

func TestParseExpiredToken(t *testing.T) {
	jwt := NewJWT(-10*time.Second, "secret")

	accessToken, refreshToken, err := jwt.GenerateTokens(123)
	if err != nil {
		t.Error(err)
	}

	if refreshToken == uuid.Nil {
		t.Error("not expected nil uuid")
	}

	_, err = jwt.ParseAccessToken(accessToken)

	assert.ErrorContains(t, err, "token is expired")
}

func TestParseInvalidToken(t *testing.T) {
	jwt1 := NewJWT(10*time.Second, "secret")

	accessToken, refreshToken, err := jwt1.GenerateTokens(123)
	if err != nil {
		t.Error(err)
	}

	if refreshToken == uuid.Nil {
		t.Error("not expected nil uuid")
	}

	jwt2 := NewJWT(10*time.Second, "another secret")

	_, err = jwt2.ParseAccessToken(accessToken)

	assert.ErrorContains(t, err, "signature is invalid")
}

func TestParseTokenWithoutSubField(t *testing.T) {
	// HS256
	// Payload
	// {
	//	"jti": "1234567890",
	//	"name": "John Doe",
	//	"iat": 22222222222
	// }
	accessToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.hRkrwwQiL33cjtZgK6H4pLY3wAPHMkYDQlutm6vOiXg"
	jwt := NewJWT(10*time.Second, "secret")

	_, err := jwt.ParseAccessToken(accessToken)

	assert.ErrorContains(t, err, "missing sub field")
}

func TestParseTokenWithUnexpectedSigningMethod(t *testing.T) {
	// RS256
	accessToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJqdGkiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoyMjIyMjIyMjIyMn0.r4qNH3w7Pa5h_9XoYylgrUavPQPUn6GvcUtLnR7Du_PfO9F_-B66-ZvhrfcIyL63bARJ35hrHFnemWzJBo9ZMJ_qTxmVLAc9OUHBHiml0aBG3Pi8fRpPXWe-_6IwMlyp8EqS1rKdjUsRj32eELD19n7lNj2DdGD0IykdG9ixDiCvykE8UMCKdO6ObweuCfvYyx-cKoiNHZzrXWUXUA-ypXME9o3lwArqxxkrf1-xlG4fnLOcPLhkE-Wfl-FVDzCILPmbjm5n4wWnwxQ65dLehVpPcXO3ehG438dPZitm6aX9IMIB6GcsUKdY3FvoIHi1PRJX68tSRzYED1CJox6qOg"
	jwt := NewJWT(10*time.Second, "secret")

	_, err := jwt.ParseAccessToken(accessToken)

	assert.ErrorContains(t, err, "unexpected signing method")
}
