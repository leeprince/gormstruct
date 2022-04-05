package model

// 用户表
type Users struct {
	ID        int64   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                          // 主键
	Name      *string `gorm:"column:name;type:varchar(255);is null;" json:"name"`                    // 名称
	Age       int64   `gorm:"column:age;type:int(11);not null;DEFAULT '18'" json:"age"`              // 年龄
	CardNo    string  `gorm:"column:card_no;type:varchar(18);not null;DEFAULT ''" json:"card_no"`    // 身份证
	HeadImg   string  `gorm:"column:head_img;type:varchar(255);not null;DEFAULT ''" json:"head_img"` // 头像
	CreatedAt int64   `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`            // 创建时间
	UpdatedAt int64   `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`            // 更新时间
	DeletedAt int64   `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"` // 删除时间
}

// 获取结构体对应的表名方法
func (m *Users) TableName() string {
	return "users"
}
