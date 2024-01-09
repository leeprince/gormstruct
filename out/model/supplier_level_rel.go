package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:57
 * @Desc:   supplier_level_rel 表
 */

// 一级供应商关联二级供应商（一个或多个一级对应一个二级）表
type SupplierLevelRel struct {
	ID          int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                    // 主键id
	OrgID       int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                // 一级供应商（平台企业）租户ID
	OpenOrgID   string `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"`   // 一级供应商（平台企业）开放租户ID
	SupplierID  string `gorm:"column:supplier_id;type:varchar(50);not null;DEFAULT ''" json:"supplier_id"`   // 二级供应商ID
	PurchaserID string `gorm:"column:purchaser_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_id"` // 采购商ID
	CreatedAt   int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`        // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`        // 更新时间
	DeletedAt   int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`        // 删除时间
}

// 获取结构体对应的表名方法
func (m *SupplierLevelRel) TableName() string {
	return "supplier_level_rel"
}

// 实例化结构体对象
func NewSupplierLevelRel() *SupplierLevelRel {
	return &SupplierLevelRel{}
}

// 获取主键的对应字段
func (m *SupplierLevelRel) PrimaryKey() string {
	return SupplierLevelRelColumns.ID
}

// 获取主键值
func (m *SupplierLevelRel) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var SupplierLevelRelColumns = struct {
	ID          string
	OrgID       string
	OpenOrgID   string
	SupplierID  string
	PurchaserID string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	OrgID:       "org_id",
	OpenOrgID:   "open_org_id",
	SupplierID:  "supplier_id",
	PurchaserID: "purchaser_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// 包含所有表字段的切片
var SupplierLevelRelAllColumns = []string{
	SupplierLevelRelColumns.ID,          // 主键id
	SupplierLevelRelColumns.OrgID,       // 一级供应商（平台企业）租户ID
	SupplierLevelRelColumns.OpenOrgID,   // 一级供应商（平台企业）开放租户ID
	SupplierLevelRelColumns.SupplierID,  // 二级供应商ID
	SupplierLevelRelColumns.PurchaserID, // 采购商ID
	SupplierLevelRelColumns.CreatedAt,   // 创建时间
	SupplierLevelRelColumns.UpdatedAt,   // 更新时间
	SupplierLevelRelColumns.DeletedAt,   // 删除时间

}

// 设置值：主键id
func (m *SupplierLevelRel) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户ID
func (m *SupplierLevelRel) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户ID
func (m *SupplierLevelRel) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：二级供应商ID
func (m *SupplierLevelRel) SetSupplierID(v string) {
	m.SupplierID = v
}

// 设置值：采购商ID
func (m *SupplierLevelRel) SetPurchaserID(v string) {
	m.PurchaserID = v
}

// 设置值：创建时间
func (m *SupplierLevelRel) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *SupplierLevelRel) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *SupplierLevelRel) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *SupplierLevelRel) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户ID
func (m *SupplierLevelRel) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户ID
func (m *SupplierLevelRel) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：二级供应商ID
func (m *SupplierLevelRel) GetSupplierID() string {
	return m.SupplierID
}

// 获取值：采购商ID
func (m *SupplierLevelRel) GetPurchaserID() string {
	return m.PurchaserID
}

// 获取值：创建时间
func (m *SupplierLevelRel) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *SupplierLevelRel) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *SupplierLevelRel) GetDeletedAt() int64 {
	return m.DeletedAt
}
