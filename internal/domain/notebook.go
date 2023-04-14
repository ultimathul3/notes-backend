package domain

import (
	"context"
	"errors"
)

type (
	Notebook struct {
		ID          int64   `json:"id"`
		Description *string `json:"description"`
		UserID      int64   `json:"user_id"`
	}
)

type NotebookUsecase interface {
	Create(ctx context.Context, notebook Notebook) (int64, error)
}

type NotebookRepository interface {
	Create(ctx context.Context, notebook Notebook) (int64, error)
}

func (cn *Notebook) Validate() error {
	if len(*cn.Description) == 0 || len(*cn.Description) > 64 {
		return errors.New("description length must be from 1 to 64 characters")
	}
	if cn.Description == nil {
		return errors.New("empty description")
	}
	return nil
}
