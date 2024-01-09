package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:11
 * @Desc:   invoice_product 表
 */

// 发票关联的商品明细
type InvoiceProduct struct {
	ID               int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                             // 主键id
	InvoicePrimaryID int64  `gorm:"column:invoice_primary_id;type:int(11);not null;DEFAULT '0'" json:"invoice_primary_id"` // 发票数据表主键id
	ProjectName      string `gorm:"column:project_name;type:varchar(50);not null;DEFAULT ''" json:"project_name"`          // 商品名称
	TaxCode          string `gorm:"column:tax_code;type:varchar(50);not null;DEFAULT ''" json:"tax_code"`                  // 税收商品编码
	Type             string `gorm:"column:type;type:varchar(50);not null;DEFAULT ''" json:"type"`                          // 税收商品类别 需要提供字典编码
	Special          string `gorm:"column:special;type:varchar(50);not null;DEFAULT ''" json:"special"`                    // 商品规格
	Unit             string `gorm:"column:unit;type:varchar(10);not null;DEFAULT ''" json:"unit"`                          // 商品单位
	AmountNoTax      int64  `gorm:"column:amount_no_tax;type:int(11);not null;DEFAULT '0'" json:"amount_no_tax"`           // 商品不含税金额（单位：分）
	Count            int64  `gorm:"column:count;type:int(11);not null;DEFAULT '0'" json:"count"`                           // 商品数量
	PriceNoTax       int64  `gorm:"column:price_no_tax;type:int(11);not null;DEFAULT '0'" json:"price_no_tax"`             // 商品不含税单价（单位：分）
	TaxRate          int64  `gorm:"column:tax_rate;type:int(11);not null;DEFAULT '0'" json:"tax_rate"`                     // 商品税率
	TaxAmount        int64  `gorm:"column:tax_amount;type:int(11);not null;DEFAULT '0'" json:"tax_amount"`                 // 商品税额（单位：分）
	DiscountHasTax   int64  `gorm:"column:discount_has_tax;type:int(11);not null;DEFAULT '0'" json:"discount_has_tax"`     // 含税折扣总金额（单位：分）
	DiscountFlag     int64  `gorm:"column:discount_flag;type:int(11);not null;DEFAULT '0'" json:"discount_flag"`           // 优惠政策标志 0：不使用优惠政策 1：使用优惠政策
	ZeroFlag         int64  `gorm:"column:zero_flag;type:int(11);not null;DEFAULT '0'" json:"zero_flag"`                   // 零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率
	AddTaxFlag       string `gorm:"column:add_tax_flag;type:varchar(50);not null;DEFAULT ''" json:"add_tax_flag"`          // 增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等
	CreatedAt        int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                 // 创建时间
	UpdatedAt        int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                 // 更新时间
	DeletedAt        int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                 // 删除时间
}

// 获取结构体对应的表名方法
func (m *InvoiceProduct) TableName() string {
	return "invoice_product"
}

// 实例化结构体对象
func NewInvoiceProduct() *InvoiceProduct {
	return &InvoiceProduct{}
}

// 获取主键的对应字段
func (m *InvoiceProduct) PrimaryKey() string {
	return InvoiceProductColumns.ID
}

// 获取主键值
func (m *InvoiceProduct) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var InvoiceProductColumns = struct {
	ID               string
	InvoicePrimaryID string
	ProjectName      string
	TaxCode          string
	Type             string
	Special          string
	Unit             string
	AmountNoTax      string
	Count            string
	PriceNoTax       string
	TaxRate          string
	TaxAmount        string
	DiscountHasTax   string
	DiscountFlag     string
	ZeroFlag         string
	AddTaxFlag       string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	InvoicePrimaryID: "invoice_primary_id",
	ProjectName:      "project_name",
	TaxCode:          "tax_code",
	Type:             "type",
	Special:          "special",
	Unit:             "unit",
	AmountNoTax:      "amount_no_tax",
	Count:            "count",
	PriceNoTax:       "price_no_tax",
	TaxRate:          "tax_rate",
	TaxAmount:        "tax_amount",
	DiscountHasTax:   "discount_has_tax",
	DiscountFlag:     "discount_flag",
	ZeroFlag:         "zero_flag",
	AddTaxFlag:       "add_tax_flag",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// 包含所有表字段的切片
var InvoiceProductAllColumns = []string{
	InvoiceProductColumns.ID,               // 主键id
	InvoiceProductColumns.InvoicePrimaryID, // 发票数据表主键id
	InvoiceProductColumns.ProjectName,      // 商品名称
	InvoiceProductColumns.TaxCode,          // 税收商品编码
	InvoiceProductColumns.Type,             // 税收商品类别 需要提供字典编码
	InvoiceProductColumns.Special,          // 商品规格
	InvoiceProductColumns.Unit,             // 商品单位
	InvoiceProductColumns.AmountNoTax,      // 商品不含税金额（单位：分）
	InvoiceProductColumns.Count,            // 商品数量
	InvoiceProductColumns.PriceNoTax,       // 商品不含税单价（单位：分）
	InvoiceProductColumns.TaxRate,          // 商品税率
	InvoiceProductColumns.TaxAmount,        // 商品税额（单位：分）
	InvoiceProductColumns.DiscountHasTax,   // 含税折扣总金额（单位：分）
	InvoiceProductColumns.DiscountFlag,     // 优惠政策标志 0：不使用优惠政策 1：使用优惠政策
	InvoiceProductColumns.ZeroFlag,         // 零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率
	InvoiceProductColumns.AddTaxFlag,       // 增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等
	InvoiceProductColumns.CreatedAt,        // 创建时间
	InvoiceProductColumns.UpdatedAt,        // 更新时间
	InvoiceProductColumns.DeletedAt,        // 删除时间

}

// 设置值：主键id
func (m *InvoiceProduct) SetID(v int64) {
	m.ID = v
}

// 设置值：发票数据表主键id
func (m *InvoiceProduct) SetInvoicePrimaryID(v int64) {
	m.InvoicePrimaryID = v
}

// 设置值：商品名称
func (m *InvoiceProduct) SetProjectName(v string) {
	m.ProjectName = v
}

// 设置值：税收商品编码
func (m *InvoiceProduct) SetTaxCode(v string) {
	m.TaxCode = v
}

// 设置值：税收商品类别 需要提供字典编码
func (m *InvoiceProduct) SetType(v string) {
	m.Type = v
}

// 设置值：商品规格
func (m *InvoiceProduct) SetSpecial(v string) {
	m.Special = v
}

// 设置值：商品单位
func (m *InvoiceProduct) SetUnit(v string) {
	m.Unit = v
}

// 设置值：商品不含税金额（单位：分）
func (m *InvoiceProduct) SetAmountNoTax(v int64) {
	m.AmountNoTax = v
}

// 设置值：商品数量
func (m *InvoiceProduct) SetCount(v int64) {
	m.Count = v
}

// 设置值：商品不含税单价（单位：分）
func (m *InvoiceProduct) SetPriceNoTax(v int64) {
	m.PriceNoTax = v
}

// 设置值：商品税率
func (m *InvoiceProduct) SetTaxRate(v int64) {
	m.TaxRate = v
}

// 设置值：商品税额（单位：分）
func (m *InvoiceProduct) SetTaxAmount(v int64) {
	m.TaxAmount = v
}

// 设置值：含税折扣总金额（单位：分）
func (m *InvoiceProduct) SetDiscountHasTax(v int64) {
	m.DiscountHasTax = v
}

// 设置值：优惠政策标志 0：不使用优惠政策 1：使用优惠政策
func (m *InvoiceProduct) SetDiscountFlag(v int64) {
	m.DiscountFlag = v
}

// 设置值：零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率
func (m *InvoiceProduct) SetZeroFlag(v int64) {
	m.ZeroFlag = v
}

// 设置值：增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等
func (m *InvoiceProduct) SetAddTaxFlag(v string) {
	m.AddTaxFlag = v
}

// 设置值：创建时间
func (m *InvoiceProduct) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *InvoiceProduct) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *InvoiceProduct) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *InvoiceProduct) GetID() int64 {
	return m.ID
}

// 获取值：发票数据表主键id
func (m *InvoiceProduct) GetInvoicePrimaryID() int64 {
	return m.InvoicePrimaryID
}

// 获取值：商品名称
func (m *InvoiceProduct) GetProjectName() string {
	return m.ProjectName
}

// 获取值：税收商品编码
func (m *InvoiceProduct) GetTaxCode() string {
	return m.TaxCode
}

// 获取值：税收商品类别 需要提供字典编码
func (m *InvoiceProduct) GetType() string {
	return m.Type
}

// 获取值：商品规格
func (m *InvoiceProduct) GetSpecial() string {
	return m.Special
}

// 获取值：商品单位
func (m *InvoiceProduct) GetUnit() string {
	return m.Unit
}

// 获取值：商品不含税金额（单位：分）
func (m *InvoiceProduct) GetAmountNoTax() int64 {
	return m.AmountNoTax
}

// 获取值：商品数量
func (m *InvoiceProduct) GetCount() int64 {
	return m.Count
}

// 获取值：商品不含税单价（单位：分）
func (m *InvoiceProduct) GetPriceNoTax() int64 {
	return m.PriceNoTax
}

// 获取值：商品税率
func (m *InvoiceProduct) GetTaxRate() int64 {
	return m.TaxRate
}

// 获取值：商品税额（单位：分）
func (m *InvoiceProduct) GetTaxAmount() int64 {
	return m.TaxAmount
}

// 获取值：含税折扣总金额（单位：分）
func (m *InvoiceProduct) GetDiscountHasTax() int64 {
	return m.DiscountHasTax
}

// 获取值：优惠政策标志 0：不使用优惠政策 1：使用优惠政策
func (m *InvoiceProduct) GetDiscountFlag() int64 {
	return m.DiscountFlag
}

// 获取值：零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率
func (m *InvoiceProduct) GetZeroFlag() int64 {
	return m.ZeroFlag
}

// 获取值：增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等
func (m *InvoiceProduct) GetAddTaxFlag() string {
	return m.AddTaxFlag
}

// 获取值：创建时间
func (m *InvoiceProduct) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *InvoiceProduct) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *InvoiceProduct) GetDeletedAt() int64 {
	return m.DeletedAt
}
