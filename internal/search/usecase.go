package search

import (
	"context"

	"github.com/ultimathul3/notes-backend/internal/domain"
)

type Usecase struct {
	repo domain.SearchRepository
}

func NewUsecase(repo domain.SearchRepository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) GetAll(ctx context.Context, userID int64, search domain.Search) (domain.SearchResult, error) {
	result := domain.SearchResult{}

	if search.ByNotes {
		notes, err := u.repo.GetAllNotes(ctx, userID, search)
		if err != nil {
			return domain.SearchResult{}, err
		}
		result.Notes = notes
		result.NotesCount = len(notes)
	}

	if search.ByTodoLists {
		lists, err := u.repo.GetAllTodoLists(ctx, userID, search)
		if err != nil {
			return domain.SearchResult{}, err
		}
		result.TodoLists = lists
		result.TodoListsCount = len(lists)
	}

	if search.BySharedNotes {
		notes, err := u.repo.GetAllAcceptedSharedNotes(ctx, userID, search)
		if err != nil {
			return domain.SearchResult{}, err
		}
		result.SharedNotes = notes
		result.SharedNotesCount = len(notes)
	}

	if search.BySharedTodoLists {
		lists, err := u.repo.GetAllAcceptedSharedTodoLists(ctx, userID, search)
		if err != nil {
			return domain.SearchResult{}, err
		}
		result.SharedTodoLists = lists
		result.SharedTodoListsCount = len(lists)
	}

	return result, nil
}
