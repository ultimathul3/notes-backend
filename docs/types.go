package docs

type SignInResponse struct {
	ID           int64  `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignUpResponse struct {
	ID int64 `json:"id"`
}

type Message struct {
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
