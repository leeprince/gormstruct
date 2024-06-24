package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:57:57
 * @Desc:   role_permission 表
 */

// 角色权限关联表
type RolePermission struct {
	ID           int64 `gorm:"column:id;primaryKey;type:bigint(20);autoIncrement" json:"id"`        // 主键ID
	RoleID       int64 `gorm:"column:role_id;type:bigint(20);not null;" json:"role_id"`             // 角色ID,对应role表主键
	PermissionID int64 `gorm:"column:permission_id;type:bigint(20);not null;" json:"permission_id"` // 权限ID,对应permission表主键
	CreatedAt    int64 `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`          // 创建时间
	UpdatedAt    int64 `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`          // 更新时间
	DeletedAt    int64 `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`          // 删除时间
}

// 获取结构体对应的表名方法
func (m *RolePermission) TableName() string {
	return "role_permission"
}

// 实例化结构体对象
func NewRolePermission() *RolePermission {
	return &RolePermission{}
}

// 获取主键的对应字段
func (m *RolePermission) PrimaryKey() string {
	return RolePermissionColumns.ID
}

// 获取主键值
func (m *RolePermission) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var RolePermissionColumns = struct {
	ID           string
	RoleID       string
	PermissionID string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	RoleID:       "role_id",
	PermissionID: "permission_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// 包含所有表字段的切片
var RolePermissionAllColumns = []string{
	RolePermissionColumns.ID,           // 主键ID
	RolePermissionColumns.RoleID,       // 角色ID,对应role表主键
	RolePermissionColumns.PermissionID, // 权限ID,对应permission表主键
	RolePermissionColumns.CreatedAt,    // 创建时间
	RolePermissionColumns.UpdatedAt,    // 更新时间
	RolePermissionColumns.DeletedAt,    // 删除时间

}

// 设置值：主键ID
func (m *RolePermission) SetID(v int64) {
	m.ID = v
}

// 设置值：角色ID,对应role表主键
func (m *RolePermission) SetRoleID(v int64) {
	m.RoleID = v
}

// 设置值：权限ID,对应permission表主键
func (m *RolePermission) SetPermissionID(v int64) {
	m.PermissionID = v
}

// 设置值：创建时间
func (m *RolePermission) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *RolePermission) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *RolePermission) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *RolePermission) GetID() int64 {
	return m.ID
}

// 获取值：角色ID,对应role表主键
func (m *RolePermission) GetRoleID() int64 {
	return m.RoleID
}

// 获取值：权限ID,对应permission表主键
func (m *RolePermission) GetPermissionID() int64 {
	return m.PermissionID
}

// 获取值：创建时间
func (m *RolePermission) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *RolePermission) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *RolePermission) GetDeletedAt() int64 {
	return m.DeletedAt
}
