package types

type QueryPagination struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=1,max=100"`

	// TODO
	Sort      string `form:"sort"`
	Order     string `form:"order"`
	SortOrder string `form:"sortOrder"`
}

// PageInfo 分页请求参数
type PaginationType struct {
	Page     int `json:"page,default=1"`
	PageSize int `json:"pageSize,default=10"`
}

type SorterItem struct {
	Field string `json:"field,optional"`
	Order string `json:"order,optional"`
}

type ListFilterType struct {
	Filter     map[string]interface{} `json:"filter,optional"`
	Pagination PaginationType         `json:"pagination"`
	Sorter     []SorterItem           `json:"sorter,optional"`
}

type BodyJsonId struct {
	Id int `json:"id" binding:"required,min=1"`
}

type QueryId struct {
	Id uint `form:"id" binding:"required,min=1"`
}
