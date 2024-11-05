package types

import (
	"time"

	"github.com/steebchen/prisma-client-go/runtime/types"
)

type Change struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type ChangeCreate struct {
	ComponentID string     `json:"componentId" validate:"required,min=1,max=25"`
	Name        string     `json:"name" validate:"required,min=1,max=32"`
	Value       string     `json:"value" validate:"required,min=1,max=255"`
	Type        string     `json:"type" validate:"required,min=1,max=32"`
	Content     types.JSON `json:"content" validate:"required,min=1,max=99999"`
}
