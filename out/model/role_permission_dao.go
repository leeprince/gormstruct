package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:57:57
 * @Desc:   role_permission 表的 DAO 层
 */

type RolePermissionDAO struct {
	*_BaseDAO
}

// NewRolePermissionDAO 初始化
func NewRolePermissionDAO(ctx context.Context, db *gorm.DB) *RolePermissionDAO {
	if db == nil {
		panic(fmt.Errorf("RolePermissionDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &RolePermissionDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&RolePermission{}),
			db:               db,
			model:            RolePermission{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          RolePermissionAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *RolePermissionDAO) GetTableName() string {
	rolePermission := &RolePermission{}
	return rolePermission.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *RolePermissionDAO) UpdateOrCreate(rolePermission *RolePermission) (rowsAffected int64, err error) {
	if rolePermission.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(rolePermission, obj.WithPrimaryKey(rolePermission.PrimaryKeyValue()))
	}
	return obj.Create(rolePermission)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *RolePermissionDAO) Save(rolePermission *RolePermission) (rowsAffected int64) {
	return obj.db.Save(rolePermission).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *RolePermissionDAO) Create(rolePermission interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(rolePermission)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *RolePermissionDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&RolePermission{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *RolePermissionDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.ID] = ids })
}

// WithRoleID 设置 role_id(角色ID,对应role表主键) 字段作为 option 条件
func (obj *RolePermissionDAO) WithRoleID(roleID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.RoleID] = roleID })
}

// WithRoleIDs 设置 role_id(角色ID,对应role表主键) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithRoleIDs(roleIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.RoleID] = roleIDs })
}

// WithPermissionID 设置 permission_id(权限ID,对应permission表主键) 字段作为 option 条件
func (obj *RolePermissionDAO) WithPermissionID(permissionID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.PermissionID] = permissionID })
}

// WithPermissionIDs 设置 permission_id(权限ID,对应permission表主键) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithPermissionIDs(permissionIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.PermissionID] = permissionIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *RolePermissionDAO) WithCreatedAt(createdAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithCreatedAts(createdAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *RolePermissionDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *RolePermissionDAO) WithDeletedAt(deletedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *RolePermissionDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RolePermissionColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *RolePermissionDAO) GetByOption(opts ...Option) (result *RolePermission, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *RolePermissionDAO) GetListByOption(opts ...Option) (results []*RolePermission, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *RolePermissionDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *RolePermissionDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     rolePermission: 要更新的数据
//	     opts: 更新的条件
func (obj *RolePermissionDAO) UpdateByOption(rolePermission *RolePermission, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(rolePermission)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *RolePermissionDAO) GetFromID(id int64) (result *RolePermission, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromID(ids []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromRoleID 通过单个 role_id(角色ID,对应role表主键) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetFromRoleID(roleID int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithRoleID(roleID))
	return
}

// GetListFromRoleID 通过多个 role_id(角色ID,对应role表主键) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromRoleID(roleIDs []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithRoleIDs(roleIDs))
	return
}

// GetFromPermissionID 通过单个 permission_id(权限ID,对应permission表主键) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetFromPermissionID(permissionID int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithPermissionID(permissionID))
	return
}

// GetListFromPermissionID 通过多个 permission_id(权限ID,对应permission表主键) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromPermissionID(permissionIDs []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithPermissionIDs(permissionIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetFromCreatedAt(createdAt int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromCreatedAt(createdAts []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetFromUpdatedAt(updatedAt int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetFromDeletedAt(deletedAt int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *RolePermissionDAO) GetListFromDeletedAt(deletedAts []int64) (results []*RolePermission, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *RolePermissionDAO) FetchByPrimaryKey(id int64) (result *RolePermission, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
