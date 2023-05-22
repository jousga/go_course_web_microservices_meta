package meta

import (
	"strconv"
)

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

func New(page, perPage, total int, pageLimitDefault string) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pageLimitDefault)

		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
	}, nil
}

// A partir de que número de resultados tenemos que devolver
func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// hasta que número de resultados tenemos que devolver
func (p *Meta) Limit() int {
	return p.PerPage
}
