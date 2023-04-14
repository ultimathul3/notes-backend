package domain

import (
	"context"
	"errors"
	"regexp"
	"unicode/utf8"
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

	GetUserIDDTO struct {
		Login    *string `json:"login"`
		Password *string `json:"password"`
	}
)

type UserUsecase interface {
	Create(ctx context.Context, input CreateUserDTO) (int64, error)
	GetID(ctx context.Context, input GetUserIDDTO) (int64, error)
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	Create(ctx context.Context, user User) (int64, error)
	GetID(ctx context.Context, login, passwordHash string) (int64, error)
}

func (cu *CreateUserDTO) Validate() error {
	if cu.Login == nil {
		return errors.New("empty login")
	}
	if cu.Name == nil {
		return errors.New("empty name")
	}
	if cu.Password == nil {
		return errors.New("empty password")
	}
	if utf8.RuneCountInString(*cu.Login) < 4 || utf8.RuneCountInString(*cu.Login) > 15 {
		return errors.New("login length must be from 4 to 15 characters")
	}
	if utf8.RuneCountInString(*cu.Name) < 2 || utf8.RuneCountInString(*cu.Name) > 50 {
		return errors.New("name length must be from 2 to 50 characters")
	}
	if utf8.RuneCountInString(*cu.Password) < 6 || utf8.RuneCountInString(*cu.Password) > 256 {
		return errors.New("password length must be from 6 to 256 characters")
	}
	if regexp.MustCompile(`^\d+$`).MatchString(*cu.Login) {
		return errors.New("login should not consist only of numbers")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(*cu.Login) {
		return errors.New("login must contain only numbers and letters")
	}
	if !regexp.MustCompile(`^[a-zA-Zа-яА-Я]+$`).MatchString(*cu.Name) {
		return errors.New("name must contain only letters")
	}
	return nil
}

func (gu *GetUserIDDTO) Validate() error {
	if gu.Login == nil {
		return errors.New("empty login")
	}
	if gu.Password == nil {
		return errors.New("empty password")
	}
	return nil
}

var (
	ErrUserAlreadyExists      = errors.New("user with such login already exists")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
)
