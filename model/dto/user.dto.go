package dto

type UserCreateDTO struct {
	Name  string                 `json:"name" binding:"required,min=5"`
	Phone string                 `json:"phone" binding:"required,min=5"`
	Age   string                 `json:"age" binding:"required,min=0"`
	Extra map[string]interface{} `json:"extra"`
}

type UserListFilterDTO struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   string `json:"age"`
}
