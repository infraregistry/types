package types

import (
	"time"

	"github.com/steebchen/prisma-client-go/runtime/types"
)

type ChangeType string

const (
	ChangeTypeSystemChange ChangeType = "system"
	ChangeTypeUserChange   ChangeType = "user"
)

type Change struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Created     time.Time   `json:"created"`
	Updated     time.Time   `json:"updated"`
	ComponentID string      `json:"componentId"`
	Type        *ChangeType `json:"type"`
	From        string      `json:"from"`
	To          string      `json:"to"`
}

type ChangeCreate struct {
	ComponentID string     `json:"componentId" validate:"required,min=1,max=25"`
	Name        string     `json:"name" validate:"required,min=1,max=32"`
	Value       string     `json:"value" validate:"required,min=1,max=255"`
	Type        ChangeType `json:"type" validate:"required"`
	Content     types.JSON `json:"content" validate:"required,min=1,max=99999"`
}
