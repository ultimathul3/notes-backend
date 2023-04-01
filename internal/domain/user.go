package domain

import "context"

type User struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (User, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (User, error)
}
