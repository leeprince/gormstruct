package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:36
 * @Desc:   settlement_order 表
 */

// 结算单表
type SettlementOrder struct {
	ID         int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                  // 主键id
	OrgID      string `gorm:"column:org_id;type:varchar(50);not null;DEFAULT ''" json:"org_id"`           // 一级供应商（平台企业）租户id
	OpenOrgID  string `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"` // 一级供应商（平台企业）开放租户id
	SettleID   string `gorm:"column:settle_id;type:varchar(50);not null;DEFAULT ''" json:"settle_id"`     // 结算单编号
	SettleTime int64  `gorm:"column:settle_time;type:int(11);not null;DEFAULT '0'" json:"settle_time"`    // 结算单创建时间
	URL        string `gorm:"column:url;type:varchar(255);not null;DEFAULT ''" json:"url"`                // 结算单文件下载url
	State      int    `gorm:"column:state;type:int(1);not null;DEFAULT '0'" json:"state"`                 // 结算单状态
	OrderIDs   string `gorm:"column:order_ids;type:varchar(255);not null;DEFAULT ''" json:"order_ids"`    // 关联采购订单号list
	CreatedAt  int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`      // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`      // 更新时间
	DeletedAt  int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`      // 删除时间
}

// 获取结构体对应的表名方法
func (m *SettlementOrder) TableName() string {
	return "settlement_order"
}

// 实例化结构体对象
func NewSettlementOrder() *SettlementOrder {
	return &SettlementOrder{}
}

// 获取主键的对应字段
func (m *SettlementOrder) PrimaryKey() string {
	return SettlementOrderColumns.ID
}

// 获取主键值
func (m *SettlementOrder) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var SettlementOrderColumns = struct {
	ID         string
	OrgID      string
	OpenOrgID  string
	SettleID   string
	SettleTime string
	URL        string
	State      string
	OrderIDs   string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	OrgID:      "org_id",
	OpenOrgID:  "open_org_id",
	SettleID:   "settle_id",
	SettleTime: "settle_time",
	URL:        "url",
	State:      "state",
	OrderIDs:   "order_ids",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// 包含所有表字段的切片
var SettlementOrderAllColumns = []string{
	SettlementOrderColumns.ID,         // 主键id
	SettlementOrderColumns.OrgID,      // 一级供应商（平台企业）租户id
	SettlementOrderColumns.OpenOrgID,  // 一级供应商（平台企业）开放租户id
	SettlementOrderColumns.SettleID,   // 结算单编号
	SettlementOrderColumns.SettleTime, // 结算单创建时间
	SettlementOrderColumns.URL,        // 结算单文件下载url
	SettlementOrderColumns.State,      // 结算单状态
	SettlementOrderColumns.OrderIDs,   // 关联采购订单号list
	SettlementOrderColumns.CreatedAt,  // 创建时间
	SettlementOrderColumns.UpdatedAt,  // 更新时间
	SettlementOrderColumns.DeletedAt,  // 删除时间

}

// 设置值：主键id
func (m *SettlementOrder) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id
func (m *SettlementOrder) SetOrgID(v string) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户id
func (m *SettlementOrder) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：结算单编号
func (m *SettlementOrder) SetSettleID(v string) {
	m.SettleID = v
}

// 设置值：结算单创建时间
func (m *SettlementOrder) SetSettleTime(v int64) {
	m.SettleTime = v
}

// 设置值：结算单文件下载url
func (m *SettlementOrder) SetURL(v string) {
	m.URL = v
}

// 设置值：结算单状态
func (m *SettlementOrder) SetState(v int) {
	m.State = v
}

// 设置值：关联采购订单号list
func (m *SettlementOrder) SetOrderIDs(v string) {
	m.OrderIDs = v
}

// 设置值：创建时间
func (m *SettlementOrder) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *SettlementOrder) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *SettlementOrder) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *SettlementOrder) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id
func (m *SettlementOrder) GetOrgID() string {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户id
func (m *SettlementOrder) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：结算单编号
func (m *SettlementOrder) GetSettleID() string {
	return m.SettleID
}

// 获取值：结算单创建时间
func (m *SettlementOrder) GetSettleTime() int64 {
	return m.SettleTime
}

// 获取值：结算单文件下载url
func (m *SettlementOrder) GetURL() string {
	return m.URL
}

// 获取值：结算单状态
func (m *SettlementOrder) GetState() int {
	return m.State
}

// 获取值：关联采购订单号list
func (m *SettlementOrder) GetOrderIDs() string {
	return m.OrderIDs
}

// 获取值：创建时间
func (m *SettlementOrder) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *SettlementOrder) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *SettlementOrder) GetDeletedAt() int64 {
	return m.DeletedAt
}
