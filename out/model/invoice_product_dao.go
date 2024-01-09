package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:11
 * @Desc:   invoice_product 表的 DAO 层
 */

type InvoiceProductDAO struct {
	*_BaseDAO
}

// InvoiceProductDAO 初始化
func NewInvoiceProductDAO(ctx context.Context, db *gorm.DB) *InvoiceProductDAO {
	if db == nil {
		panic(fmt.Errorf("InvoiceProductDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &InvoiceProductDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&InvoiceProduct{}),
			db:               db,
			model:            InvoiceProduct{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          InvoiceProductAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *InvoiceProductDAO) GetTableName() string {
	invoiceProduct := &InvoiceProduct{}
	return invoiceProduct.TableName()
}

// Save 存在则更新，否则插入
func (obj *InvoiceProductDAO) Save(invoiceProduct *InvoiceProduct) (rowsAffected int64, err error) {
	if invoiceProduct.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(invoiceProduct, obj.WithID(invoiceProduct.ID))
	}
	return obj.Create(invoiceProduct)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *InvoiceProductDAO) Create(invoiceProduct interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(invoiceProduct)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ID] = ids })
}

// WithInvoicePrimaryID 设置 invoice_primary_id(发票数据表主键id) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithInvoicePrimaryID(invoicePrimaryID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.InvoicePrimaryID] = invoicePrimaryID })
}

// WithInvoicePrimaryIDs 设置 invoice_primary_id(发票数据表主键id) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithInvoicePrimaryIDs(invoicePrimaryIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.InvoicePrimaryID] = invoicePrimaryIDs })
}

// WithProjectName 设置 project_name(商品名称) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithProjectName(projectName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ProjectName] = projectName })
}

// WithProjectNames 设置 project_name(商品名称) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithProjectNames(projectNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ProjectName] = projectNames })
}

// WithTaxCode 设置 tax_code(税收商品编码) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithTaxCode(taxCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxCode] = taxCode })
}

// WithTaxCodes 设置 tax_code(税收商品编码) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithTaxCodes(taxCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxCode] = taxCodes })
}

// WithType 设置 type(税收商品类别 需要提供字典编码) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithType(_type string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Type] = _type })
}

// WithTypes 设置 type(税收商品类别 需要提供字典编码) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithTypes(_types []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Type] = _types })
}

// WithSpecial 设置 special(商品规格) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithSpecial(special string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Special] = special })
}

// WithSpecials 设置 special(商品规格) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithSpecials(specials []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Special] = specials })
}

// WithUnit 设置 unit(商品单位) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithUnit(unit string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Unit] = unit })
}

// WithUnits 设置 unit(商品单位) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithUnits(units []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Unit] = units })
}

// WithAmountNoTax 设置 amount_no_tax(商品不含税金额（单位：分）) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithAmountNoTax(amountNoTax int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.AmountNoTax] = amountNoTax })
}

// WithAmountNoTaxs 设置 amount_no_tax(商品不含税金额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithAmountNoTaxs(amountNoTaxs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.AmountNoTax] = amountNoTaxs })
}

// WithCount 设置 count(商品数量) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithCount(count int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Count] = count })
}

// WithCounts 设置 count(商品数量) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithCounts(counts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.Count] = counts })
}

// WithPriceNoTax 设置 price_no_tax(商品不含税单价（单位：分）) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithPriceNoTax(priceNoTax int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.PriceNoTax] = priceNoTax })
}

// WithPriceNoTaxs 设置 price_no_tax(商品不含税单价（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithPriceNoTaxs(priceNoTaxs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.PriceNoTax] = priceNoTaxs })
}

// WithTaxRate 设置 tax_rate(商品税率) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithTaxRate(taxRate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxRate] = taxRate })
}

// WithTaxRates 设置 tax_rate(商品税率) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithTaxRates(taxRates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxRate] = taxRates })
}

// WithTaxAmount 设置 tax_amount(商品税额（单位：分）) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithTaxAmount(taxAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxAmount] = taxAmount })
}

// WithTaxAmounts 设置 tax_amount(商品税额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithTaxAmounts(taxAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.TaxAmount] = taxAmounts })
}

// WithDiscountHasTax 设置 discount_has_tax(含税折扣总金额（单位：分）) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithDiscountHasTax(discountHasTax int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DiscountHasTax] = discountHasTax })
}

// WithDiscountHasTaxs 设置 discount_has_tax(含税折扣总金额（单位：分）) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithDiscountHasTaxs(discountHasTaxs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DiscountHasTax] = discountHasTaxs })
}

// WithDiscountFlag 设置 discount_flag(优惠政策标志 0：不使用优惠政策 1：使用优惠政策) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithDiscountFlag(discountFlag int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DiscountFlag] = discountFlag })
}

// WithDiscountFlags 设置 discount_flag(优惠政策标志 0：不使用优惠政策 1：使用优惠政策) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithDiscountFlags(discountFlags []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DiscountFlag] = discountFlags })
}

// WithZeroFlag 设置 zero_flag(零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithZeroFlag(zeroFlag int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ZeroFlag] = zeroFlag })
}

// WithZeroFlags 设置 zero_flag(零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithZeroFlags(zeroFlags []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.ZeroFlag] = zeroFlags })
}

// WithAddTaxFlag 设置 add_tax_flag(增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithAddTaxFlag(addTaxFlag string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.AddTaxFlag] = addTaxFlag })
}

// WithAddTaxFlags 设置 add_tax_flag(增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithAddTaxFlags(addTaxFlags []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.AddTaxFlag] = addTaxFlags })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *InvoiceProductDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *InvoiceProductDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[InvoiceProductColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *InvoiceProductDAO) GetByOption(opts ...Option) (result *InvoiceProduct, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *InvoiceProductDAO) GetListByOption(opts ...Option) (results []*InvoiceProduct, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *InvoiceProductDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *InvoiceProductDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      invoiceProduct: 要更新的数据
//      opts: 更新的条件
func (obj *InvoiceProductDAO) UpdateByOption(invoiceProduct *InvoiceProduct, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(invoiceProduct)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *InvoiceProductDAO) GetFromID(id int64) (result *InvoiceProduct, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromID(ids []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromInvoicePrimaryID 通过单个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromInvoicePrimaryID(invoicePrimaryID int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryID(invoicePrimaryID))
	return
}

// GetsFromInvoicePrimaryID 通过多个 invoice_primary_id(发票数据表主键id) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromInvoicePrimaryID(invoicePrimaryIDs []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithInvoicePrimaryIDs(invoicePrimaryIDs))
	return
}

// GetFromProjectName 通过单个 project_name(商品名称) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromProjectName(projectName string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithProjectName(projectName))
	return
}

// GetsFromProjectName 通过多个 project_name(商品名称) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromProjectName(projectNames []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithProjectNames(projectNames))
	return
}

// GetFromTaxCode 通过单个 tax_code(税收商品编码) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromTaxCode(taxCode string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxCode(taxCode))
	return
}

// GetsFromTaxCode 通过多个 tax_code(税收商品编码) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromTaxCode(taxCodes []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxCodes(taxCodes))
	return
}

// GetFromType 通过单个 type(税收商品类别 需要提供字典编码) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromType(_type string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithType(_type))
	return
}

// GetsFromType 通过多个 type(税收商品类别 需要提供字典编码) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromType(_types []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTypes(_types))
	return
}

// GetFromSpecial 通过单个 special(商品规格) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromSpecial(special string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithSpecial(special))
	return
}

// GetsFromSpecial 通过多个 special(商品规格) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromSpecial(specials []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithSpecials(specials))
	return
}

// GetFromUnit 通过单个 unit(商品单位) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromUnit(unit string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUnit(unit))
	return
}

// GetsFromUnit 通过多个 unit(商品单位) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromUnit(units []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUnits(units))
	return
}

// GetFromAmountNoTax 通过单个 amount_no_tax(商品不含税金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromAmountNoTax(amountNoTax int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithAmountNoTax(amountNoTax))
	return
}

// GetsFromAmountNoTax 通过多个 amount_no_tax(商品不含税金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromAmountNoTax(amountNoTaxs []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithAmountNoTaxs(amountNoTaxs))
	return
}

// GetFromCount 通过单个 count(商品数量) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromCount(count int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCount(count))
	return
}

// GetsFromCount 通过多个 count(商品数量) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromCount(counts []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCounts(counts))
	return
}

// GetFromPriceNoTax 通过单个 price_no_tax(商品不含税单价（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromPriceNoTax(priceNoTax int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithPriceNoTax(priceNoTax))
	return
}

// GetsFromPriceNoTax 通过多个 price_no_tax(商品不含税单价（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromPriceNoTax(priceNoTaxs []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithPriceNoTaxs(priceNoTaxs))
	return
}

// GetFromTaxRate 通过单个 tax_rate(商品税率) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromTaxRate(taxRate int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxRate(taxRate))
	return
}

// GetsFromTaxRate 通过多个 tax_rate(商品税率) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromTaxRate(taxRates []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxRates(taxRates))
	return
}

// GetFromTaxAmount 通过单个 tax_amount(商品税额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromTaxAmount(taxAmount int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxAmount(taxAmount))
	return
}

// GetsFromTaxAmount 通过多个 tax_amount(商品税额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromTaxAmount(taxAmounts []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithTaxAmounts(taxAmounts))
	return
}

// GetFromDiscountHasTax 通过单个 discount_has_tax(含税折扣总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromDiscountHasTax(discountHasTax int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDiscountHasTax(discountHasTax))
	return
}

// GetsFromDiscountHasTax 通过多个 discount_has_tax(含税折扣总金额（单位：分）) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromDiscountHasTax(discountHasTaxs []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDiscountHasTaxs(discountHasTaxs))
	return
}

// GetFromDiscountFlag 通过单个 discount_flag(优惠政策标志 0：不使用优惠政策 1：使用优惠政策) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromDiscountFlag(discountFlag int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDiscountFlag(discountFlag))
	return
}

// GetsFromDiscountFlag 通过多个 discount_flag(优惠政策标志 0：不使用优惠政策 1：使用优惠政策) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromDiscountFlag(discountFlags []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDiscountFlags(discountFlags))
	return
}

// GetFromZeroFlag 通过单个 zero_flag(零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromZeroFlag(zeroFlag int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithZeroFlag(zeroFlag))
	return
}

// GetsFromZeroFlag 通过多个 zero_flag(零税率标识 零税率标识： 空：非零税率 1：免税、出口零税 2：不征税 3：普通零税率) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromZeroFlag(zeroFlags []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithZeroFlags(zeroFlags))
	return
}

// GetFromAddTaxFlag 通过单个 add_tax_flag(增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromAddTaxFlag(addTaxFlag string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithAddTaxFlag(addTaxFlag))
	return
}

// GetsFromAddTaxFlag 通过多个 add_tax_flag(增值税特殊管理 增值税特殊管理：优惠政策标志为1时该字段必传，目前支持： 免税、按5%（3%、1.5%）简易征收、不征税等) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromAddTaxFlag(addTaxFlags []string) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithAddTaxFlags(addTaxFlags))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromCreatedAt(createdAt int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromCreatedAt(createdAts []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromUpdatedAt(updatedAt int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetFromDeletedAt(deletedAt int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *InvoiceProductDAO) GetsFromDeletedAt(deletedAts []int64) (results []*InvoiceProduct, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *InvoiceProductDAO) FetchByPrimaryKey(id int64) (result *InvoiceProduct, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
