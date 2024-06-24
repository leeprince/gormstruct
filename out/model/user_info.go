package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-21 15:15:44
 * @Desc:   user_info 表
 */

// 用户基础信息表
type UserInfo struct {
	ID        int64  `gorm:"column:id;primaryKey;type:bigint(20);autoIncrement" json:"id"` // 主键ID
	Username  string `gorm:"column:username;type:varchar(64);not null;" json:"username"`   // 用户名
	Phone     string `gorm:"column:phone;type:varchar(20);not null;DEFAULT:" json:"phone"` // 手机号，可用于第三方同一用户的判断标准
	Password  string `gorm:"column:password;type:varchar(64);not null;" json:"password"`   // 密码
	Email     string `gorm:"column:email;type:varchar(64);not null;" json:"email"`         // 邮箱
	Nickname  string `gorm:"column:nickname;type:varchar(64);not null;" json:"nickname"`   // 昵称
	Avatar    string `gorm:"column:avatar;type:varchar(1024);not null;" json:"avatar"`     // 头像地址
	CreatedAt int64  `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`   // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`   // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`   // 删除时间
}

// 获取结构体对应的表名方法
func (m *UserInfo) TableName() string {
	return "user_info"
}

// 实例化结构体对象
func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

// 获取主键的对应字段
func (m *UserInfo) PrimaryKey() string {
	return UserInfoColumns.ID
}

// 获取主键值
func (m *UserInfo) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var UserInfoColumns = struct {
	ID        string
	Username  string
	Phone     string
	Password  string
	Email     string
	Nickname  string
	Avatar    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Username:  "username",
	Phone:     "phone",
	Password:  "password",
	Email:     "email",
	Nickname:  "nickname",
	Avatar:    "avatar",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var UserInfoAllColumns = []string{
	UserInfoColumns.ID,        // 主键ID
	UserInfoColumns.Username,  // 用户名
	UserInfoColumns.Phone,     // 手机号，可用于第三方同一用户的判断标准
	UserInfoColumns.Password,  // 密码
	UserInfoColumns.Email,     // 邮箱
	UserInfoColumns.Nickname,  // 昵称
	UserInfoColumns.Avatar,    // 头像地址
	UserInfoColumns.CreatedAt, // 创建时间
	UserInfoColumns.UpdatedAt, // 更新时间
	UserInfoColumns.DeletedAt, // 删除时间

}

// 设置值：主键ID
func (m *UserInfo) SetID(v int64) {
	m.ID = v
}

// 设置值：用户名
func (m *UserInfo) SetUsername(v string) {
	m.Username = v
}

// 设置值：手机号，可用于第三方同一用户的判断标准
func (m *UserInfo) SetPhone(v string) {
	m.Phone = v
}

// 设置值：密码
func (m *UserInfo) SetPassword(v string) {
	m.Password = v
}

// 设置值：邮箱
func (m *UserInfo) SetEmail(v string) {
	m.Email = v
}

// 设置值：昵称
func (m *UserInfo) SetNickname(v string) {
	m.Nickname = v
}

// 设置值：头像地址
func (m *UserInfo) SetAvatar(v string) {
	m.Avatar = v
}

// 设置值：创建时间
func (m *UserInfo) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *UserInfo) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *UserInfo) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *UserInfo) GetID() int64 {
	return m.ID
}

// 获取值：用户名
func (m *UserInfo) GetUsername() string {
	return m.Username
}

// 获取值：手机号，可用于第三方同一用户的判断标准
func (m *UserInfo) GetPhone() string {
	return m.Phone
}

// 获取值：密码
func (m *UserInfo) GetPassword() string {
	return m.Password
}

// 获取值：邮箱
func (m *UserInfo) GetEmail() string {
	return m.Email
}

// 获取值：昵称
func (m *UserInfo) GetNickname() string {
	return m.Nickname
}

// 获取值：头像地址
func (m *UserInfo) GetAvatar() string {
	return m.Avatar
}

// 获取值：创建时间
func (m *UserInfo) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *UserInfo) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *UserInfo) GetDeletedAt() int64 {
	return m.DeletedAt
}
