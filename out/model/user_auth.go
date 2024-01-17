package model

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-17 18:49:00
 * @Desc:   user_auth 表
 */

//
type UserAuth struct {
	ID         int64      `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`             // 主键
	UserID     int64      `gorm:"column:user_id;type:int(11);not null;DEFAULT:0" json:"user_id"`         // 用户ID
	ExpireTime time.Time  `gorm:"column:expire_time;type:datetime;is null;" json:"expire_time"`          // 授权过期时间
	AccessTime int64      `gorm:"column:access_time;type:int(11);not null;DEFAULT:0" json:"access_time"` // 访问时间
	UpdatedAt  time.Time  `gorm:"column:updated_at;type:datetime(3);is null;" json:"updated_at"`         // 更新时间
	CreatedAt  time.Time  `gorm:"column:created_at;type:datetime(3);is null;" json:"created_at"`         // 创建时间
	DeletedAt  *time.Time `gorm:"column:deleted_at;type:datetime(3);is null;" json:"deleted_at"`         // 删除时间
}

// 获取结构体对应的表名方法
func (m *UserAuth) TableName() string {
	return "user_auth"
}

// 实例化结构体对象
func NewUserAuth() *UserAuth {
	return &UserAuth{}
}

// 获取主键的对应字段
func (m *UserAuth) PrimaryKey() string {
	return UserAuthColumns.ID
}

// 获取主键值
func (m *UserAuth) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UserAuthColumns = struct {
	ID         string
	UserID     string
	ExpireTime string
	AccessTime string
	UpdatedAt  string
	CreatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	UserID:     "user_id",
	ExpireTime: "expire_time",
	AccessTime: "access_time",
	UpdatedAt:  "updated_at",
	CreatedAt:  "created_at",
	DeletedAt:  "deleted_at",
}

// 包含所有表字段的切片
var UserAuthAllColumns = []string{
	UserAuthColumns.ID,         // 主键
	UserAuthColumns.UserID,     // 用户ID
	UserAuthColumns.ExpireTime, // 授权过期时间
	UserAuthColumns.AccessTime, // 访问时间
	UserAuthColumns.UpdatedAt,  // 更新时间
	UserAuthColumns.CreatedAt,  // 创建时间
	UserAuthColumns.DeletedAt,  // 删除时间

}

// 设置值：主键
func (m *UserAuth) SetID(v int64) {
	m.ID = v
}

// 设置值：用户ID
func (m *UserAuth) SetUserID(v int64) {
	m.UserID = v
}

// 设置值：授权过期时间
func (m *UserAuth) SetExpireTime(v time.Time) {
	m.ExpireTime = v
}

// 设置值：访问时间
func (m *UserAuth) SetAccessTime(v int64) {
	m.AccessTime = v
}

// 设置值：更新时间
func (m *UserAuth) SetUpdatedAt(v time.Time) {
	m.UpdatedAt = v
}

// 设置值：创建时间
func (m *UserAuth) SetCreatedAt(v time.Time) {
	m.CreatedAt = v
}

// 设置值：删除时间
func (m *UserAuth) SetDeletedAt(v *time.Time) {
	m.DeletedAt = v
}

// 获取值：主键
func (m *UserAuth) GetID() int64 {
	return m.ID
}

// 获取值：用户ID
func (m *UserAuth) GetUserID() int64 {
	return m.UserID
}

// 获取值：授权过期时间
func (m *UserAuth) GetExpireTime() time.Time {
	return m.ExpireTime
}

// 获取值：访问时间
func (m *UserAuth) GetAccessTime() int64 {
	return m.AccessTime
}

// 获取值：更新时间
func (m *UserAuth) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// 获取值：创建时间
func (m *UserAuth) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// 获取值：删除时间
func (m *UserAuth) GetDeletedAt() *time.Time {
	return m.DeletedAt
}
