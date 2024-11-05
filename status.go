package types

type Status string

const (
	StatusActive  Status = "active"
	StatusDraft   Status = "draft"
	StatusDeleted Status = "deleted"
)

func (s Status) String() string {
	return string(s)
}
