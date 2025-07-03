package pagination

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	HasNext  bool `json:"has_next"`
}

func GetPagination(r *http.Request) Pagination {
	q := r.URL.Query()

	page, err := strconv.Atoi(q.Get("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(q.Get("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

func (p Pagination) GetLimitOffset(page Pagination) (int, int) {
	limit := page.PageSize + 1
	offset := (page.Page - 1) * page.PageSize

	return limit, offset
}
