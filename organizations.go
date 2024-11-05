package types

type OrganizationCreate struct {
	Name        string `json:"name" validate:"required,min=1,max=32"`
	Description string `json:"description" validate:"max=255"`
}
