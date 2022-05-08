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

// 设置：
func (m *UserBase) SetID(v int32) {
	m.ID = v
}

// 设置：
func (m *UserBase) SetName(v string) {
	m.Name = v
}

// 设置：
func (m *UserBase) SetAge(v int32) {
	m.Age = v
}
