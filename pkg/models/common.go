package models

const (
	LimitDefault  = 5
	LimitMax      = 100
	OffsetDefault = 0
)

type CommonQueryParams struct {
	Offset int
	Limit  int
}
