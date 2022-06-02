package common

import "math"

type Pagination struct {
	Total     int64 `json:"total"`
	Offset    int64 `json:"offset"`
	Limit     int64 `json:"limit"`
	Page      int64 `json:"page"`
	TotalPage int64 `json:"totalPage"`
	PageSize  int64 `json:"pageSize"`
}

func GetPagination(total, page, pageSize int64) Pagination {
	maxPage := math.Ceil(float64(total) / float64(pageSize))
	//if page > int64(maxPage) {
	//	page = int64(maxPage)
	//}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * pageSize

	limit := offset + pageSize
	if limit > total {
		limit = total
	}

	return Pagination{
		Total:     total,
		Offset:    offset,
		Limit:     limit,
		Page:      page,
		TotalPage: int64(maxPage),
		PageSize:  pageSize,
	}
}
