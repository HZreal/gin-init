package entity

type UserModel struct {
	Id    int32                  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name  string                 `gorm:"column:name"`
	Phone string                 `gorm:"column:phone"`
	Age   int                    `gorm:"column:age"`
	Extra map[string]interface{} `gorm:"column:extra;serializer:json"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (t *UserModel) TableName() string {
	return "tb_user"
}
