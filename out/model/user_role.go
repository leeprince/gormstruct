package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-19 01:19:51
 * @Desc:   user_role 表
 */

// 用户角色关联表
type UserRole struct {
	ID        int64 `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`  // 主键ID
	UserID    int64 `gorm:"column:user_id;type:int(11);not null;" json:"user_id"`       // 用户ID,user_info表主键
	RoleID    int64 `gorm:"column:role_id;type:int(11);not null;" json:"role_id"`       // 角色ID,role表主键
	CreatedAt int64 `gorm:"column:created_at;type:int(11);not null;" json:"created_at"` // 创建时间
	UpdatedAt int64 `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"` // 更新时间
	DeletedAt int64 `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"` // 删除时间
}

// 获取结构体对应的表名方法
func (m *UserRole) TableName() string {
	return "user_role"
}

// 实例化结构体对象
func NewUserRole() *UserRole {
	return &UserRole{}
}

// 获取主键的对应字段
func (m *UserRole) PrimaryKey() string {
	return UserRoleColumns.ID
}

// 获取主键值
func (m *UserRole) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UserRoleColumns = struct {
	ID        string
	UserID    string
	RoleID    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	RoleID:    "role_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var UserRoleAllColumns = []string{
	UserRoleColumns.ID,        // 主键ID
	UserRoleColumns.UserID,    // 用户ID,user_info表主键
	UserRoleColumns.RoleID,    // 角色ID,role表主键
	UserRoleColumns.CreatedAt, // 创建时间
	UserRoleColumns.UpdatedAt, // 更新时间
	UserRoleColumns.DeletedAt, // 删除时间

}

// 设置值：主键ID
func (m *UserRole) SetID(v int64) {
	m.ID = v
}

// 设置值：用户ID,user_info表主键
func (m *UserRole) SetUserID(v int64) {
	m.UserID = v
}

// 设置值：角色ID,role表主键
func (m *UserRole) SetRoleID(v int64) {
	m.RoleID = v
}

// 设置值：创建时间
func (m *UserRole) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *UserRole) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *UserRole) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *UserRole) GetID() int64 {
	return m.ID
}

// 获取值：用户ID,user_info表主键
func (m *UserRole) GetUserID() int64 {
	return m.UserID
}

// 获取值：角色ID,role表主键
func (m *UserRole) GetRoleID() int64 {
	return m.RoleID
}

// 获取值：创建时间
func (m *UserRole) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *UserRole) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *UserRole) GetDeletedAt() int64 {
	return m.DeletedAt
}
