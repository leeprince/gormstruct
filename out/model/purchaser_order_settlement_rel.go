package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:14
 * @Desc:   purchaser_order_settlement_rel 表
 */

// 采购订单表关联结算单（一个或多个采购订单关联一个结算单）表
type PurchaserOrderSettlementRel struct {
	ID        int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`              // 主键id
	OrderID   string `gorm:"column:order_id;type:varchar(50);not null;DEFAULT ''" json:"order_id"`   // 订单编号
	SettleID  string `gorm:"column:settle_id;type:varchar(50);not null;DEFAULT ''" json:"settle_id"` // 结算单ID
	CreatedAt int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`  // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`  // 更新时间
	DeletedAt int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`  // 删除时间
}

// 获取结构体对应的表名方法
func (m *PurchaserOrderSettlementRel) TableName() string {
	return "purchaser_order_settlement_rel"
}

// 实例化结构体对象
func NewPurchaserOrderSettlementRel() *PurchaserOrderSettlementRel {
	return &PurchaserOrderSettlementRel{}
}

// 获取主键的对应字段
func (m *PurchaserOrderSettlementRel) PrimaryKey() string {
	return PurchaserOrderSettlementRelColumns.ID
}

// 获取主键值
func (m *PurchaserOrderSettlementRel) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrderSettlementRelColumns = struct {
	ID        string
	OrderID   string
	SettleID  string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	OrderID:   "order_id",
	SettleID:  "settle_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// 包含所有表字段的切片
var PurchaserOrderSettlementRelAllColumns = []string{
	PurchaserOrderSettlementRelColumns.ID,        // 主键id
	PurchaserOrderSettlementRelColumns.OrderID,   // 订单编号
	PurchaserOrderSettlementRelColumns.SettleID,  // 结算单ID
	PurchaserOrderSettlementRelColumns.CreatedAt, // 创建时间
	PurchaserOrderSettlementRelColumns.UpdatedAt, // 更新时间
	PurchaserOrderSettlementRelColumns.DeletedAt, // 删除时间

}

// 设置值：主键id
func (m *PurchaserOrderSettlementRel) SetID(v int64) {
	m.ID = v
}

// 设置值：订单编号
func (m *PurchaserOrderSettlementRel) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：结算单ID
func (m *PurchaserOrderSettlementRel) SetSettleID(v string) {
	m.SettleID = v
}

// 设置值：创建时间
func (m *PurchaserOrderSettlementRel) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *PurchaserOrderSettlementRel) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *PurchaserOrderSettlementRel) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *PurchaserOrderSettlementRel) GetID() int64 {
	return m.ID
}

// 获取值：订单编号
func (m *PurchaserOrderSettlementRel) GetOrderID() string {
	return m.OrderID
}

// 获取值：结算单ID
func (m *PurchaserOrderSettlementRel) GetSettleID() string {
	return m.SettleID
}

// 获取值：创建时间
func (m *PurchaserOrderSettlementRel) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *PurchaserOrderSettlementRel) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *PurchaserOrderSettlementRel) GetDeletedAt() int64 {
	return m.DeletedAt
}
