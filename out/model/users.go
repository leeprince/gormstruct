package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-03-12 11:42:04
 * @Desc:   users 表
 */

// 用户表
type Users struct {
	ID        int64   `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`              // 主键
	Name      *string `gorm:"column:name;type:varchar(255);is null;" json:"name"`                     // 名称
	Age       int64   `gorm:"column:age;type:int(11);not null;DEFAULT:18" json:"age"`                 // 年龄
	CardNo    string  `gorm:"column:card_no;type:varchar(18);not null;DEFAULT:" json:"card_no"`       // 身份证
	HeadImg   string  `gorm:"column:head_img;type:varchar(255);not null;DEFAULT:xxx" json:"head_img"` // 头像
	School    *string `gorm:"column:school;type:varchar(255);is null;DEFAULT:" json:"school"`
	CreatedAt int64   `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`          // 创建时间
	UpdatedAt int64   `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`          // 更新时间
	DeletedAt int64   `gorm:"column:deleted_at;type:int(11);not null;DEFAULT:0" json:"deleted_at"` // 删除时间
}

// 获取结构体对应的表名方法
func (m *Users) TableName() string {
	return "users"
}

// 实例化结构体对象
func NewUsers() *Users {
	return &Users{}
}

// 获取主键的对应字段
func (m *Users) PrimaryKey() string {
	return UsersColumns.ID
}

// 获取主键值
func (m *Users) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UsersColumns = struct {
	ID        string
	Name      string
	Age       string
	CardNo    string
	HeadImg   string
	School    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Name:      "name",
	Age:       "age",
	CardNo:    "card_no",
	HeadImg:   "head_img",
	School:    "school",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var UsersAllColumns = []string{
	UsersColumns.ID,        // 主键
	UsersColumns.Name,      // 名称
	UsersColumns.Age,       // 年龄
	UsersColumns.CardNo,    // 身份证
	UsersColumns.HeadImg,   // 头像
	UsersColumns.School,    //
	UsersColumns.CreatedAt, // 创建时间
	UsersColumns.UpdatedAt, // 更新时间
	UsersColumns.DeletedAt, // 删除时间

}

// 设置值：主键
func (m *Users) SetID(v int64) {
	m.ID = v
}

// 设置值：名称
func (m *Users) SetName(v *string) {
	m.Name = v
}

// 设置值：年龄
func (m *Users) SetAge(v int64) {
	m.Age = v
}

// 设置值：身份证
func (m *Users) SetCardNo(v string) {
	m.CardNo = v
}

// 设置值：头像
func (m *Users) SetHeadImg(v string) {
	m.HeadImg = v
}

// 设置值：
func (m *Users) SetSchool(v *string) {
	m.School = v
}

// 设置值：创建时间
func (m *Users) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *Users) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *Users) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键
func (m *Users) GetID() int64 {
	return m.ID
}

// 获取值：名称
func (m *Users) GetName() *string {
	return m.Name
}

// 获取值：年龄
func (m *Users) GetAge() int64 {
	return m.Age
}

// 获取值：身份证
func (m *Users) GetCardNo() string {
	return m.CardNo
}

// 获取值：头像
func (m *Users) GetHeadImg() string {
	return m.HeadImg
}

// 获取值：
func (m *Users) GetSchool() *string {
	return m.School
}

// 获取值：创建时间
func (m *Users) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *Users) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *Users) GetDeletedAt() int64 {
	return m.DeletedAt
}
