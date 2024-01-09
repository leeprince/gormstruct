package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:18
 * @Desc:   voucher_info 表
 */

// 支付凭证表
type VoucherInfo struct {
	ID                 int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                    // 主键id
	OrgID              int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                                // 一级供应商（平台企业）租户id
	OpenOrgID          string `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"`                   // 一级供应商（平台企业）开放租户id
	PayID              string `gorm:"column:pay_id;type:varchar(50);not null;DEFAULT ''" json:"pay_id"`                             // 支付单编号
	Amount             int64  `gorm:"column:amount;type:int(11);not null;DEFAULT '0'" json:"amount"`                                // 金额（单位：分）
	Channel            string `gorm:"column:channel;type:varchar(50);not null;DEFAULT ''" json:"channel"`                           // 支付通道 0-银行卡 1-微信 2-支付宝 3-其他
	PayName            string `gorm:"column:pay_name;type:varchar(50);not null;DEFAULT ''" json:"pay_name"`                         // 付款方名称
	PayBankName        string `gorm:"column:pay_bank_name;type:varchar(50);not null;DEFAULT ''" json:"pay_bank_name"`               // 付款方银行名称
	PayBankAccount     string `gorm:"column:pay_bank_account;type:varchar(50);not null;DEFAULT ''" json:"pay_bank_account"`         // 付款方账号
	ReceiveName        string `gorm:"column:receive_name;type:varchar(50);not null;DEFAULT ''" json:"receive_name"`                 // 收款方名称
	ReceiveBankName    string `gorm:"column:receive_bank_name;type:varchar(50);not null;DEFAULT ''" json:"receive_bank_name"`       // 收款方银行名称
	ReceiveBankAccount string `gorm:"column:receive_bank_account;type:varchar(50);not null;DEFAULT ''" json:"receive_bank_account"` // 收款方账号
	URL                string `gorm:"column:url;type:varchar(255);not null;DEFAULT ''" json:"url"`                                  // 支付凭证下载URL
	StartTime          int64  `gorm:"column:start_time;type:int(11);not null;DEFAULT '0'" json:"start_time"`                        // 支付开始时间
	SuccessTime        int64  `gorm:"column:success_time;type:int(11);not null;DEFAULT '0'" json:"success_time"`                    // 支付成功时间
	VoucherStatus      int    `gorm:"column:voucher_status;type:int(1);not null;DEFAULT '0'" json:"voucher_status"`                 // 支付状态 0-支付成功 1-支付失败 2-支付中
	Remark             string `gorm:"column:remark;type:varchar(255);not null;DEFAULT ''" json:"remark"`                            // 业务备注
	CreatedAt          int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                        // 创建时间
	UpdatedAt          int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                        // 更新时间
	DeletedAt          int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                        // 删除时间
}

// 获取结构体对应的表名方法
func (m *VoucherInfo) TableName() string {
	return "voucher_info"
}

// 实例化结构体对象
func NewVoucherInfo() *VoucherInfo {
	return &VoucherInfo{}
}

// 获取主键的对应字段
func (m *VoucherInfo) PrimaryKey() string {
	return VoucherInfoColumns.ID
}

// 获取主键值
func (m *VoucherInfo) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var VoucherInfoColumns = struct {
	ID                 string
	OrgID              string
	OpenOrgID          string
	PayID              string
	Amount             string
	Channel            string
	PayName            string
	PayBankName        string
	PayBankAccount     string
	ReceiveName        string
	ReceiveBankName    string
	ReceiveBankAccount string
	URL                string
	StartTime          string
	SuccessTime        string
	VoucherStatus      string
	Remark             string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
}{
	ID:                 "id",
	OrgID:              "org_id",
	OpenOrgID:          "open_org_id",
	PayID:              "pay_id",
	Amount:             "amount",
	Channel:            "channel",
	PayName:            "pay_name",
	PayBankName:        "pay_bank_name",
	PayBankAccount:     "pay_bank_account",
	ReceiveName:        "receive_name",
	ReceiveBankName:    "receive_bank_name",
	ReceiveBankAccount: "receive_bank_account",
	URL:                "url",
	StartTime:          "start_time",
	SuccessTime:        "success_time",
	VoucherStatus:      "voucher_status",
	Remark:             "remark",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// 包含所有表字段的切片
var VoucherInfoAllColumns = []string{
	VoucherInfoColumns.ID,                 // 主键id
	VoucherInfoColumns.OrgID,              // 一级供应商（平台企业）租户id
	VoucherInfoColumns.OpenOrgID,          // 一级供应商（平台企业）开放租户id
	VoucherInfoColumns.PayID,              // 支付单编号
	VoucherInfoColumns.Amount,             // 金额（单位：分）
	VoucherInfoColumns.Channel,            // 支付通道 0-银行卡 1-微信 2-支付宝 3-其他
	VoucherInfoColumns.PayName,            // 付款方名称
	VoucherInfoColumns.PayBankName,        // 付款方银行名称
	VoucherInfoColumns.PayBankAccount,     // 付款方账号
	VoucherInfoColumns.ReceiveName,        // 收款方名称
	VoucherInfoColumns.ReceiveBankName,    // 收款方银行名称
	VoucherInfoColumns.ReceiveBankAccount, // 收款方账号
	VoucherInfoColumns.URL,                // 支付凭证下载URL
	VoucherInfoColumns.StartTime,          // 支付开始时间
	VoucherInfoColumns.SuccessTime,        // 支付成功时间
	VoucherInfoColumns.VoucherStatus,      // 支付状态 0-支付成功 1-支付失败 2-支付中
	VoucherInfoColumns.Remark,             // 业务备注
	VoucherInfoColumns.CreatedAt,          // 创建时间
	VoucherInfoColumns.UpdatedAt,          // 更新时间
	VoucherInfoColumns.DeletedAt,          // 删除时间

}

// 设置值：主键id
func (m *VoucherInfo) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id
func (m *VoucherInfo) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户id
func (m *VoucherInfo) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：支付单编号
func (m *VoucherInfo) SetPayID(v string) {
	m.PayID = v
}

// 设置值：金额（单位：分）
func (m *VoucherInfo) SetAmount(v int64) {
	m.Amount = v
}

// 设置值：支付通道 0-银行卡 1-微信 2-支付宝 3-其他
func (m *VoucherInfo) SetChannel(v string) {
	m.Channel = v
}

// 设置值：付款方名称
func (m *VoucherInfo) SetPayName(v string) {
	m.PayName = v
}

// 设置值：付款方银行名称
func (m *VoucherInfo) SetPayBankName(v string) {
	m.PayBankName = v
}

// 设置值：付款方账号
func (m *VoucherInfo) SetPayBankAccount(v string) {
	m.PayBankAccount = v
}

// 设置值：收款方名称
func (m *VoucherInfo) SetReceiveName(v string) {
	m.ReceiveName = v
}

// 设置值：收款方银行名称
func (m *VoucherInfo) SetReceiveBankName(v string) {
	m.ReceiveBankName = v
}

// 设置值：收款方账号
func (m *VoucherInfo) SetReceiveBankAccount(v string) {
	m.ReceiveBankAccount = v
}

// 设置值：支付凭证下载URL
func (m *VoucherInfo) SetURL(v string) {
	m.URL = v
}

// 设置值：支付开始时间
func (m *VoucherInfo) SetStartTime(v int64) {
	m.StartTime = v
}

// 设置值：支付成功时间
func (m *VoucherInfo) SetSuccessTime(v int64) {
	m.SuccessTime = v
}

// 设置值：支付状态 0-支付成功 1-支付失败 2-支付中
func (m *VoucherInfo) SetVoucherStatus(v int) {
	m.VoucherStatus = v
}

// 设置值：业务备注
func (m *VoucherInfo) SetRemark(v string) {
	m.Remark = v
}

// 设置值：创建时间
func (m *VoucherInfo) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *VoucherInfo) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *VoucherInfo) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *VoucherInfo) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id
func (m *VoucherInfo) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户id
func (m *VoucherInfo) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：支付单编号
func (m *VoucherInfo) GetPayID() string {
	return m.PayID
}

// 获取值：金额（单位：分）
func (m *VoucherInfo) GetAmount() int64 {
	return m.Amount
}

// 获取值：支付通道 0-银行卡 1-微信 2-支付宝 3-其他
func (m *VoucherInfo) GetChannel() string {
	return m.Channel
}

// 获取值：付款方名称
func (m *VoucherInfo) GetPayName() string {
	return m.PayName
}

// 获取值：付款方银行名称
func (m *VoucherInfo) GetPayBankName() string {
	return m.PayBankName
}

// 获取值：付款方账号
func (m *VoucherInfo) GetPayBankAccount() string {
	return m.PayBankAccount
}

// 获取值：收款方名称
func (m *VoucherInfo) GetReceiveName() string {
	return m.ReceiveName
}

// 获取值：收款方银行名称
func (m *VoucherInfo) GetReceiveBankName() string {
	return m.ReceiveBankName
}

// 获取值：收款方账号
func (m *VoucherInfo) GetReceiveBankAccount() string {
	return m.ReceiveBankAccount
}

// 获取值：支付凭证下载URL
func (m *VoucherInfo) GetURL() string {
	return m.URL
}

// 获取值：支付开始时间
func (m *VoucherInfo) GetStartTime() int64 {
	return m.StartTime
}

// 获取值：支付成功时间
func (m *VoucherInfo) GetSuccessTime() int64 {
	return m.SuccessTime
}

// 获取值：支付状态 0-支付成功 1-支付失败 2-支付中
func (m *VoucherInfo) GetVoucherStatus() int {
	return m.VoucherStatus
}

// 获取值：业务备注
func (m *VoucherInfo) GetRemark() string {
	return m.Remark
}

// 获取值：创建时间
func (m *VoucherInfo) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *VoucherInfo) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *VoucherInfo) GetDeletedAt() int64 {
	return m.DeletedAt
}
