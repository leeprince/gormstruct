package model

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-09-05 01:42:43
 * @Desc:   licence 表
 */

// 授权表
type Licence struct {
	ID          int64      `gorm:"column:id;primaryKey;type:bigint(11);autoIncrement" json:"id"`       // 主键ID
	ClientIP    *string    `gorm:"column:client_ip;type:varchar(32);is null;" json:"client_ip"`        // 客户端IP
	MachineName *string    `gorm:"column:machine_name;type:varchar(128);is null;" json:"machine_name"` // 设备名称
	ExpireTime  time.Time  `gorm:"column:expire_time;type:datetime(3);is null;" json:"expire_time"`    // 授权设备过期时间
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:datetime(3);is null;" json:"deleted_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime(3);is null;" json:"updated_at"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime(3);is null;" json:"created_at"`
}

// 获取结构体对应的表名方法
func (m *Licence) TableName() string {
	return "licence"
}

// 实例化结构体对象
func NewLicence() *Licence {
	return &Licence{}
}

// 获取主键的对应字段
func (m *Licence) PrimaryKey() string {
	return LicenceColumns.ID
}

// 获取主键值
func (m *Licence) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var LicenceColumns = struct {
	ID          string
	ClientIP    string
	MachineName string
	ExpireTime  string
	DeletedAt   string
	UpdatedAt   string
	CreatedAt   string
}{
	ID:          "id",
	ClientIP:    "client_ip",
	MachineName: "machine_name",
	ExpireTime:  "expire_time",
	DeletedAt:   "deleted_at",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
}

// 包含所有表字段的切片
var LicenceAllColumns = []string{
	LicenceColumns.ID,          // 主键ID
	LicenceColumns.ClientIP,    // 客户端IP
	LicenceColumns.MachineName, // 设备名称
	LicenceColumns.ExpireTime,  // 授权设备过期时间
	LicenceColumns.DeletedAt,   //
	LicenceColumns.UpdatedAt,   //
	LicenceColumns.CreatedAt,   //

}

// 设置值：主键ID
func (m *Licence) SetID(v int64) {
	m.ID = v
}

// 设置值：客户端IP
func (m *Licence) SetClientIP(v *string) {
	m.ClientIP = v
}

// 设置值：设备名称
func (m *Licence) SetMachineName(v *string) {
	m.MachineName = v
}

// 设置值：授权设备过期时间
func (m *Licence) SetExpireTime(v time.Time) {
	m.ExpireTime = v
}

// 设置值：
func (m *Licence) SetDeletedAt(v *time.Time) {
	m.DeletedAt = v
}

// 设置值：
func (m *Licence) SetUpdatedAt(v time.Time) {
	m.UpdatedAt = v
}

// 设置值：
func (m *Licence) SetCreatedAt(v time.Time) {
	m.CreatedAt = v
}

// 获取值：主键ID
func (m *Licence) GetID() int64 {
	return m.ID
}

// 获取值：客户端IP
func (m *Licence) GetClientIP() *string {
	return m.ClientIP
}

// 获取值：设备名称
func (m *Licence) GetMachineName() *string {
	return m.MachineName
}

// 获取值：授权设备过期时间
func (m *Licence) GetExpireTime() time.Time {
	return m.ExpireTime
}

// 获取值：
func (m *Licence) GetDeletedAt() *time.Time {
	return m.DeletedAt
}

// 获取值：
func (m *Licence) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// 获取值：
func (m *Licence) GetCreatedAt() time.Time {
	return m.CreatedAt
}
