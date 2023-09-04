package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-09-05 01:42:00
 * @Desc:   order_info 表的 DAO 层
 */

type OrderInfoDAO struct {
	*_BaseDAO
}

// OrderInfoDAO 初始化
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

// GetTableName 获取表名称
func (obj *OrderInfoDAO) GetTableName() string {
	orderInfo := &OrderInfo{}
	return orderInfo.TableName()
}

// Save 存在则更新，否则插入
func (obj *OrderInfoDAO) Save(orderInfo *OrderInfo) (rowsAffected int64, err error) {
	if orderInfo.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(orderInfo, obj.WithID(orderInfo.ID))
	}
	return obj.Create(orderInfo)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *OrderInfoDAO) Create(orderInfo interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(orderInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *OrderInfoDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ID] = ids })
}

// WithOrderID 设置 order_id(订单ID) 字段作为 option 条件
func (obj *OrderInfoDAO) WithOrderID(orderID *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(订单ID) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithOrderIDs(orderIDs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.OrderID] = orderIDs })
}

// WithTicketNumber 设置 ticket_number(订单取票号) 字段作为 option 条件
func (obj *OrderInfoDAO) WithTicketNumber(ticketNumber *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.TicketNumber] = ticketNumber })
}

// WithTicketNumbers 设置 ticket_number(订单取票号) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithTicketNumbers(ticketNumbers []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.TicketNumber] = ticketNumbers })
}

// WithContactPhone 设置 contact_phone(订单联系电话) 字段作为 option 条件
func (obj *OrderInfoDAO) WithContactPhone(contactPhone *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ContactPhone] = contactPhone })
}

// WithContactPhones 设置 contact_phone(订单联系电话) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithContactPhones(contactPhones []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.ContactPhone] = contactPhones })
}

// WithIsTransfer 设置 is_transfer(是否换乘.1:true;2:false;) 字段作为 option 条件
func (obj *OrderInfoDAO) WithIsTransfer(isTransfer *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsTransfer] = isTransfer })
}

// WithIsTransfers 设置 is_transfer(是否换乘.1:true;2:false;) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIsTransfers(isTransfers []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsTransfer] = isTransfers })
}

// WithIsOccupySeat 设置 is_occupy_seat(是否占座.1:true;2:false;) 字段作为 option 条件
func (obj *OrderInfoDAO) WithIsOccupySeat(isOccupySeat *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsOccupySeat] = isOccupySeat })
}

// WithIsOccupySeats 设置 is_occupy_seat(是否占座.1:true;2:false;) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithIsOccupySeats(isOccupySeats []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.IsOccupySeat] = isOccupySeats })
}

// WithCompleteStatus 设置 complete_status(完成状态.1:未完成;2:出票成功;3:出票失败;) 字段作为 option 条件
func (obj *OrderInfoDAO) WithCompleteStatus(completeStatus *int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CompleteStatus] = completeStatus })
}

// WithCompleteStatuss 设置 complete_status(完成状态.1:未完成;2:出票成功;3:出票失败;) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithCompleteStatuss(completeStatuss []*int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CompleteStatus] = completeStatuss })
}

// WithFailReason 设置 fail_reason(出票失败原因) 字段作为 option 条件
func (obj *OrderInfoDAO) WithFailReason(failReason *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.FailReason] = failReason })
}

// WithFailReasons 设置 fail_reason(出票失败原因) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithFailReasons(failReasons []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.FailReason] = failReasons })
}

// WithMachineName 设置 machine_name(设备名称) 字段作为 option 条件
func (obj *OrderInfoDAO) WithMachineName(machineName *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.MachineName] = machineName })
}

// WithMachineNames 设置 machine_name(设备名称) 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithMachineNames(machineNames []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.MachineName] = machineNames })
}

// WithDeletedAt 设置 deleted_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithDeletedAt(deletedAt *time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithDeletedAts(deletedAts []*time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.DeletedAt] = deletedAts })
}

// WithUpdatedAt 设置 updated_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.UpdatedAt] = updatedAts })
}

// WithCreatedAt 设置 created_at() 字段作为 option 条件
func (obj *OrderInfoDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at() 字段的切片作为 option 条件
func (obj *OrderInfoDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[OrderInfoColumns.CreatedAt] = createdAts })
}

// WithDeletedAtIsNull 设置 DeletedAt(删除标记) 字段为NULL作为 option 条件
func (obj *OrderInfoDAO) WithDeletedAtIsNull() Option {
	return queryArgOptionFunc(func(o *options) { o.queryArg.query = fmt.Sprintf("%s IS NULL", OrderInfoColumns.DeletedAt) })
}

// GetByOption 函数选项模式获取单条记录
func (obj *OrderInfoDAO) GetByOption(opts ...Option) (result *OrderInfo, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *OrderInfoDAO) GetListByOption(opts ...Option) (results []*OrderInfo, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *OrderInfoDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *OrderInfoDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     orderInfo: 要更新的数据
//	     opts: 更新的条件
func (obj *OrderInfoDAO) UpdateByOption(orderInfo *OrderInfo, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(orderInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *OrderInfoDAO) GetFromID(id int64) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromID(ids []int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrderID 通过单个 order_id(订单ID) 字段值，获取单条记录
func (obj *OrderInfoDAO) GetFromOrderID(orderID *string) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(订单ID) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromOrderID(orderIDs []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromTicketNumber 通过单个 ticket_number(订单取票号) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromTicketNumber(ticketNumber *string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithTicketNumber(ticketNumber))
	return
}

// GetsFromTicketNumber 通过多个 ticket_number(订单取票号) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromTicketNumber(ticketNumbers []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithTicketNumbers(ticketNumbers))
	return
}

// GetFromContactPhone 通过单个 contact_phone(订单联系电话) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromContactPhone(contactPhone *string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithContactPhone(contactPhone))
	return
}

// GetsFromContactPhone 通过多个 contact_phone(订单联系电话) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromContactPhone(contactPhones []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithContactPhones(contactPhones))
	return
}

// GetFromIsTransfer 通过单个 is_transfer(是否换乘.1:true;2:false;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromIsTransfer(isTransfer *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIsTransfer(isTransfer))
	return
}

// GetsFromIsTransfer 通过多个 is_transfer(是否换乘.1:true;2:false;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromIsTransfer(isTransfers []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIsTransfers(isTransfers))
	return
}

// GetFromIsOccupySeat 通过单个 is_occupy_seat(是否占座.1:true;2:false;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromIsOccupySeat(isOccupySeat *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIsOccupySeat(isOccupySeat))
	return
}

// GetsFromIsOccupySeat 通过多个 is_occupy_seat(是否占座.1:true;2:false;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromIsOccupySeat(isOccupySeats []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIsOccupySeats(isOccupySeats))
	return
}

// GetFromCompleteStatus 通过单个 complete_status(完成状态.1:未完成;2:出票成功;3:出票失败;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromCompleteStatus(completeStatus *int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCompleteStatus(completeStatus))
	return
}

// GetsFromCompleteStatus 通过多个 complete_status(完成状态.1:未完成;2:出票成功;3:出票失败;) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromCompleteStatus(completeStatuss []*int64) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCompleteStatuss(completeStatuss))
	return
}

// GetFromFailReason 通过单个 fail_reason(出票失败原因) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromFailReason(failReason *string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithFailReason(failReason))
	return
}

// GetsFromFailReason 通过多个 fail_reason(出票失败原因) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromFailReason(failReasons []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithFailReasons(failReasons))
	return
}

// GetFromMachineName 通过单个 machine_name(设备名称) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromMachineName(machineName *string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithMachineName(machineName))
	return
}

// GetsFromMachineName 通过多个 machine_name(设备名称) 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromMachineName(machineNames []*string) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithMachineNames(machineNames))
	return
}

// GetFromDeletedAt 通过单个 deleted_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromDeletedAt(deletedAt *time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromDeletedAt(deletedAts []*time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromUpdatedAt(updatedAts []time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromCreatedAt 通过单个 created_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetFromCreatedAt(createdAt time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at() 字段值，获取多条记录
func (obj *OrderInfoDAO) GetsFromCreatedAt(createdAts []time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *OrderInfoDAO) FetchByPrimaryKey(id int64) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchUniqueByIDxOrderInfoOrderID 通过 order_id 字段值，获取单条记录
func (obj *OrderInfoDAO) FetchUniqueByIDxOrderInfoOrderID(orderID *string) (result *OrderInfo, err error) {
	result, err = obj.GetByOption(obj.WithOrderID(orderID))
	return
}

// FetchIndexByIDxOrderInfoDeletedAt 通过 deleted_at 字段值，获取多条记录
func (obj *OrderInfoDAO) FetchIndexByIDxOrderInfoDeletedAt(deletedAt *time.Time) (results []*OrderInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
