package types

type OrganizationCreate struct {
	Name        string `json:"name" validate:"required,min=1,max=32"`
	Description string `json:"description" validate:"max=255"`
}

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
