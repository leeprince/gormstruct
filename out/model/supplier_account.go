package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:43
 * @Desc:   supplier_account 表
 */

// 二级供应商的账户列表
type SupplierAccount struct {
	ID          int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                    // 主键id
	SupplierID  string `gorm:"column:supplier_id;type:varchar(50);not null;DEFAULT ''" json:"supplier_id"`   // 供应商id
	Name        string `gorm:"column:name;type:varchar(50);not null;DEFAULT ''" json:"name"`                 // 账户名称
	Bank        string `gorm:"column:bank;type:varchar(50);not null;DEFAULT ''" json:"bank"`                 // 开户行
	AccountType string `gorm:"column:account_type;type:varchar(50);not null;DEFAULT ''" json:"account_type"` // 账户类型：0 支付宝 1 银行卡 2 微信
	Number      string `gorm:"column:number;type:varchar(50);not null;DEFAULT ''" json:"number"`             // 账号
	CreatedAt   int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`        // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`        // 更新时间
	DeletedAt   int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`        // 删除时间
}

// 获取结构体对应的表名方法
func (m *SupplierAccount) TableName() string {
	return "supplier_account"
}

// 实例化结构体对象
func NewSupplierAccount() *SupplierAccount {
	return &SupplierAccount{}
}

// 获取主键的对应字段
func (m *SupplierAccount) PrimaryKey() string {
	return SupplierAccountColumns.ID
}

// 获取主键值
func (m *SupplierAccount) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var SupplierAccountColumns = struct {
	ID          string
	SupplierID  string
	Name        string
	Bank        string
	AccountType string
	Number      string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}{
	ID:          "id",
	SupplierID:  "supplier_id",
	Name:        "name",
	Bank:        "bank",
	AccountType: "account_type",
	Number:      "number",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// 包含所有表字段的切片
var SupplierAccountAllColumns = []string{
	SupplierAccountColumns.ID,          // 主键id
	SupplierAccountColumns.SupplierID,  // 供应商id
	SupplierAccountColumns.Name,        // 账户名称
	SupplierAccountColumns.Bank,        // 开户行
	SupplierAccountColumns.AccountType, // 账户类型：0 支付宝 1 银行卡 2 微信
	SupplierAccountColumns.Number,      // 账号
	SupplierAccountColumns.CreatedAt,   // 创建时间
	SupplierAccountColumns.UpdatedAt,   // 更新时间
	SupplierAccountColumns.DeletedAt,   // 删除时间

}

// 设置值：主键id
func (m *SupplierAccount) SetID(v int64) {
	m.ID = v
}

// 设置值：供应商id
func (m *SupplierAccount) SetSupplierID(v string) {
	m.SupplierID = v
}

// 设置值：账户名称
func (m *SupplierAccount) SetName(v string) {
	m.Name = v
}

// 设置值：开户行
func (m *SupplierAccount) SetBank(v string) {
	m.Bank = v
}

// 设置值：账户类型：0 支付宝 1 银行卡 2 微信
func (m *SupplierAccount) SetAccountType(v string) {
	m.AccountType = v
}

// 设置值：账号
func (m *SupplierAccount) SetNumber(v string) {
	m.Number = v
}

// 设置值：创建时间
func (m *SupplierAccount) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *SupplierAccount) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *SupplierAccount) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *SupplierAccount) GetID() int64 {
	return m.ID
}

// 获取值：供应商id
func (m *SupplierAccount) GetSupplierID() string {
	return m.SupplierID
}

// 获取值：账户名称
func (m *SupplierAccount) GetName() string {
	return m.Name
}

// 获取值：开户行
func (m *SupplierAccount) GetBank() string {
	return m.Bank
}

// 获取值：账户类型：0 支付宝 1 银行卡 2 微信
func (m *SupplierAccount) GetAccountType() string {
	return m.AccountType
}

// 获取值：账号
func (m *SupplierAccount) GetNumber() string {
	return m.Number
}

// 获取值：创建时间
func (m *SupplierAccount) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *SupplierAccount) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *SupplierAccount) GetDeletedAt() int64 {
	return m.DeletedAt
}
