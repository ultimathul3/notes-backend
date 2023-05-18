package sharedtodolist

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.SharedTodoListRepository
}

func NewUsecase(repo domain.SharedTodoListRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, whoseID, whomID, noteID int64) (int64, error) {
	if whoseID == whomID {
		return 0, domain.ErrImpossibleToShareNoteWithYourself
	}

	return u.repo.Create(ctx, domain.SharedTodoList{
		WhoseID:    whoseID,
		WhomID:     whomID,
		TodoListID: noteID,
		Accepted:   false,
	})
}

func (u *Usecase) Delete(ctx context.Context, id, whomID int64) error {
	return u.repo.Delete(ctx, id, whomID)
}

func (u *Usecase) GetAllInfo(ctx context.Context, whomID int64) ([]domain.SharedTodoListInfo, error) {
	return u.repo.GetAllInfo(ctx, whomID)
}

func (u *Usecase) Accept(ctx context.Context, id, whomID int64) error {
	return u.repo.Accept(ctx, id, whomID)
}

func (u *Usecase) GetDataByID(ctx context.Context, id, whomID int64) (domain.SharedTodoListData, error) {
	createdAt, updatedAt, err := u.repo.GetTimestampsByID(ctx, id, whomID)
	if err != nil {
		return domain.SharedTodoListData{}, err
	}

	items, err := u.repo.GetItemsByID(ctx, id, whomID)
	if err != nil {
		return domain.SharedTodoListData{}, err
	}

	return domain.SharedTodoListData{
		Items:     items,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (u *Usecase) GetOutgoingInfoByTodoListID(ctx context.Context, noteID, whoseID int64) ([]domain.OutgoingSharedTodoListInfo, error) {
	return u.repo.GetOutgoingInfoByTodoListID(ctx, noteID, whoseID)
}
