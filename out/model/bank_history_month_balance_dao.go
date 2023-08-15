package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-15 09:37:52
 * @Desc:   bank_history_month_balance 表的 DAO 层
 */

type BankHistoryMonthBalanceDAO struct {
	*_BaseDAO
}

// 初始化 BankHistoryMonthBalanceDAO
func NewBankHistoryMonthBalanceDAO(ctx context.Context, db *gorm.DB) *BankHistoryMonthBalanceDAO {
	if db == nil {
		panic(fmt.Errorf("BankHistoryMonthBalanceDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &BankHistoryMonthBalanceDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&BankHistoryMonthBalance{}),
			db:               db,
			model:            BankHistoryMonthBalance{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          BankHistoryMonthBalanceAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *BankHistoryMonthBalanceDAO) GetTableName() string {
	bankHistoryMonthBalance := &BankHistoryMonthBalance{}
	return bankHistoryMonthBalance.TableName()
}

// 存在则更新，否则插入
func (obj *BankHistoryMonthBalanceDAO) Save(bankHistoryMonthBalance *BankHistoryMonthBalance) (rowsAffected int64, err error) {
	if bankHistoryMonthBalance.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(bankHistoryMonthBalance, obj.WithID(bankHistoryMonthBalance.ID))
	}
	return obj.Create(bankHistoryMonthBalance)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *BankHistoryMonthBalanceDAO) Create(bankHistoryMonthBalance interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(bankHistoryMonthBalance)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithID(id uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithIDs(ids []uint64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.ID] = ids })
}

// 设置 org_id(租户ID) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.OrgID] = orgID })
}

// 设置 org_id(租户ID) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.OrgID] = orgIDs })
}

// 设置 bank_number(银行卡号) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBankNumber(bankNumber string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.BankNumber] = bankNumber })
}

// 设置 bank_number(银行卡号) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBankNumbers(bankNumbers []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.BankNumber] = bankNumbers })
}

// 设置 bank_account_name(银行账户名称) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBankAccountName(bankAccountName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.BankAccountName] = bankAccountName })
}

// 设置 bank_account_name(银行账户名称) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBankAccountNames(bankAccountNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.BankAccountName] = bankAccountNames })
}

// 设置 statistics_date(统计:年月,202308) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithStatisticsDate(statisticsDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.StatisticsDate] = statisticsDate })
}

// 设置 statistics_date(统计:年月,202308) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithStatisticsDates(statisticsDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.StatisticsDate] = statisticsDates })
}

// 设置 balance(资产总额（余额），单位分) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBalance(balance int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.Balance] = balance })
}

// 设置 balance(资产总额（余额），单位分) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithBalances(balances []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.Balance] = balances })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *BankHistoryMonthBalanceDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[BankHistoryMonthBalanceColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *BankHistoryMonthBalanceDAO) GetByOption(opts ...Option) (result *BankHistoryMonthBalance, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *BankHistoryMonthBalanceDAO) GetByOptions(opts ...Option) (results []*BankHistoryMonthBalance, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *BankHistoryMonthBalanceDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      bankHistoryMonthBalance: 要更新的数据
//      opts: 更新的条件
func (obj *BankHistoryMonthBalanceDAO) UpdateByOption(bankHistoryMonthBalance *BankHistoryMonthBalance, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(bankHistoryMonthBalance)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromID(id uint64) (result *BankHistoryMonthBalance, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromID(ids []uint64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 org_id(租户ID) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromOrgID(orgID int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithOrgID(orgID))
	return
}

// 通过多个 org_id(租户ID) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromOrgID(orgIDs []int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithOrgIDs(orgIDs))
	return
}

// 通过单个 bank_number(银行卡号) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromBankNumber(bankNumber string) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBankNumber(bankNumber))
	return
}

// 通过多个 bank_number(银行卡号) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromBankNumber(bankNumbers []string) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBankNumbers(bankNumbers))
	return
}

// 通过单个 bank_account_name(银行账户名称) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromBankAccountName(bankAccountName string) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBankAccountName(bankAccountName))
	return
}

// 通过多个 bank_account_name(银行账户名称) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromBankAccountName(bankAccountNames []string) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBankAccountNames(bankAccountNames))
	return
}

// 通过单个 statistics_date(统计:年月,202308) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromStatisticsDate(statisticsDate int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithStatisticsDate(statisticsDate))
	return
}

// 通过多个 statistics_date(统计:年月,202308) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromStatisticsDate(statisticsDates []int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithStatisticsDates(statisticsDates))
	return
}

// 通过单个 balance(资产总额（余额），单位分) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromBalance(balance int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBalance(balance))
	return
}

// 通过多个 balance(资产总额（余额），单位分) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromBalance(balances []int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithBalances(balances))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetFromCreatedAt(createdAt int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) GetsFromCreatedAt(createdAts []int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *BankHistoryMonthBalanceDAO) FetchByPrimaryKey(id uint64) (result *BankHistoryMonthBalance, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 org_id, statistics_date 字段值，获取多条记录
func (obj *BankHistoryMonthBalanceDAO) FetchIndexByIDxOrgidDate(orgID int64, statisticsDate int64) (results []*BankHistoryMonthBalance, err error) {
	results, err = obj.GetByOptions(
		obj.WithOrgID(orgID),
		obj.WithStatisticsDate(statisticsDate))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
