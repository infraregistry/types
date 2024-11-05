package types

import "time"

type Tag struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type TagCreate struct {
	Name string `json:"name" validate:"required,min=1,max=32"`
}
