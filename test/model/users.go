package model

// 用户表
type Users struct {
	ID        int64  `gorm:"primaryKey;column:id" json:"id"`      // 主键
	Name      string `gorm:"column:name" json:"name"`             // 名称
	Age       int64  `gorm:"column:age" json:"age"`               // 年龄
	CardNo    string `gorm:"column:card_no" json:"card_no"`       // 身份证
	HeadImg   string `gorm:"column:head_img" json:"head_img"`     // 头像
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at" json:"deleted_at"` // 删除 时间
}

// 获取结构体对应的表名方法
func (m *Users) TableName() string {
	return "users"
}

// 表字段的映射
var UsersColumns = struct {
	ID        string
	Name      string
	Age       string
	CardNo    string
	HeadImg   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Age:       "age",
	CardNo:    "card_no",
	HeadImg:   "head_img",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}
