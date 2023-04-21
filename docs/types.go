package docs

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

type GetAllNotebooksResponse struct {
	Notebooks []struct {
		ID          int64   `json:"id"`
		Description *string `json:"description"`
	} `json:"notebooks,omitempty"`
	Count int `json:"count"`
}
