package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-06-16 20:58:08
 * @Desc:   permission 表的 DAO 层
 */

type PermissionDAO struct {
	*_BaseDAO
}

// NewPermissionDAO 初始化
func NewPermissionDAO(ctx context.Context, db *gorm.DB) *PermissionDAO {
	if db == nil {
		panic(fmt.Errorf("PermissionDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PermissionDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Permission{}),
			db:               db,
			model:            Permission{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PermissionAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PermissionDAO) GetTableName() string {
	permission := &Permission{}
	return permission.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *PermissionDAO) UpdateOrCreate(permission *Permission) (rowsAffected int64, err error) {
	if permission.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(permission, obj.WithPrimaryKey(permission.PrimaryKeyValue()))
	}
	return obj.Create(permission)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *PermissionDAO) Save(permission *Permission) (rowsAffected int64) {
	return obj.db.Save(permission).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PermissionDAO) Create(permission interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(permission)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *PermissionDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&Permission{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID，权限ID) 字段作为 option 条件
func (obj *PermissionDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.ID] = id })
}

// WithIDs 设置 id(主键ID，权限ID) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.ID] = ids })
}

// WithParentID 设置 parent_id(父权限ID,该表的主键ID) 字段作为 option 条件
func (obj *PermissionDAO) WithParentID(parentID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.ParentID] = parentID })
}

// WithParentIDs 设置 parent_id(父权限ID,该表的主键ID) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithParentIDs(parentIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.ParentID] = parentIDs })
}

// WithName 设置 name(权限名称) 字段作为 option 条件
func (obj *PermissionDAO) WithName(name string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Name] = name })
}

// WithNames 设置 name(权限名称) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithNames(names []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Name] = names })
}

// WithPathID 设置 path_id(权限路径ID) 字段作为 option 条件
func (obj *PermissionDAO) WithPathID(pathID string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.PathID] = pathID })
}

// WithPathIDs 设置 path_id(权限路径ID) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithPathIDs(pathIDs []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.PathID] = pathIDs })
}

// WithPathName 设置 path_name(权限路径名称) 字段作为 option 条件
func (obj *PermissionDAO) WithPathName(pathName string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.PathName] = pathName })
}

// WithPathNames 设置 path_name(权限路径名称) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithPathNames(pathNames []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.PathName] = pathNames })
}

// WithCode 设置 code(权限编码) 字段作为 option 条件
func (obj *PermissionDAO) WithCode(code string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Code] = code })
}

// WithCodes 设置 code(权限编码) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithCodes(codes []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Code] = codes })
}

// WithData 设置 data(权限前端渲染数据) 字段作为 option 条件
func (obj *PermissionDAO) WithData(data datatypes.JSON) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Data] = data })
}

// WithDatas 设置 data(权限前端渲染数据) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithDatas(datas []datatypes.JSON) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.Data] = datas })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PermissionDAO) WithCreatedAt(createdAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithCreatedAts(createdAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PermissionDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PermissionDAO) WithDeletedAt(deletedAt int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PermissionDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[PermissionColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PermissionDAO) GetByOption(opts ...Option) (result *Permission, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PermissionDAO) GetListByOption(opts ...Option) (results []*Permission, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PermissionDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PermissionDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     permission: 要更新的数据
//	     opts: 更新的条件
func (obj *PermissionDAO) UpdateByOption(permission *Permission, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(permission)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID，权限ID) 字段值，获取单条记录
func (obj *PermissionDAO) GetFromID(id int64) (result *Permission, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID，权限ID) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromID(ids []int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromParentID 通过单个 parent_id(父权限ID,该表的主键ID) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromParentID(parentID int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithParentID(parentID))
	return
}

// GetListFromParentID 通过多个 parent_id(父权限ID,该表的主键ID) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromParentID(parentIDs []int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithParentIDs(parentIDs))
	return
}

// GetFromName 通过单个 name(权限名称) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromName(name string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetListFromName 通过多个 name(权限名称) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromName(names []string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromPathID 通过单个 path_id(权限路径ID) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromPathID(pathID string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithPathID(pathID))
	return
}

// GetListFromPathID 通过多个 path_id(权限路径ID) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromPathID(pathIDs []string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithPathIDs(pathIDs))
	return
}

// GetFromPathName 通过单个 path_name(权限路径名称) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromPathName(pathName string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithPathName(pathName))
	return
}

// GetListFromPathName 通过多个 path_name(权限路径名称) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromPathName(pathNames []string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithPathNames(pathNames))
	return
}

// GetFromCode 通过单个 code(权限编码) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromCode(code string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithCode(code))
	return
}

// GetListFromCode 通过多个 code(权限编码) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromCode(codes []string) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithCodes(codes))
	return
}

// GetFromData 通过单个 data(权限前端渲染数据) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromData(data datatypes.JSON) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithData(data))
	return
}

// GetListFromData 通过多个 data(权限前端渲染数据) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromData(datas []datatypes.JSON) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithDatas(datas))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromCreatedAt(createdAt int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromCreatedAt(createdAts []int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromUpdatedAt(updatedAt int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetFromDeletedAt(deletedAt int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PermissionDAO) GetListFromDeletedAt(deletedAts []int64) (results []*Permission, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PermissionDAO) FetchByPrimaryKey(id int64) (result *Permission, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
