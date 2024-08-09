package vo

/**
 * @Author elastic·H
 * @Date 2024-08-08
 * @File: user.vo.go
 * @Description:
 */

type UserSingle struct {
	Id    int32                  `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name  string                 `json:"name" gorm:"column:name"`
	Phone string                 `json:"phone" gorm:"column:phone"`
	Age   int                    `json:"age" gorm:"column:age"`
	Extra map[string]interface{} `json:"extra" gorm:"column:extra;serializer:json"`
}
