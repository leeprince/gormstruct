// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserInfo = "user_info"

// UserInfo mapped from table <user_info>
type UserInfo struct {
	ID        int64          `gorm:"column:id;primaryKey;comment:用户ID" json:"id"`                                // 用户ID
	Username  string         `gorm:"column:username;not null;comment:用户名" json:"username"`                       // 用户名
	Password  string         `gorm:"column:password;not null;comment:密码" json:"password"`                        // 密码
	Email     *string        `gorm:"column:email;comment:邮箱" json:"email"`                                       // 邮箱
	Phone     *string        `gorm:"column:phone;comment:手机号" json:"phone"`                                      // 手机号
	Nickname  *string        `gorm:"column:nickname;comment:昵称" json:"nickname"`                                 // 昵称
	Avatar    *string        `gorm:"column:avatar;comment:头像" json:"avatar"`                                     // 头像
	CreatedAt *time.Time     `gorm:"column:created_at;comment:创建时间" json:"created_at"`                           // 创建时间
	UpdatedAt *time.Time     `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName UserInfo's table name
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}
