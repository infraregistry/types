package types
type Session struct {
	ID           string `json:"id"`
	Organization string `json:"organization"`
	User         string `json:"user"`
}
type SessionLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}

type SessionLoginResponse struct {
	Token string `json:"token"`
}
