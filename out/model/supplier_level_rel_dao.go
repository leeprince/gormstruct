package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:57
 * @Desc:   supplier_level_rel 表的 DAO 层
 */

type SupplierLevelRelDAO struct {
	*_BaseDAO
}

// SupplierLevelRelDAO 初始化
func NewSupplierLevelRelDAO(ctx context.Context, db *gorm.DB) *SupplierLevelRelDAO {
	if db == nil {
		panic(fmt.Errorf("SupplierLevelRelDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SupplierLevelRelDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&SupplierLevelRel{}),
			db:               db,
			model:            SupplierLevelRel{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          SupplierLevelRelAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *SupplierLevelRelDAO) GetTableName() string {
	supplierLevelRel := &SupplierLevelRel{}
	return supplierLevelRel.TableName()
}

// Save 存在则更新，否则插入
func (obj *SupplierLevelRelDAO) Save(supplierLevelRel *SupplierLevelRel) (rowsAffected int64, err error) {
	if supplierLevelRel.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(supplierLevelRel, obj.WithID(supplierLevelRel.ID))
	}
	return obj.Create(supplierLevelRel)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *SupplierLevelRelDAO) Create(supplierLevelRel interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(supplierLevelRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户ID) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户ID) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户ID) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户ID) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.OpenOrgID] = openOrgIDs })
}

// WithSupplierID 设置 supplier_id(二级供应商ID) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithSupplierID(supplierID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.SupplierID] = supplierID })
}

// WithSupplierIDs 设置 supplier_id(二级供应商ID) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithSupplierIDs(supplierIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.SupplierID] = supplierIDs })
}

// WithPurchaserID 设置 purchaser_id(采购商ID) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithPurchaserID(purchaserID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.PurchaserID] = purchaserID })
}

// WithPurchaserIDs 设置 purchaser_id(采购商ID) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithPurchaserIDs(purchaserIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.PurchaserID] = purchaserIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *SupplierLevelRelDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *SupplierLevelRelDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierLevelRelColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *SupplierLevelRelDAO) GetByOption(opts ...Option) (result *SupplierLevelRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *SupplierLevelRelDAO) GetListByOption(opts ...Option) (results []*SupplierLevelRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *SupplierLevelRelDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *SupplierLevelRelDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      supplierLevelRel: 要更新的数据
//      opts: 更新的条件
func (obj *SupplierLevelRelDAO) UpdateByOption(supplierLevelRel *SupplierLevelRel, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(supplierLevelRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *SupplierLevelRelDAO) GetFromID(id int64) (result *SupplierLevelRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromID(ids []int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromOrgID(orgID int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromOrgID(orgIDs []int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromOpenOrgID(openOrgID string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromSupplierID 通过单个 supplier_id(二级供应商ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromSupplierID(supplierID string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierID(supplierID))
	return
}

// GetsFromSupplierID 通过多个 supplier_id(二级供应商ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromSupplierID(supplierIDs []string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierIDs(supplierIDs))
	return
}

// GetFromPurchaserID 通过单个 purchaser_id(采购商ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromPurchaserID(purchaserID string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserID(purchaserID))
	return
}

// GetsFromPurchaserID 通过多个 purchaser_id(采购商ID) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromPurchaserID(purchaserIDs []string) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserIDs(purchaserIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromCreatedAt(createdAt int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromCreatedAt(createdAts []int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromUpdatedAt(updatedAt int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetFromDeletedAt(deletedAt int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierLevelRelDAO) GetsFromDeletedAt(deletedAts []int64) (results []*SupplierLevelRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *SupplierLevelRelDAO) FetchByPrimaryKey(id int64) (result *SupplierLevelRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
