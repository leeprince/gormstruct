package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:18
 * @Desc:   voucher_info 表的 DAO 层
 */

type VoucherInfoDAO struct {
	*_BaseDAO
}

// VoucherInfoDAO 初始化
func NewVoucherInfoDAO(ctx context.Context, db *gorm.DB) *VoucherInfoDAO {
	if db == nil {
		panic(fmt.Errorf("VoucherInfoDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &VoucherInfoDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&VoucherInfo{}),
			db:               db,
			model:            VoucherInfo{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          VoucherInfoAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *VoucherInfoDAO) GetTableName() string {
	voucherInfo := &VoucherInfo{}
	return voucherInfo.TableName()
}

// Save 存在则更新，否则插入
func (obj *VoucherInfoDAO) Save(voucherInfo *VoucherInfo) (rowsAffected int64, err error) {
	if voucherInfo.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(voucherInfo, obj.WithID(voucherInfo.ID))
	}
	return obj.Create(voucherInfo)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *VoucherInfoDAO) Create(voucherInfo interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(voucherInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户id) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户id) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.OpenOrgID] = openOrgIDs })
}

// WithPayID 设置 pay_id(支付单编号) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithPayID(payID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayID] = payID })
}

// WithPayIDs 设置 pay_id(支付单编号) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithPayIDs(payIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayID] = payIDs })
}

// WithAmount 设置 amount(金额（单位：分）) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithAmount(amount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Amount] = amount })
}

// WithAmounts 设置 amount(金额（单位：分）) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithAmounts(amounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Amount] = amounts })
}

// WithChannel 设置 channel(支付通道 0-银行卡 1-微信 2-支付宝 3-其他) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithChannel(channel string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Channel] = channel })
}

// WithChannels 设置 channel(支付通道 0-银行卡 1-微信 2-支付宝 3-其他) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithChannels(channels []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Channel] = channels })
}

// WithPayName 设置 pay_name(付款方名称) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithPayName(payName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayName] = payName })
}

// WithPayNames 设置 pay_name(付款方名称) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithPayNames(payNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayName] = payNames })
}

// WithPayBankName 设置 pay_bank_name(付款方银行名称) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithPayBankName(payBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayBankName] = payBankName })
}

// WithPayBankNames 设置 pay_bank_name(付款方银行名称) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithPayBankNames(payBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayBankName] = payBankNames })
}

// WithPayBankAccount 设置 pay_bank_account(付款方账号) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithPayBankAccount(payBankAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayBankAccount] = payBankAccount })
}

// WithPayBankAccounts 设置 pay_bank_account(付款方账号) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithPayBankAccounts(payBankAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.PayBankAccount] = payBankAccounts })
}

// WithReceiveName 设置 receive_name(收款方名称) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveName(receiveName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveName] = receiveName })
}

// WithReceiveNames 设置 receive_name(收款方名称) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveNames(receiveNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveName] = receiveNames })
}

// WithReceiveBankName 设置 receive_bank_name(收款方银行名称) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveBankName(receiveBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveBankName] = receiveBankName })
}

// WithReceiveBankNames 设置 receive_bank_name(收款方银行名称) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveBankNames(receiveBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveBankName] = receiveBankNames })
}

// WithReceiveBankAccount 设置 receive_bank_account(收款方账号) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveBankAccount(receiveBankAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveBankAccount] = receiveBankAccount })
}

// WithReceiveBankAccounts 设置 receive_bank_account(收款方账号) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithReceiveBankAccounts(receiveBankAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.ReceiveBankAccount] = receiveBankAccounts })
}

// WithURL 设置 url(支付凭证下载URL) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithURL(url string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.URL] = url })
}

// WithURLs 设置 url(支付凭证下载URL) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithURLs(urls []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.URL] = urls })
}

// WithStartTime 设置 start_time(支付开始时间) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithStartTime(startTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.StartTime] = startTime })
}

// WithStartTimes 设置 start_time(支付开始时间) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithStartTimes(startTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.StartTime] = startTimes })
}

// WithSuccessTime 设置 success_time(支付成功时间) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithSuccessTime(successTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.SuccessTime] = successTime })
}

// WithSuccessTimes 设置 success_time(支付成功时间) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithSuccessTimes(successTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.SuccessTime] = successTimes })
}

// WithVoucherStatus 设置 voucher_status(支付状态 0-支付成功 1-支付失败 2-支付中) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithVoucherStatus(voucherStatus int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.VoucherStatus] = voucherStatus })
}

// WithVoucherStatuss 设置 voucher_status(支付状态 0-支付成功 1-支付失败 2-支付中) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithVoucherStatuss(voucherStatuss []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.VoucherStatus] = voucherStatuss })
}

// WithRemark 设置 remark(业务备注) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithRemark(remark string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Remark] = remark })
}

// WithRemarks 设置 remark(业务备注) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithRemarks(remarks []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.Remark] = remarks })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *VoucherInfoDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *VoucherInfoDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[VoucherInfoColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *VoucherInfoDAO) GetByOption(opts ...Option) (result *VoucherInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *VoucherInfoDAO) GetListByOption(opts ...Option) (results []*VoucherInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *VoucherInfoDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *VoucherInfoDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      voucherInfo: 要更新的数据
//      opts: 更新的条件
func (obj *VoucherInfoDAO) UpdateByOption(voucherInfo *VoucherInfo, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(voucherInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *VoucherInfoDAO) GetFromID(id int64) (result *VoucherInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromID(ids []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromOrgID(orgID int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromOrgID(orgIDs []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromOpenOrgID(openOrgID string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromPayID 通过单个 pay_id(支付单编号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromPayID(payID string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayID(payID))
	return
}

// GetsFromPayID 通过多个 pay_id(支付单编号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromPayID(payIDs []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayIDs(payIDs))
	return
}

// GetFromAmount 通过单个 amount(金额（单位：分）) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromAmount(amount int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmount(amount))
	return
}

// GetsFromAmount 通过多个 amount(金额（单位：分）) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromAmount(amounts []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmounts(amounts))
	return
}

// GetFromChannel 通过单个 channel(支付通道 0-银行卡 1-微信 2-支付宝 3-其他) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromChannel(channel string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithChannel(channel))
	return
}

// GetsFromChannel 通过多个 channel(支付通道 0-银行卡 1-微信 2-支付宝 3-其他) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromChannel(channels []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithChannels(channels))
	return
}

// GetFromPayName 通过单个 pay_name(付款方名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromPayName(payName string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayName(payName))
	return
}

// GetsFromPayName 通过多个 pay_name(付款方名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromPayName(payNames []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayNames(payNames))
	return
}

// GetFromPayBankName 通过单个 pay_bank_name(付款方银行名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromPayBankName(payBankName string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayBankName(payBankName))
	return
}

// GetsFromPayBankName 通过多个 pay_bank_name(付款方银行名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromPayBankName(payBankNames []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayBankNames(payBankNames))
	return
}

// GetFromPayBankAccount 通过单个 pay_bank_account(付款方账号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromPayBankAccount(payBankAccount string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayBankAccount(payBankAccount))
	return
}

// GetsFromPayBankAccount 通过多个 pay_bank_account(付款方账号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromPayBankAccount(payBankAccounts []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayBankAccounts(payBankAccounts))
	return
}

// GetFromReceiveName 通过单个 receive_name(收款方名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromReceiveName(receiveName string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveName(receiveName))
	return
}

// GetsFromReceiveName 通过多个 receive_name(收款方名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromReceiveName(receiveNames []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveNames(receiveNames))
	return
}

// GetFromReceiveBankName 通过单个 receive_bank_name(收款方银行名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromReceiveBankName(receiveBankName string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveBankName(receiveBankName))
	return
}

// GetsFromReceiveBankName 通过多个 receive_bank_name(收款方银行名称) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromReceiveBankName(receiveBankNames []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveBankNames(receiveBankNames))
	return
}

// GetFromReceiveBankAccount 通过单个 receive_bank_account(收款方账号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromReceiveBankAccount(receiveBankAccount string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveBankAccount(receiveBankAccount))
	return
}

// GetsFromReceiveBankAccount 通过多个 receive_bank_account(收款方账号) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromReceiveBankAccount(receiveBankAccounts []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveBankAccounts(receiveBankAccounts))
	return
}

// GetFromURL 通过单个 url(支付凭证下载URL) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromURL(url string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithURL(url))
	return
}

// GetsFromURL 通过多个 url(支付凭证下载URL) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromURL(urls []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithURLs(urls))
	return
}

// GetFromStartTime 通过单个 start_time(支付开始时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromStartTime(startTime int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithStartTime(startTime))
	return
}

// GetsFromStartTime 通过多个 start_time(支付开始时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromStartTime(startTimes []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithStartTimes(startTimes))
	return
}

// GetFromSuccessTime 通过单个 success_time(支付成功时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromSuccessTime(successTime int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSuccessTime(successTime))
	return
}

// GetsFromSuccessTime 通过多个 success_time(支付成功时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromSuccessTime(successTimes []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSuccessTimes(successTimes))
	return
}

// GetFromVoucherStatus 通过单个 voucher_status(支付状态 0-支付成功 1-支付失败 2-支付中) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromVoucherStatus(voucherStatus int) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithVoucherStatus(voucherStatus))
	return
}

// GetsFromVoucherStatus 通过多个 voucher_status(支付状态 0-支付成功 1-支付失败 2-支付中) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromVoucherStatus(voucherStatuss []int) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithVoucherStatuss(voucherStatuss))
	return
}

// GetFromRemark 通过单个 remark(业务备注) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromRemark(remark string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithRemark(remark))
	return
}

// GetsFromRemark 通过多个 remark(业务备注) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromRemark(remarks []string) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithRemarks(remarks))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromCreatedAt(createdAt int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromCreatedAt(createdAts []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromUpdatedAt(updatedAt int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetFromDeletedAt(deletedAt int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *VoucherInfoDAO) GetsFromDeletedAt(deletedAts []int64) (results []*VoucherInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *VoucherInfoDAO) FetchByPrimaryKey(id int64) (result *VoucherInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
