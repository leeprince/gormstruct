package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:34
 * @Desc:   purchaser 表的 DAO 层
 */

type PurchaserDAO struct {
	*_BaseDAO
}

// PurchaserDAO 初始化
func NewPurchaserDAO(ctx context.Context, db *gorm.DB) *PurchaserDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Purchaser{}),
			db:               db,
			model:            Purchaser{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserDAO) GetTableName() string {
	purchaser := &Purchaser{}
	return purchaser.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserDAO) Save(purchaser *Purchaser) (rowsAffected int64, err error) {
	if purchaser.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaser, obj.WithID(purchaser.ID))
	}
	return obj.Create(purchaser)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserDAO) Create(purchaser interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaser)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *PurchaserDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ID] = ids })
}

// WithOpenOrgIDs 设置 open_org_ids(平台企业租户id集合) 字段作为 option 条件
func (obj *PurchaserDAO) WithOpenOrgIDs(openOrgIDs string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.OpenOrgIDs] = openOrgIDs })
}

// WithOpenOrgIDss 设置 open_org_ids(平台企业租户id集合) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithOpenOrgIDss(openOrgIDss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.OpenOrgIDs] = openOrgIDss })
}

// WithPurchaserID 设置 purchaser_id(企业id（采购商ID）) 字段作为 option 条件
func (obj *PurchaserDAO) WithPurchaserID(purchaserID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.PurchaserID] = purchaserID })
}

// WithPurchaserIDs 设置 purchaser_id(企业id（采购商ID）) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithPurchaserIDs(purchaserIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.PurchaserID] = purchaserIDs })
}

// WithRegisterDate 设置 register_date(企业注册时间) 字段作为 option 条件
func (obj *PurchaserDAO) WithRegisterDate(registerDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterDate] = registerDate })
}

// WithRegisterDates 设置 register_date(企业注册时间) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithRegisterDates(registerDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterDate] = registerDates })
}

// WithState 设置 state(采购商状态 0-异常 1-正常 2-退出) 字段作为 option 条件
func (obj *PurchaserDAO) WithState(state int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.State] = state })
}

// WithStates 设置 state(采购商状态 0-异常 1-正常 2-退出) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithStates(states []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.State] = states })
}

// WithStateDesc 设置 state_desc(采购商状态说明) 字段作为 option 条件
func (obj *PurchaserDAO) WithStateDesc(stateDesc string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.StateDesc] = stateDesc })
}

// WithStateDescs 设置 state_desc(采购商状态说明) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithStateDescs(stateDescs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.StateDesc] = stateDescs })
}

// WithPurchaserName 设置 purchaser_name(企业（采购商）名称) 字段作为 option 条件
func (obj *PurchaserDAO) WithPurchaserName(purchaserName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.PurchaserName] = purchaserName })
}

// WithPurchaserNames 设置 purchaser_name(企业（采购商）名称) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithPurchaserNames(purchaserNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.PurchaserName] = purchaserNames })
}

// WithCreditNo 设置 credit_no(统一社会信用代码) 字段作为 option 条件
func (obj *PurchaserDAO) WithCreditNo(creditNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CreditNo] = creditNo })
}

// WithCreditNos 设置 credit_no(统一社会信用代码) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithCreditNos(creditNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CreditNo] = creditNos })
}

// WithBusinessLicenseImgURL 设置 business_license_img_url(营业执照) 字段作为 option 条件
func (obj *PurchaserDAO) WithBusinessLicenseImgURL(businessLicenseImgURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessLicenseImgURL] = businessLicenseImgURL })
}

// WithBusinessLicenseImgURLs 设置 business_license_img_url(营业执照) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithBusinessLicenseImgURLs(businessLicenseImgURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessLicenseImgURL] = businessLicenseImgURLs })
}

// WithIndustryCode 设置 industry_code(行业代码) 字段作为 option 条件
func (obj *PurchaserDAO) WithIndustryCode(industryCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.IndustryCode] = industryCode })
}

// WithIndustryCodes 设置 industry_code(行业代码) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithIndustryCodes(industryCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.IndustryCode] = industryCodes })
}

// WithRegistrationOrg 设置 registration_org(登记机关) 字段作为 option 条件
func (obj *PurchaserDAO) WithRegistrationOrg(registrationOrg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegistrationOrg] = registrationOrg })
}

// WithRegistrationOrgs 设置 registration_org(登记机关) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithRegistrationOrgs(registrationOrgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegistrationOrg] = registrationOrgs })
}

// WithBuildDate 设置 build_date(成立日期) 字段作为 option 条件
func (obj *PurchaserDAO) WithBuildDate(buildDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BuildDate] = buildDate })
}

// WithBuildDates 设置 build_date(成立日期) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithBuildDates(buildDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BuildDate] = buildDates })
}

// WithRegisterMoney 设置 register_money(注册资本（单位分）) 字段作为 option 条件
func (obj *PurchaserDAO) WithRegisterMoney(registerMoney int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterMoney] = registerMoney })
}

// WithRegisterMoneys 设置 register_money(注册资本（单位分）) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithRegisterMoneys(registerMoneys []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterMoney] = registerMoneys })
}

// WithEmployeeNum 设置 employee_num(从业人数) 字段作为 option 条件
func (obj *PurchaserDAO) WithEmployeeNum(employeeNum int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.EmployeeNum] = employeeNum })
}

// WithEmployeeNums 设置 employee_num(从业人数) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithEmployeeNums(employeeNums []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.EmployeeNum] = employeeNums })
}

// WithProcureAmt 设置 procure_amt(月均采购预算) 字段作为 option 条件
func (obj *PurchaserDAO) WithProcureAmt(procureAmt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ProcureAmt] = procureAmt })
}

// WithProcureAmts 设置 procure_amt(月均采购预算) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithProcureAmts(procureAmts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ProcureAmt] = procureAmts })
}

// WithRegisterAreaCode 设置 register_area_code(注册地址行政区划代码（省市区）) 字段作为 option 条件
func (obj *PurchaserDAO) WithRegisterAreaCode(registerAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterAreaCode] = registerAreaCode })
}

// WithRegisterAreaCodes 设置 register_area_code(注册地址行政区划代码（省市区）) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithRegisterAreaCodes(registerAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterAreaCode] = registerAreaCodes })
}

// WithRegisterAddress 设置 register_address(注册地址详细信息) 字段作为 option 条件
func (obj *PurchaserDAO) WithRegisterAddress(registerAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterAddress] = registerAddress })
}

// WithRegisterAddresss 设置 register_address(注册地址详细信息) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithRegisterAddresss(registerAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.RegisterAddress] = registerAddresss })
}

// WithBusinessAreaCode 设置 business_area_code(生产经营地址行政区划代码（省市区）) 字段作为 option 条件
func (obj *PurchaserDAO) WithBusinessAreaCode(businessAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessAreaCode] = businessAreaCode })
}

// WithBusinessAreaCodes 设置 business_area_code(生产经营地址行政区划代码（省市区）) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithBusinessAreaCodes(businessAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessAreaCode] = businessAreaCodes })
}

// WithBusinessAddress 设置 business_address(生产经营地址详细信息) 字段作为 option 条件
func (obj *PurchaserDAO) WithBusinessAddress(businessAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessAddress] = businessAddress })
}

// WithBusinessAddresss 设置 business_address(生产经营地址详细信息) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithBusinessAddresss(businessAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessAddress] = businessAddresss })
}

// WithEntType 设置 ent_type(公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社) 字段作为 option 条件
func (obj *PurchaserDAO) WithEntType(entType int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.EntType] = entType })
}

// WithEntTypes 设置 ent_type(公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithEntTypes(entTypes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.EntType] = entTypes })
}

// WithBusinessScope 设置 business_scope(经营范围) 字段作为 option 条件
func (obj *PurchaserDAO) WithBusinessScope(businessScope string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessScope] = businessScope })
}

// WithBusinessScopes 设置 business_scope(经营范围) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithBusinessScopes(businessScopes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.BusinessScope] = businessScopes })
}

// WithLegalName 设置 legal_name(法人姓名) 字段作为 option 条件
func (obj *PurchaserDAO) WithLegalName(legalName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalName] = legalName })
}

// WithLegalNames 设置 legal_name(法人姓名) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithLegalNames(legalNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalName] = legalNames })
}

// WithLegalCardType 设置 legal_card_type(法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证) 字段作为 option 条件
func (obj *PurchaserDAO) WithLegalCardType(legalCardType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalCardType] = legalCardType })
}

// WithLegalCardTypes 设置 legal_card_type(法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithLegalCardTypes(legalCardTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalCardType] = legalCardTypes })
}

// WithLegalCardNo 设置 legal_card_no(法人身份证号) 字段作为 option 条件
func (obj *PurchaserDAO) WithLegalCardNo(legalCardNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalCardNo] = legalCardNo })
}

// WithLegalCardNos 设置 legal_card_no(法人身份证号) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithLegalCardNos(legalCardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalCardNo] = legalCardNos })
}

// WithLegalPhone 设置 legal_phone(法人手机号码) 字段作为 option 条件
func (obj *PurchaserDAO) WithLegalPhone(legalPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalPhone] = legalPhone })
}

// WithLegalPhones 设置 legal_phone(法人手机号码) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithLegalPhones(legalPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.LegalPhone] = legalPhones })
}

// WithMail 设置 mail(企业邮箱) 字段作为 option 条件
func (obj *PurchaserDAO) WithMail(mail string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.Mail] = mail })
}

// WithMails 设置 mail(企业邮箱) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithMails(mails []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.Mail] = mails })
}

// WithPhone 设置 phone(企业电话) 字段作为 option 条件
func (obj *PurchaserDAO) WithPhone(phone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.Phone] = phone })
}

// WithPhones 设置 phone(企业电话) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithPhones(phones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.Phone] = phones })
}

// WithContactPersonName 设置 contact_person_name(企业联系人姓名) 字段作为 option 条件
func (obj *PurchaserDAO) WithContactPersonName(contactPersonName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ContactPersonName] = contactPersonName })
}

// WithContactPersonNames 设置 contact_person_name(企业联系人姓名) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithContactPersonNames(contactPersonNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ContactPersonName] = contactPersonNames })
}

// WithContactPersonPhone 设置 contact_person_phone(联系人电话) 字段作为 option 条件
func (obj *PurchaserDAO) WithContactPersonPhone(contactPersonPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ContactPersonPhone] = contactPersonPhone })
}

// WithContactPersonPhones 设置 contact_person_phone(联系人电话) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithContactPersonPhones(contactPersonPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ContactPersonPhone] = contactPersonPhones })
}

// WithManageTax 设置 manage_tax(主管税务局) 字段作为 option 条件
func (obj *PurchaserDAO) WithManageTax(manageTax string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ManageTax] = manageTax })
}

// WithManageTaxs 设置 manage_tax(主管税务局) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithManageTaxs(manageTaxs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ManageTax] = manageTaxs })
}

// WithManageTaxOrg 设置 manage_tax_org(主管税务机关（科、分局）) 字段作为 option 条件
func (obj *PurchaserDAO) WithManageTaxOrg(manageTaxOrg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ManageTaxOrg] = manageTaxOrg })
}

// WithManageTaxOrgs 设置 manage_tax_org(主管税务机关（科、分局）) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithManageTaxOrgs(manageTaxOrgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.ManageTaxOrg] = manageTaxOrgs })
}

// WithCleanProtocalFile 设置 clean_protocal_file(是否清理协议文件 1：是；其他否) 字段作为 option 条件
func (obj *PurchaserDAO) WithCleanProtocalFile(cleanProtocalFile int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CleanProtocalFile] = cleanProtocalFile })
}

// WithCleanProtocalFiles 设置 clean_protocal_file(是否清理协议文件 1：是；其他否) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithCleanProtocalFiles(cleanProtocalFiles []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CleanProtocalFile] = cleanProtocalFiles })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *PurchaserDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *PurchaserDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *PurchaserDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *PurchaserDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserDAO) GetByOption(opts ...Option) (result *Purchaser, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserDAO) GetListByOption(opts ...Option) (results []*Purchaser, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaser: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserDAO) UpdateByOption(purchaser *Purchaser, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaser)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *PurchaserDAO) GetFromID(id int64) (result *Purchaser, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromID(ids []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOpenOrgIDs 通过单个 open_org_ids(平台企业租户id集合) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromOpenOrgIDs(openOrgIDs string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetsFromOpenOrgIDs 通过多个 open_org_ids(平台企业租户id集合) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromOpenOrgIDs(openOrgIDss []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDss(openOrgIDss))
	return
}

// GetFromPurchaserID 通过单个 purchaser_id(企业id（采购商ID）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromPurchaserID(purchaserID string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserID(purchaserID))
	return
}

// GetsFromPurchaserID 通过多个 purchaser_id(企业id（采购商ID）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromPurchaserID(purchaserIDs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserIDs(purchaserIDs))
	return
}

// GetFromRegisterDate 通过单个 register_date(企业注册时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromRegisterDate(registerDate int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterDate(registerDate))
	return
}

// GetsFromRegisterDate 通过多个 register_date(企业注册时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromRegisterDate(registerDates []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterDates(registerDates))
	return
}

// GetFromState 通过单个 state(采购商状态 0-异常 1-正常 2-退出) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromState(state int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithState(state))
	return
}

// GetsFromState 通过多个 state(采购商状态 0-异常 1-正常 2-退出) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromState(states []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithStates(states))
	return
}

// GetFromStateDesc 通过单个 state_desc(采购商状态说明) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromStateDesc(stateDesc string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithStateDesc(stateDesc))
	return
}

// GetsFromStateDesc 通过多个 state_desc(采购商状态说明) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromStateDesc(stateDescs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithStateDescs(stateDescs))
	return
}

// GetFromPurchaserName 通过单个 purchaser_name(企业（采购商）名称) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromPurchaserName(purchaserName string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserName(purchaserName))
	return
}

// GetsFromPurchaserName 通过多个 purchaser_name(企业（采购商）名称) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromPurchaserName(purchaserNames []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserNames(purchaserNames))
	return
}

// GetFromCreditNo 通过单个 credit_no(统一社会信用代码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromCreditNo(creditNo string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCreditNo(creditNo))
	return
}

// GetsFromCreditNo 通过多个 credit_no(统一社会信用代码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromCreditNo(creditNos []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCreditNos(creditNos))
	return
}

// GetFromBusinessLicenseImgURL 通过单个 business_license_img_url(营业执照) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromBusinessLicenseImgURL(businessLicenseImgURL string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessLicenseImgURL(businessLicenseImgURL))
	return
}

// GetsFromBusinessLicenseImgURL 通过多个 business_license_img_url(营业执照) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromBusinessLicenseImgURL(businessLicenseImgURLs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessLicenseImgURLs(businessLicenseImgURLs))
	return
}

// GetFromIndustryCode 通过单个 industry_code(行业代码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromIndustryCode(industryCode string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithIndustryCode(industryCode))
	return
}

// GetsFromIndustryCode 通过多个 industry_code(行业代码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromIndustryCode(industryCodes []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithIndustryCodes(industryCodes))
	return
}

// GetFromRegistrationOrg 通过单个 registration_org(登记机关) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromRegistrationOrg(registrationOrg string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegistrationOrg(registrationOrg))
	return
}

// GetsFromRegistrationOrg 通过多个 registration_org(登记机关) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromRegistrationOrg(registrationOrgs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegistrationOrgs(registrationOrgs))
	return
}

// GetFromBuildDate 通过单个 build_date(成立日期) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromBuildDate(buildDate int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBuildDate(buildDate))
	return
}

// GetsFromBuildDate 通过多个 build_date(成立日期) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromBuildDate(buildDates []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBuildDates(buildDates))
	return
}

// GetFromRegisterMoney 通过单个 register_money(注册资本（单位分）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromRegisterMoney(registerMoney int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterMoney(registerMoney))
	return
}

// GetsFromRegisterMoney 通过多个 register_money(注册资本（单位分）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromRegisterMoney(registerMoneys []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterMoneys(registerMoneys))
	return
}

// GetFromEmployeeNum 通过单个 employee_num(从业人数) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromEmployeeNum(employeeNum int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithEmployeeNum(employeeNum))
	return
}

// GetsFromEmployeeNum 通过多个 employee_num(从业人数) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromEmployeeNum(employeeNums []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithEmployeeNums(employeeNums))
	return
}

// GetFromProcureAmt 通过单个 procure_amt(月均采购预算) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromProcureAmt(procureAmt int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithProcureAmt(procureAmt))
	return
}

// GetsFromProcureAmt 通过多个 procure_amt(月均采购预算) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromProcureAmt(procureAmts []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithProcureAmts(procureAmts))
	return
}

// GetFromRegisterAreaCode 通过单个 register_area_code(注册地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromRegisterAreaCode(registerAreaCode string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAreaCode(registerAreaCode))
	return
}

// GetsFromRegisterAreaCode 通过多个 register_area_code(注册地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromRegisterAreaCode(registerAreaCodes []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAreaCodes(registerAreaCodes))
	return
}

// GetFromRegisterAddress 通过单个 register_address(注册地址详细信息) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromRegisterAddress(registerAddress string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAddress(registerAddress))
	return
}

// GetsFromRegisterAddress 通过多个 register_address(注册地址详细信息) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromRegisterAddress(registerAddresss []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAddresss(registerAddresss))
	return
}

// GetFromBusinessAreaCode 通过单个 business_area_code(生产经营地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromBusinessAreaCode(businessAreaCode string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAreaCode(businessAreaCode))
	return
}

// GetsFromBusinessAreaCode 通过多个 business_area_code(生产经营地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromBusinessAreaCode(businessAreaCodes []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAreaCodes(businessAreaCodes))
	return
}

// GetFromBusinessAddress 通过单个 business_address(生产经营地址详细信息) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromBusinessAddress(businessAddress string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAddress(businessAddress))
	return
}

// GetsFromBusinessAddress 通过多个 business_address(生产经营地址详细信息) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromBusinessAddress(businessAddresss []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAddresss(businessAddresss))
	return
}

// GetFromEntType 通过单个 ent_type(公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromEntType(entType int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithEntType(entType))
	return
}

// GetsFromEntType 通过多个 ent_type(公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromEntType(entTypes []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithEntTypes(entTypes))
	return
}

// GetFromBusinessScope 通过单个 business_scope(经营范围) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromBusinessScope(businessScope string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessScope(businessScope))
	return
}

// GetsFromBusinessScope 通过多个 business_scope(经营范围) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromBusinessScope(businessScopes []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessScopes(businessScopes))
	return
}

// GetFromLegalName 通过单个 legal_name(法人姓名) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromLegalName(legalName string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalName(legalName))
	return
}

// GetsFromLegalName 通过多个 legal_name(法人姓名) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromLegalName(legalNames []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalNames(legalNames))
	return
}

// GetFromLegalCardType 通过单个 legal_card_type(法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromLegalCardType(legalCardType string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardType(legalCardType))
	return
}

// GetsFromLegalCardType 通过多个 legal_card_type(法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromLegalCardType(legalCardTypes []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardTypes(legalCardTypes))
	return
}

// GetFromLegalCardNo 通过单个 legal_card_no(法人身份证号) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromLegalCardNo(legalCardNo string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardNo(legalCardNo))
	return
}

// GetsFromLegalCardNo 通过多个 legal_card_no(法人身份证号) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromLegalCardNo(legalCardNos []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardNos(legalCardNos))
	return
}

// GetFromLegalPhone 通过单个 legal_phone(法人手机号码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromLegalPhone(legalPhone string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalPhone(legalPhone))
	return
}

// GetsFromLegalPhone 通过多个 legal_phone(法人手机号码) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromLegalPhone(legalPhones []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithLegalPhones(legalPhones))
	return
}

// GetFromMail 通过单个 mail(企业邮箱) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromMail(mail string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithMail(mail))
	return
}

// GetsFromMail 通过多个 mail(企业邮箱) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromMail(mails []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithMails(mails))
	return
}

// GetFromPhone 通过单个 phone(企业电话) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromPhone(phone string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPhone(phone))
	return
}

// GetsFromPhone 通过多个 phone(企业电话) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromPhone(phones []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithPhones(phones))
	return
}

// GetFromContactPersonName 通过单个 contact_person_name(企业联系人姓名) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromContactPersonName(contactPersonName string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonName(contactPersonName))
	return
}

// GetsFromContactPersonName 通过多个 contact_person_name(企业联系人姓名) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromContactPersonName(contactPersonNames []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonNames(contactPersonNames))
	return
}

// GetFromContactPersonPhone 通过单个 contact_person_phone(联系人电话) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromContactPersonPhone(contactPersonPhone string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonPhone(contactPersonPhone))
	return
}

// GetsFromContactPersonPhone 通过多个 contact_person_phone(联系人电话) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromContactPersonPhone(contactPersonPhones []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonPhones(contactPersonPhones))
	return
}

// GetFromManageTax 通过单个 manage_tax(主管税务局) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromManageTax(manageTax string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithManageTax(manageTax))
	return
}

// GetsFromManageTax 通过多个 manage_tax(主管税务局) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromManageTax(manageTaxs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxs(manageTaxs))
	return
}

// GetFromManageTaxOrg 通过单个 manage_tax_org(主管税务机关（科、分局）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromManageTaxOrg(manageTaxOrg string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxOrg(manageTaxOrg))
	return
}

// GetsFromManageTaxOrg 通过多个 manage_tax_org(主管税务机关（科、分局）) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromManageTaxOrg(manageTaxOrgs []string) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxOrgs(manageTaxOrgs))
	return
}

// GetFromCleanProtocalFile 通过单个 clean_protocal_file(是否清理协议文件 1：是；其他否) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromCleanProtocalFile(cleanProtocalFile int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCleanProtocalFile(cleanProtocalFile))
	return
}

// GetsFromCleanProtocalFile 通过多个 clean_protocal_file(是否清理协议文件 1：是；其他否) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromCleanProtocalFile(cleanProtocalFiles []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCleanProtocalFiles(cleanProtocalFiles))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromCreatedAt(createdAt int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromCreatedAt(createdAts []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromUpdatedAt(updatedAt int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetFromDeletedAt(deletedAt int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *PurchaserDAO) GetsFromDeletedAt(deletedAts []int64) (results []*Purchaser, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserDAO) FetchByPrimaryKey(id int64) (result *Purchaser, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
