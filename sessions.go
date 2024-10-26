package types

type Session struct {
	ID           string `json:"id"`
	Organization string `json:"organization"`
	User         string `json:"user"`
}
