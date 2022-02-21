package models

import (
	"net/url"
	"strconv"
)

const (
	LimitDefault  = 5
	LimitMax      = 100
	OffsetDefault = 0
)

type CommonQueryParams struct {
	Offset int
	Limit  int
}

func (a *CommonQueryParams) ValidateLimit(q url.Values) error {
	limit, err := strconv.Atoi(q.Get("limit"))
	if err != nil {
		return err
	}

	if limit > LimitMax {
		a.Limit = LimitMax
	} else if limit > 0 {
		a.Limit = limit
	}

	return nil
}

func (a *CommonQueryParams) ValidateOffset(q url.Values) error {
	offset, err := strconv.Atoi(q.Get("offset"))
	if err != nil {
		return err
	}

	if offset > 0 {
		a.Offset = offset
	}

	return nil
}
