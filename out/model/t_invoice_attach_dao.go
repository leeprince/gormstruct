package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-07-21 15:59:19
 * @Desc:   t_invoice_attach 表的 DAO 层
 */

type TInvoiceAttachDAO struct {
	*_BaseDAO
}

// 初始化 TInvoiceAttachDAO
func NewTInvoiceAttachDAO(ctx context.Context, db *gorm.DB) *TInvoiceAttachDAO {
	if db == nil {
		panic(fmt.Errorf("TInvoiceAttachDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &TInvoiceAttachDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&TInvoiceAttach{}),
			db:               db,
			model:            TInvoiceAttach{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          TInvoiceAttachAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *TInvoiceAttachDAO) GetTableName() string {
	tInvoiceAttach := &TInvoiceAttach{}
	return tInvoiceAttach.TableName()
}

// 存在则更新，否则插入
func (obj *TInvoiceAttachDAO) Save(tInvoiceAttach *TInvoiceAttach) (rowsAffected int64, err error) {
	if tInvoiceAttach.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(tInvoiceAttach, obj.WithID(tInvoiceAttach.ID))
	}
	return obj.Create(tInvoiceAttach)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *TInvoiceAttachDAO) Create(tInvoiceAttach interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(tInvoiceAttach)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id() 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithID(id uint64) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ID] = id })
}

// 设置 id() 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithIDs(ids []uint64) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ID] = ids })
}

// 设置 invoice_id(发票id) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithInvoiceID(invoiceID string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.InvoiceID] = invoiceID })
}

// 设置 invoice_id(发票id) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithInvoiceIDs(invoiceIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.InvoiceID] = invoiceIDs })
}

// 设置 order_sn(订单id) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithOrderSn(orderSn string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.OrderSn] = orderSn })
}

// 设置 order_sn(订单id) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithOrderSns(orderSns []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.OrderSn] = orderSns })
}

// 设置 attachment_url(附件链接,JSON["a","b"]) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithAttachmentURL(attachmentURL string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.AttachmentURL] = attachmentURL })
}

// 设置 attachment_url(附件链接,JSON["a","b"]) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithAttachmentURLs(attachmentURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.AttachmentURL] = attachmentURLs })
}

// 设置 ocr_modify(OCR修改记录,JSON) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithOcrModify(ocrModify string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.OcrModify] = ocrModify })
}

// 设置 ocr_modify(OCR修改记录,JSON) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithOcrModifys(ocrModifys []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.OcrModify] = ocrModifys })
}

// 设置 custom_tax_rate(税率默认值-1) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithCustomTaxRate(customTaxRate float64) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CustomTaxRate] = customTaxRate })
}

// 设置 custom_tax_rate(税率默认值-1) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithCustomTaxRates(customTaxRates []float64) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CustomTaxRate] = customTaxRates })
}

// 设置 cert_status(凭证入账状态) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithCertStatus(certStatus int8) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CertStatus] = certStatus })
}

// 设置 cert_status(凭证入账状态) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithCertStatuss(certStatuss []int8) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CertStatus] = certStatuss })
}

// 设置 pdf_url(ofd转pdf的url) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithPdfURL(pdfURL string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.PdfURL] = pdfURL })
}

// 设置 pdf_url(ofd转pdf的url) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithPdfURLs(pdfURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.PdfURL] = pdfURLs })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.UpdatedAt] = updatedAts })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.CreatedAt] = createdAts })
}

// 设置 external_order_id(外部订单号（多个订单用逗号隔开）) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithExternalOrderID(externalOrderID string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ExternalOrderID] = externalOrderID })
}

// 设置 external_order_id(外部订单号（多个订单用逗号隔开）) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithExternalOrderIDs(externalOrderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ExternalOrderID] = externalOrderIDs })
}

// 设置 external_ticket_type(外部票据类型（taxi, flight, hotel, train）) 字段作为 option 条件
func (obj *TInvoiceAttachDAO) WithExternalTicketType(externalTicketType string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ExternalTicketType] = externalTicketType })
}

// 设置 external_ticket_type(外部票据类型（taxi, flight, hotel, train）) 字段的切片作为 option 条件
func (obj *TInvoiceAttachDAO) WithExternalTicketTypes(externalTicketTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.query[TInvoiceAttachColumns.ExternalTicketType] = externalTicketTypes })
}

// 函数选项模式获取单条记录
func (obj *TInvoiceAttachDAO) GetByOption(opts ...Option) (result *TInvoiceAttach, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *TInvoiceAttachDAO) GetByOptions(opts ...Option) (results []*TInvoiceAttach, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *TInvoiceAttachDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      tInvoiceAttach: 要更新的数据
//      opts: 更新的条件
func (obj *TInvoiceAttachDAO) UpdateByOption(tInvoiceAttach *TInvoiceAttach, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(tInvoiceAttach)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id() 字段值，获取单条记录
func (obj *TInvoiceAttachDAO) GetFromID(id uint64) (result *TInvoiceAttach, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id() 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromID(ids []uint64) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 invoice_id(发票id) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromInvoiceID(invoiceID string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithInvoiceID(invoiceID))
	return
}

// 通过多个 invoice_id(发票id) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromInvoiceID(invoiceIDs []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithInvoiceIDs(invoiceIDs))
	return
}

// 通过单个 order_sn(订单id) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromOrderSn(orderSn string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithOrderSn(orderSn))
	return
}

// 通过多个 order_sn(订单id) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromOrderSn(orderSns []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithOrderSns(orderSns))
	return
}

// 通过单个 attachment_url(附件链接,JSON["a","b"]) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromAttachmentURL(attachmentURL string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithAttachmentURL(attachmentURL))
	return
}

// 通过多个 attachment_url(附件链接,JSON["a","b"]) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromAttachmentURL(attachmentURLs []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithAttachmentURLs(attachmentURLs))
	return
}

// 通过单个 ocr_modify(OCR修改记录,JSON) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromOcrModify(ocrModify string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithOcrModify(ocrModify))
	return
}

// 通过多个 ocr_modify(OCR修改记录,JSON) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromOcrModify(ocrModifys []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithOcrModifys(ocrModifys))
	return
}

// 通过单个 custom_tax_rate(税率默认值-1) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromCustomTaxRate(customTaxRate float64) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCustomTaxRate(customTaxRate))
	return
}

// 通过多个 custom_tax_rate(税率默认值-1) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromCustomTaxRate(customTaxRates []float64) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCustomTaxRates(customTaxRates))
	return
}

// 通过单个 cert_status(凭证入账状态) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromCertStatus(certStatus int8) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCertStatus(certStatus))
	return
}

// 通过多个 cert_status(凭证入账状态) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromCertStatus(certStatuss []int8) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCertStatuss(certStatuss))
	return
}

// 通过单个 pdf_url(ofd转pdf的url) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromPdfURL(pdfURL string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithPdfURL(pdfURL))
	return
}

// 通过多个 pdf_url(ofd转pdf的url) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromPdfURL(pdfURLs []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithPdfURLs(pdfURLs))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromUpdatedAt(updatedAts []time.Time) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromCreatedAt(createdAt time.Time) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromCreatedAt(createdAts []time.Time) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// 通过单个 external_order_id(外部订单号（多个订单用逗号隔开）) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromExternalOrderID(externalOrderID string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithExternalOrderID(externalOrderID))
	return
}

// 通过多个 external_order_id(外部订单号（多个订单用逗号隔开）) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromExternalOrderID(externalOrderIDs []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithExternalOrderIDs(externalOrderIDs))
	return
}

// 通过单个 external_ticket_type(外部票据类型（taxi, flight, hotel, train）) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetFromExternalTicketType(externalTicketType string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithExternalTicketType(externalTicketType))
	return
}

// 通过多个 external_ticket_type(外部票据类型（taxi, flight, hotel, train）) 字段值，获取多条记录
func (obj *TInvoiceAttachDAO) GetsFromExternalTicketType(externalTicketTypes []string) (results []*TInvoiceAttach, err error) {
	results, err = obj.GetByOptions(obj.WithExternalTicketTypes(externalTicketTypes))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *TInvoiceAttachDAO) FetchByPrimaryKey(id uint64) (result *TInvoiceAttach, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 invoice_id, order_sn 字段值，获取单条记录
func (obj *TInvoiceAttachDAO) FetchUniqueIndexByIDSn(invoiceID string, orderSn string) (result *TInvoiceAttach, err error) {
	result, err = obj.GetByOption(
		obj.WithInvoiceID(invoiceID),
		obj.WithOrderSn(orderSn))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
