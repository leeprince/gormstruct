package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-05-20 14:36:31
 * @Desc:   vipcard 表的 DAO 层
 */

type VipcardDAO struct {
	*_BaseDAO
}

// NewVipcardDAO 初始化
func NewVipcardDAO(ctx context.Context, db *gorm.DB) *VipcardDAO {
	if db == nil {
		panic(fmt.Errorf("VipcardDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &VipcardDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Vipcard{}),
			db:               db,
			model:            Vipcard{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          VipcardAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *VipcardDAO) GetTableName() string {
	vipcard := &Vipcard{}
	return vipcard.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *VipcardDAO) UpdateOrCreate(vipcard *Vipcard) (rowsAffected int64, err error) {
	if vipcard.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(vipcard, obj.WithPrimaryKey(vipcard.PrimaryKeyValue()))
	}
	return obj.Create(vipcard)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *VipcardDAO) Save(vipcard *Vipcard) (rowsAffected int64) {
	return obj.db.Save(vipcard).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *VipcardDAO) Create(vipcard interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(vipcard)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *VipcardDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[(&Vipcard{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(id) 字段作为 option 条件
func (obj *VipcardDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.ID] = id })
}

// WithIDs 设置 id(id) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.ID] = ids })
}

// WithVipID 设置 vip_id(vip id) 字段作为 option 条件
func (obj *VipcardDAO) WithVipID(vipID *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.VipID] = vipID })
}

// WithVipIDs 设置 vip_id(vip id) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithVipIDs(vipIDs []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.VipID] = vipIDs })
}

// WithV2TypeID 设置 v2_type_id(类型id) 字段作为 option 条件
func (obj *VipcardDAO) WithV2TypeID(v2TypeID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.V2TypeID] = v2TypeID })
}

// WithV2TypeIDs 设置 v2_type_id(类型id) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithV2TypeIDs(v2TypeIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.V2TypeID] = v2TypeIDs })
}

// WithIsPublicSale 设置 is_public_sale(0 非公开售卖 1公开售卖) 字段作为 option 条件
func (obj *VipcardDAO) WithIsPublicSale(isPublicSale int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.IsPublicSale] = isPublicSale })
}

// WithIsPublicSales 设置 is_public_sale(0 非公开售卖 1公开售卖) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithIsPublicSales(isPublicSales []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.IsPublicSale] = isPublicSales })
}

// WithCreateTime 设置 create_time(创建时间) 字段作为 option 条件
func (obj *VipcardDAO) WithCreateTime(createTime time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.CreateTime] = createTime })
}

// WithCreateTimes 设置 create_time(创建时间) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithCreateTimes(createTimes []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.CreateTime] = createTimes })
}

// WithUpdateTime 设置 update_time(更新时间) 字段作为 option 条件
func (obj *VipcardDAO) WithUpdateTime(updateTime time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.UpdateTime] = updateTime })
}

// WithUpdateTimes 设置 update_time(更新时间) 字段的切片作为 option 条件
func (obj *VipcardDAO) WithUpdateTimes(updateTimes []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VipcardColumns.UpdateTime] = updateTimes })
}

// GetByOption 函数选项模式获取单条记录
func (obj *VipcardDAO) GetByOption(opts ...Option) (result *Vipcard, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *VipcardDAO) GetListByOption(opts ...Option) (results []*Vipcard, err error) {

	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *VipcardDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)

	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *VipcardDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      vipcard: 要更新的数据
//      opts: 更新的条件
func (obj *VipcardDAO) UpdateByOption(vipcard *Vipcard, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(vipcard)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(id) 字段值，获取单条记录
func (obj *VipcardDAO) GetFromID(id int64) (result *Vipcard, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(id) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromID(ids []int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromVipID 通过单个 vip_id(vip id) 字段值，获取多条记录
func (obj *VipcardDAO) GetFromVipID(vipID *int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithVipID(vipID))
	return
}

// GetListFromVipID 通过多个 vip_id(vip id) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromVipID(vipIDs []*int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithVipIDs(vipIDs))
	return
}

// GetFromV2TypeID 通过单个 v2_type_id(类型id) 字段值，获取多条记录
func (obj *VipcardDAO) GetFromV2TypeID(v2TypeID int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithV2TypeID(v2TypeID))
	return
}

// GetListFromV2TypeID 通过多个 v2_type_id(类型id) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromV2TypeID(v2TypeIDs []int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithV2TypeIDs(v2TypeIDs))
	return
}

// GetFromIsPublicSale 通过单个 is_public_sale(0 非公开售卖 1公开售卖) 字段值，获取多条记录
func (obj *VipcardDAO) GetFromIsPublicSale(isPublicSale int8) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithIsPublicSale(isPublicSale))
	return
}

// GetListFromIsPublicSale 通过多个 is_public_sale(0 非公开售卖 1公开售卖) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromIsPublicSale(isPublicSales []int8) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithIsPublicSales(isPublicSales))
	return
}

// GetFromCreateTime 通过单个 create_time(创建时间) 字段值，获取多条记录
func (obj *VipcardDAO) GetFromCreateTime(createTime time.Time) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithCreateTime(createTime))
	return
}

// GetListFromCreateTime 通过多个 create_time(创建时间) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromCreateTime(createTimes []time.Time) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithCreateTimes(createTimes))
	return
}

// GetFromUpdateTime 通过单个 update_time(更新时间) 字段值，获取多条记录
func (obj *VipcardDAO) GetFromUpdateTime(updateTime time.Time) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithUpdateTime(updateTime))
	return
}

// GetListFromUpdateTime 通过多个 update_time(更新时间) 字段值，获取多条记录
func (obj *VipcardDAO) GetListFromUpdateTime(updateTimes []time.Time) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithUpdateTimes(updateTimes))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *VipcardDAO) FetchByPrimaryKey(id int64) (result *Vipcard, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchIndexByQVipID 通过 vip_id 字段值，获取多条记录
func (obj *VipcardDAO) FetchIndexByQVipID(vipID *int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithVipID(vipID))
	return
}

// FetchIndexByIDxV2TypeID 通过 v2_type_id 字段值，获取多条记录
func (obj *VipcardDAO) FetchIndexByIDxV2TypeID(v2TypeID int64) (results []*Vipcard, err error) {
	results, err = obj.GetListByOption(obj.WithV2TypeID(v2TypeID))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
