package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:41:29
 * @Desc:   invoice_info 表
 */

// 发票数据表
type InvoiceInfo struct {
	ID                int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                  // 主键id
	OrgID             int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                              // 一级供应商（平台企业）租户id
	OpenOrgID         string `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"`                 // 一级供应商（平台企业）开放租户id
	InvoiceCode       string `gorm:"column:invoice_code;type:varchar(50);not null;DEFAULT ''" json:"invoice_code"`               // 发票代码
	InvoiceNo         string `gorm:"column:invoice_no;type:varchar(50);not null;DEFAULT ''" json:"invoice_no"`                   // 发票号码
	AmountHasTax      int64  `gorm:"column:amount_has_tax;type:int(11);not null;DEFAULT '0'" json:"amount_has_tax"`              // 含税总金额（单位：分）
	TaxAmount         int64  `gorm:"column:tax_amount;type:int(11);not null;DEFAULT '0'" json:"tax_amount"`                      // 总税额（单位：分）
	AmountNoTax       int64  `gorm:"column:amount_no_tax;type:int(11);not null;DEFAULT '0'" json:"amount_no_tax"`                // 不含税总金额（单位：分）
	SellerName        string `gorm:"column:seller_name;type:varchar(50);not null;DEFAULT ''" json:"seller_name"`                 // 销方名称
	SellerTaxpayerNum string `gorm:"column:seller_taxpayer_num;type:varchar(50);not null;DEFAULT ''" json:"seller_taxpayer_num"` // 销方纳税人识别号
	SellerAddress     string `gorm:"column:seller_address;type:varchar(50);not null;DEFAULT ''" json:"seller_address"`           // 销方地址
	SellerPhone       string `gorm:"column:seller_phone;type:varchar(50);not null;DEFAULT ''" json:"seller_phone"`               // 销方电话
	SellerBankName    string `gorm:"column:seller_bank_name;type:varchar(50);not null;DEFAULT ''" json:"seller_bank_name"`       // 销方银行名称
	SellerBankAccount string `gorm:"column:seller_bank_account;type:varchar(50);not null;DEFAULT ''" json:"seller_bank_account"` // 销方银行账号
	HeaderType        string `gorm:"column:header_type;type:varchar(50);not null;DEFAULT ''" json:"header_type"`                 // 抬头类型抬头类型：1：个人、政府事业单位；2：企业
	BuyerName         string `gorm:"column:buyer_name;type:varchar(50);not null;DEFAULT ''" json:"buyer_name"`                   // 购方名称
	BuyerTaxpayerNum  string `gorm:"column:buyer_taxpayer_num;type:varchar(50);not null;DEFAULT ''" json:"buyer_taxpayer_num"`   // 购方纳税人识别号
	BuyerAddress      string `gorm:"column:buyer_address;type:varchar(50);not null;DEFAULT ''" json:"buyer_address"`             // 购方地址
	BuyerPhone        string `gorm:"column:buyer_phone;type:varchar(11);not null;DEFAULT ''" json:"buyer_phone"`                 // 购方电话
	BuyerBankName     string `gorm:"column:buyer_bank_name;type:varchar(50);not null;DEFAULT ''" json:"buyer_bank_name"`         // 购方银行名称
	BuyerBankAccount  string `gorm:"column:buyer_bank_account;type:varchar(50);not null;DEFAULT ''" json:"buyer_bank_account"`   // 购方银行账号
	InvoiceTypeCode   string `gorm:"column:invoice_type_code;type:varchar(11);not null;DEFAULT ''" json:"invoice_type_code"`     // 发票类型 081：全电电子专票 082：全电电子普票
	InvoiceTime       int64  `gorm:"column:invoice_time;type:int(11);not null;DEFAULT '0'" json:"invoice_time"`                  // 开票时间
	PdfURL            string `gorm:"column:pdf_url;type:varchar(255);not null;DEFAULT ''" json:"pdf_url"`                        // 发票文件PDF下载地址
	DrawerName        string `gorm:"column:drawer_name;type:varchar(50);not null;DEFAULT ''" json:"drawer_name"`                 // 开票人姓名
	PayeeName         string `gorm:"column:payee_name;type:varchar(50);not null;DEFAULT ''" json:"payee_name"`                   // 收款人姓名
	ReviewerName      string `gorm:"column:reviewer_name;type:varchar(50);not null;DEFAULT ''" json:"reviewer_name"`             // 复核人姓名
	Remark            string `gorm:"column:remark;type:varchar(255);not null;DEFAULT ''" json:"remark"`                          // 备注
	CreatedAt         int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                      // 创建时间
	UpdatedAt         int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                      // 更新时间
	DeletedAt         int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                      // 删除时间
}

// 获取结构体对应的表名方法
func (m *InvoiceInfo) TableName() string {
	return "invoice_info"
}

// 实例化结构体对象
func NewInvoiceInfo() *InvoiceInfo {
	return &InvoiceInfo{}
}

// 获取主键的对应字段
func (m *InvoiceInfo) PrimaryKey() string {
	return InvoiceInfoColumns.ID
}

// 获取主键值
func (m *InvoiceInfo) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var InvoiceInfoColumns = struct {
	ID                string
	OrgID             string
	OpenOrgID         string
	InvoiceCode       string
	InvoiceNo         string
	AmountHasTax      string
	TaxAmount         string
	AmountNoTax       string
	SellerName        string
	SellerTaxpayerNum string
	SellerAddress     string
	SellerPhone       string
	SellerBankName    string
	SellerBankAccount string
	HeaderType        string
	BuyerName         string
	BuyerTaxpayerNum  string
	BuyerAddress      string
	BuyerPhone        string
	BuyerBankName     string
	BuyerBankAccount  string
	InvoiceTypeCode   string
	InvoiceTime       string
	PdfURL            string
	DrawerName        string
	PayeeName         string
	ReviewerName      string
	Remark            string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
}{
	ID:                "id",
	OrgID:             "org_id",
	OpenOrgID:         "open_org_id",
	InvoiceCode:       "invoice_code",
	InvoiceNo:         "invoice_no",
	AmountHasTax:      "amount_has_tax",
	TaxAmount:         "tax_amount",
	AmountNoTax:       "amount_no_tax",
	SellerName:        "seller_name",
	SellerTaxpayerNum: "seller_taxpayer_num",
	SellerAddress:     "seller_address",
	SellerPhone:       "seller_phone",
	SellerBankName:    "seller_bank_name",
	SellerBankAccount: "seller_bank_account",
	HeaderType:        "header_type",
	BuyerName:         "buyer_name",
	BuyerTaxpayerNum:  "buyer_taxpayer_num",
	BuyerAddress:      "buyer_address",
	BuyerPhone:        "buyer_phone",
	BuyerBankName:     "buyer_bank_name",
	BuyerBankAccount:  "buyer_bank_account",
	InvoiceTypeCode:   "invoice_type_code",
	InvoiceTime:       "invoice_time",
	PdfURL:            "pdf_url",
	DrawerName:        "drawer_name",
	PayeeName:         "payee_name",
	ReviewerName:      "reviewer_name",
	Remark:            "remark",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
}

// 包含所有表字段的切片
var InvoiceInfoAllColumns = []string{
	InvoiceInfoColumns.ID,                // 主键id
	InvoiceInfoColumns.OrgID,             // 一级供应商（平台企业）租户id
	InvoiceInfoColumns.OpenOrgID,         // 一级供应商（平台企业）开放租户id
	InvoiceInfoColumns.InvoiceCode,       // 发票代码
	InvoiceInfoColumns.InvoiceNo,         // 发票号码
	InvoiceInfoColumns.AmountHasTax,      // 含税总金额（单位：分）
	InvoiceInfoColumns.TaxAmount,         // 总税额（单位：分）
	InvoiceInfoColumns.AmountNoTax,       // 不含税总金额（单位：分）
	InvoiceInfoColumns.SellerName,        // 销方名称
	InvoiceInfoColumns.SellerTaxpayerNum, // 销方纳税人识别号
	InvoiceInfoColumns.SellerAddress,     // 销方地址
	InvoiceInfoColumns.SellerPhone,       // 销方电话
	InvoiceInfoColumns.SellerBankName,    // 销方银行名称
	InvoiceInfoColumns.SellerBankAccount, // 销方银行账号
	InvoiceInfoColumns.HeaderType,        // 抬头类型抬头类型：1：个人、政府事业单位；2：企业
	InvoiceInfoColumns.BuyerName,         // 购方名称
	InvoiceInfoColumns.BuyerTaxpayerNum,  // 购方纳税人识别号
	InvoiceInfoColumns.BuyerAddress,      // 购方地址
	InvoiceInfoColumns.BuyerPhone,        // 购方电话
	InvoiceInfoColumns.BuyerBankName,     // 购方银行名称
	InvoiceInfoColumns.BuyerBankAccount,  // 购方银行账号
	InvoiceInfoColumns.InvoiceTypeCode,   // 发票类型 081：全电电子专票 082：全电电子普票
	InvoiceInfoColumns.InvoiceTime,       // 开票时间
	InvoiceInfoColumns.PdfURL,            // 发票文件PDF下载地址
	InvoiceInfoColumns.DrawerName,        // 开票人姓名
	InvoiceInfoColumns.PayeeName,         // 收款人姓名
	InvoiceInfoColumns.ReviewerName,      // 复核人姓名
	InvoiceInfoColumns.Remark,            // 备注
	InvoiceInfoColumns.CreatedAt,         // 创建时间
	InvoiceInfoColumns.UpdatedAt,         // 更新时间
	InvoiceInfoColumns.DeletedAt,         // 删除时间

}

// 设置值：主键id
func (m *InvoiceInfo) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id
func (m *InvoiceInfo) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户id
func (m *InvoiceInfo) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：发票代码
func (m *InvoiceInfo) SetInvoiceCode(v string) {
	m.InvoiceCode = v
}

// 设置值：发票号码
func (m *InvoiceInfo) SetInvoiceNo(v string) {
	m.InvoiceNo = v
}

// 设置值：含税总金额（单位：分）
func (m *InvoiceInfo) SetAmountHasTax(v int64) {
	m.AmountHasTax = v
}

// 设置值：总税额（单位：分）
func (m *InvoiceInfo) SetTaxAmount(v int64) {
	m.TaxAmount = v
}

// 设置值：不含税总金额（单位：分）
func (m *InvoiceInfo) SetAmountNoTax(v int64) {
	m.AmountNoTax = v
}

// 设置值：销方名称
func (m *InvoiceInfo) SetSellerName(v string) {
	m.SellerName = v
}

// 设置值：销方纳税人识别号
func (m *InvoiceInfo) SetSellerTaxpayerNum(v string) {
	m.SellerTaxpayerNum = v
}

// 设置值：销方地址
func (m *InvoiceInfo) SetSellerAddress(v string) {
	m.SellerAddress = v
}

// 设置值：销方电话
func (m *InvoiceInfo) SetSellerPhone(v string) {
	m.SellerPhone = v
}

// 设置值：销方银行名称
func (m *InvoiceInfo) SetSellerBankName(v string) {
	m.SellerBankName = v
}

// 设置值：销方银行账号
func (m *InvoiceInfo) SetSellerBankAccount(v string) {
	m.SellerBankAccount = v
}

// 设置值：抬头类型抬头类型：1：个人、政府事业单位；2：企业
func (m *InvoiceInfo) SetHeaderType(v string) {
	m.HeaderType = v
}

// 设置值：购方名称
func (m *InvoiceInfo) SetBuyerName(v string) {
	m.BuyerName = v
}

// 设置值：购方纳税人识别号
func (m *InvoiceInfo) SetBuyerTaxpayerNum(v string) {
	m.BuyerTaxpayerNum = v
}

// 设置值：购方地址
func (m *InvoiceInfo) SetBuyerAddress(v string) {
	m.BuyerAddress = v
}

// 设置值：购方电话
func (m *InvoiceInfo) SetBuyerPhone(v string) {
	m.BuyerPhone = v
}

// 设置值：购方银行名称
func (m *InvoiceInfo) SetBuyerBankName(v string) {
	m.BuyerBankName = v
}

// 设置值：购方银行账号
func (m *InvoiceInfo) SetBuyerBankAccount(v string) {
	m.BuyerBankAccount = v
}

// 设置值：发票类型 081：全电电子专票 082：全电电子普票
func (m *InvoiceInfo) SetInvoiceTypeCode(v string) {
	m.InvoiceTypeCode = v
}

// 设置值：开票时间
func (m *InvoiceInfo) SetInvoiceTime(v int64) {
	m.InvoiceTime = v
}

// 设置值：发票文件PDF下载地址
func (m *InvoiceInfo) SetPdfURL(v string) {
	m.PdfURL = v
}

// 设置值：开票人姓名
func (m *InvoiceInfo) SetDrawerName(v string) {
	m.DrawerName = v
}

// 设置值：收款人姓名
func (m *InvoiceInfo) SetPayeeName(v string) {
	m.PayeeName = v
}

// 设置值：复核人姓名
func (m *InvoiceInfo) SetReviewerName(v string) {
	m.ReviewerName = v
}

// 设置值：备注
func (m *InvoiceInfo) SetRemark(v string) {
	m.Remark = v
}

// 设置值：创建时间
func (m *InvoiceInfo) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *InvoiceInfo) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *InvoiceInfo) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *InvoiceInfo) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id
func (m *InvoiceInfo) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户id
func (m *InvoiceInfo) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：发票代码
func (m *InvoiceInfo) GetInvoiceCode() string {
	return m.InvoiceCode
}

// 获取值：发票号码
func (m *InvoiceInfo) GetInvoiceNo() string {
	return m.InvoiceNo
}

// 获取值：含税总金额（单位：分）
func (m *InvoiceInfo) GetAmountHasTax() int64 {
	return m.AmountHasTax
}

// 获取值：总税额（单位：分）
func (m *InvoiceInfo) GetTaxAmount() int64 {
	return m.TaxAmount
}

// 获取值：不含税总金额（单位：分）
func (m *InvoiceInfo) GetAmountNoTax() int64 {
	return m.AmountNoTax
}

// 获取值：销方名称
func (m *InvoiceInfo) GetSellerName() string {
	return m.SellerName
}

// 获取值：销方纳税人识别号
func (m *InvoiceInfo) GetSellerTaxpayerNum() string {
	return m.SellerTaxpayerNum
}

// 获取值：销方地址
func (m *InvoiceInfo) GetSellerAddress() string {
	return m.SellerAddress
}

// 获取值：销方电话
func (m *InvoiceInfo) GetSellerPhone() string {
	return m.SellerPhone
}

// 获取值：销方银行名称
func (m *InvoiceInfo) GetSellerBankName() string {
	return m.SellerBankName
}

// 获取值：销方银行账号
func (m *InvoiceInfo) GetSellerBankAccount() string {
	return m.SellerBankAccount
}

// 获取值：抬头类型抬头类型：1：个人、政府事业单位；2：企业
func (m *InvoiceInfo) GetHeaderType() string {
	return m.HeaderType
}

// 获取值：购方名称
func (m *InvoiceInfo) GetBuyerName() string {
	return m.BuyerName
}

// 获取值：购方纳税人识别号
func (m *InvoiceInfo) GetBuyerTaxpayerNum() string {
	return m.BuyerTaxpayerNum
}

// 获取值：购方地址
func (m *InvoiceInfo) GetBuyerAddress() string {
	return m.BuyerAddress
}

// 获取值：购方电话
func (m *InvoiceInfo) GetBuyerPhone() string {
	return m.BuyerPhone
}

// 获取值：购方银行名称
func (m *InvoiceInfo) GetBuyerBankName() string {
	return m.BuyerBankName
}

// 获取值：购方银行账号
func (m *InvoiceInfo) GetBuyerBankAccount() string {
	return m.BuyerBankAccount
}

// 获取值：发票类型 081：全电电子专票 082：全电电子普票
func (m *InvoiceInfo) GetInvoiceTypeCode() string {
	return m.InvoiceTypeCode
}

// 获取值：开票时间
func (m *InvoiceInfo) GetInvoiceTime() int64 {
	return m.InvoiceTime
}

// 获取值：发票文件PDF下载地址
func (m *InvoiceInfo) GetPdfURL() string {
	return m.PdfURL
}

// 获取值：开票人姓名
func (m *InvoiceInfo) GetDrawerName() string {
	return m.DrawerName
}

// 获取值：收款人姓名
func (m *InvoiceInfo) GetPayeeName() string {
	return m.PayeeName
}

// 获取值：复核人姓名
func (m *InvoiceInfo) GetReviewerName() string {
	return m.ReviewerName
}

// 获取值：备注
func (m *InvoiceInfo) GetRemark() string {
	return m.Remark
}

// 获取值：创建时间
func (m *InvoiceInfo) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *InvoiceInfo) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *InvoiceInfo) GetDeletedAt() int64 {
	return m.DeletedAt
}
