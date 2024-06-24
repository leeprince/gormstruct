package model

import (
	"gorm.io/datatypes"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:58:08
 * @Desc:   permission 表
 */

// 权限表
type Permission struct {
	ID        int64          `gorm:"column:id;primaryKey;type:bigint(20);autoIncrement" json:"id"`   // 主键ID，权限ID
	ParentID  int64          `gorm:"column:parent_id;type:bigint(20);not null;" json:"parent_id"`    // 父权限ID,该表的主键ID
	Name      string         `gorm:"column:name;type:varchar(64);not null;" json:"name"`             // 权限名称
	PathID    string         `gorm:"column:path_id;type:varchar(1024);not null;" json:"path_id"`     // 权限路径ID
	PathName  string         `gorm:"column:path_name;type:varchar(1024);not null;" json:"path_name"` // 权限路径名称
	Code      string         `gorm:"column:code;type:varchar(32);not null;" json:"code"`             // 权限编码
	Data      datatypes.JSON `gorm:"column:data;type:json;not null;" json:"data"`                    // 权限前端渲染数据
	CreatedAt int64          `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`     // 创建时间
	UpdatedAt int64          `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`     // 更新时间
	DeletedAt int64          `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`     // 删除时间
}

// 获取结构体对应的表名方法
func (m *Permission) TableName() string {
	return "permission"
}

// 实例化结构体对象
func NewPermission() *Permission {
	return &Permission{}
}

// 获取主键的对应字段
func (m *Permission) PrimaryKey() string {
	return PermissionColumns.ID
}

// 获取主键值
func (m *Permission) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PermissionColumns = struct {
	ID        string
	ParentID  string
	Name      string
	PathID    string
	PathName  string
	Code      string
	Data      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	ParentID:  "parent_id",
	Name:      "name",
	PathID:    "path_id",
	PathName:  "path_name",
	Code:      "code",
	Data:      "data",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var PermissionAllColumns = []string{
	PermissionColumns.ID,        // 主键ID，权限ID
	PermissionColumns.ParentID,  // 父权限ID,该表的主键ID
	PermissionColumns.Name,      // 权限名称
	PermissionColumns.PathID,    // 权限路径ID
	PermissionColumns.PathName,  // 权限路径名称
	PermissionColumns.Code,      // 权限编码
	PermissionColumns.Data,      // 权限前端渲染数据
	PermissionColumns.CreatedAt, // 创建时间
	PermissionColumns.UpdatedAt, // 更新时间
	PermissionColumns.DeletedAt, // 删除时间

}

// 设置值：主键ID，权限ID
func (m *Permission) SetID(v int64) {
	m.ID = v
}

// 设置值：父权限ID,该表的主键ID
func (m *Permission) SetParentID(v int64) {
	m.ParentID = v
}

// 设置值：权限名称
func (m *Permission) SetName(v string) {
	m.Name = v
}

// 设置值：权限路径ID
func (m *Permission) SetPathID(v string) {
	m.PathID = v
}

// 设置值：权限路径名称
func (m *Permission) SetPathName(v string) {
	m.PathName = v
}

// 设置值：权限编码
func (m *Permission) SetCode(v string) {
	m.Code = v
}

// 设置值：权限前端渲染数据
func (m *Permission) SetData(v datatypes.JSON) {
	m.Data = v
}

// 设置值：创建时间
func (m *Permission) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *Permission) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *Permission) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID，权限ID
func (m *Permission) GetID() int64 {
	return m.ID
}

// 获取值：父权限ID,该表的主键ID
func (m *Permission) GetParentID() int64 {
	return m.ParentID
}

// 获取值：权限名称
func (m *Permission) GetName() string {
	return m.Name
}

// 获取值：权限路径ID
func (m *Permission) GetPathID() string {
	return m.PathID
}

// 获取值：权限路径名称
func (m *Permission) GetPathName() string {
	return m.PathName
}

// 获取值：权限编码
func (m *Permission) GetCode() string {
	return m.Code
}

// 获取值：权限前端渲染数据
func (m *Permission) GetData() datatypes.JSON {
	return m.Data
}

// 获取值：创建时间
func (m *Permission) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *Permission) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *Permission) GetDeletedAt() int64 {
	return m.DeletedAt
}
