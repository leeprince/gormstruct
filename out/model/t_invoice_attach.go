package model

import (
	"time"
)

// 发票附加信息表
type TInvoiceAttach struct {
	ID                 uint64    `gorm:"column:id;primaryKey;type:bigint(20) unsigned;autoIncrement" json:"id"`
	InvoiceID          string    `gorm:"column:invoice_id;type:varchar(64);not null;DEFAULT '0'" json:"invoice_id"`                    // 发票id
	OrderSn            string    `gorm:"column:order_sn;type:varchar(64);not null;DEFAULT '0'" json:"order_sn"`                        // 订单id
	AttachmentURL      string    `gorm:"column:attachment_url;type:varchar(2048);not null;DEFAULT ''" json:"attachment_url"`           // 附件链接,JSON["a","b"]
	OcrModify          string    `gorm:"column:ocr_modify;type:varchar(2048);not null;DEFAULT ''" json:"ocr_modify"`                   // OCR修改记录,JSON
	CustomTaxRate      float64   `gorm:"column:custom_tax_rate;type:decimal(20,2);not null;DEFAULT '-1.00'" json:"custom_tax_rate"`    // 税率默认值-1
	CertStatus         int8      `gorm:"column:cert_status;type:tinyint(2);not null;DEFAULT '0'" json:"cert_status"`                   // 凭证入账状态
	PdfURL             string    `gorm:"column:pdf_url;type:varchar(2048);not null;DEFAULT ''" json:"pdf_url"`                         // ofd转pdf的url
	UpdatedAt          time.Time `gorm:"column:updated_at;type:timestamp;not null;DEFAULT 'CURRENT_TIMESTAMP'" json:"updated_at"`      // 更新时间
	CreatedAt          time.Time `gorm:"column:created_at;type:timestamp;not null;DEFAULT 'CURRENT_TIMESTAMP'" json:"created_at"`      // 创建时间
	ExternalOrderID    string    `gorm:"column:external_order_id;type:varchar(1000);not null;DEFAULT ''" json:"external_order_id"`     // 外部订单号（多个订单用逗号隔开）
	ExternalTicketType string    `gorm:"column:external_ticket_type;type:varchar(30);not null;DEFAULT ''" json:"external_ticket_type"` // 外部票据类型（taxi, flight, hotel, train）
}

// 获取结构体对应的表名方法
func (m *TInvoiceAttach) TableName() string {
	return "t_invoice_attach"
}

// 实例化结构体对象
func NewTInvoiceAttach() *TInvoiceAttach {
	return &TInvoiceAttach{}
}

// 获取主键的对应字段
func (m *TInvoiceAttach) PrimaryKey() string {
	return TInvoiceAttachColumns.ID
}

// 获取主键值
func (m *TInvoiceAttach) PrimaryKeyValue() uint64 {
	return m.ID
}

// 表字段的映射
var TInvoiceAttachColumns = struct {
	ID                 string
	InvoiceID          string
	OrderSn            string
	AttachmentURL      string
	OcrModify          string
	CustomTaxRate      string
	CertStatus         string
	PdfURL             string
	UpdatedAt          string
	CreatedAt          string
	ExternalOrderID    string
	ExternalTicketType string
}{
	ID:                 "id",
	InvoiceID:          "invoice_id",
	OrderSn:            "order_sn",
	AttachmentURL:      "attachment_url",
	OcrModify:          "ocr_modify",
	CustomTaxRate:      "custom_tax_rate",
	CertStatus:         "cert_status",
	PdfURL:             "pdf_url",
	UpdatedAt:          "updated_at",
	CreatedAt:          "created_at",
	ExternalOrderID:    "external_order_id",
	ExternalTicketType: "external_ticket_type",
}

// 包含所有表字段的切片
var TInvoiceAttachAllColumns = []string{
	TInvoiceAttachColumns.ID,                 //
	TInvoiceAttachColumns.InvoiceID,          // 发票id
	TInvoiceAttachColumns.OrderSn,            // 订单id
	TInvoiceAttachColumns.AttachmentURL,      // 附件链接,JSON[&#34;a&#34;,&#34;b&#34;]
	TInvoiceAttachColumns.OcrModify,          // OCR修改记录,JSON
	TInvoiceAttachColumns.CustomTaxRate,      // 税率默认值-1
	TInvoiceAttachColumns.CertStatus,         // 凭证入账状态
	TInvoiceAttachColumns.PdfURL,             // ofd转pdf的url
	TInvoiceAttachColumns.UpdatedAt,          // 更新时间
	TInvoiceAttachColumns.CreatedAt,          // 创建时间
	TInvoiceAttachColumns.ExternalOrderID,    // 外部订单号（多个订单用逗号隔开）
	TInvoiceAttachColumns.ExternalTicketType, // 外部票据类型（taxi, flight, hotel, train）

}

// 设置值：
func (m *TInvoiceAttach) SetID(v uint64) {
	m.ID = v
}

// 设置值：发票id
func (m *TInvoiceAttach) SetInvoiceID(v string) {
	m.InvoiceID = v
}

// 设置值：订单id
func (m *TInvoiceAttach) SetOrderSn(v string) {
	m.OrderSn = v
}

// 设置值：附件链接,JSON[&#34;a&#34;,&#34;b&#34;]
func (m *TInvoiceAttach) SetAttachmentURL(v string) {
	m.AttachmentURL = v
}

// 设置值：OCR修改记录,JSON
func (m *TInvoiceAttach) SetOcrModify(v string) {
	m.OcrModify = v
}

// 设置值：税率默认值-1
func (m *TInvoiceAttach) SetCustomTaxRate(v float64) {
	m.CustomTaxRate = v
}

// 设置值：凭证入账状态
func (m *TInvoiceAttach) SetCertStatus(v int8) {
	m.CertStatus = v
}

// 设置值：ofd转pdf的url
func (m *TInvoiceAttach) SetPdfURL(v string) {
	m.PdfURL = v
}

// 设置值：更新时间
func (m *TInvoiceAttach) SetUpdatedAt(v time.Time) {
	m.UpdatedAt = v
}

// 设置值：创建时间
func (m *TInvoiceAttach) SetCreatedAt(v time.Time) {
	m.CreatedAt = v
}

// 设置值：外部订单号（多个订单用逗号隔开）
func (m *TInvoiceAttach) SetExternalOrderID(v string) {
	m.ExternalOrderID = v
}

// 设置值：外部票据类型（taxi, flight, hotel, train）
func (m *TInvoiceAttach) SetExternalTicketType(v string) {
	m.ExternalTicketType = v
}

// 获取值：
func (m *TInvoiceAttach) GetID() uint64 {
	return m.ID
}

// 获取值：发票id
func (m *TInvoiceAttach) GetInvoiceID() string {
	return m.InvoiceID
}

// 获取值：订单id
func (m *TInvoiceAttach) GetOrderSn() string {
	return m.OrderSn
}

// 获取值：附件链接,JSON[&#34;a&#34;,&#34;b&#34;]
func (m *TInvoiceAttach) GetAttachmentURL() string {
	return m.AttachmentURL
}

// 获取值：OCR修改记录,JSON
func (m *TInvoiceAttach) GetOcrModify() string {
	return m.OcrModify
}

// 获取值：税率默认值-1
func (m *TInvoiceAttach) GetCustomTaxRate() float64 {
	return m.CustomTaxRate
}

// 获取值：凭证入账状态
func (m *TInvoiceAttach) GetCertStatus() int8 {
	return m.CertStatus
}

// 获取值：ofd转pdf的url
func (m *TInvoiceAttach) GetPdfURL() string {
	return m.PdfURL
}

// 获取值：更新时间
func (m *TInvoiceAttach) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

// 获取值：创建时间
func (m *TInvoiceAttach) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// 获取值：外部订单号（多个订单用逗号隔开）
func (m *TInvoiceAttach) GetExternalOrderID() string {
	return m.ExternalOrderID
}

// 获取值：外部票据类型（taxi, flight, hotel, train）
func (m *TInvoiceAttach) GetExternalTicketType() string {
	return m.ExternalTicketType
}
