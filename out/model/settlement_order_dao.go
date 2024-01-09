package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:36
 * @Desc:   settlement_order 表的 DAO 层
 */

type SettlementOrderDAO struct {
	*_BaseDAO
}

// SettlementOrderDAO 初始化
func NewSettlementOrderDAO(ctx context.Context, db *gorm.DB) *SettlementOrderDAO {
	if db == nil {
		panic(fmt.Errorf("SettlementOrderDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SettlementOrderDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&SettlementOrder{}),
			db:               db,
			model:            SettlementOrder{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          SettlementOrderAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *SettlementOrderDAO) GetTableName() string {
	settlementOrder := &SettlementOrder{}
	return settlementOrder.TableName()
}

// Save 存在则更新，否则插入
func (obj *SettlementOrderDAO) Save(settlementOrder *SettlementOrder) (rowsAffected int64, err error) {
	if settlementOrder.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(settlementOrder, obj.WithID(settlementOrder.ID))
	}
	return obj.Create(settlementOrder)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *SettlementOrderDAO) Create(settlementOrder interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(settlementOrder)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户id) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithOrgID(orgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户id) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithOrgIDs(orgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OpenOrgID] = openOrgIDs })
}

// WithSettleID 设置 settle_id(结算单编号) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithSettleID(settleID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.SettleID] = settleID })
}

// WithSettleIDs 设置 settle_id(结算单编号) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithSettleIDs(settleIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.SettleID] = settleIDs })
}

// WithSettleTime 设置 settle_time(结算单创建时间) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithSettleTime(settleTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.SettleTime] = settleTime })
}

// WithSettleTimes 设置 settle_time(结算单创建时间) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithSettleTimes(settleTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.SettleTime] = settleTimes })
}

// WithURL 设置 url(结算单文件下载url) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithURL(url string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.URL] = url })
}

// WithURLs 设置 url(结算单文件下载url) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithURLs(urls []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.URL] = urls })
}

// WithState 设置 state(结算单状态) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithState(state int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.State] = state })
}

// WithStates 设置 state(结算单状态) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithStates(states []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.State] = states })
}

// WithOrderIDs 设置 order_ids(关联采购订单号list) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithOrderIDs(orderIDs string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OrderIDs] = orderIDs })
}

// WithOrderIDss 设置 order_ids(关联采购订单号list) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithOrderIDss(orderIDss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.OrderIDs] = orderIDss })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *SettlementOrderDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *SettlementOrderDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SettlementOrderColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *SettlementOrderDAO) GetByOption(opts ...Option) (result *SettlementOrder, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *SettlementOrderDAO) GetListByOption(opts ...Option) (results []*SettlementOrder, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *SettlementOrderDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *SettlementOrderDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      settlementOrder: 要更新的数据
//      opts: 更新的条件
func (obj *SettlementOrderDAO) UpdateByOption(settlementOrder *SettlementOrder, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(settlementOrder)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *SettlementOrderDAO) GetFromID(id int64) (result *SettlementOrder, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromID(ids []int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromOrgID(orgID string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromOrgID(orgIDs []string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromOpenOrgID(openOrgID string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromSettleID 通过单个 settle_id(结算单编号) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromSettleID(settleID string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSettleID(settleID))
	return
}

// GetsFromSettleID 通过多个 settle_id(结算单编号) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromSettleID(settleIDs []string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSettleIDs(settleIDs))
	return
}

// GetFromSettleTime 通过单个 settle_time(结算单创建时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromSettleTime(settleTime int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSettleTime(settleTime))
	return
}

// GetsFromSettleTime 通过多个 settle_time(结算单创建时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromSettleTime(settleTimes []int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSettleTimes(settleTimes))
	return
}

// GetFromURL 通过单个 url(结算单文件下载url) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromURL(url string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithURL(url))
	return
}

// GetsFromURL 通过多个 url(结算单文件下载url) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromURL(urls []string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithURLs(urls))
	return
}

// GetFromState 通过单个 state(结算单状态) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromState(state int) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithState(state))
	return
}

// GetsFromState 通过多个 state(结算单状态) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromState(states []int) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithStates(states))
	return
}

// GetFromOrderIDs 通过单个 order_ids(关联采购订单号list) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromOrderIDs(orderIDs string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetsFromOrderIDs 通过多个 order_ids(关联采购订单号list) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromOrderIDs(orderIDss []string) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDss(orderIDss))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromCreatedAt(createdAt int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromCreatedAt(createdAts []int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromUpdatedAt(updatedAt int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetFromDeletedAt(deletedAt int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SettlementOrderDAO) GetsFromDeletedAt(deletedAts []int64) (results []*SettlementOrder, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *SettlementOrderDAO) FetchByPrimaryKey(id int64) (result *SettlementOrder, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
