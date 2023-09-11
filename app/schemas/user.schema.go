package schemas

type SignIn struct {
	Signature string `json:"signature" validate:"required"`
	Message   string `json:"message" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

type UserSignInResponse struct {
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
	Success     bool   `json:"success"`
}
