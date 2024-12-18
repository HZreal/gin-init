package entity

type TbUser struct {
	Id       int                    `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL" json:"id" binding:"omitempty"`
	Username string                 `gorm:"column:username" json:"username" binding:"omitempty,min=5,max=20,alphanum"`
	Password string                 `gorm:"column:password" json:"password" binding:"omitempty,min=6,max=30"`
	Phone    string                 `gorm:"column:phone" json:"phone" binding:"omitempty,len=11,numeric"`
	Age      int                    `gorm:"column:age" json:"age" binding:"omitempty,min=0,max=120"`
	Status   bool                   `gorm:"column:status" json:"status" binding:"omitempty"`
	Extra    map[string]interface{} `gorm:"column:extra;serializer:json" json:"extra" binding:"omitempty"`
}

func NewUserModel() *TbUser {
	return &TbUser{}
}

func (t *TbUser) TableName() string {
	return "tb_user"
}
