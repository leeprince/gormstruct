package model

import (
	"gorm.io/datatypes"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-21 15:28:01
 * @Desc:   user_auth_provider 表
 */

// 第三方授权登录用户表
type UserAuthProvider struct {
	ID                  int64          `gorm:"column:id;primaryKey;type:bigint(20);autoIncrement" json:"id"`                // 主键ID
	UserID              int64          `gorm:"column:user_id;type:bigint(20);not null;" json:"user_id"`                     // 用户ID,user_info的主键
	Provider            string         `gorm:"column:provider;type:varchar(32);not null;" json:"provider"`                  // 第三方登录提供商。wechatopen：微信开放平台；
	ProviderUserID      string         `gorm:"column:provider_user_id;type:varchar(128);not null;" json:"provider_user_id"` // 第三方登录提供商用户ID
	ProviderUserUnionid string         `gorm:"column:provider_user_unionid;type:varchar(128);not null;DEFAULT:" json:"provider_user_unionid"`
	ProviderData        datatypes.JSON `gorm:"column:provider_data;type:json;not null;" json:"provider_data"` // 第三方登录提供商用户数据
	CreatedAt           int64          `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`    // 创建时间
	UpdatedAt           int64          `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`    // 更新时间
	DeletedAt           int64          `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`    // 删除时间
}

// 获取结构体对应的表名方法
func (m *UserAuthProvider) TableName() string {
	return "user_auth_provider"
}

// 实例化结构体对象
func NewUserAuthProvider() *UserAuthProvider {
	return &UserAuthProvider{}
}

// 获取主键的对应字段
func (m *UserAuthProvider) PrimaryKey() string {
	return UserAuthProviderColumns.ID
}

// 获取主键值
func (m *UserAuthProvider) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UserAuthProviderColumns = struct {
	ID                  string
	UserID              string
	Provider            string
	ProviderUserID      string
	ProviderUserUnionid string
	ProviderData        string
	CreatedAt           string
	UpdatedAt           string
	DeletedAt           string
}{
	ID:                  "id",
	UserID:              "user_id",
	Provider:            "provider",
	ProviderUserID:      "provider_user_id",
	ProviderUserUnionid: "provider_user_unionid",
	ProviderData:        "provider_data",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
}

// 包含所有表字段的切片
var UserAuthProviderAllColumns = []string{
	UserAuthProviderColumns.ID,                  // 主键ID
	UserAuthProviderColumns.UserID,              // 用户ID,user_info的主键
	UserAuthProviderColumns.Provider,            // 第三方登录提供商。wechatopen：微信开放平台；
	UserAuthProviderColumns.ProviderUserID,      // 第三方登录提供商用户ID
	UserAuthProviderColumns.ProviderUserUnionid, //
	UserAuthProviderColumns.ProviderData,        // 第三方登录提供商用户数据
	UserAuthProviderColumns.CreatedAt,           // 创建时间
	UserAuthProviderColumns.UpdatedAt,           // 更新时间
	UserAuthProviderColumns.DeletedAt,           // 删除时间

}

// 设置值：主键ID
func (m *UserAuthProvider) SetID(v int64) {
	m.ID = v
}

// 设置值：用户ID,user_info的主键
func (m *UserAuthProvider) SetUserID(v int64) {
	m.UserID = v
}

// 设置值：第三方登录提供商。wechatopen：微信开放平台；
func (m *UserAuthProvider) SetProvider(v string) {
	m.Provider = v
}

// 设置值：第三方登录提供商用户ID
func (m *UserAuthProvider) SetProviderUserID(v string) {
	m.ProviderUserID = v
}

// 设置值：
func (m *UserAuthProvider) SetProviderUserUnionid(v string) {
	m.ProviderUserUnionid = v
}

// 设置值：第三方登录提供商用户数据
func (m *UserAuthProvider) SetProviderData(v datatypes.JSON) {
	m.ProviderData = v
}

// 设置值：创建时间
func (m *UserAuthProvider) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *UserAuthProvider) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *UserAuthProvider) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *UserAuthProvider) GetID() int64 {
	return m.ID
}

// 获取值：用户ID,user_info的主键
func (m *UserAuthProvider) GetUserID() int64 {
	return m.UserID
}

// 获取值：第三方登录提供商。wechatopen：微信开放平台；
func (m *UserAuthProvider) GetProvider() string {
	return m.Provider
}

// 获取值：第三方登录提供商用户ID
func (m *UserAuthProvider) GetProviderUserID() string {
	return m.ProviderUserID
}

// 获取值：
func (m *UserAuthProvider) GetProviderUserUnionid() string {
	return m.ProviderUserUnionid
}

// 获取值：第三方登录提供商用户数据
func (m *UserAuthProvider) GetProviderData() datatypes.JSON {
	return m.ProviderData
}

// 获取值：创建时间
func (m *UserAuthProvider) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *UserAuthProvider) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *UserAuthProvider) GetDeletedAt() int64 {
	return m.DeletedAt
}
