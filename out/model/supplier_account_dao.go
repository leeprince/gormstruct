package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:43
 * @Desc:   supplier_account 表的 DAO 层
 */

type SupplierAccountDAO struct {
	*_BaseDAO
}

// SupplierAccountDAO 初始化
func NewSupplierAccountDAO(ctx context.Context, db *gorm.DB) *SupplierAccountDAO {
	if db == nil {
		panic(fmt.Errorf("SupplierAccountDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SupplierAccountDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&SupplierAccount{}),
			db:               db,
			model:            SupplierAccount{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          SupplierAccountAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *SupplierAccountDAO) GetTableName() string {
	supplierAccount := &SupplierAccount{}
	return supplierAccount.TableName()
}

// Save 存在则更新，否则插入
func (obj *SupplierAccountDAO) Save(supplierAccount *SupplierAccount) (rowsAffected int64, err error) {
	if supplierAccount.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(supplierAccount, obj.WithID(supplierAccount.ID))
	}
	return obj.Create(supplierAccount)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *SupplierAccountDAO) Create(supplierAccount interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(supplierAccount)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.ID] = ids })
}

// WithSupplierID 设置 supplier_id(供应商id) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithSupplierID(supplierID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.SupplierID] = supplierID })
}

// WithSupplierIDs 设置 supplier_id(供应商id) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithSupplierIDs(supplierIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.SupplierID] = supplierIDs })
}

// WithName 设置 name(账户名称) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Name] = name })
}

// WithNames 设置 name(账户名称) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Name] = names })
}

// WithBank 设置 bank(开户行) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithBank(bank string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Bank] = bank })
}

// WithBanks 设置 bank(开户行) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithBanks(banks []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Bank] = banks })
}

// WithAccountType 设置 account_type(账户类型：0 支付宝 1 银行卡 2 微信) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithAccountType(accountType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.AccountType] = accountType })
}

// WithAccountTypes 设置 account_type(账户类型：0 支付宝 1 银行卡 2 微信) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithAccountTypes(accountTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.AccountType] = accountTypes })
}

// WithNumber 设置 number(账号) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithNumber(number string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Number] = number })
}

// WithNumbers 设置 number(账号) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithNumbers(numbers []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.Number] = numbers })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *SupplierAccountDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *SupplierAccountDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierAccountColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *SupplierAccountDAO) GetByOption(opts ...Option) (result *SupplierAccount, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *SupplierAccountDAO) GetListByOption(opts ...Option) (results []*SupplierAccount, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *SupplierAccountDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *SupplierAccountDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      supplierAccount: 要更新的数据
//      opts: 更新的条件
func (obj *SupplierAccountDAO) UpdateByOption(supplierAccount *SupplierAccount, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(supplierAccount)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *SupplierAccountDAO) GetFromID(id int64) (result *SupplierAccount, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromID(ids []int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromSupplierID 通过单个 supplier_id(供应商id) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromSupplierID(supplierID string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierID(supplierID))
	return
}

// GetsFromSupplierID 通过多个 supplier_id(供应商id) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromSupplierID(supplierIDs []string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierIDs(supplierIDs))
	return
}

// GetFromName 通过单个 name(账户名称) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromName(name string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name(账户名称) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromName(names []string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromBank 通过单个 bank(开户行) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromBank(bank string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithBank(bank))
	return
}

// GetsFromBank 通过多个 bank(开户行) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromBank(banks []string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithBanks(banks))
	return
}

// GetFromAccountType 通过单个 account_type(账户类型：0 支付宝 1 银行卡 2 微信) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromAccountType(accountType string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithAccountType(accountType))
	return
}

// GetsFromAccountType 通过多个 account_type(账户类型：0 支付宝 1 银行卡 2 微信) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromAccountType(accountTypes []string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithAccountTypes(accountTypes))
	return
}

// GetFromNumber 通过单个 number(账号) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromNumber(number string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithNumber(number))
	return
}

// GetsFromNumber 通过多个 number(账号) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromNumber(numbers []string) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithNumbers(numbers))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromCreatedAt(createdAt int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromCreatedAt(createdAts []int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromUpdatedAt(updatedAt int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetFromDeletedAt(deletedAt int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierAccountDAO) GetsFromDeletedAt(deletedAts []int64) (results []*SupplierAccount, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *SupplierAccountDAO) FetchByPrimaryKey(id int64) (result *SupplierAccount, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
