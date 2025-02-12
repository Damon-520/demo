package consts

type PageInfo struct { // List接口返回的分页信息
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type Pagination struct { // List接口的分页参数
	Page  int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

func DefaultPagination(p *Pagination) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 50 {
		p.Limit = 10
	}
}
