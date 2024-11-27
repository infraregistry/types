package types

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

type RegisterConfirmRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
type RegisterConfirmResponse struct {
	Token string `json:"token"`
}
