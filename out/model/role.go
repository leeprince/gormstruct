package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:57:46
 * @Desc:   role 表
 */

// 角色表
type Role struct {
	ID        int64  `gorm:"column:id;primaryKey;type:bigint(20);autoIncrement" json:"id"` // 主键ID
	Name      string `gorm:"column:name;type:varchar(64);not null;" json:"name"`           // 角色名称
	Code      string `gorm:"column:code;type:varchar(32);not null;" json:"code"`           // 角色编码
	Desc      string `gorm:"column:desc;type:varchar(255);not null;" json:"desc"`          // 角色描述
	CreatedAt int64  `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`   // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`   // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`   // 删除时间
}

// 获取结构体对应的表名方法
func (m *Role) TableName() string {
	return "role"
}

// 实例化结构体对象
func NewRole() *Role {
	return &Role{}
}

// 获取主键的对应字段
func (m *Role) PrimaryKey() string {
	return RoleColumns.ID
}

// 获取主键值
func (m *Role) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var RoleColumns = struct {
	ID        string
	Name      string
	Code      string
	Desc      string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Code:      "code",
	Desc:      "desc",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var RoleAllColumns = []string{
	RoleColumns.ID,        // 主键ID
	RoleColumns.Name,      // 角色名称
	RoleColumns.Code,      // 角色编码
	RoleColumns.Desc,      // 角色描述
	RoleColumns.CreatedAt, // 创建时间
	RoleColumns.UpdatedAt, // 更新时间
	RoleColumns.DeletedAt, // 删除时间

}

// 设置值：主键ID
func (m *Role) SetID(v int64) {
	m.ID = v
}

// 设置值：角色名称
func (m *Role) SetName(v string) {
	m.Name = v
}

// 设置值：角色编码
func (m *Role) SetCode(v string) {
	m.Code = v
}

// 设置值：角色描述
func (m *Role) SetDesc(v string) {
	m.Desc = v
}

// 设置值：创建时间
func (m *Role) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *Role) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *Role) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *Role) GetID() int64 {
	return m.ID
}

// 获取值：角色名称
func (m *Role) GetName() string {
	return m.Name
}

// 获取值：角色编码
func (m *Role) GetCode() string {
	return m.Code
}

// 获取值：角色描述
func (m *Role) GetDesc() string {
	return m.Desc
}

// 获取值：创建时间
func (m *Role) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *Role) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *Role) GetDeletedAt() int64 {
	return m.DeletedAt
}
