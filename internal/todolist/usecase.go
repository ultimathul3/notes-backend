package todolist

import (
	"context"
	"time"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.TodoListRepository
}

func NewUsecase(repo domain.TodoListRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, userID, notebookID int64, input domain.CreateTodoListDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	return u.repo.Create(ctx, domain.TodoList{
		Title:      *input.Title,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		UserID:     userID,
		NotebookID: notebookID,
	})
}

func (u *Usecase) Update(ctx context.Context, todoListID, userID, notebookID int64, input domain.UpdateTodoListDTO) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return u.repo.Update(ctx, todoListID, userID, notebookID, input)
}

func (u *Usecase) GetAllByNotebookID(ctx context.Context, userID, notebookID int64) ([]domain.TodoList, error) {
	return u.repo.GetAllByNotebookID(ctx, userID, notebookID)
}

func (u *Usecase) Delete(ctx context.Context, todoListID, userID, notebookID int64) error {
	return u.repo.Delete(ctx, todoListID, userID, notebookID)
}
