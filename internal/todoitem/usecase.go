package todoitem

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.TodoItemRepository
}

func NewUsecase(repo domain.TodoItemRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, userID, notebookID, listID int64, input domain.CreateTodoItemDTO) (int64, error) {
	if err := input.Validate(); err != nil {
		return 0, err
	}

	return u.repo.Create(ctx, userID, notebookID, domain.TodoItem{
		Body:       *input.Body,
		Done:       false,
		TodoListID: listID,
	})
}

func (u *Usecase) GetAllByListID(ctx context.Context, userID, notebookID, listID int64) ([]domain.TodoItem, error) {
	return u.repo.GetAllByListID(ctx, userID, notebookID, listID)
}

func (u *Usecase) Patch(ctx context.Context, itemID, userID, notebookID, listID int64, input domain.PatchTodoItemDTO) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return u.repo.Patch(ctx, itemID, userID, notebookID, listID, input)
}

func (u *Usecase) Delete(ctx context.Context, itemID, userID, notebookID, listID int64) error {
	return u.repo.Delete(ctx, itemID, userID, notebookID, listID)
}
