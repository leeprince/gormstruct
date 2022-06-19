package model

//
type UserBase struct {
	ID   int32  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(30);not null;" json:"name"`
	Age  int32  `gorm:"column:age;type:int(11);not null;" json:"age"`
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
func (m *UserBase) PrimaryKeyValue() int32 {
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
func (m *UserBase) SetID(v int32) {
	m.ID = v
}

// 设置值：
func (m *UserBase) SetName(v string) {
	m.Name = v
}

// 设置值：
func (m *UserBase) SetAge(v int32) {
	m.Age = v
}

// 获取值：
func (m *UserBase) GetID(v int32) {
	m.ID = v
}

// 获取值：
func (m *UserBase) GetName(v string) {
	m.Name = v
}

// 获取值：
func (m *UserBase) GetAge(v int32) {
	m.Age = v
}
