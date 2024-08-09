package dto

type QueryPagination struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=1,max=100"`

	// TODO
	Sort      string `form:"sort"`
	Order     string `form:"order"`
	SortOrder string `form:"sortOrder"`
}

type BodyId struct {
	Id int `form:"id" binding:"required,min=1"`
}
