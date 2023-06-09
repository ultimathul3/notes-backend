package sharednote

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.SharedNoteRepository
}

func NewUsecase(repo domain.SharedNoteRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, whoseID, whomID, noteID int64) (int64, error) {
	if whoseID == whomID {
		return 0, domain.ErrImpossibleToShareNoteWithYourself
	}

	return u.repo.Create(ctx, domain.SharedNote{
		WhoseID:  whoseID,
		WhomID:   whomID,
		NoteID:   noteID,
		Accepted: false,
	})
}

func (u *Usecase) Delete(ctx context.Context, id, whomID int64) error {
	return u.repo.Delete(ctx, id, whomID)
}

func (u *Usecase) GetAllInfo(ctx context.Context, whomID int64) ([]domain.SharedNoteInfo, error) {
	return u.repo.GetAllInfo(ctx, whomID)
}

func (u *Usecase) Accept(ctx context.Context, id, whomID int64) error {
	return u.repo.Accept(ctx, id, whomID)
}

func (u *Usecase) GetDataByID(ctx context.Context, id, whomID int64) (domain.SharedNoteData, error) {
	return u.repo.GetDataByID(ctx, id, whomID)
}

func (u *Usecase) GetOutgoingInfoByNoteID(ctx context.Context, noteID, whoseID int64) ([]domain.OutgoingSharedNoteInfo, error) {
	return u.repo.GetOutgoingInfoByNoteID(ctx, noteID, whoseID)
}
