package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:00
 * @Desc:   purchaser_order_pre 表的 DAO 层
 */

type PurchaserOrderPreDAO struct {
	*_BaseDAO
}

// PurchaserOrderPreDAO 初始化
func NewPurchaserOrderPreDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderPreDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderPreDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderPreDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrderPre{}),
			db:               db,
			model:            PurchaserOrderPre{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderPreAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderPreDAO) GetTableName() string {
	purchaserOrderPre := &PurchaserOrderPre{}
	return purchaserOrderPre.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderPreDAO) Save(purchaserOrderPre *PurchaserOrderPre) (rowsAffected int64, err error) {
	if purchaserOrderPre.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrderPre, obj.WithID(purchaserOrderPre.ID))
	}
	return obj.Create(purchaserOrderPre)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderPreDAO) Create(purchaserOrderPre interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrderPre)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户id) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户id) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OpenOrgID] = openOrgIDs })
}

// WithPurchaserID 设置 purchaser_id(采购商id) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserID(purchaserID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserID] = purchaserID })
}

// WithPurchaserIDs 设置 purchaser_id(采购商id) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserIDs(purchaserIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserID] = purchaserIDs })
}

// WithPurchaserOrderID 设置 purchaser_order_id(采购单编号) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserOrderID(purchaserOrderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserOrderID] = purchaserOrderID })
}

// WithPurchaserOrderIDs 设置 purchaser_order_id(采购单编号) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserOrderIDs(purchaserOrderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserOrderID] = purchaserOrderIDs })
}

// WithPurchaserOrderName 设置 purchaser_order_name(采购单名称) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserOrderName(purchaserOrderName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserOrderName] = purchaserOrderName })
}

// WithPurchaserOrderNames 设置 purchaser_order_name(采购单名称) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithPurchaserOrderNames(purchaserOrderNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.PurchaserOrderName] = purchaserOrderNames })
}

// WithOrderTime 设置 order_time(订单创建时间) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOrderTime(orderTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OrderTime] = orderTime })
}

// WithOrderTimes 设置 order_time(订单创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithOrderTimes(orderTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.OrderTime] = orderTimes })
}

// WithState 设置 state(状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithState(state int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.State] = state })
}

// WithStates 设置 state(状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithStates(states []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.State] = states })
}

// WithURL 设置 url(采购单文件下载URL) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithURL(url string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.URL] = url })
}

// WithURLs 设置 url(采购单文件下载URL) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithURLs(urls []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.URL] = urls })
}

// WithCycle 设置 cycle(采购周期) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCycle(cycle string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Cycle] = cycle })
}

// WithCycles 设置 cycle(采购周期) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCycles(cycles []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Cycle] = cycles })
}

// WithClassName 设置 class_name(商品品类) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithClassName(className string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ClassName] = className })
}

// WithClassNames 设置 class_name(商品品类) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithClassNames(classNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ClassName] = classNames })
}

// WithCount 设置 count(采购数量) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCount(count int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Count] = count })
}

// WithCounts 设置 count(采购数量) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCounts(counts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Count] = counts })
}

// WithAmount 设置 amount(采购金额) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithAmount(amount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Amount] = amount })
}

// WithAmounts 设置 amount(采购金额) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithAmounts(amounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Amount] = amounts })
}

// WithReceiveAreaCode 设置 receive_area_code(收货所在地省市区) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithReceiveAreaCode(receiveAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ReceiveAreaCode] = receiveAreaCode })
}

// WithReceiveAreaCodes 设置 receive_area_code(收货所在地省市区) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithReceiveAreaCodes(receiveAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ReceiveAreaCode] = receiveAreaCodes })
}

// WithReceiveAddress 设置 receive_address(收货所在地详细地址) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithReceiveAddress(receiveAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ReceiveAddress] = receiveAddress })
}

// WithReceiveAddresss 设置 receive_address(收货所在地详细地址) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithReceiveAddresss(receiveAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.ReceiveAddress] = receiveAddresss })
}

// WithSpecial 设置 special(品质规格) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithSpecial(special string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Special] = special })
}

// WithSpecials 设置 special(品质规格) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithSpecials(specials []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Special] = specials })
}

// WithRequireText 设置 require_text(采购要求) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRequireText(requireText string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.RequireText] = requireText })
}

// WithRequireTexts 设置 require_text(采购要求) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRequireTexts(requireTexts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.RequireText] = requireTexts })
}

// WithRate 设置 rate(签约服务费率) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRate(rate string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Rate] = rate })
}

// WithRates 设置 rate(签约服务费率) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRates(rates []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.Rate] = rates })
}

// WithFeeAmount 设置 fee_amount(服务费【单位：分】) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithFeeAmount(feeAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.FeeAmount] = feeAmount })
}

// WithFeeAmounts 设置 fee_amount(服务费【单位：分】) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithFeeAmounts(feeAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.FeeAmount] = feeAmounts })
}

// WithBillAmount 设置 bill_amount(账单金额【单位：分】) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithBillAmount(billAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.BillAmount] = billAmount })
}

// WithBillAmounts 设置 bill_amount(账单金额【单位：分】) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithBillAmounts(billAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.BillAmount] = billAmounts })
}

// WithRealAmount 设置 real_amount(实际支付采购金额【单位分】) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRealAmount(realAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.RealAmount] = realAmount })
}

// WithRealAmounts 设置 real_amount(实际支付采购金额【单位分】) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithRealAmounts(realAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.RealAmount] = realAmounts })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserOrderPreDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderPreDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderPreColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderPreDAO) GetByOption(opts ...Option) (result *PurchaserOrderPre, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderPreDAO) GetListByOption(opts ...Option) (results []*PurchaserOrderPre, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderPreDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderPreDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrderPre: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderPreDAO) UpdateByOption(purchaserOrderPre *PurchaserOrderPre, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrderPre)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderPreDAO) GetFromID(id int64) (result *PurchaserOrderPre, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromID(ids []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromOrgID(orgID int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromOrgID(orgIDs []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromOpenOrgID(openOrgID string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromPurchaserID 通过单个 purchaser_id(采购商id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromPurchaserID(purchaserID string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserID(purchaserID))
	return
}

// GetsFromPurchaserID 通过多个 purchaser_id(采购商id) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromPurchaserID(purchaserIDs []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserIDs(purchaserIDs))
	return
}

// GetFromPurchaserOrderID 通过单个 purchaser_order_id(采购单编号) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromPurchaserOrderID(purchaserOrderID string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderID(purchaserOrderID))
	return
}

// GetsFromPurchaserOrderID 通过多个 purchaser_order_id(采购单编号) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromPurchaserOrderID(purchaserOrderIDs []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderIDs(purchaserOrderIDs))
	return
}

// GetFromPurchaserOrderName 通过单个 purchaser_order_name(采购单名称) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromPurchaserOrderName(purchaserOrderName string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderName(purchaserOrderName))
	return
}

// GetsFromPurchaserOrderName 通过多个 purchaser_order_name(采购单名称) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromPurchaserOrderName(purchaserOrderNames []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderNames(purchaserOrderNames))
	return
}

// GetFromOrderTime 通过单个 order_time(订单创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromOrderTime(orderTime int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOrderTime(orderTime))
	return
}

// GetsFromOrderTime 通过多个 order_time(订单创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromOrderTime(orderTimes []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithOrderTimes(orderTimes))
	return
}

// GetFromState 通过单个 state(状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromState(state int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithState(state))
	return
}

// GetsFromState 通过多个 state(状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromState(states []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithStates(states))
	return
}

// GetFromURL 通过单个 url(采购单文件下载URL) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromURL(url string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithURL(url))
	return
}

// GetsFromURL 通过多个 url(采购单文件下载URL) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromURL(urls []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithURLs(urls))
	return
}

// GetFromCycle 通过单个 cycle(采购周期) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromCycle(cycle string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCycle(cycle))
	return
}

// GetsFromCycle 通过多个 cycle(采购周期) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromCycle(cycles []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCycles(cycles))
	return
}

// GetFromClassName 通过单个 class_name(商品品类) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromClassName(className string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithClassName(className))
	return
}

// GetsFromClassName 通过多个 class_name(商品品类) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromClassName(classNames []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithClassNames(classNames))
	return
}

// GetFromCount 通过单个 count(采购数量) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromCount(count int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCount(count))
	return
}

// GetsFromCount 通过多个 count(采购数量) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromCount(counts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCounts(counts))
	return
}

// GetFromAmount 通过单个 amount(采购金额) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromAmount(amount int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithAmount(amount))
	return
}

// GetsFromAmount 通过多个 amount(采购金额) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromAmount(amounts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithAmounts(amounts))
	return
}

// GetFromReceiveAreaCode 通过单个 receive_area_code(收货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromReceiveAreaCode(receiveAreaCode string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAreaCode(receiveAreaCode))
	return
}

// GetsFromReceiveAreaCode 通过多个 receive_area_code(收货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromReceiveAreaCode(receiveAreaCodes []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAreaCodes(receiveAreaCodes))
	return
}

// GetFromReceiveAddress 通过单个 receive_address(收货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromReceiveAddress(receiveAddress string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAddress(receiveAddress))
	return
}

// GetsFromReceiveAddress 通过多个 receive_address(收货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromReceiveAddress(receiveAddresss []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAddresss(receiveAddresss))
	return
}

// GetFromSpecial 通过单个 special(品质规格) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromSpecial(special string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithSpecial(special))
	return
}

// GetsFromSpecial 通过多个 special(品质规格) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromSpecial(specials []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithSpecials(specials))
	return
}

// GetFromRequireText 通过单个 require_text(采购要求) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromRequireText(requireText string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRequireText(requireText))
	return
}

// GetsFromRequireText 通过多个 require_text(采购要求) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromRequireText(requireTexts []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRequireTexts(requireTexts))
	return
}

// GetFromRate 通过单个 rate(签约服务费率) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromRate(rate string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRate(rate))
	return
}

// GetsFromRate 通过多个 rate(签约服务费率) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromRate(rates []string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRates(rates))
	return
}

// GetFromFeeAmount 通过单个 fee_amount(服务费【单位：分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromFeeAmount(feeAmount int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithFeeAmount(feeAmount))
	return
}

// GetsFromFeeAmount 通过多个 fee_amount(服务费【单位：分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromFeeAmount(feeAmounts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithFeeAmounts(feeAmounts))
	return
}

// GetFromBillAmount 通过单个 bill_amount(账单金额【单位：分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromBillAmount(billAmount int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithBillAmount(billAmount))
	return
}

// GetsFromBillAmount 通过多个 bill_amount(账单金额【单位：分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromBillAmount(billAmounts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithBillAmounts(billAmounts))
	return
}

// GetFromRealAmount 通过单个 real_amount(实际支付采购金额【单位分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromRealAmount(realAmount int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRealAmount(realAmount))
	return
}

// GetsFromRealAmount 通过多个 real_amount(实际支付采购金额【单位分】) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromRealAmount(realAmounts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithRealAmounts(realAmounts))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromCreatedAt(createdAt int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromCreatedAt(createdAts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromUpdatedAt(updatedAt int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetFromDeletedAt(deletedAt int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) GetsFromDeletedAt(deletedAts []int64) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderPreDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrderPre, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchIndexByUdxPoid 通过 purchaser_order_id 字段值，获取多条记录
func (obj *PurchaserOrderPreDAO) FetchIndexByUdxPoid(purchaserOrderID string) (results []*PurchaserOrderPre, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderID(purchaserOrderID))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
