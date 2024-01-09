package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:22
 * @Desc:   purchaser_order_voucher_rel 表的 DAO 层
 */

type PurchaserOrderVoucherRelDAO struct {
	*_BaseDAO
}

// PurchaserOrderVoucherRelDAO 初始化
func NewPurchaserOrderVoucherRelDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderVoucherRelDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderVoucherRelDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderVoucherRelDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrderVoucherRel{}),
			db:               db,
			model:            PurchaserOrderVoucherRel{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderVoucherRelAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderVoucherRelDAO) GetTableName() string {
	purchaserOrderVoucherRel := &PurchaserOrderVoucherRel{}
	return purchaserOrderVoucherRel.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderVoucherRelDAO) Save(purchaserOrderVoucherRel *PurchaserOrderVoucherRel) (rowsAffected int64, err error) {
	if purchaserOrderVoucherRel.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrderVoucherRel, obj.WithID(purchaserOrderVoucherRel.ID))
	}
	return obj.Create(purchaserOrderVoucherRel)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderVoucherRelDAO) Create(purchaserOrderVoucherRel interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrderVoucherRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.ID] = ids })
}

// WithInvoicePrimaryID 设置 invoice_primary_id(发票数据表主键id) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithInvoicePrimaryID(invoicePrimaryID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.InvoicePrimaryID] = invoicePrimaryID })
}

// WithInvoicePrimaryIDs 设置 invoice_primary_id(发票数据表主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithInvoicePrimaryIDs(invoicePrimaryIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.InvoicePrimaryID] = invoicePrimaryIDs })
}

// WithOrderID 设置 order_id(采购订单ID) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(采购订单ID) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.OrderID] = orderIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderVoucherRelDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderVoucherRelColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderVoucherRelDAO) GetByOption(opts ...Option) (result *PurchaserOrderVoucherRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderVoucherRelDAO) GetListByOption(opts ...Option) (results []*PurchaserOrderVoucherRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderVoucherRelDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderVoucherRelDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrderVoucherRel: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderVoucherRelDAO) UpdateByOption(purchaserOrderVoucherRel *PurchaserOrderVoucherRel, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrderVoucherRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromID(id int64) (result *PurchaserOrderVoucherRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromID(ids []int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromInvoicePrimaryID 通过单个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromInvoicePrimaryID(invoicePrimaryID int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryID(invoicePrimaryID))
	return
}

// GetsFromInvoicePrimaryID 通过多个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromInvoicePrimaryID(invoicePrimaryIDs []int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryIDs(invoicePrimaryIDs))
	return
}

// GetFromOrderID 通过单个 order_id(采购订单ID) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromOrderID(orderID string) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(采购订单ID) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromOrderID(orderIDs []string) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderVoucherRelDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrderVoucherRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderVoucherRelDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrderVoucherRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
