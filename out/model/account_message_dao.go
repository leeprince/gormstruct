package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-09-29 01:56:42
 * @Desc:   account_message 表的 DAO 层
 */

type AccountMessageDAO struct {
	*_BaseDAO
}

// 初始化 AccountMessageDAO
func NewAccountMessageDAO(ctx context.Context, db *gorm.DB) *AccountMessageDAO {
	if db == nil {
		panic(fmt.Errorf("AccountMessageDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &AccountMessageDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&AccountMessage{}),
			db:               db,
			model:            AccountMessage{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          AccountMessageAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *AccountMessageDAO) GetTableName() string {
	accountMessage := &AccountMessage{}
	return accountMessage.TableName()
}

// 存在则更新，否则插入
func (obj *AccountMessageDAO) Save(accountMessage *AccountMessage) (rowsAffected int64, err error) {
	if accountMessage.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(accountMessage, obj.WithID(accountMessage.ID))
	}
	return obj.Create(accountMessage)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *AccountMessageDAO) Create(accountMessage interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(accountMessage)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键ID,消息ID) 字段作为 option 条件
func (obj *AccountMessageDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ID] = id })
}

// 设置 id(主键ID,消息ID) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ID] = ids })
}

// 设置 org_id(租户ID) 字段作为 option 条件
func (obj *AccountMessageDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.OrgID] = orgID })
}

// 设置 org_id(租户ID) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.OrgID] = orgIDs })
}

// 设置 client_id(商户ID,租户ID) 字段作为 option 条件
func (obj *AccountMessageDAO) WithClientID(clientID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ClientID] = clientID })
}

// 设置 client_id(商户ID,租户ID) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithClientIDs(clientIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ClientID] = clientIDs })
}

// 设置 client_name(商户名称) 字段作为 option 条件
func (obj *AccountMessageDAO) WithClientName(clientName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ClientName] = clientName })
}

// 设置 client_name(商户名称) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithClientNames(clientNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.ClientName] = clientNames })
}

// 设置 order_id(高灯云订单号) 字段作为 option 条件
func (obj *AccountMessageDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.OrderID] = orderID })
}

// 设置 order_id(高灯云订单号) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.OrderID] = orderIDs })
}

// 设置 message_type(消息类型(1:到账; 2:系统; 3:活动)；固定为1) 字段作为 option 条件
func (obj *AccountMessageDAO) WithMessageType(messageType int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.MessageType] = messageType })
}

// 设置 message_type(消息类型(1:到账; 2:系统; 3:活动)；固定为1) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithMessageTypes(messageTypes []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.MessageType] = messageTypes })
}

// 设置 real_amount(实付金额。单位分) 字段作为 option 条件
func (obj *AccountMessageDAO) WithRealAmount(realAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.RealAmount] = realAmount })
}

// 设置 real_amount(实付金额。单位分) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithRealAmounts(realAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.RealAmount] = realAmounts })
}

// 设置 status(打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功) 字段作为 option 条件
func (obj *AccountMessageDAO) WithStatus(status int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.Status] = status })
}

// 设置 status(打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithStatuss(statuss []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.Status] = statuss })
}

// 设置 fail_reason(失败原因) 字段作为 option 条件
func (obj *AccountMessageDAO) WithFailReason(failReason string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.FailReason] = failReason })
}

// 设置 fail_reason(失败原因) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithFailReasons(failReasons []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.FailReason] = failReasons })
}

// 设置 payee_bank_account(打款银行卡号) 字段作为 option 条件
func (obj *AccountMessageDAO) WithPayeeBankAccount(payeeBankAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.PayeeBankAccount] = payeeBankAccount })
}

// 设置 payee_bank_account(打款银行卡号) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithPayeeBankAccounts(payeeBankAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.PayeeBankAccount] = payeeBankAccounts })
}

// 设置 is_read(消息是否已读；0：未读；1：已读) 字段作为 option 条件
func (obj *AccountMessageDAO) WithIsRead(isRead int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.IsRead] = isRead })
}

// 设置 is_read(消息是否已读；0：未读；1：已读) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithIsReads(isReads []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.IsRead] = isReads })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *AccountMessageDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.CreatedAt] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *AccountMessageDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.UpdatedAt] = updatedAts })
}

// 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *AccountMessageDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *AccountMessageDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[AccountMessageColumns.DeletedAt] = deletedAts })
}

// 函数选项模式获取单条记录
func (obj *AccountMessageDAO) GetByOption(opts ...Option) (result *AccountMessage, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *AccountMessageDAO) GetByOptions(opts ...Option) (results []*AccountMessage, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *AccountMessageDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      accountMessage: 要更新的数据
//      opts: 更新的条件
func (obj *AccountMessageDAO) UpdateByOption(accountMessage *AccountMessage, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(accountMessage)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键ID,消息ID) 字段值，获取单条记录
func (obj *AccountMessageDAO) GetFromID(id int64) (result *AccountMessage, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键ID,消息ID) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromID(ids []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 org_id(租户ID) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromOrgID(orgID int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithOrgID(orgID))
	return
}

// 通过多个 org_id(租户ID) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromOrgID(orgIDs []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithOrgIDs(orgIDs))
	return
}

// 通过单个 client_id(商户ID,租户ID) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromClientID(clientID int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithClientID(clientID))
	return
}

// 通过多个 client_id(商户ID,租户ID) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromClientID(clientIDs []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithClientIDs(clientIDs))
	return
}

// 通过单个 client_name(商户名称) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromClientName(clientName string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithClientName(clientName))
	return
}

// 通过多个 client_name(商户名称) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromClientName(clientNames []string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithClientNames(clientNames))
	return
}

// 通过单个 order_id(高灯云订单号) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromOrderID(orderID string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithOrderID(orderID))
	return
}

// 通过多个 order_id(高灯云订单号) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromOrderID(orderIDs []string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithOrderIDs(orderIDs))
	return
}

// 通过单个 message_type(消息类型(1:到账; 2:系统; 3:活动)；固定为1) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromMessageType(messageType int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithMessageType(messageType))
	return
}

// 通过多个 message_type(消息类型(1:到账; 2:系统; 3:活动)；固定为1) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromMessageType(messageTypes []int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithMessageTypes(messageTypes))
	return
}

// 通过单个 real_amount(实付金额。单位分) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromRealAmount(realAmount int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithRealAmount(realAmount))
	return
}

// 通过多个 real_amount(实付金额。单位分) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromRealAmount(realAmounts []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithRealAmounts(realAmounts))
	return
}

// 通过单个 status(打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromStatus(status int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithStatus(status))
	return
}

// 通过多个 status(打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromStatus(statuss []int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithStatuss(statuss))
	return
}

// 通过单个 fail_reason(失败原因) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromFailReason(failReason string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithFailReason(failReason))
	return
}

// 通过多个 fail_reason(失败原因) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromFailReason(failReasons []string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithFailReasons(failReasons))
	return
}

// 通过单个 payee_bank_account(打款银行卡号) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromPayeeBankAccount(payeeBankAccount string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithPayeeBankAccount(payeeBankAccount))
	return
}

// 通过多个 payee_bank_account(打款银行卡号) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromPayeeBankAccount(payeeBankAccounts []string) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithPayeeBankAccounts(payeeBankAccounts))
	return
}

// 通过单个 is_read(消息是否已读；0：未读；1：已读) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromIsRead(isRead int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithIsRead(isRead))
	return
}

// 通过多个 is_read(消息是否已读；0：未读；1：已读) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromIsRead(isReads []int8) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithIsReads(isReads))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromCreatedAt(createdAt int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromCreatedAt(createdAts []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromUpdatedAt(updatedAt int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetFromDeletedAt(deletedAt int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *AccountMessageDAO) GetsFromDeletedAt(deletedAts []int64) (results []*AccountMessage, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *AccountMessageDAO) FetchByPrimaryKey(id int64) (result *AccountMessage, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
