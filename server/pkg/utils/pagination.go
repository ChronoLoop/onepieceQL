package utils

type PaginationInfo struct {
	Count int  `json:"count"`
	Pages int  `json:"pages"`
	Next  *int `json:"next"`
	Prev  *int `json:"prev"`
}

const PAGE_LIMIT = 40

func CreatePaginationInfo(count int, page int) PaginationInfo {
	totalPages := (count + PAGE_LIMIT - 1) / PAGE_LIMIT
	var nextPage, prevPage *int

	if page < totalPages {
		nextPageValue := page + 1
		nextPage = &nextPageValue
	}
	if page > 1 {
		prevPageValue := page - 1
		prevPage = &prevPageValue
	}

	return PaginationInfo{
		Count: count,
		Pages: totalPages,
		Prev:  prevPage,
		Next:  nextPage,
	}
}
