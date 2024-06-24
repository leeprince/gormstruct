package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-21 15:15:44
 * @Desc:   user_info 表的 DAO 层
 */

type UserInfoDAO struct {
	*_BaseDAO
}

// NewUserInfoDAO 初始化
func NewUserInfoDAO(ctx context.Context, db *gorm.DB) *UserInfoDAO {
	if db == nil {
		panic(fmt.Errorf("UserInfoDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserInfoDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&UserInfo{}),
			db:               db,
			model:            UserInfo{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserInfoAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UserInfoDAO) GetTableName() string {
	userInfo := &UserInfo{}
	return userInfo.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *UserInfoDAO) UpdateOrCreate(userInfo *UserInfo) (rowsAffected int64, err error) {
	if userInfo.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(userInfo, obj.WithPrimaryKey(userInfo.PrimaryKeyValue()))
	}
	return obj.Create(userInfo)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *UserInfoDAO) Save(userInfo *UserInfo) (rowsAffected int64, err error) {
	tx := obj.db.Save(userInfo)
	return tx.RowsAffected, tx.Error
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UserInfoDAO) Create(userInfo interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(userInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *UserInfoDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&UserInfo{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *UserInfoDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.ID] = ids })
}

// WithUsername 设置 username(用户名) 字段作为 option 条件
func (obj *UserInfoDAO) WithUsername(username string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Username] = username })
}

// WithUsernames 设置 username(用户名) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithUsernames(usernames []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Username] = usernames })
}

// WithPhone 设置 phone(手机号，可用于第三方同一用户的判断标准) 字段作为 option 条件
func (obj *UserInfoDAO) WithPhone(phone string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Phone] = phone })
}

// WithPhones 设置 phone(手机号，可用于第三方同一用户的判断标准) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithPhones(phones []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Phone] = phones })
}

// WithPassword 设置 password(密码) 字段作为 option 条件
func (obj *UserInfoDAO) WithPassword(password string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Password] = password })
}

// WithPasswords 设置 password(密码) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithPasswords(passwords []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Password] = passwords })
}

// WithEmail 设置 email(邮箱) 字段作为 option 条件
func (obj *UserInfoDAO) WithEmail(email string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Email] = email })
}

// WithEmails 设置 email(邮箱) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithEmails(emails []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Email] = emails })
}

// WithNickname 设置 nickname(昵称) 字段作为 option 条件
func (obj *UserInfoDAO) WithNickname(nickname string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Nickname] = nickname })
}

// WithNicknames 设置 nickname(昵称) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithNicknames(nicknames []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Nickname] = nicknames })
}

// WithAvatar 设置 avatar(头像地址) 字段作为 option 条件
func (obj *UserInfoDAO) WithAvatar(avatar string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Avatar] = avatar })
}

// WithAvatars 设置 avatar(头像地址) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithAvatars(avatars []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.Avatar] = avatars })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UserInfoDAO) WithCreatedAt(createdAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithCreatedAts(createdAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UserInfoDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UserInfoDAO) WithDeletedAt(deletedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UserInfoDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserInfoColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *UserInfoDAO) GetByOption(opts ...Option) (result *UserInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UserInfoDAO) GetListByOption(opts ...Option) (results []*UserInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserInfoDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *UserInfoDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     userInfo: 要更新的数据
//	     opts: 更新的条件
func (obj *UserInfoDAO) UpdateByOption(userInfo *UserInfo, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *UserInfoDAO) GetFromID(id int64) (result *UserInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromID(ids []int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromUsername 通过单个 username(用户名) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromUsername(username string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUsername(username))
	return
}

// GetListFromUsername 通过多个 username(用户名) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromUsername(usernames []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUsernames(usernames))
	return
}

// GetFromPhone 通过单个 phone(手机号，可用于第三方同一用户的判断标准) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromPhone(phone string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPhone(phone))
	return
}

// GetListFromPhone 通过多个 phone(手机号，可用于第三方同一用户的判断标准) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromPhone(phones []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPhones(phones))
	return
}

// GetFromPassword 通过单个 password(密码) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromPassword(password string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPassword(password))
	return
}

// GetListFromPassword 通过多个 password(密码) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromPassword(passwords []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPasswords(passwords))
	return
}

// GetFromEmail 通过单个 email(邮箱) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromEmail(email string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithEmail(email))
	return
}

// GetListFromEmail 通过多个 email(邮箱) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromEmail(emails []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithEmails(emails))
	return
}

// GetFromNickname 通过单个 nickname(昵称) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromNickname(nickname string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithNickname(nickname))
	return
}

// GetListFromNickname 通过多个 nickname(昵称) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromNickname(nicknames []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithNicknames(nicknames))
	return
}

// GetFromAvatar 通过单个 avatar(头像地址) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromAvatar(avatar string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAvatar(avatar))
	return
}

// GetListFromAvatar 通过多个 avatar(头像地址) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromAvatar(avatars []string) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAvatars(avatars))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromCreatedAt(createdAt int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromCreatedAt(createdAts []int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromUpdatedAt(updatedAt int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetFromDeletedAt(deletedAt int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserInfoDAO) GetListFromDeletedAt(deletedAts []int64) (results []*UserInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UserInfoDAO) FetchByPrimaryKey(id int64) (result *UserInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
