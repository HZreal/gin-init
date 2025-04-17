package types

type PaginationResult struct {
	Total       int         `json:"total"`
	Pages       int         `json:"pages"`
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Records     interface{} `json:"records"`
}

// PageResult 分页结果
type PageResult[T any] struct {
	Total    int64 `json:"total"`    // 总数
	Page     int   `json:"page"`     // 当前页码
	PageSize int   `json:"pageSize"` // 每页数量
	Items    []T   `json:"items"`    // 数据列表
}
