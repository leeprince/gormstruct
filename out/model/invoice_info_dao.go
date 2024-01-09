package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:41:29
 * @Desc:   invoice_info 表的 DAO 层
 */

type InvoiceInfoDAO struct {
	*_BaseDAO
}

// InvoiceInfoDAO 初始化
func NewInvoiceInfoDAO(ctx context.Context, db *gorm.DB) *InvoiceInfoDAO {
	if db == nil {
		panic(fmt.Errorf("InvoiceInfoDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &InvoiceInfoDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&InvoiceInfo{}),
			db:               db,
			model:            InvoiceInfo{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          InvoiceInfoAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *InvoiceInfoDAO) GetTableName() string {
	invoiceInfo := &InvoiceInfo{}
	return invoiceInfo.TableName()
}

// Save 存在则更新，否则插入
func (obj *InvoiceInfoDAO) Save(invoiceInfo *InvoiceInfo) (rowsAffected int64, err error) {
	if invoiceInfo.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(invoiceInfo, obj.WithID(invoiceInfo.ID))
	}
	return obj.Create(invoiceInfo)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *InvoiceInfoDAO) Create(invoiceInfo interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(invoiceInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户id) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户id) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.OpenOrgID] = openOrgIDs })
}

// WithInvoiceCode 设置 invoice_code(发票代码) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceCode(invoiceCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceCode] = invoiceCode })
}

// WithInvoiceCodes 设置 invoice_code(发票代码) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceCodes(invoiceCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceCode] = invoiceCodes })
}

// WithInvoiceNo 设置 invoice_no(发票号码) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceNo(invoiceNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceNo] = invoiceNo })
}

// WithInvoiceNos 设置 invoice_no(发票号码) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceNos(invoiceNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceNo] = invoiceNos })
}

// WithAmountHasTax 设置 amount_has_tax(含税总金额（单位：分）) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithAmountHasTax(amountHasTax int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.AmountHasTax] = amountHasTax })
}

// WithAmountHasTaxs 设置 amount_has_tax(含税总金额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithAmountHasTaxs(amountHasTaxs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.AmountHasTax] = amountHasTaxs })
}

// WithTaxAmount 设置 tax_amount(总税额（单位：分）) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithTaxAmount(taxAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.TaxAmount] = taxAmount })
}

// WithTaxAmounts 设置 tax_amount(总税额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithTaxAmounts(taxAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.TaxAmount] = taxAmounts })
}

// WithAmountNoTax 设置 amount_no_tax(不含税总金额（单位：分）) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithAmountNoTax(amountNoTax int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.AmountNoTax] = amountNoTax })
}

// WithAmountNoTaxs 设置 amount_no_tax(不含税总金额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithAmountNoTaxs(amountNoTaxs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.AmountNoTax] = amountNoTaxs })
}

// WithSellerName 设置 seller_name(销方名称) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerName(sellerName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerName] = sellerName })
}

// WithSellerNames 设置 seller_name(销方名称) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerNames(sellerNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerName] = sellerNames })
}

// WithSellerTaxpayerNum 设置 seller_taxpayer_num(销方纳税人识别号) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerTaxpayerNum(sellerTaxpayerNum string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerTaxpayerNum] = sellerTaxpayerNum })
}

// WithSellerTaxpayerNums 设置 seller_taxpayer_num(销方纳税人识别号) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerTaxpayerNums(sellerTaxpayerNums []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerTaxpayerNum] = sellerTaxpayerNums })
}

// WithSellerAddress 设置 seller_address(销方地址) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerAddress(sellerAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerAddress] = sellerAddress })
}

// WithSellerAddresss 设置 seller_address(销方地址) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerAddresss(sellerAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerAddress] = sellerAddresss })
}

// WithSellerPhone 设置 seller_phone(销方电话) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerPhone(sellerPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerPhone] = sellerPhone })
}

// WithSellerPhones 设置 seller_phone(销方电话) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerPhones(sellerPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerPhone] = sellerPhones })
}

// WithSellerBankName 设置 seller_bank_name(销方银行名称) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerBankName(sellerBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerBankName] = sellerBankName })
}

// WithSellerBankNames 设置 seller_bank_name(销方银行名称) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerBankNames(sellerBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerBankName] = sellerBankNames })
}

// WithSellerBankAccount 设置 seller_bank_account(销方银行账号) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerBankAccount(sellerBankAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerBankAccount] = sellerBankAccount })
}

// WithSellerBankAccounts 设置 seller_bank_account(销方银行账号) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithSellerBankAccounts(sellerBankAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.SellerBankAccount] = sellerBankAccounts })
}

// WithHeaderType 设置 header_type(抬头类型抬头类型：1：个人、政府事业单位；2：企业) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithHeaderType(headerType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.HeaderType] = headerType })
}

// WithHeaderTypes 设置 header_type(抬头类型抬头类型：1：个人、政府事业单位；2：企业) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithHeaderTypes(headerTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.HeaderType] = headerTypes })
}

// WithBuyerName 设置 buyer_name(购方名称) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerName(buyerName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerName] = buyerName })
}

// WithBuyerNames 设置 buyer_name(购方名称) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerNames(buyerNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerName] = buyerNames })
}

// WithBuyerTaxpayerNum 设置 buyer_taxpayer_num(购方纳税人识别号) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerTaxpayerNum(buyerTaxpayerNum string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerTaxpayerNum] = buyerTaxpayerNum })
}

// WithBuyerTaxpayerNums 设置 buyer_taxpayer_num(购方纳税人识别号) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerTaxpayerNums(buyerTaxpayerNums []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerTaxpayerNum] = buyerTaxpayerNums })
}

// WithBuyerAddress 设置 buyer_address(购方地址) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerAddress(buyerAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerAddress] = buyerAddress })
}

// WithBuyerAddresss 设置 buyer_address(购方地址) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerAddresss(buyerAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerAddress] = buyerAddresss })
}

// WithBuyerPhone 设置 buyer_phone(购方电话) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerPhone(buyerPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerPhone] = buyerPhone })
}

// WithBuyerPhones 设置 buyer_phone(购方电话) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerPhones(buyerPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerPhone] = buyerPhones })
}

// WithBuyerBankName 设置 buyer_bank_name(购方银行名称) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerBankName(buyerBankName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerBankName] = buyerBankName })
}

// WithBuyerBankNames 设置 buyer_bank_name(购方银行名称) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerBankNames(buyerBankNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerBankName] = buyerBankNames })
}

// WithBuyerBankAccount 设置 buyer_bank_account(购方银行账号) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerBankAccount(buyerBankAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerBankAccount] = buyerBankAccount })
}

// WithBuyerBankAccounts 设置 buyer_bank_account(购方银行账号) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithBuyerBankAccounts(buyerBankAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.BuyerBankAccount] = buyerBankAccounts })
}

// WithInvoiceTypeCode 设置 invoice_type_code(发票类型 081：全电电子专票 082：全电电子普票) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceTypeCode(invoiceTypeCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceTypeCode] = invoiceTypeCode })
}

// WithInvoiceTypeCodes 设置 invoice_type_code(发票类型 081：全电电子专票 082：全电电子普票) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceTypeCodes(invoiceTypeCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceTypeCode] = invoiceTypeCodes })
}

// WithInvoiceTime 设置 invoice_time(开票时间) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceTime(invoiceTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceTime] = invoiceTime })
}

// WithInvoiceTimes 设置 invoice_time(开票时间) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithInvoiceTimes(invoiceTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.InvoiceTime] = invoiceTimes })
}

// WithPdfURL 设置 pdf_url(发票文件PDF下载地址) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithPdfURL(pdfURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.PdfURL] = pdfURL })
}

// WithPdfURLs 设置 pdf_url(发票文件PDF下载地址) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithPdfURLs(pdfURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.PdfURL] = pdfURLs })
}

// WithDrawerName 设置 drawer_name(开票人姓名) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithDrawerName(drawerName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.DrawerName] = drawerName })
}

// WithDrawerNames 设置 drawer_name(开票人姓名) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithDrawerNames(drawerNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.DrawerName] = drawerNames })
}

// WithPayeeName 设置 payee_name(收款人姓名) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithPayeeName(payeeName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.PayeeName] = payeeName })
}

// WithPayeeNames 设置 payee_name(收款人姓名) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithPayeeNames(payeeNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.PayeeName] = payeeNames })
}

// WithReviewerName 设置 reviewer_name(复核人姓名) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithReviewerName(reviewerName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.ReviewerName] = reviewerName })
}

// WithReviewerNames 设置 reviewer_name(复核人姓名) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithReviewerNames(reviewerNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.ReviewerName] = reviewerNames })
}

// WithRemark 设置 remark(备注) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithRemark(remark string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.Remark] = remark })
}

// WithRemarks 设置 remark(备注) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithRemarks(remarks []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.Remark] = remarks })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *InvoiceInfoDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *InvoiceInfoDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceInfoColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *InvoiceInfoDAO) GetByOption(opts ...Option) (result *InvoiceInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *InvoiceInfoDAO) GetListByOption(opts ...Option) (results []*InvoiceInfo, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *InvoiceInfoDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *InvoiceInfoDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      invoiceInfo: 要更新的数据
//      opts: 更新的条件
func (obj *InvoiceInfoDAO) UpdateByOption(invoiceInfo *InvoiceInfo, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(invoiceInfo)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *InvoiceInfoDAO) GetFromID(id int64) (result *InvoiceInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromID(ids []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromOrgID(orgID int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromOrgID(orgIDs []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromOpenOrgID(openOrgID string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromInvoiceCode 通过单个 invoice_code(发票代码) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromInvoiceCode(invoiceCode string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceCode(invoiceCode))
	return
}

// GetsFromInvoiceCode 通过多个 invoice_code(发票代码) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromInvoiceCode(invoiceCodes []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceCodes(invoiceCodes))
	return
}

// GetFromInvoiceNo 通过单个 invoice_no(发票号码) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromInvoiceNo(invoiceNo string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceNo(invoiceNo))
	return
}

// GetsFromInvoiceNo 通过多个 invoice_no(发票号码) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromInvoiceNo(invoiceNos []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceNos(invoiceNos))
	return
}

// GetFromAmountHasTax 通过单个 amount_has_tax(含税总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromAmountHasTax(amountHasTax int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmountHasTax(amountHasTax))
	return
}

// GetsFromAmountHasTax 通过多个 amount_has_tax(含税总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromAmountHasTax(amountHasTaxs []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmountHasTaxs(amountHasTaxs))
	return
}

// GetFromTaxAmount 通过单个 tax_amount(总税额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromTaxAmount(taxAmount int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithTaxAmount(taxAmount))
	return
}

// GetsFromTaxAmount 通过多个 tax_amount(总税额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromTaxAmount(taxAmounts []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithTaxAmounts(taxAmounts))
	return
}

// GetFromAmountNoTax 通过单个 amount_no_tax(不含税总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromAmountNoTax(amountNoTax int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmountNoTax(amountNoTax))
	return
}

// GetsFromAmountNoTax 通过多个 amount_no_tax(不含税总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromAmountNoTax(amountNoTaxs []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithAmountNoTaxs(amountNoTaxs))
	return
}

// GetFromSellerName 通过单个 seller_name(销方名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerName(sellerName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerName(sellerName))
	return
}

// GetsFromSellerName 通过多个 seller_name(销方名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerName(sellerNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerNames(sellerNames))
	return
}

// GetFromSellerTaxpayerNum 通过单个 seller_taxpayer_num(销方纳税人识别号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerTaxpayerNum(sellerTaxpayerNum string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerTaxpayerNum(sellerTaxpayerNum))
	return
}

// GetsFromSellerTaxpayerNum 通过多个 seller_taxpayer_num(销方纳税人识别号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerTaxpayerNum(sellerTaxpayerNums []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerTaxpayerNums(sellerTaxpayerNums))
	return
}

// GetFromSellerAddress 通过单个 seller_address(销方地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerAddress(sellerAddress string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerAddress(sellerAddress))
	return
}

// GetsFromSellerAddress 通过多个 seller_address(销方地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerAddress(sellerAddresss []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerAddresss(sellerAddresss))
	return
}

// GetFromSellerPhone 通过单个 seller_phone(销方电话) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerPhone(sellerPhone string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerPhone(sellerPhone))
	return
}

// GetsFromSellerPhone 通过多个 seller_phone(销方电话) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerPhone(sellerPhones []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerPhones(sellerPhones))
	return
}

// GetFromSellerBankName 通过单个 seller_bank_name(销方银行名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerBankName(sellerBankName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerBankName(sellerBankName))
	return
}

// GetsFromSellerBankName 通过多个 seller_bank_name(销方银行名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerBankName(sellerBankNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerBankNames(sellerBankNames))
	return
}

// GetFromSellerBankAccount 通过单个 seller_bank_account(销方银行账号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromSellerBankAccount(sellerBankAccount string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerBankAccount(sellerBankAccount))
	return
}

// GetsFromSellerBankAccount 通过多个 seller_bank_account(销方银行账号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromSellerBankAccount(sellerBankAccounts []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithSellerBankAccounts(sellerBankAccounts))
	return
}

// GetFromHeaderType 通过单个 header_type(抬头类型抬头类型：1：个人、政府事业单位；2：企业) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromHeaderType(headerType string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithHeaderType(headerType))
	return
}

// GetsFromHeaderType 通过多个 header_type(抬头类型抬头类型：1：个人、政府事业单位；2：企业) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromHeaderType(headerTypes []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithHeaderTypes(headerTypes))
	return
}

// GetFromBuyerName 通过单个 buyer_name(购方名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerName(buyerName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerName(buyerName))
	return
}

// GetsFromBuyerName 通过多个 buyer_name(购方名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerName(buyerNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerNames(buyerNames))
	return
}

// GetFromBuyerTaxpayerNum 通过单个 buyer_taxpayer_num(购方纳税人识别号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerTaxpayerNum(buyerTaxpayerNum string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerTaxpayerNum(buyerTaxpayerNum))
	return
}

// GetsFromBuyerTaxpayerNum 通过多个 buyer_taxpayer_num(购方纳税人识别号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerTaxpayerNum(buyerTaxpayerNums []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerTaxpayerNums(buyerTaxpayerNums))
	return
}

// GetFromBuyerAddress 通过单个 buyer_address(购方地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerAddress(buyerAddress string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerAddress(buyerAddress))
	return
}

// GetsFromBuyerAddress 通过多个 buyer_address(购方地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerAddress(buyerAddresss []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerAddresss(buyerAddresss))
	return
}

// GetFromBuyerPhone 通过单个 buyer_phone(购方电话) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerPhone(buyerPhone string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerPhone(buyerPhone))
	return
}

// GetsFromBuyerPhone 通过多个 buyer_phone(购方电话) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerPhone(buyerPhones []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerPhones(buyerPhones))
	return
}

// GetFromBuyerBankName 通过单个 buyer_bank_name(购方银行名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerBankName(buyerBankName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerBankName(buyerBankName))
	return
}

// GetsFromBuyerBankName 通过多个 buyer_bank_name(购方银行名称) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerBankName(buyerBankNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerBankNames(buyerBankNames))
	return
}

// GetFromBuyerBankAccount 通过单个 buyer_bank_account(购方银行账号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromBuyerBankAccount(buyerBankAccount string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerBankAccount(buyerBankAccount))
	return
}

// GetsFromBuyerBankAccount 通过多个 buyer_bank_account(购方银行账号) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromBuyerBankAccount(buyerBankAccounts []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithBuyerBankAccounts(buyerBankAccounts))
	return
}

// GetFromInvoiceTypeCode 通过单个 invoice_type_code(发票类型 081：全电电子专票 082：全电电子普票) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromInvoiceTypeCode(invoiceTypeCode string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceTypeCode(invoiceTypeCode))
	return
}

// GetsFromInvoiceTypeCode 通过多个 invoice_type_code(发票类型 081：全电电子专票 082：全电电子普票) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromInvoiceTypeCode(invoiceTypeCodes []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceTypeCodes(invoiceTypeCodes))
	return
}

// GetFromInvoiceTime 通过单个 invoice_time(开票时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromInvoiceTime(invoiceTime int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceTime(invoiceTime))
	return
}

// GetsFromInvoiceTime 通过多个 invoice_time(开票时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromInvoiceTime(invoiceTimes []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithInvoiceTimes(invoiceTimes))
	return
}

// GetFromPdfURL 通过单个 pdf_url(发票文件PDF下载地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromPdfURL(pdfURL string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPdfURL(pdfURL))
	return
}

// GetsFromPdfURL 通过多个 pdf_url(发票文件PDF下载地址) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromPdfURL(pdfURLs []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPdfURLs(pdfURLs))
	return
}

// GetFromDrawerName 通过单个 drawer_name(开票人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromDrawerName(drawerName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDrawerName(drawerName))
	return
}

// GetsFromDrawerName 通过多个 drawer_name(开票人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromDrawerName(drawerNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDrawerNames(drawerNames))
	return
}

// GetFromPayeeName 通过单个 payee_name(收款人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromPayeeName(payeeName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayeeName(payeeName))
	return
}

// GetsFromPayeeName 通过多个 payee_name(收款人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromPayeeName(payeeNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithPayeeNames(payeeNames))
	return
}

// GetFromReviewerName 通过单个 reviewer_name(复核人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromReviewerName(reviewerName string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReviewerName(reviewerName))
	return
}

// GetsFromReviewerName 通过多个 reviewer_name(复核人姓名) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromReviewerName(reviewerNames []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithReviewerNames(reviewerNames))
	return
}

// GetFromRemark 通过单个 remark(备注) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromRemark(remark string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithRemark(remark))
	return
}

// GetsFromRemark 通过多个 remark(备注) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromRemark(remarks []string) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithRemarks(remarks))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromCreatedAt(createdAt int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromCreatedAt(createdAts []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromUpdatedAt(updatedAt int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetFromDeletedAt(deletedAt int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *InvoiceInfoDAO) GetsFromDeletedAt(deletedAts []int64) (results []*InvoiceInfo, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *InvoiceInfoDAO) FetchByPrimaryKey(id int64) (result *InvoiceInfo, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
