package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-31 23:21:05
 * @Desc:   user_base 表
 */

type UserBase struct {
	ID   int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(30);not null;" json:"name"`
	Age  int64  `gorm:"column:age;type:int(11);not null;" json:"age"`
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
	ID   string
	Name string
	Age  string
}{
	ID:   "id",
	Name: "name",
	Age:  "age",
}

// 包含所有表字段的切片
var UserBaseAllColumns = []string{
	UserBaseColumns.ID,   //
	UserBaseColumns.Name, //
	UserBaseColumns.Age,  //

}

// 设置值：
func (m *UserBase) SetID(v int64) {
	m.ID = v
}

// 设置值：
func (m *UserBase) SetName(v string) {
	m.Name = v
}

// 设置值：
func (m *UserBase) SetAge(v int64) {
	m.Age = v
}

// 获取值：
func (m *UserBase) GetID() int64 {
	return m.ID
}

// 获取值：
func (m *UserBase) GetName() string {
	return m.Name
}

// 获取值：
func (m *UserBase) GetAge() int64 {
	return m.Age
}
