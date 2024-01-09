package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:53
 * @Desc:   purchaser_order_invoice_rel 表
 */

// 发票关联的采购订单表
type PurchaserOrderInvoiceRel struct {
	ID               int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                             // 主键id
	InvoicePrimaryID int64  `gorm:"column:invoice_primary_id;type:int(11);not null;DEFAULT '0'" json:"invoice_primary_id"` // 发票数据表主键id
	OrderID          string `gorm:"column:order_id;type:varchar(50);not null;DEFAULT ''" json:"order_id"`                  // 采购订单ID
	CreatedAt        int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                 // 创建时间
	UpdatedAt        int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                 // 更新时间
	DeletedAt        int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                 // 删除时间
}

// 获取结构体对应的表名方法
func (m *PurchaserOrderInvoiceRel) TableName() string {
	return "purchaser_order_invoice_rel"
}

// 实例化结构体对象
func NewPurchaserOrderInvoiceRel() *PurchaserOrderInvoiceRel {
	return &PurchaserOrderInvoiceRel{}
}

// 获取主键的对应字段
func (m *PurchaserOrderInvoiceRel) PrimaryKey() string {
	return PurchaserOrderInvoiceRelColumns.ID
}

// 获取主键值
func (m *PurchaserOrderInvoiceRel) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrderInvoiceRelColumns = struct {
	ID               string
	InvoicePrimaryID string
	OrderID          string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	InvoicePrimaryID: "invoice_primary_id",
	OrderID:          "order_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// 包含所有表字段的切片
var PurchaserOrderInvoiceRelAllColumns = []string{
	PurchaserOrderInvoiceRelColumns.ID,               // 主键id
	PurchaserOrderInvoiceRelColumns.InvoicePrimaryID, // 发票数据表主键id
	PurchaserOrderInvoiceRelColumns.OrderID,          // 采购订单ID
	PurchaserOrderInvoiceRelColumns.CreatedAt,        // 创建时间
	PurchaserOrderInvoiceRelColumns.UpdatedAt,        // 更新时间
	PurchaserOrderInvoiceRelColumns.DeletedAt,        // 删除时间

}

// 设置值：主键id
func (m *PurchaserOrderInvoiceRel) SetID(v int64) {
	m.ID = v
}

// 设置值：发票数据表主键id
func (m *PurchaserOrderInvoiceRel) SetInvoicePrimaryID(v int64) {
	m.InvoicePrimaryID = v
}

// 设置值：采购订单ID
func (m *PurchaserOrderInvoiceRel) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：创建时间
func (m *PurchaserOrderInvoiceRel) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *PurchaserOrderInvoiceRel) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *PurchaserOrderInvoiceRel) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *PurchaserOrderInvoiceRel) GetID() int64 {
	return m.ID
}

// 获取值：发票数据表主键id
func (m *PurchaserOrderInvoiceRel) GetInvoicePrimaryID() int64 {
	return m.InvoicePrimaryID
}

// 获取值：采购订单ID
func (m *PurchaserOrderInvoiceRel) GetOrderID() string {
	return m.OrderID
}

// 获取值：创建时间
func (m *PurchaserOrderInvoiceRel) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *PurchaserOrderInvoiceRel) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *PurchaserOrderInvoiceRel) GetDeletedAt() int64 {
	return m.DeletedAt
}
