package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:29
 * @Desc:   purchaser_org_rel 表
 */

// 高灯云租户ID关联采购商（一个或多个高灯云租户ID对应一个采购商）表
type PurchaserOrgRel struct {
	ID          int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                    // ID
	OrgID       int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                // 服务商租户ID
	OpenOrgID   int64  `gorm:"column:open_org_id;type:int(11);not null;DEFAULT '0'" json:"open_org_id"`      // 开放服务商租户ID
	PurchaserID string `gorm:"column:purchaser_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_id"` // 采购商ID
	CreatedAt   int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`        // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`        // 更新时间
	DeletedAt   int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`        // 删除时间
}

// 获取结构体对应的表名方法
func (m *PurchaserOrgRel) TableName() string {
	return "purchaser_org_rel"
}

// 实例化结构体对象
func NewPurchaserOrgRel() *PurchaserOrgRel {
	return &PurchaserOrgRel{}
}

// 获取主键的对应字段
func (m *PurchaserOrgRel) PrimaryKey() string {
	return PurchaserOrgRelColumns.ID
}

// 获取主键值
func (m *PurchaserOrgRel) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrgRelColumns = struct {
	ID          string
	OrgID       string
	OpenOrgID   string
	PurchaserID string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	OrgID:       "org_id",
	OpenOrgID:   "open_org_id",
	PurchaserID: "purchaser_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// 包含所有表字段的切片
var PurchaserOrgRelAllColumns = []string{
	PurchaserOrgRelColumns.ID,          // ID
	PurchaserOrgRelColumns.OrgID,       // 服务商租户ID
	PurchaserOrgRelColumns.OpenOrgID,   // 开放服务商租户ID
	PurchaserOrgRelColumns.PurchaserID, // 采购商ID
	PurchaserOrgRelColumns.CreatedAt,   // 创建时间
	PurchaserOrgRelColumns.UpdatedAt,   // 更新时间
	PurchaserOrgRelColumns.DeletedAt,   // 删除时间

}

// 设置值：ID
func (m *PurchaserOrgRel) SetID(v int64) {
	m.ID = v
}

// 设置值：服务商租户ID
func (m *PurchaserOrgRel) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：开放服务商租户ID
func (m *PurchaserOrgRel) SetOpenOrgID(v int64) {
	m.OpenOrgID = v
}

// 设置值：采购商ID
func (m *PurchaserOrgRel) SetPurchaserID(v string) {
	m.PurchaserID = v
}

// 设置值：创建时间
func (m *PurchaserOrgRel) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *PurchaserOrgRel) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *PurchaserOrgRel) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：ID
func (m *PurchaserOrgRel) GetID() int64 {
	return m.ID
}

// 获取值：服务商租户ID
func (m *PurchaserOrgRel) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：开放服务商租户ID
func (m *PurchaserOrgRel) GetOpenOrgID() int64 {
	return m.OpenOrgID
}

// 获取值：采购商ID
func (m *PurchaserOrgRel) GetPurchaserID() string {
	return m.PurchaserID
}

// 获取值：创建时间
func (m *PurchaserOrgRel) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *PurchaserOrgRel) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *PurchaserOrgRel) GetDeletedAt() int64 {
	return m.DeletedAt
}
