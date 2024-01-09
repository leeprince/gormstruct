package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:14
 * @Desc:   purchaser_order_settlement_rel 表的 DAO 层
 */

type PurchaserOrderSettlementRelDAO struct {
	*_BaseDAO
}

// PurchaserOrderSettlementRelDAO 初始化
func NewPurchaserOrderSettlementRelDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderSettlementRelDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderSettlementRelDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderSettlementRelDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrderSettlementRel{}),
			db:               db,
			model:            PurchaserOrderSettlementRel{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderSettlementRelAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderSettlementRelDAO) GetTableName() string {
	purchaserOrderSettlementRel := &PurchaserOrderSettlementRel{}
	return purchaserOrderSettlementRel.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderSettlementRelDAO) Save(purchaserOrderSettlementRel *PurchaserOrderSettlementRel) (rowsAffected int64, err error) {
	if purchaserOrderSettlementRel.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrderSettlementRel, obj.WithID(purchaserOrderSettlementRel.ID))
	}
	return obj.Create(purchaserOrderSettlementRel)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderSettlementRelDAO) Create(purchaserOrderSettlementRel interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrderSettlementRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.ID] = ids })
}

// WithOrderID 设置 order_id(订单编号) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(订单编号) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.OrderID] = orderIDs })
}

// WithSettleID 设置 settle_id(结算单ID) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithSettleID(settleID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.SettleID] = settleID })
}

// WithSettleIDs 设置 settle_id(结算单ID) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithSettleIDs(settleIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.SettleID] = settleIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderSettlementRelDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderSettlementRelColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderSettlementRelDAO) GetByOption(opts ...Option) (result *PurchaserOrderSettlementRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderSettlementRelDAO) GetListByOption(opts ...Option) (results []*PurchaserOrderSettlementRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderSettlementRelDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderSettlementRelDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrderSettlementRel: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderSettlementRelDAO) UpdateByOption(purchaserOrderSettlementRel *PurchaserOrderSettlementRel, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrderSettlementRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromID(id int64) (result *PurchaserOrderSettlementRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromID(ids []int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrderID 通过单个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromOrderID(orderID string) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromOrderID(orderIDs []string) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromSettleID 通过单个 settle_id(结算单ID) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromSettleID(settleID string) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithSettleID(settleID))
	return
}

// GetsFromSettleID 通过多个 settle_id(结算单ID) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromSettleID(settleIDs []string) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithSettleIDs(settleIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderSettlementRelDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrderSettlementRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderSettlementRelDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrderSettlementRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
