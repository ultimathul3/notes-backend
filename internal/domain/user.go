package domain

import (
	"context"
	"errors"
)

type (
	User struct {
		ID           int64  `json:"id"`
		Login        string `json:"login"`
		Name         string `json:"name"`
		PasswordHash string `json:"password_hash"`
	}

	CreateUserDTO struct {
		Login    *string `json:"login"`
		Name     *string `json:"name"`
		Password *string `json:"password"`
	}
)

type (
	UserUsecase interface {
		Create(ctx context.Context, user *CreateUserDTO) (int64, error)
	}

	UserRepository interface {
		Create(ctx context.Context, user *User) (int64, error)
		GetByID(ctx context.Context, id int64) (*User, error)
	}
)

var (
	ErrUserAlreadyExists = errors.New("user with such login already exists")
)
