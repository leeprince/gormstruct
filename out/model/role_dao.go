package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:57:46
 * @Desc:   role 表的 DAO 层
 */

type RoleDAO struct {
	*_BaseDAO
}

// NewRoleDAO 初始化
func NewRoleDAO(ctx context.Context, db *gorm.DB) *RoleDAO {
	if db == nil {
		panic(fmt.Errorf("RoleDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &RoleDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Role{}),
			db:               db,
			model:            Role{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          RoleAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *RoleDAO) GetTableName() string {
	role := &Role{}
	return role.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *RoleDAO) UpdateOrCreate(role *Role) (rowsAffected int64, err error) {
	if role.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(role, obj.WithPrimaryKey(role.PrimaryKeyValue()))
	}
	return obj.Create(role)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *RoleDAO) Save(role *Role) (rowsAffected int64) {
	return obj.db.Save(role).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *RoleDAO) Create(role interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(role)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *RoleDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&Role{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *RoleDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *RoleDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.ID] = ids })
}

// WithName 设置 name(角色名称) 字段作为 option 条件
func (obj *RoleDAO) WithName(name string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Name] = name })
}

// WithNames 设置 name(角色名称) 字段的切片作为 option 条件
func (obj *RoleDAO) WithNames(names []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Name] = names })
}

// WithCode 设置 code(角色编码) 字段作为 option 条件
func (obj *RoleDAO) WithCode(code string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Code] = code })
}

// WithCodes 设置 code(角色编码) 字段的切片作为 option 条件
func (obj *RoleDAO) WithCodes(codes []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Code] = codes })
}

// WithDesc 设置 desc(角色描述) 字段作为 option 条件
func (obj *RoleDAO) WithDesc(desc string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Desc] = desc })
}

// WithDescs 设置 desc(角色描述) 字段的切片作为 option 条件
func (obj *RoleDAO) WithDescs(descs []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.Desc] = descs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *RoleDAO) WithCreatedAt(createdAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *RoleDAO) WithCreatedAts(createdAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *RoleDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *RoleDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *RoleDAO) WithDeletedAt(deletedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *RoleDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[RoleColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *RoleDAO) GetByOption(opts ...Option) (result *Role, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *RoleDAO) GetListByOption(opts ...Option) (results []*Role, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *RoleDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *RoleDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     role: 要更新的数据
//	     opts: 更新的条件
func (obj *RoleDAO) UpdateByOption(role *Role, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(role)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *RoleDAO) GetFromID(id int64) (result *Role, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromID(ids []int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromName 通过单个 name(角色名称) 字段值，获取多条记录
func (obj *RoleDAO) GetFromName(name string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetListFromName 通过多个 name(角色名称) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromName(names []string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromCode 通过单个 code(角色编码) 字段值，获取多条记录
func (obj *RoleDAO) GetFromCode(code string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithCode(code))
	return
}

// GetListFromCode 通过多个 code(角色编码) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromCode(codes []string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithCodes(codes))
	return
}

// GetFromDesc 通过单个 desc(角色描述) 字段值，获取多条记录
func (obj *RoleDAO) GetFromDesc(desc string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithDesc(desc))
	return
}

// GetListFromDesc 通过多个 desc(角色描述) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromDesc(descs []string) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithDescs(descs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *RoleDAO) GetFromCreatedAt(createdAt int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromCreatedAt(createdAts []int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *RoleDAO) GetFromUpdatedAt(updatedAt int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *RoleDAO) GetFromDeletedAt(deletedAt int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *RoleDAO) GetListFromDeletedAt(deletedAts []int64) (results []*Role, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *RoleDAO) FetchByPrimaryKey(id int64) (result *Role, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
