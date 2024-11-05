package types

type Search struct {
	ID     string   `json:"id" validate:"omitempty,min=35,max=35"`
	Status []Status `json:"status" validate:"omitempty,oneof=active inactive"`
	Terms  string   `json:"terms" validate:"omitempty,min=1,max=100"`
	Limit  int      `json:"limit" validate:"omitempty,min=1,max=100"`
	Offset int      `json:"offset" validate:"omitempty,min=0"`
	Sort   string   `json:"sort" validate:"omitempty,oneof=name description created updated"`
	Order  string   `json:"order" validate:"omitempty,oneof=asc desc"`
}

func (s *Search) GetLimit() int {
	if s.Limit == 0 {
		return 100
	}
	return s.Limit
}

func (s *Search) GetOffset() int {
	if s.Offset == 0 {
		return 0
	}
	return s.Offset
}

func (s *Search) GetSort() string {
	if s.Sort == "" {
		return "created"
	}
	return s.Sort
}
func (s *Search) GetOrder() string {
	if s.Order == "" {
		return "desc"
	}
	return s.Order
}
