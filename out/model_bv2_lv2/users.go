package model_bv2_lv2

// 用户表
type Users struct {
	ID      int    `gorm:"primaryKey;column:id" json:"id"`  // 主键
	Name    string `gorm:"column:name" json:"name"`         // 名称
	Age     int    `gorm:"column:age" json:"age"`           // 年龄
	CardNo  string `gorm:"column:card_no" json:"card_no"`   // 身份证
	HeadImg string `gorm:"column:head_img" json:"head_img"` // 头像
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}
