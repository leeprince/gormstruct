package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-19 01:19:51
 * @Desc:   user_role 表的 DAO 层
 */

type UserRoleDAO struct {
	*_BaseDAO
}

// NewUserRoleDAO 初始化
func NewUserRoleDAO(ctx context.Context, db *gorm.DB) *UserRoleDAO {
	if db == nil {
		panic(fmt.Errorf("UserRoleDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserRoleDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&UserRole{}),
			db:               db,
			model:            UserRole{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserRoleAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UserRoleDAO) GetTableName() string {
	userRole := &UserRole{}
	return userRole.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *UserRoleDAO) UpdateOrCreate(userRole *UserRole) (rowsAffected int64, err error) {
	if userRole.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(userRole, obj.WithPrimaryKey(userRole.PrimaryKeyValue()))
	}
	return obj.Create(userRole)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *UserRoleDAO) Save(userRole *UserRole) (rowsAffected int64) {
	return obj.db.Save(userRole).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UserRoleDAO) Create(userRole interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(userRole)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *UserRoleDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&UserRole{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *UserRoleDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.ID] = ids })
}

// WithUserID 设置 user_id(用户ID,user_info表主键) 字段作为 option 条件
func (obj *UserRoleDAO) WithUserID(userID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.UserID] = userID })
}

// WithUserIDs 设置 user_id(用户ID,user_info表主键) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithUserIDs(userIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.UserID] = userIDs })
}

// WithRoleID 设置 role_id(角色ID,role表主键) 字段作为 option 条件
func (obj *UserRoleDAO) WithRoleID(roleID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.RoleID] = roleID })
}

// WithRoleIDs 设置 role_id(角色ID,role表主键) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithRoleIDs(roleIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.RoleID] = roleIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UserRoleDAO) WithCreatedAt(createdAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithCreatedAts(createdAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UserRoleDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UserRoleDAO) WithDeletedAt(deletedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UserRoleDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserRoleColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *UserRoleDAO) GetByOption(opts ...Option) (result *UserRole, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UserRoleDAO) GetListByOption(opts ...Option) (results []*UserRole, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserRoleDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *UserRoleDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     userRole: 要更新的数据
//	     opts: 更新的条件
func (obj *UserRoleDAO) UpdateByOption(userRole *UserRole, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userRole)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *UserRoleDAO) GetFromID(id int64) (result *UserRole, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromID(ids []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromUserID 通过单个 user_id(用户ID,user_info表主键) 字段值，获取多条记录
func (obj *UserRoleDAO) GetFromUserID(userID int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithUserID(userID))
	return
}

// GetListFromUserID 通过多个 user_id(用户ID,user_info表主键) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromUserID(userIDs []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithUserIDs(userIDs))
	return
}

// GetFromRoleID 通过单个 role_id(角色ID,role表主键) 字段值，获取多条记录
func (obj *UserRoleDAO) GetFromRoleID(roleID int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithRoleID(roleID))
	return
}

// GetListFromRoleID 通过多个 role_id(角色ID,role表主键) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromRoleID(roleIDs []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithRoleIDs(roleIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetFromCreatedAt(createdAt int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromCreatedAt(createdAts []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetFromUpdatedAt(updatedAt int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetFromDeletedAt(deletedAt int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserRoleDAO) GetListFromDeletedAt(deletedAts []int64) (results []*UserRole, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UserRoleDAO) FetchByPrimaryKey(id int64) (result *UserRole, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
