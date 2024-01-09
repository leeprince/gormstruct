package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:53
 * @Desc:   purchaser_order_invoice_rel 表的 DAO 层
 */

type PurchaserOrderInvoiceRelDAO struct {
	*_BaseDAO
}

// PurchaserOrderInvoiceRelDAO 初始化
func NewPurchaserOrderInvoiceRelDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderInvoiceRelDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderInvoiceRelDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderInvoiceRelDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrderInvoiceRel{}),
			db:               db,
			model:            PurchaserOrderInvoiceRel{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderInvoiceRelAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderInvoiceRelDAO) GetTableName() string {
	purchaserOrderInvoiceRel := &PurchaserOrderInvoiceRel{}
	return purchaserOrderInvoiceRel.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderInvoiceRelDAO) Save(purchaserOrderInvoiceRel *PurchaserOrderInvoiceRel) (rowsAffected int64, err error) {
	if purchaserOrderInvoiceRel.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrderInvoiceRel, obj.WithID(purchaserOrderInvoiceRel.ID))
	}
	return obj.Create(purchaserOrderInvoiceRel)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderInvoiceRelDAO) Create(purchaserOrderInvoiceRel interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrderInvoiceRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.ID] = ids })
}

// WithInvoicePrimaryID 设置 invoice_primary_id(发票数据表主键id) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithInvoicePrimaryID(invoicePrimaryID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.InvoicePrimaryID] = invoicePrimaryID })
}

// WithInvoicePrimaryIDs 设置 invoice_primary_id(发票数据表主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithInvoicePrimaryIDs(invoicePrimaryIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.InvoicePrimaryID] = invoicePrimaryIDs })
}

// WithOrderID 设置 order_id(采购订单ID) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(采购订单ID) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.OrderID] = orderIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderInvoiceRelDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderInvoiceRelColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetByOption(opts ...Option) (result *PurchaserOrderInvoiceRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderInvoiceRelDAO) GetListByOption(opts ...Option) (results []*PurchaserOrderInvoiceRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderInvoiceRelDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderInvoiceRelDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrderInvoiceRel: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderInvoiceRelDAO) UpdateByOption(purchaserOrderInvoiceRel *PurchaserOrderInvoiceRel, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrderInvoiceRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromID(id int64) (result *PurchaserOrderInvoiceRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromID(ids []int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromInvoicePrimaryID 通过单个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromInvoicePrimaryID(invoicePrimaryID int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryID(invoicePrimaryID))
	return
}

// GetsFromInvoicePrimaryID 通过多个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromInvoicePrimaryID(invoicePrimaryIDs []int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryIDs(invoicePrimaryIDs))
	return
}

// GetFromOrderID 通过单个 order_id(采购订单ID) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromOrderID(orderID string) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(采购订单ID) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromOrderID(orderIDs []string) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderInvoiceRelDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrderInvoiceRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderInvoiceRelDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrderInvoiceRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
