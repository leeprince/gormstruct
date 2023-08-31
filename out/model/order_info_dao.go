package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-29 23:48:38
 * @Desc:   order_info 表的 DAO 层
 */

type OrderInfoDAO struct {
	*_BaseDAO
}

// 初始化 OrderInfoDAO
func NewOrderInfoDAO(ctx context.Context, db *gorm.DB) *OrderInfoDAO {
	if db == nil {
		panic(fmt.Errorf("OrderInfoDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &OrderInfoDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&OrderInfo{}),
			db:               db,
			model:            OrderInfo{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          OrderInfoAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *OrderInfoDAO) GetTableName() string {
	orderInfo := &OrderInfo{}
	return orderInfo.TableName()
}

// 存在则更新，否则插入
func (obj *OrderInfoDAO) Save(orderInfo *OrderInfo) (rowsAffected int64, err error) {
	if orderInfo.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(orderInfo, obj.WithID(orderInfo.ID))
	}
	return obj.Create(orderInfo)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *OrderInfoDAO) Create(orderInfo interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(orderInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键ID) 字段作为 option 条件
func (obj *OrderInfoDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ID] = id })
}

// 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ID] = ids })
}

// 设置 order_id(订单ID) 字段作为 option 条件
func (obj *OrderInfoDAO) WithOrderID(orderID *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.OrderID] = orderID })
}

// 设置 order_id(订单ID) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithOrderIDs(orderIDs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.OrderID] = orderIDs })
}

// 设置 ticket_number(订单取票号) 字段作为 option 条件
func (obj *OrderInfoDAO) WithTicketNumber(ticketNumber *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.TicketNumber] = ticketNumber })
}

// 设置 ticket_number(订单取票号) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithTicketNumbers(ticketNumbers []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.TicketNumber] = ticketNumbers })
}

// 设置 contact_phone(订单联系电话) 字段作为 option 条件
func (obj *OrderInfoDAO) WithContactPhone(contactPhone *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ContactPhone] = contactPhone })
}

// 设置 contact_phone(订单联系电话) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithContactPhones(contactPhones []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ContactPhone] = contactPhones })
}

// 设置 is_transfer(是否换乘.0:false) 字段作为 option 条件
func (obj *OrderInfoDAO) WithIsTransfer(isTransfer *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsTransfer] = isTransfer })
}

// 设置 is_transfer(是否换乘.0:false) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIsTransfers(isTransfers []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsTransfer] = isTransfers })
}

// 设置 is_occupy_seat(是否占座.0:false) 字段作为 option 条件
func (obj *OrderInfoDAO) WithIsOccupySeat(isOccupySeat *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsOccupySeat] = isOccupySeat })
}

// 设置 is_occupy_seat(是否占座.0:false) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIsOccupySeats(isOccupySeats []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsOccupySeat] = isOccupySeats })
}

// 设置 complete_status(完成状态.0:未完成;1:出票成功;2:出票失败;) 字段作为 option 条件
func (obj *OrderInfoDAO) WithCompleteStatus(completeStatus *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CompleteStatus] = completeStatus })
}

// 设置 complete_status(完成状态.0:未完成;1:出票成功;2:出票失败;) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithCompleteStatuss(completeStatuss []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CompleteStatus] = completeStatuss })
}

// 设置 fail_reason(出票失败原因) 字段作为 option 条件
func (obj *OrderInfoDAO) WithFailReason(failReason *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.FailReason] = failReason })
}

// 设置 fail_reason(出票失败原因) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithFailReasons(failReasons []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.FailReason] = failReasons })
}

// 设置 machine_name(设备名称) 字段作为 option 条件
func (obj *OrderInfoDAO) WithMachineName(machineName *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.MachineName] = machineName })
}

// 设置 machine_name(设备名称) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithMachineNames(machineNames []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.MachineName] = machineNames })
}

// 设置 deleted_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithDeletedAt(deletedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithDeletedAts(deletedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.DeletedAt] = deletedAts })
}

// 设置 updated_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.UpdatedAt] = updatedAts })
}

// 设置 created_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CreatedAt] = createdAt })
}

// 设置 created_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *OrderInfoDAO) GetByOption(opts ...Option) (result *OrderInfo, err error) {
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *OrderInfoDAO) GetByOptions(opts ...Option) (results []*OrderInfo, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *OrderInfoDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//
//	参数说明：
//	    orderInfo: 要更新的数据
//	    opts: 更新的条件
func (obj *OrderInfoDAO) UpdateByOption(orderInfo *OrderInfo, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(orderInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *OrderInfoDAO) GetFromID(id int64) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromID(ids []int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 order_id(订单ID) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromOrderID(orderID *string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithOrderID(orderID))
	return
}

// 通过多个 order_id(订单ID) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromOrderID(orderIDs []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithOrderIDs(orderIDs))
	return
}

// 通过单个 ticket_number(订单取票号) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromTicketNumber(ticketNumber *string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithTicketNumber(ticketNumber))
	return
}

// 通过多个 ticket_number(订单取票号) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromTicketNumber(ticketNumbers []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithTicketNumbers(ticketNumbers))
	return
}

// 通过单个 contact_phone(订单联系电话) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromContactPhone(contactPhone *string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithContactPhone(contactPhone))
	return
}

// 通过多个 contact_phone(订单联系电话) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromContactPhone(contactPhones []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithContactPhones(contactPhones))
	return
}

// 通过单个 is_transfer(是否换乘.0:false) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromIsTransfer(isTransfer *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithIsTransfer(isTransfer))
	return
}

// 通过多个 is_transfer(是否换乘.0:false) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromIsTransfer(isTransfers []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithIsTransfers(isTransfers))
	return
}

// 通过单个 is_occupy_seat(是否占座.0:false) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromIsOccupySeat(isOccupySeat *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithIsOccupySeat(isOccupySeat))
	return
}

// 通过多个 is_occupy_seat(是否占座.0:false) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromIsOccupySeat(isOccupySeats []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithIsOccupySeats(isOccupySeats))
	return
}

// 通过单个 complete_status(完成状态.0:未完成;1:出票成功;2:出票失败;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromCompleteStatus(completeStatus *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithCompleteStatus(completeStatus))
	return
}

// 通过多个 complete_status(完成状态.0:未完成;1:出票成功;2:出票失败;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromCompleteStatus(completeStatuss []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithCompleteStatuss(completeStatuss))
	return
}

// 通过单个 fail_reason(出票失败原因) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromFailReason(failReason *string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithFailReason(failReason))
	return
}

// 通过多个 fail_reason(出票失败原因) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromFailReason(failReasons []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithFailReasons(failReasons))
	return
}

// 通过单个 machine_name(设备名称) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromMachineName(machineName *string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithMachineName(machineName))
	return
}

// 通过多个 machine_name(设备名称) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromMachineName(machineNames []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithMachineNames(machineNames))
	return
}

// 通过单个 deleted_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromDeletedAt(deletedAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromDeletedAt(deletedAts []time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// 通过单个 updated_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromUpdatedAt(updatedAts []time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 created_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromCreatedAt(createdAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromCreatedAt(createdAts []time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *OrderInfoDAO) FetchByPrimaryKey(id int64) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 deleted_at 字段值，获取多条记录
func (obj *OrderInfoDAO) FetchIndexByIDxOrderInfoDeletedAt(deletedAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
