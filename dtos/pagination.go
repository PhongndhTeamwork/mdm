package dtos

type PaginationQuery struct {
	Page int `form:"page"`
	Take int `form:"take"`
}

type ReturningPagination struct {
	Total       int64       `json:"total"`        // total items
	PerPage     int         `json:"per_page"`     // items per page
	CurrentPage int         `json:"current_page"` // current page number
	TotalPages  int         `json:"total_pages"`  // total number of pages
	Data        interface{} `json:"data"`         // paginated data
}
