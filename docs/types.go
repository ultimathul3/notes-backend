package docs

type SignInResponse struct {
	ID           int64  `json:"id"`
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
	UserID       int64  `json:"user_id"`
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

type GetAllNotebooksResponse struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
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
