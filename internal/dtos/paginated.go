package dtos

import "strconv"

type PaginatedResponseDto struct {
	Total int64       `json:"total_count"`
	Data  interface{} `json:"data"`
}

func NewPaginatedResponse(total int64, data interface{}) *PaginatedResponseDto {
	return &PaginatedResponseDto{
		Total: total,
		Data:  data,
	}
}

type PaginationOpts struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type PaginationQueryDto struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
	page  int64
	limit int64
}

func (p PaginationQueryDto) Validate() error {
	page, err := strconv.ParseInt(p.Page, 10, 64)
	if err != nil {
		return err
	}

	p.page = page

	limit, err := strconv.ParseInt(p.Limit, 10, 64)
	if err != nil {
		return err
	}

	if limit == 0 || limit < 0 || limit > 10 {
		limit = 10
	}

	p.limit = limit

	return nil
}

func FromPaginationQueryDtoToPaginationOpts(query *PaginationQueryDto) *PaginationOpts {
	return &PaginationOpts{
		Page:  query.page,
		Limit: query.limit,
	}
}
