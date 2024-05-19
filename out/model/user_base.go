package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-05-20 02:57:56
 * @Desc:   user_base 表
 */

// 用户基础表\n用户基础表换行
type UserBase struct {
	ID      int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"` // 主键id
	Name    string `gorm:"column:name;type:varchar(30);not null;" json:"name"`        // 姓名 \r 姓名
	Age     int64  `gorm:"column:age;type:int(11);not null;" json:"age"`              // 年龄 \n 年龄
	Schoool string `gorm:"column:schoool;type:varchar(255);not null;" json:"schoool"` // 学校 \r\n 学校
	Name1   string `gorm:"column:name_1;type:varchar(255);not null;" json:"name_1"`   // 姓名\n姓名换行
}

// 获取结构体对应的表名方法
func (m *UserBase) TableName() string {
	return "user_base"
}

// 实例化结构体对象
func NewUserBase() *UserBase {
	return &UserBase{}
}

// 获取主键的对应字段
func (m *UserBase) PrimaryKey() string {
	return UserBaseColumns.ID
}

// 获取主键值
func (m *UserBase) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UserBaseColumns = struct {
	ID      string
	Name    string
	Age     string
	Schoool string
	Name1   string
}{
	ID:      "id",
	Name:    "name",
	Age:     "age",
	Schoool: "schoool",
	Name1:   "name_1",
}

// 包含所有表字段的切片
var UserBaseAllColumns = []string{
	UserBaseColumns.ID,      // 主键id
	UserBaseColumns.Name,    // 姓名 \r 姓名
	UserBaseColumns.Age,     // 年龄 \n 年龄
	UserBaseColumns.Schoool, // 学校 \r\n 学校
	UserBaseColumns.Name1,   // 姓名\n姓名换行

}

// 设置值：主键id
func (m *UserBase) SetID(v int64) {
	m.ID = v
}

// 设置值：姓名 \r 姓名
func (m *UserBase) SetName(v string) {
	m.Name = v
}

// 设置值：年龄 \n 年龄
func (m *UserBase) SetAge(v int64) {
	m.Age = v
}

// 设置值：学校 \r\n 学校
func (m *UserBase) SetSchoool(v string) {
	m.Schoool = v
}

// 设置值：姓名\n姓名换行
func (m *UserBase) SetName1(v string) {
	m.Name1 = v
}

// 获取值：主键id
func (m *UserBase) GetID() int64 {
	return m.ID
}

// 获取值：姓名 \r 姓名
func (m *UserBase) GetName() string {
	return m.Name
}

// 获取值：年龄 \n 年龄
func (m *UserBase) GetAge() int64 {
	return m.Age
}

// 获取值：学校 \r\n 学校
func (m *UserBase) GetSchoool() string {
	return m.Schoool
}

// 获取值：姓名\n姓名换行
func (m *UserBase) GetName1() string {
	return m.Name1
}
