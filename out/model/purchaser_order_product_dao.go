package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:07
 * @Desc:   purchaser_order_product 表的 DAO 层
 */

type PurchaserOrderProductDAO struct {
	*_BaseDAO
}

// PurchaserOrderProductDAO 初始化
func NewPurchaserOrderProductDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderProductDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderProductDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderProductDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrderProduct{}),
			db:               db,
			model:            PurchaserOrderProduct{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderProductAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderProductDAO) GetTableName() string {
	purchaserOrderProduct := &PurchaserOrderProduct{}
	return purchaserOrderProduct.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderProductDAO) Save(purchaserOrderProduct *PurchaserOrderProduct) (rowsAffected int64, err error) {
	if purchaserOrderProduct.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrderProduct, obj.WithID(purchaserOrderProduct.ID))
	}
	return obj.Create(purchaserOrderProduct)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderProductDAO) Create(purchaserOrderProduct interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrderProduct)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.ID] = ids })
}

// WithOrderID 设置 order_id(订单编号) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(订单编号) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.OrderID] = orderIDs })
}

// WithClassName 设置 class_name(采购商品分类) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithClassName(className string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.ClassName] = className })
}

// WithClassNames 设置 class_name(采购商品分类) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithClassNames(classNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.ClassName] = classNames })
}

// WithName 设置 name(采购商品名称) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Name] = name })
}

// WithNames 设置 name(采购商品名称) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Name] = names })
}

// WithSpecial 设置 special(规格型号) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithSpecial(special string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Special] = special })
}

// WithSpecials 设置 special(规格型号) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithSpecials(specials []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Special] = specials })
}

// WithUnit 设置 unit(单位) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithUnit(unit string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Unit] = unit })
}

// WithUnits 设置 unit(单位) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithUnits(units []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Unit] = units })
}

// WithCount 设置 count(数量) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithCount(count int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Count] = count })
}

// WithCounts 设置 count(数量) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithCounts(counts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Count] = counts })
}

// WithPrice 设置 price(单价（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithPrice(price int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Price] = price })
}

// WithPrices 设置 price(单价（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithPrices(prices []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Price] = prices })
}

// WithRealCount 设置 real_count(实际采购数量) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealCount(realCount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealCount] = realCount })
}

// WithRealCounts 设置 real_count(实际采购数量) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealCounts(realCounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealCount] = realCounts })
}

// WithRealPrice 设置 real_price(实际采购单价（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealPrice(realPrice int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealPrice] = realPrice })
}

// WithRealPrices 设置 real_price(实际采购单价（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealPrices(realPrices []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealPrice] = realPrices })
}

// WithRealAmount 设置 real_amount(实际采购金额（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealAmount(realAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealAmount] = realAmount })
}

// WithRealAmounts 设置 real_amount(实际采购金额（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRealAmounts(realAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.RealAmount] = realAmounts })
}

// WithRemark 设置 remark(采购商品描述) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRemark(remark string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Remark] = remark })
}

// WithRemarks 设置 remark(采购商品描述) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithRemarks(remarks []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.Remark] = remarks })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrderProductDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderProductDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderProductColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderProductDAO) GetByOption(opts ...Option) (result *PurchaserOrderProduct, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderProductDAO) GetListByOption(opts ...Option) (results []*PurchaserOrderProduct, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderProductDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderProductDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrderProduct: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderProductDAO) UpdateByOption(purchaserOrderProduct *PurchaserOrderProduct, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrderProduct)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderProductDAO) GetFromID(id int64) (result *PurchaserOrderProduct, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromID(ids []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrderID 通过单个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromOrderID(orderID string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromOrderID(orderIDs []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromClassName 通过单个 class_name(采购商品分类) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromClassName(className string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithClassName(className))
	return
}

// GetsFromClassName 通过多个 class_name(采购商品分类) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromClassName(classNames []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithClassNames(classNames))
	return
}

// GetFromName 通过单个 name(采购商品名称) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromName(name string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name(采购商品名称) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromName(names []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromSpecial 通过单个 special(规格型号) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromSpecial(special string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithSpecial(special))
	return
}

// GetsFromSpecial 通过多个 special(规格型号) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromSpecial(specials []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithSpecials(specials))
	return
}

// GetFromUnit 通过单个 unit(单位) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromUnit(unit string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUnit(unit))
	return
}

// GetsFromUnit 通过多个 unit(单位) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromUnit(units []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUnits(units))
	return
}

// GetFromCount 通过单个 count(数量) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromCount(count int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCount(count))
	return
}

// GetsFromCount 通过多个 count(数量) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromCount(counts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCounts(counts))
	return
}

// GetFromPrice 通过单个 price(单价（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromPrice(price int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithPrice(price))
	return
}

// GetsFromPrice 通过多个 price(单价（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromPrice(prices []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithPrices(prices))
	return
}

// GetFromRealCount 通过单个 real_count(实际采购数量) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromRealCount(realCount int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealCount(realCount))
	return
}

// GetsFromRealCount 通过多个 real_count(实际采购数量) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromRealCount(realCounts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealCounts(realCounts))
	return
}

// GetFromRealPrice 通过单个 real_price(实际采购单价（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromRealPrice(realPrice int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealPrice(realPrice))
	return
}

// GetsFromRealPrice 通过多个 real_price(实际采购单价（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromRealPrice(realPrices []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealPrices(realPrices))
	return
}

// GetFromRealAmount 通过单个 real_amount(实际采购金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromRealAmount(realAmount int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealAmount(realAmount))
	return
}

// GetsFromRealAmount 通过多个 real_amount(实际采购金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromRealAmount(realAmounts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRealAmounts(realAmounts))
	return
}

// GetFromRemark 通过单个 remark(采购商品描述) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromRemark(remark string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRemark(remark))
	return
}

// GetsFromRemark 通过多个 remark(采购商品描述) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromRemark(remarks []string) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithRemarks(remarks))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderProductDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrderProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderProductDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrderProduct, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
