package model

// 用户表
type Users struct {
	ID        int    `gorm:"primaryKey;column:id" json:"id"`      // 主键
	Name      string `gorm:"column:name" json:"name"`             // 名称
	Age       int    `gorm:"column:age" json:"age"`               // 年龄
	CardNo    string `gorm:"column:card_no" json:"card_no"`       // 身份证
	HeadImg   string `gorm:"column:head_img" json:"head_img"`     // 头像
	CreatedAt int    `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	DeletedAt *int   `gorm:"column:deleted_at" json:"deleted_at"` // 删除 时间
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}
