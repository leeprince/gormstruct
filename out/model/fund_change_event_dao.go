package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-15 19:53:04
 * @Desc:   fund_change_event 表的 DAO 层
 */

type FundChangeEventDAO struct {
	*_BaseDAO
}

// 初始化 FundChangeEventDAO
func NewFundChangeEventDAO(ctx context.Context, db *gorm.DB) *FundChangeEventDAO {
	if db == nil {
		panic(fmt.Errorf("FundChangeEventDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &FundChangeEventDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&FundChangeEvent{}),
			db:               db,
			model:            FundChangeEvent{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          FundChangeEventAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *FundChangeEventDAO) GetTableName() string {
	fundChangeEvent := &FundChangeEvent{}
	return fundChangeEvent.TableName()
}

// 存在则更新，否则插入
func (obj *FundChangeEventDAO) Save(fundChangeEvent *FundChangeEvent) (rowsAffected int64, err error) {
	if fundChangeEvent.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(fundChangeEvent, obj.WithID(fundChangeEvent.ID))
	}
	return obj.Create(fundChangeEvent)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *FundChangeEventDAO) Create(fundChangeEvent interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(fundChangeEvent)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithID(id uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithIDs(ids []uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ID] = ids })
}

// 设置 event_id(事件ID) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithEventID(eventID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventID] = eventID })
}

// 设置 event_id(事件ID) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithEventIDs(eventIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventID] = eventIDs })
}

// 设置 event_type(事件类型，区分大小写：Income(收入)、Expenditure(支出)) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithEventType(eventType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventType] = eventType })
}

// 设置 event_type(事件类型，区分大小写：Income(收入)、Expenditure(支出)) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithEventTypes(eventTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventType] = eventTypes })
}

// 设置 event_msg(事件说明) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithEventMsg(eventMsg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventMsg] = eventMsg })
}

// 设置 event_msg(事件说明) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithEventMsgs(eventMsgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.EventMsg] = eventMsgs })
}

// 设置 order_id(订单ID) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.OrderID] = orderID })
}

// 设置 order_id(订单ID) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.OrderID] = orderIDs })
}

// 设置 org_id(租户ID) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.OrgID] = orgID })
}

// 设置 org_id(租户ID) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.OrgID] = orgIDs })
}

// 设置 expenditure_bank_number(付款银行卡号) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankNumber(expenditureBankNumber string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ExpenditureBankNumber] = expenditureBankNumber })
}

// 设置 expenditure_bank_number(付款银行卡号) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankNumbers(expenditureBankNumbers []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ExpenditureBankNumber] = expenditureBankNumbers })
}

// 设置 expenditure_bank_account_name(付款银行账户名称) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankAccountName(expenditureBankAccountName string) Option {
	return queryOptionFunc(func(o *options) {
		o.queryMap[FundChangeEventColumns.ExpenditureBankAccountName] = expenditureBankAccountName
	})
}

// 设置 expenditure_bank_account_name(付款银行账户名称) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankAccountNames(expenditureBankAccountNames []string) Option {
	return queryOptionFunc(func(o *options) {
		o.queryMap[FundChangeEventColumns.ExpenditureBankAccountName] = expenditureBankAccountNames
	})
}

// 设置 expenditure_bank_name(付款银行名称) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankName(expenditureBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ExpenditureBankName] = expenditureBankName })
}

// 设置 expenditure_bank_name(付款银行名称) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithExpenditureBankNames(expenditureBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.ExpenditureBankName] = expenditureBankNames })
}

// 设置 income_bank_number(收款银行卡号) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankNumber(incomeBankNumber string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankNumber] = incomeBankNumber })
}

// 设置 income_bank_number(收款银行卡号) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankNumbers(incomeBankNumbers []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankNumber] = incomeBankNumbers })
}

// 设置 income_bank_account_name(收款银行账户名称) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankAccountName(incomeBankAccountName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankAccountName] = incomeBankAccountName })
}

// 设置 income_bank_account_name(收款银行账户名称) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankAccountNames(incomeBankAccountNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankAccountName] = incomeBankAccountNames })
}

// 设置 income_bank_name(收款银行名称) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankName(incomeBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankName] = incomeBankName })
}

// 设置 income_bank_name(收款银行名称) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithIncomeBankNames(incomeBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IncomeBankName] = incomeBankNames })
}

// 设置 trading_amount(交易金额，单位分) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithTradingAmount(tradingAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.TradingAmount] = tradingAmount })
}

// 设置 trading_amount(交易金额，单位分) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithTradingAmounts(tradingAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.TradingAmount] = tradingAmounts })
}

// 设置 trading_time(交易时间戳) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithTradingTime(tradingTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.TradingTime] = tradingTime })
}

// 设置 trading_time(交易时间戳) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithTradingTimes(tradingTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.TradingTime] = tradingTimes })
}

// 设置 is_statistic(是否统计过。0:未统计；1:已统计) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithIsStatistic(isStatistic int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IsStatistic] = isStatistic })
}

// 设置 is_statistic(是否统计过。0:未统计；1:已统计) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithIsStatistics(isStatistics []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.IsStatistic] = isStatistics })
}

// 设置 created_at(事件创建时间戳) 字段作为 option 条件
func (obj *FundChangeEventDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.CreatedAt] = createdAt })
}

// 设置 created_at(事件创建时间戳) 字段的切片作为 option 条件
func (obj *FundChangeEventDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[FundChangeEventColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *FundChangeEventDAO) GetByOption(opts ...Option) (result *FundChangeEvent, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *FundChangeEventDAO) GetByOptions(opts ...Option) (results []*FundChangeEvent, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *FundChangeEventDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      fundChangeEvent: 要更新的数据
//      opts: 更新的条件
func (obj *FundChangeEventDAO) UpdateByOption(fundChangeEvent *FundChangeEvent, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(fundChangeEvent)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *FundChangeEventDAO) GetFromID(id uint64) (result *FundChangeEvent, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromID(ids []uint64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 event_id(事件ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromEventID(eventID string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventID(eventID))
	return
}

// 通过多个 event_id(事件ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromEventID(eventIDs []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventIDs(eventIDs))
	return
}

// 通过单个 event_type(事件类型，区分大小写：Income(收入)、Expenditure(支出)) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromEventType(eventType string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventType(eventType))
	return
}

// 通过多个 event_type(事件类型，区分大小写：Income(收入)、Expenditure(支出)) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromEventType(eventTypes []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventTypes(eventTypes))
	return
}

// 通过单个 event_msg(事件说明) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromEventMsg(eventMsg string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventMsg(eventMsg))
	return
}

// 通过多个 event_msg(事件说明) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromEventMsg(eventMsgs []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithEventMsgs(eventMsgs))
	return
}

// 通过单个 order_id(订单ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromOrderID(orderID string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithOrderID(orderID))
	return
}

// 通过多个 order_id(订单ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromOrderID(orderIDs []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithOrderIDs(orderIDs))
	return
}

// 通过单个 org_id(租户ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromOrgID(orgID int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithOrgID(orgID))
	return
}

// 通过多个 org_id(租户ID) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromOrgID(orgIDs []int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithOrgIDs(orgIDs))
	return
}

// 通过单个 expenditure_bank_number(付款银行卡号) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromExpenditureBankNumber(expenditureBankNumber string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankNumber(expenditureBankNumber))
	return
}

// 通过多个 expenditure_bank_number(付款银行卡号) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromExpenditureBankNumber(expenditureBankNumbers []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankNumbers(expenditureBankNumbers))
	return
}

// 通过单个 expenditure_bank_account_name(付款银行账户名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromExpenditureBankAccountName(expenditureBankAccountName string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankAccountName(expenditureBankAccountName))
	return
}

// 通过多个 expenditure_bank_account_name(付款银行账户名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromExpenditureBankAccountName(expenditureBankAccountNames []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankAccountNames(expenditureBankAccountNames))
	return
}

// 通过单个 expenditure_bank_name(付款银行名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromExpenditureBankName(expenditureBankName string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankName(expenditureBankName))
	return
}

// 通过多个 expenditure_bank_name(付款银行名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromExpenditureBankName(expenditureBankNames []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithExpenditureBankNames(expenditureBankNames))
	return
}

// 通过单个 income_bank_number(收款银行卡号) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromIncomeBankNumber(incomeBankNumber string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankNumber(incomeBankNumber))
	return
}

// 通过多个 income_bank_number(收款银行卡号) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromIncomeBankNumber(incomeBankNumbers []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankNumbers(incomeBankNumbers))
	return
}

// 通过单个 income_bank_account_name(收款银行账户名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromIncomeBankAccountName(incomeBankAccountName string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankAccountName(incomeBankAccountName))
	return
}

// 通过多个 income_bank_account_name(收款银行账户名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromIncomeBankAccountName(incomeBankAccountNames []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankAccountNames(incomeBankAccountNames))
	return
}

// 通过单个 income_bank_name(收款银行名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromIncomeBankName(incomeBankName string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankName(incomeBankName))
	return
}

// 通过多个 income_bank_name(收款银行名称) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromIncomeBankName(incomeBankNames []string) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeBankNames(incomeBankNames))
	return
}

// 通过单个 trading_amount(交易金额，单位分) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromTradingAmount(tradingAmount int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithTradingAmount(tradingAmount))
	return
}

// 通过多个 trading_amount(交易金额，单位分) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromTradingAmount(tradingAmounts []int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithTradingAmounts(tradingAmounts))
	return
}

// 通过单个 trading_time(交易时间戳) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromTradingTime(tradingTime int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithTradingTime(tradingTime))
	return
}

// 通过多个 trading_time(交易时间戳) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromTradingTime(tradingTimes []int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithTradingTimes(tradingTimes))
	return
}

// 通过单个 is_statistic(是否统计过。0:未统计；1:已统计) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromIsStatistic(isStatistic int8) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIsStatistic(isStatistic))
	return
}

// 通过多个 is_statistic(是否统计过。0:未统计；1:已统计) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromIsStatistic(isStatistics []int8) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithIsStatistics(isStatistics))
	return
}

// 通过单个 created_at(事件创建时间戳) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetFromCreatedAt(createdAt int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(事件创建时间戳) 字段值，获取多条记录
func (obj *FundChangeEventDAO) GetsFromCreatedAt(createdAts []int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *FundChangeEventDAO) FetchByPrimaryKey(id uint64) (result *FundChangeEvent, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 order_id, org_id 字段值，获取单条记录
func (obj *FundChangeEventDAO) FetchUniqueIndexByUniqOrgidOrderid(orderID string, orgID int64) (result *FundChangeEvent, err error) {
	result, err = obj.GetByOption(
		obj.WithOrderID(orderID),
		obj.WithOrgID(orgID))
	return
}

// 通过 event_type, org_id, trading_time 字段值，获取多条记录
func (obj *FundChangeEventDAO) FetchIndexByIDxOrgidTimeType(eventType string, orgID int64, tradingTime int64) (results []*FundChangeEvent, err error) {
	results, err = obj.GetByOptions(
		obj.WithEventType(eventType),
		obj.WithOrgID(orgID),
		obj.WithTradingTime(tradingTime))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
