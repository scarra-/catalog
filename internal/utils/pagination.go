package utils

type Pagination[T any] struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	Rows       []*T  `json:"rows"`
	sort       string
}

func (p *Pagination[T]) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination[T]) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination[T]) GetSort() string {
	if p.sort == "" {
		p.sort = "created_at desc"
	}
	return p.sort
}

type CursorPagination[T any] struct {
	Limit int    `json:"limit"`
	Rel   string `json:"rel"`
	Next  string `json:"next"`
	Rows  []*T   `json:"rows"`
}

func (p *CursorPagination[T]) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
