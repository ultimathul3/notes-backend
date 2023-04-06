package domain

import (
	"context"
	"errors"
	"regexp"
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

type UserUsecase interface {
	Create(ctx context.Context, user *CreateUserDTO) (int64, error)
}

//go:generate mockery --name UserRepository
type UserRepository interface {
	Create(ctx context.Context, user *User) (int64, error)
	GetByID(ctx context.Context, id int64) (*User, error)
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
	if len(*cu.Login) < 4 || len(*cu.Login) > 15 {
		return errors.New("login length must be from 4 to 15 characters")
	}
	if len(*cu.Name) < 2 || len(*cu.Name) > 50 {
		return errors.New("name length must be from 2 to 50 characters")
	}
	if len(*cu.Password) < 6 || len(*cu.Password) > 256 {
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

var (
	ErrUserAlreadyExists = errors.New("user with such login already exists")
)
