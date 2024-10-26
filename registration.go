package types

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

type RegisterResponse struct {
	ID string
}
