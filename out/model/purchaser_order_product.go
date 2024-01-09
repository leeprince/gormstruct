package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:07
 * @Desc:   purchaser_order_product 表
 */

// 采购订单关联商品表
type PurchaserOrderProduct struct {
	ID         int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                // 主键id
	OrderID    string `gorm:"column:order_id;type:varchar(50);not null;DEFAULT ''" json:"order_id"`     // 订单编号
	ClassName  string `gorm:"column:class_name;type:varchar(50);not null;DEFAULT ''" json:"class_name"` // 采购商品分类
	Name       string `gorm:"column:name;type:varchar(50);not null;DEFAULT ''" json:"name"`             // 采购商品名称
	Special    string `gorm:"column:special;type:varchar(50);not null;DEFAULT ''" json:"special"`       // 规格型号
	Unit       string `gorm:"column:unit;type:varchar(10);not null;DEFAULT ''" json:"unit"`             // 单位
	Count      int64  `gorm:"column:count;type:int(11);not null;DEFAULT '0'" json:"count"`              // 数量
	Price      int64  `gorm:"column:price;type:int(11);not null;DEFAULT '0'" json:"price"`              // 单价（单位：分）
	RealCount  int64  `gorm:"column:real_count;type:int(11);not null;DEFAULT '0'" json:"real_count"`    // 实际采购数量
	RealPrice  int64  `gorm:"column:real_price;type:int(11);not null;DEFAULT '0'" json:"real_price"`    // 实际采购单价（单位：分）
	RealAmount int64  `gorm:"column:real_amount;type:int(11);not null;DEFAULT '0'" json:"real_amount"`  // 实际采购金额（单位：分）
	Remark     string `gorm:"column:remark;type:varchar(255);not null;DEFAULT ''" json:"remark"`        // 采购商品描述
	CreatedAt  int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`    // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`    // 更新时间
	DeletedAt  int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`    // 删除时间
}

// 获取结构体对应的表名方法
func (m *PurchaserOrderProduct) TableName() string {
	return "purchaser_order_product"
}

// 实例化结构体对象
func NewPurchaserOrderProduct() *PurchaserOrderProduct {
	return &PurchaserOrderProduct{}
}

// 获取主键的对应字段
func (m *PurchaserOrderProduct) PrimaryKey() string {
	return PurchaserOrderProductColumns.ID
}

// 获取主键值
func (m *PurchaserOrderProduct) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrderProductColumns = struct {
	ID         string
	OrderID    string
	ClassName  string
	Name       string
	Special    string
	Unit       string
	Count      string
	Price      string
	RealCount  string
	RealPrice  string
	RealAmount string
	Remark     string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}{
	ID:         "id",
	OrderID:    "order_id",
	ClassName:  "class_name",
	Name:       "name",
	Special:    "special",
	Unit:       "unit",
	Count:      "count",
	Price:      "price",
	RealCount:  "real_count",
	RealPrice:  "real_price",
	RealAmount: "real_amount",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// 包含所有表字段的切片
var PurchaserOrderProductAllColumns = []string{
	PurchaserOrderProductColumns.ID,         // 主键id
	PurchaserOrderProductColumns.OrderID,    // 订单编号
	PurchaserOrderProductColumns.ClassName,  // 采购商品分类
	PurchaserOrderProductColumns.Name,       // 采购商品名称
	PurchaserOrderProductColumns.Special,    // 规格型号
	PurchaserOrderProductColumns.Unit,       // 单位
	PurchaserOrderProductColumns.Count,      // 数量
	PurchaserOrderProductColumns.Price,      // 单价（单位：分）
	PurchaserOrderProductColumns.RealCount,  // 实际采购数量
	PurchaserOrderProductColumns.RealPrice,  // 实际采购单价（单位：分）
	PurchaserOrderProductColumns.RealAmount, // 实际采购金额（单位：分）
	PurchaserOrderProductColumns.Remark,     // 采购商品描述
	PurchaserOrderProductColumns.CreatedAt,  // 创建时间
	PurchaserOrderProductColumns.UpdatedAt,  // 更新时间
	PurchaserOrderProductColumns.DeletedAt,  // 删除时间

}

// 设置值：主键id
func (m *PurchaserOrderProduct) SetID(v int64) {
	m.ID = v
}

// 设置值：订单编号
func (m *PurchaserOrderProduct) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：采购商品分类
func (m *PurchaserOrderProduct) SetClassName(v string) {
	m.ClassName = v
}

// 设置值：采购商品名称
func (m *PurchaserOrderProduct) SetName(v string) {
	m.Name = v
}

// 设置值：规格型号
func (m *PurchaserOrderProduct) SetSpecial(v string) {
	m.Special = v
}

// 设置值：单位
func (m *PurchaserOrderProduct) SetUnit(v string) {
	m.Unit = v
}

// 设置值：数量
func (m *PurchaserOrderProduct) SetCount(v int64) {
	m.Count = v
}

// 设置值：单价（单位：分）
func (m *PurchaserOrderProduct) SetPrice(v int64) {
	m.Price = v
}

// 设置值：实际采购数量
func (m *PurchaserOrderProduct) SetRealCount(v int64) {
	m.RealCount = v
}

// 设置值：实际采购单价（单位：分）
func (m *PurchaserOrderProduct) SetRealPrice(v int64) {
	m.RealPrice = v
}

// 设置值：实际采购金额（单位：分）
func (m *PurchaserOrderProduct) SetRealAmount(v int64) {
	m.RealAmount = v
}

// 设置值：采购商品描述
func (m *PurchaserOrderProduct) SetRemark(v string) {
	m.Remark = v
}

// 设置值：创建时间
func (m *PurchaserOrderProduct) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *PurchaserOrderProduct) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *PurchaserOrderProduct) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *PurchaserOrderProduct) GetID() int64 {
	return m.ID
}

// 获取值：订单编号
func (m *PurchaserOrderProduct) GetOrderID() string {
	return m.OrderID
}

// 获取值：采购商品分类
func (m *PurchaserOrderProduct) GetClassName() string {
	return m.ClassName
}

// 获取值：采购商品名称
func (m *PurchaserOrderProduct) GetName() string {
	return m.Name
}

// 获取值：规格型号
func (m *PurchaserOrderProduct) GetSpecial() string {
	return m.Special
}

// 获取值：单位
func (m *PurchaserOrderProduct) GetUnit() string {
	return m.Unit
}

// 获取值：数量
func (m *PurchaserOrderProduct) GetCount() int64 {
	return m.Count
}

// 获取值：单价（单位：分）
func (m *PurchaserOrderProduct) GetPrice() int64 {
	return m.Price
}

// 获取值：实际采购数量
func (m *PurchaserOrderProduct) GetRealCount() int64 {
	return m.RealCount
}

// 获取值：实际采购单价（单位：分）
func (m *PurchaserOrderProduct) GetRealPrice() int64 {
	return m.RealPrice
}

// 获取值：实际采购金额（单位：分）
func (m *PurchaserOrderProduct) GetRealAmount() int64 {
	return m.RealAmount
}

// 获取值：采购商品描述
func (m *PurchaserOrderProduct) GetRemark() string {
	return m.Remark
}

// 获取值：创建时间
func (m *PurchaserOrderProduct) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *PurchaserOrderProduct) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *PurchaserOrderProduct) GetDeletedAt() int64 {
	return m.DeletedAt
}
