package docs

import "time"

type SignInResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpResponse struct {
	ID int64 `json:"id"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type RefreshSessionDTO struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CreateUpdateNotebookDTO struct {
	Description string `json:"description"`
}

type CreateNotebookResponse struct {
	ID int64 `json:"id"`
}

type OkStatusResponse struct {
	Status string `json:"status" example:"ok"`
}

type CreateNoteResponse struct {
	ID int64 `json:"id"`
}

type CreateSharedNoteResponse struct {
	ID int64 `json:"id"`
}

type CreateSharedTodoListResponse struct {
	ID int64 `json:"id"`
}

type GetAllNotebooksResponse struct {
	Notebooks []struct {
		ID          int64   `json:"id"`
		Description *string `json:"description"`
	} `json:"notebooks,omitempty"`
	Count int `json:"count"`
}

type GetAllNotesResponse struct {
	Notes []struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"notes,omitempty"`
	Count int `json:"count"`
}

type GetNoteResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTodoListResponse struct {
	ID int64 `json:"id"`
}

type GetAllTodoListsResponse struct {
	TodoLists []struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"todo_lists,omitempty"`
	Count int `json:"count"`
}

type CreateTodoItemResponse struct {
	ID int64 `json:"id"`
}

type GetAllTodoItemsResponse struct {
	TodoItems []struct {
		ID   int64  `json:"id"`
		Body string `json:"body"`
		Done bool   `json:"done"`
	} `json:"todo_items,omitempty"`
	Count int `json:"count"`
}
