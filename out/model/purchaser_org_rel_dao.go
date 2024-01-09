package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:29
 * @Desc:   purchaser_org_rel 表的 DAO 层
 */

type PurchaserOrgRelDAO struct {
	*_BaseDAO
}

// PurchaserOrgRelDAO 初始化
func NewPurchaserOrgRelDAO(ctx context.Context, db *gorm.DB) *PurchaserOrgRelDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrgRelDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrgRelDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrgRel{}),
			db:               db,
			model:            PurchaserOrgRel{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrgRelAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrgRelDAO) GetTableName() string {
	purchaserOrgRel := &PurchaserOrgRel{}
	return purchaserOrgRel.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrgRelDAO) Save(purchaserOrgRel *PurchaserOrgRel) (rowsAffected int64, err error) {
	if purchaserOrgRel.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrgRel, obj.WithID(purchaserOrgRel.ID))
	}
	return obj.Create(purchaserOrgRel)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrgRelDAO) Create(purchaserOrgRel interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrgRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(ID) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.ID] = id })
}

// WithIDs 设置 id(ID) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.ID] = ids })
}

// WithOrgID 设置 org_id(服务商租户ID) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(服务商租户ID) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(开放服务商租户ID) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithOpenOrgID(openOrgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(开放服务商租户ID) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithOpenOrgIDs(openOrgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.OpenOrgID] = openOrgIDs })
}

// WithPurchaserID 设置 purchaser_id(采购商ID) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithPurchaserID(purchaserID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.PurchaserID] = purchaserID })
}

// WithPurchaserIDs 设置 purchaser_id(采购商ID) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithPurchaserIDs(purchaserIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.PurchaserID] = purchaserIDs })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrgRelDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrgRelDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrgRelColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrgRelDAO) GetByOption(opts ...Option) (result *PurchaserOrgRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrgRelDAO) GetListByOption(opts ...Option) (results []*PurchaserOrgRel, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrgRelDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrgRelDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrgRel: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrgRelDAO) UpdateByOption(purchaserOrgRel *PurchaserOrgRel, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrgRel)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(ID) 字段值，获取单条记录
func (obj *PurchaserOrgRelDAO) GetFromID(id int64) (result *PurchaserOrgRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromID(ids []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(服务商租户ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromOrgID(orgID int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(服务商租户ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromOrgID(orgIDs []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(开放服务商租户ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromOpenOrgID(openOrgID int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(开放服务商租户ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromOpenOrgID(openOrgIDs []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromPurchaserID 通过单个 purchaser_id(采购商ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromPurchaserID(purchaserID string) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserID(purchaserID))
	return
}

// GetsFromPurchaserID 通过多个 purchaser_id(采购商ID) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromPurchaserID(purchaserIDs []string) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserIDs(purchaserIDs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrgRelDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrgRel, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrgRelDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrgRel, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
