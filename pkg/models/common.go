package models

import (
	"errors"
	"net/http"
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
	Id     *int
}

func NewCommonQueryParamsFromRequest(r *http.Request) (*CommonQueryParams, error) {
	qResult := CommonQueryParams{
		Limit:  LimitDefault,
		Offset: OffsetDefault,
	}

	qRequest := r.URL.Query()

	if qRequest.Has("limit") {
		err := qResult.validateLimit(qRequest)
		if err != nil {
			return nil, err
		}
	}
	if qRequest.Has("offset") {
		err := qResult.validateOffset(qRequest)
		if err != nil {
			return nil, err
		}
	}
	if qRequest.Has("id") {
		err := qResult.validateId(qRequest)
		if err != nil {
			return nil, err
		}
	}

	return &qResult, nil
}

func (a *CommonQueryParams) validateLimit(q url.Values) error {
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

func (a *CommonQueryParams) validateOffset(q url.Values) error {
	offset, err := strconv.Atoi(q.Get("offset"))
	if err != nil {
		return err
	}

	if offset > 0 {
		a.Offset = offset
	}

	return nil
}

func (a *CommonQueryParams) validateId(q url.Values) error {
	id, err := strconv.Atoi(q.Get("id"))
	if err != nil {
		return err
	}

	if id <= 0 {
		return errors.New(`"id" must be a positive integer`)
	}

	a.Id = &id

	return nil
}
