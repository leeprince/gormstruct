package 	model 	

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-12-02 02:13:49
 * @Desc:   users 表
 */
 	
//  	 	
type 	Users 	struct { 	 	
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
	return UsersColumns.
}

// 获取主键值
func (m *Users) PrimaryKeyValue()  {
	return m.
}

// 表字段的映射
var UsersColumns = struct {

} {

}

// 包含所有表字段的切片
var UsersAllColumns = []string{
    
}




 	

