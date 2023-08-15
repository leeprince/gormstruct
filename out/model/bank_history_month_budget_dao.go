package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-15 19:59:34
 * @Desc:   bank_history_month_budget 表的 DAO 层
 */

type BankHistoryMonthBudgetDAO struct {
	*_BaseDAO
}

// 初始化 BankHistoryMonthBudgetDAO
func NewBankHistoryMonthBudgetDAO(ctx context.Context, db *gorm.DB) *BankHistoryMonthBudgetDAO {
	if db == nil {
		panic(fmt.Errorf("BankHistoryMonthBudgetDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &BankHistoryMonthBudgetDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&BankHistoryMonthBudget{}),
			db:               db,
			model:            BankHistoryMonthBudget{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          BankHistoryMonthBudgetAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *BankHistoryMonthBudgetDAO) GetTableName() string {
	bankHistoryMonthBudget := &BankHistoryMonthBudget{}
	return bankHistoryMonthBudget.TableName()
}

// 存在则更新，否则插入
func (obj *BankHistoryMonthBudgetDAO) Save(bankHistoryMonthBudget *BankHistoryMonthBudget) (rowsAffected int64, err error) {
	if bankHistoryMonthBudget.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(bankHistoryMonthBudget, obj.WithID(bankHistoryMonthBudget.ID))
	}
	return obj.Create(bankHistoryMonthBudget)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *BankHistoryMonthBudgetDAO) Create(bankHistoryMonthBudget interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(bankHistoryMonthBudget)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithID(id uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithIDs(ids []uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.ID] = ids })
}

// 设置 org_id(租户ID) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.OrgID] = orgID })
}

// 设置 org_id(租户ID) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.OrgID] = orgIDs })
}

// 设置 income_expenditure_type(收支类型：1(收入)、2(支出)) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithIncomeExpenditureType(incomeExpenditureType int8) Option {
	return queryOptionFunc(func(o *options) {
		o.queryMap[BankHistoryMonthBudgetColumns.IncomeExpenditureType] = incomeExpenditureType
	})
}

// 设置 income_expenditure_type(收支类型：1(收入)、2(支出)) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithIncomeExpenditureTypes(incomeExpenditureTypes []int8) Option {
	return queryOptionFunc(func(o *options) {
		o.queryMap[BankHistoryMonthBudgetColumns.IncomeExpenditureType] = incomeExpenditureTypes
	})
}

// 设置 bank_number(银行卡号) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithBankNumber(bankNumber string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.BankNumber] = bankNumber })
}

// 设置 bank_number(银行卡号) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithBankNumbers(bankNumbers []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.BankNumber] = bankNumbers })
}

// 设置 bank_account_name(银行账户名称) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithBankAccountName(bankAccountName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.BankAccountName] = bankAccountName })
}

// 设置 bank_account_name(银行账户名称) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithBankAccountNames(bankAccountNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.BankAccountName] = bankAccountNames })
}

// 设置 statistics_date(统计:年月,202308) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithStatisticsDate(statisticsDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.StatisticsDate] = statisticsDate })
}

// 设置 statistics_date(统计:年月,202308) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithStatisticsDates(statisticsDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.StatisticsDate] = statisticsDates })
}

// 设置 trading_total_amount(交易总金额，单位分) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithTradingTotalAmount(tradingTotalAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.TradingTotalAmount] = tradingTotalAmount })
}

// 设置 trading_total_amount(交易总金额，单位分) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithTradingTotalAmounts(tradingTotalAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.TradingTotalAmount] = tradingTotalAmounts })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBudgetDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBudgetColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *BankHistoryMonthBudgetDAO) GetByOption(opts ...Option) (result *BankHistoryMonthBudget, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *BankHistoryMonthBudgetDAO) GetByOptions(opts ...Option) (results []*BankHistoryMonthBudget, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *BankHistoryMonthBudgetDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      bankHistoryMonthBudget: 要更新的数据
//      opts: 更新的条件
func (obj *BankHistoryMonthBudgetDAO) UpdateByOption(bankHistoryMonthBudget *BankHistoryMonthBudget, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(bankHistoryMonthBudget)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromID(id uint64) (result *BankHistoryMonthBudget, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromID(ids []uint64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 org_id(租户ID) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromOrgID(orgID int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithOrgID(orgID))
	return
}

// 通过多个 org_id(租户ID) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromOrgID(orgIDs []int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithOrgIDs(orgIDs))
	return
}

// 通过单个 income_expenditure_type(收支类型：1(收入)、2(支出)) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromIncomeExpenditureType(incomeExpenditureType int8) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeExpenditureType(incomeExpenditureType))
	return
}

// 通过多个 income_expenditure_type(收支类型：1(收入)、2(支出)) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromIncomeExpenditureType(incomeExpenditureTypes []int8) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithIncomeExpenditureTypes(incomeExpenditureTypes))
	return
}

// 通过单个 bank_number(银行卡号) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromBankNumber(bankNumber string) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithBankNumber(bankNumber))
	return
}

// 通过多个 bank_number(银行卡号) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromBankNumber(bankNumbers []string) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithBankNumbers(bankNumbers))
	return
}

// 通过单个 bank_account_name(银行账户名称) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromBankAccountName(bankAccountName string) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithBankAccountName(bankAccountName))
	return
}

// 通过多个 bank_account_name(银行账户名称) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromBankAccountName(bankAccountNames []string) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithBankAccountNames(bankAccountNames))
	return
}

// 通过单个 statistics_date(统计:年月,202308) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromStatisticsDate(statisticsDate int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithStatisticsDate(statisticsDate))
	return
}

// 通过多个 statistics_date(统计:年月,202308) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromStatisticsDate(statisticsDates []int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithStatisticsDates(statisticsDates))
	return
}

// 通过单个 trading_total_amount(交易总金额，单位分) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromTradingTotalAmount(tradingTotalAmount int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithTradingTotalAmount(tradingTotalAmount))
	return
}

// 通过多个 trading_total_amount(交易总金额，单位分) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromTradingTotalAmount(tradingTotalAmounts []int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithTradingTotalAmounts(tradingTotalAmounts))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetFromCreatedAt(createdAt int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) GetsFromCreatedAt(createdAts []int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *BankHistoryMonthBudgetDAO) FetchByPrimaryKey(id uint64) (result *BankHistoryMonthBudget, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 org_id, income_expenditure_type, statistics_date 字段值，获取多条记录
func (obj *BankHistoryMonthBudgetDAO) FetchIndexByIDxOrgidDateType(orgID int64, incomeExpenditureType int8, statisticsDate int64) (results []*BankHistoryMonthBudget, err error) {
	results, err = obj.GetByOptions(
		obj.WithOrgID(orgID),
		obj.WithIncomeExpenditureType(incomeExpenditureType),
		obj.WithStatisticsDate(statisticsDate))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
