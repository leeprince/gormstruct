package model

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-29 23:48:38
 * @Desc:   order_info 表
 */

// 订单统计表
type OrderInfo struct {
	ID             int64     `gorm:"column:id;primaryKey;type:bigint(11);autoIncrement" json:"id"`           // 主键ID
	OrderID        *string   `gorm:"column:order_id;type:varchar(32);is null;" json:"order_id"`              // 订单ID
	TicketNumber   *string   `gorm:"column:ticket_number;type:varchar(12);is null;" json:"ticket_number"`    // 订单取票号
	ContactPhone   *string   `gorm:"column:contact_phone;type:varchar(11);is null;" json:"contact_phone"`    // 订单联系电话
	IsTransfer     *int64    `gorm:"column:is_transfer;type:bigint(20);is null;" json:"is_transfer"`         // 是否换乘.0:false
	IsOccupySeat   *int64    `gorm:"column:is_occupy_seat;type:bigint(20);is null;" json:"is_occupy_seat"`   // 是否占座.0:false
	CompleteStatus *int64    `gorm:"column:complete_status;type:bigint(20);is null;" json:"complete_status"` // 完成状态.0:未完成;1:出票成功;2:出票失败;
	FailReason     *string   `gorm:"column:fail_reason;type:varchar(255);is null;" json:"fail_reason"`       // 出票失败原因
	MachineName    *string   `gorm:"column:machine_name;type:varchar(128);is null;" json:"machine_name"`     // 设备名称
	DeletedAt      time.Time `gorm:"column:deleted_at;type:datetime(3);is null;" json:"deleted_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime(3);is null;" json:"updated_at"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime(3);is null;" json:"created_at"`
}

// 获取结构体对应的表名方法
func (m *OrderInfo) TableName() string {
	return "order_info"
}

// 实例化结构体对象
func NewOrderInfo() *OrderInfo {
	return &OrderInfo{}
}

// 获取主键的对应字段
func (m *OrderInfo) PrimaryKey() string {
	return OrderInfoColumns.ID
}

// 获取主键值
func (m *OrderInfo) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var OrderInfoColumns = struct {
	ID             string
	OrderID        string
	TicketNumber   string
	ContactPhone   string
	IsTransfer     string
	IsOccupySeat   string
	CompleteStatus string
	FailReason     string
	MachineName    string
	DeletedAt      string
	UpdatedAt      string
	CreatedAt      string
}{
	ID:             "id",
	OrderID:        "order_id",
	TicketNumber:   "ticket_number",
	ContactPhone:   "contact_phone",
	IsTransfer:     "is_transfer",
	IsOccupySeat:   "is_occupy_seat",
	CompleteStatus: "complete_status",
	FailReason:     "fail_reason",
	MachineName:    "machine_name",
	DeletedAt:      "deleted_at",
	UpdatedAt:      "updated_at",
	CreatedAt:      "created_at",
}

// 包含所有表字段的切片
var OrderInfoAllColumns = []string{
	OrderInfoColumns.ID,             // 主键ID
	OrderInfoColumns.OrderID,        // 订单ID
	OrderInfoColumns.TicketNumber,   // 订单取票号
	OrderInfoColumns.ContactPhone,   // 订单联系电话
	OrderInfoColumns.IsTransfer,     // 是否换乘.0:false
	OrderInfoColumns.IsOccupySeat,   // 是否占座.0:false
	OrderInfoColumns.CompleteStatus, // 完成状态.0:未完成;1:出票成功;2:出票失败;
	OrderInfoColumns.FailReason,     // 出票失败原因
	OrderInfoColumns.MachineName,    // 设备名称
	OrderInfoColumns.DeletedAt,      //
	OrderInfoColumns.UpdatedAt,      //
	OrderInfoColumns.CreatedAt,      //

}

// 设置值：主键ID
func (m *OrderInfo) SetID(v int64) {
	m.ID = v
}

// 设置值：订单ID
func (m *OrderInfo) SetOrderID(v *string) {
	m.OrderID = v
}

// 设置值：订单取票号
func (m *OrderInfo) SetTicketNumber(v *string) {
	m.TicketNumber = v
}

// 设置值：订单联系电话
func (m *OrderInfo) SetContactPhone(v *string) {
	m.ContactPhone = v
}

// 设置值：是否换乘.0:false
func (m *OrderInfo) SetIsTransfer(v *int64) {
	m.IsTransfer = v
}

// 设置值：是否占座.0:false
func (m *OrderInfo) SetIsOccupySeat(v *int64) {
	m.IsOccupySeat = v
}

// 设置值：完成状态.0:未完成;1:出票成功;2:出票失败;
func (m *OrderInfo) SetCompleteStatus(v *int64) {
	m.CompleteStatus = v
}

// 设置值：出票失败原因
func (m *OrderInfo) SetFailReason(v *string) {
	m.FailReason = v
}

// 设置值：设备名称
func (m *OrderInfo) SetMachineName(v *string) {
	m.MachineName = v
}

// 设置值：
func (m *OrderInfo) SetDeletedAt(v time.Time) {
	m.DeletedAt = v
}

// 设置值：
func (m *OrderInfo) SetUpdatedAt(v time.Time) {
	m.UpdatedAt = v
}

// 设置值：
func (m *OrderInfo) SetCreatedAt(v time.Time) {
	m.CreatedAt = v
}

// 获取值：主键ID
func (m *OrderInfo) GetID() int64 {
	return m.ID
}

// 获取值：订单ID
func (m *OrderInfo) GetOrderID() *string {
	return m.OrderID
}

// 获取值：订单取票号
func (m *OrderInfo) GetTicketNumber() *string {
	return m.TicketNumber
}

// 获取值：订单联系电话
func (m *OrderInfo) GetContactPhone() *string {
	return m.ContactPhone
}

// 获取值：是否换乘.0:false
func (m *OrderInfo) GetIsTransfer() *int64 {
	return m.IsTransfer
}

// 获取值：是否占座.0:false
func (m *OrderInfo) GetIsOccupySeat() *int64 {
	return m.IsOccupySeat
}

// 获取值：完成状态.0:未完成;1:出票成功;2:出票失败;
func (m *OrderInfo) GetCompleteStatus() *int64 {
	return m.CompleteStatus
}

// 获取值：出票失败原因
func (m *OrderInfo) GetFailReason() *string {
	return m.FailReason
}

// 获取值：设备名称
func (m *OrderInfo) GetMachineName() *string {
	return m.MachineName
}

// 获取值：
func (m *OrderInfo) GetDeletedAt() time.Time {
	return m.DeletedAt
}

// 获取值：
func (m *OrderInfo) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// 获取值：
func (m *OrderInfo) GetCreatedAt() time.Time {
	return m.CreatedAt
}
