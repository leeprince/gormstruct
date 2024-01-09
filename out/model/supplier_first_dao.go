package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 16:48:53
 * @Desc:   supplier_first 表的 DAO 层
 */

type SupplierFirstDAO struct {
	*_BaseDAO
}

// SupplierFirstDAO 初始化
func NewSupplierFirstDAO(ctx context.Context, db *gorm.DB) *SupplierFirstDAO {
	if db == nil {
		panic(fmt.Errorf("SupplierFirstDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SupplierFirstDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&SupplierFirst{}),
			db:               db,
			model:            SupplierFirst{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          SupplierFirstAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *SupplierFirstDAO) GetTableName() string {
	supplierFirst := &SupplierFirst{}
	return supplierFirst.TableName()
}

// Save 存在则更新，否则插入
func (obj *SupplierFirstDAO) Save(supplierFirst *SupplierFirst) (rowsAffected int64, err error) {
	if supplierFirst.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(supplierFirst, obj.WithID(supplierFirst.ID))
	}
	return obj.Create(supplierFirst)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *SupplierFirstDAO) Create(supplierFirst interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(supplierFirst)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithOrgID 设置 org_id(一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商高灯云开放租户ID) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商高灯云开放租户ID) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.OpenOrgID] = openOrgIDs })
}

// WithName 设置 name(平台方企业名称) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Name] = name })
}

// WithNames 设置 name(平台方企业名称) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Name] = names })
}

// WithCreditNo 设置 credit_no(统一社会信用代码) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithCreditNo(creditNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CreditNo] = creditNo })
}

// WithCreditNos 设置 credit_no(统一社会信用代码) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithCreditNos(creditNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CreditNo] = creditNos })
}

// WithBusinessLicenseImgURL 设置 business_license_img_url(营业执照) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessLicenseImgURL(businessLicenseImgURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessLicenseImgURL] = businessLicenseImgURL })
}

// WithBusinessLicenseImgURLs 设置 business_license_img_url(营业执照) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessLicenseImgURLs(businessLicenseImgURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessLicenseImgURL] = businessLicenseImgURLs })
}

// WithIndustryCode 设置 industry_code(行业代码) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithIndustryCode(industryCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.IndustryCode] = industryCode })
}

// WithIndustryCodes 设置 industry_code(行业代码) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithIndustryCodes(industryCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.IndustryCode] = industryCodes })
}

// WithRegistrationOrg 设置 registration_org(登记机关) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithRegistrationOrg(registrationOrg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegistrationOrg] = registrationOrg })
}

// WithRegistrationOrgs 设置 registration_org(登记机关) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithRegistrationOrgs(registrationOrgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegistrationOrg] = registrationOrgs })
}

// WithBuildDate 设置 build_date(成立日期) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithBuildDate(buildDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BuildDate] = buildDate })
}

// WithBuildDates 设置 build_date(成立日期) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithBuildDates(buildDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BuildDate] = buildDates })
}

// WithRegisterMoney 设置 register_money(注册资本（单位：元，小数点后两位）) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterMoney(registerMoney float64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterMoney] = registerMoney })
}

// WithRegisterMoneys 设置 register_money(注册资本（单位：元，小数点后两位）) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterMoneys(registerMoneys []float64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterMoney] = registerMoneys })
}

// WithEmployeeNum 设置 employee_num(从业人数) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithEmployeeNum(employeeNum int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.EmployeeNum] = employeeNum })
}

// WithEmployeeNums 设置 employee_num(从业人数) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithEmployeeNums(employeeNums []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.EmployeeNum] = employeeNums })
}

// WithRegisterAreaCode 设置 register_area_code(注册地址行政区划代码（省市区）) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterAreaCode(registerAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterAreaCode] = registerAreaCode })
}

// WithRegisterAreaCodes 设置 register_area_code(注册地址行政区划代码（省市区）) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterAreaCodes(registerAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterAreaCode] = registerAreaCodes })
}

// WithRegisterAddress 设置 register_address(注册地址详细信息) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterAddress(registerAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterAddress] = registerAddress })
}

// WithRegisterAddresss 设置 register_address(注册地址详细信息) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithRegisterAddresss(registerAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.RegisterAddress] = registerAddresss })
}

// WithBusinessAreaCode 设置 business_area_code(生产经营地址行政区划代码（省市区）) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessAreaCode(businessAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessAreaCode] = businessAreaCode })
}

// WithBusinessAreaCodes 设置 business_area_code(生产经营地址行政区划代码（省市区）) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessAreaCodes(businessAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessAreaCode] = businessAreaCodes })
}

// WithBusinessAddress 设置 business_address(生产经营地址详细信息) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessAddress(businessAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessAddress] = businessAddress })
}

// WithBusinessAddresss 设置 business_address(生产经营地址详细信息) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessAddresss(businessAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessAddress] = businessAddresss })
}

// WithCompanyType 设置 company_type(公司类型) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithCompanyType(companyType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CompanyType] = companyType })
}

// WithCompanyTypes 设置 company_type(公司类型) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithCompanyTypes(companyTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CompanyType] = companyTypes })
}

// WithBusinessScope 设置 business_scope(经营范围) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessScope(businessScope string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessScope] = businessScope })
}

// WithBusinessScopes 设置 business_scope(经营范围) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithBusinessScopes(businessScopes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.BusinessScope] = businessScopes })
}

// WithLegalName 设置 legal_name(法人姓名) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithLegalName(legalName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalName] = legalName })
}

// WithLegalNames 设置 legal_name(法人姓名) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithLegalNames(legalNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalName] = legalNames })
}

// WithLegalCardType 设置 legal_card_type(法人身份证件种类) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithLegalCardType(legalCardType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalCardType] = legalCardType })
}

// WithLegalCardTypes 设置 legal_card_type(法人身份证件种类) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithLegalCardTypes(legalCardTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalCardType] = legalCardTypes })
}

// WithLegalCardNo 设置 legal_card_no(法人身份证号) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithLegalCardNo(legalCardNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalCardNo] = legalCardNo })
}

// WithLegalCardNos 设置 legal_card_no(法人身份证号) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithLegalCardNos(legalCardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalCardNo] = legalCardNos })
}

// WithLegalPhone 设置 legal_phone(法人手机号码) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithLegalPhone(legalPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalPhone] = legalPhone })
}

// WithLegalPhones 设置 legal_phone(法人手机号码) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithLegalPhones(legalPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.LegalPhone] = legalPhones })
}

// WithMail 设置 mail(企业邮箱) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithMail(mail string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Mail] = mail })
}

// WithMails 设置 mail(企业邮箱) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithMails(mails []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Mail] = mails })
}

// WithPhone 设置 phone(企业电话) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithPhone(phone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Phone] = phone })
}

// WithPhones 设置 phone(企业电话) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithPhones(phones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.Phone] = phones })
}

// WithContactPersonName 设置 contact_person_name(企业联系人姓名) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithContactPersonName(contactPersonName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ContactPersonName] = contactPersonName })
}

// WithContactPersonNames 设置 contact_person_name(企业联系人姓名) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithContactPersonNames(contactPersonNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ContactPersonName] = contactPersonNames })
}

// WithContactPersonPhone 设置 contact_person_phone(联系人电话) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithContactPersonPhone(contactPersonPhone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ContactPersonPhone] = contactPersonPhone })
}

// WithContactPersonPhones 设置 contact_person_phone(联系人电话) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithContactPersonPhones(contactPersonPhones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ContactPersonPhone] = contactPersonPhones })
}

// WithPlatformURL 设置 platform_url(平台网址) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithPlatformURL(platformURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.PlatformURL] = platformURL })
}

// WithPlatformURLs 设置 platform_url(平台网址) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithPlatformURLs(platformURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.PlatformURL] = platformURLs })
}

// WithManageTax 设置 manage_tax(主管税务局) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithManageTax(manageTax string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ManageTax] = manageTax })
}

// WithManageTaxs 设置 manage_tax(主管税务局) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithManageTaxs(manageTaxs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ManageTax] = manageTaxs })
}

// WithManageTaxOrg 设置 manage_tax_org(主管税务机关（科、分局）) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithManageTaxOrg(manageTaxOrg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ManageTaxOrg] = manageTaxOrg })
}

// WithManageTaxOrgs 设置 manage_tax_org(主管税务机关（科、分局）) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithManageTaxOrgs(manageTaxOrgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.ManageTaxOrg] = manageTaxOrgs })
}

// WithQualificationURL1 设置 qualification_url1(资质文件1) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL1(qualificationURL1 string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL1] = qualificationURL1 })
}

// WithQualificationURL1s 设置 qualification_url1(资质文件1) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL1s(qualificationURL1s []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL1] = qualificationURL1s })
}

// WithQualificationURL2 设置 qualification_url2(资质文件2) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL2(qualificationURL2 string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL2] = qualificationURL2 })
}

// WithQualificationURL2s 设置 qualification_url2(资质文件2) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL2s(qualificationURL2s []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL2] = qualificationURL2s })
}

// WithQualificationURL3 设置 qualification_url3(资质文件3) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL3(qualificationURL3 string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL3] = qualificationURL3 })
}

// WithQualificationURL3s 设置 qualification_url3(资质文件3) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL3s(qualificationURL3s []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL3] = qualificationURL3s })
}

// WithQualificationURL4 设置 qualification_url4(资质文件4) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL4(qualificationURL4 string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL4] = qualificationURL4 })
}

// WithQualificationURL4s 设置 qualification_url4(资质文件4) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithQualificationURL4s(qualificationURL4s []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.QualificationURL4] = qualificationURL4s })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *SupplierFirstDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *SupplierFirstDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierFirstColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *SupplierFirstDAO) GetByOption(opts ...Option) (result *SupplierFirst, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *SupplierFirstDAO) GetListByOption(opts ...Option) (results []*SupplierFirst, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *SupplierFirstDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *SupplierFirstDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      supplierFirst: 要更新的数据
//      opts: 更新的条件
func (obj *SupplierFirstDAO) UpdateByOption(supplierFirst *SupplierFirst, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(supplierFirst)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromOrgID 通过单个 org_id(一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入) 字段值，获取单条记录
func (obj *SupplierFirstDAO) GetFromOrgID(orgID int64) (result *SupplierFirst, err error) {
	result, err = obj.GetByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromOrgID(orgIDs []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商高灯云开放租户ID) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromOpenOrgID(openOrgID string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商高灯云开放租户ID) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromName 通过单个 name(平台方企业名称) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromName(name string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name(平台方企业名称) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromName(names []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromCreditNo 通过单个 credit_no(统一社会信用代码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromCreditNo(creditNo string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCreditNo(creditNo))
	return
}

// GetsFromCreditNo 通过多个 credit_no(统一社会信用代码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromCreditNo(creditNos []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCreditNos(creditNos))
	return
}

// GetFromBusinessLicenseImgURL 通过单个 business_license_img_url(营业执照) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromBusinessLicenseImgURL(businessLicenseImgURL string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessLicenseImgURL(businessLicenseImgURL))
	return
}

// GetsFromBusinessLicenseImgURL 通过多个 business_license_img_url(营业执照) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromBusinessLicenseImgURL(businessLicenseImgURLs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessLicenseImgURLs(businessLicenseImgURLs))
	return
}

// GetFromIndustryCode 通过单个 industry_code(行业代码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromIndustryCode(industryCode string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithIndustryCode(industryCode))
	return
}

// GetsFromIndustryCode 通过多个 industry_code(行业代码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromIndustryCode(industryCodes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithIndustryCodes(industryCodes))
	return
}

// GetFromRegistrationOrg 通过单个 registration_org(登记机关) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromRegistrationOrg(registrationOrg string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegistrationOrg(registrationOrg))
	return
}

// GetsFromRegistrationOrg 通过多个 registration_org(登记机关) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromRegistrationOrg(registrationOrgs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegistrationOrgs(registrationOrgs))
	return
}

// GetFromBuildDate 通过单个 build_date(成立日期) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromBuildDate(buildDate int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBuildDate(buildDate))
	return
}

// GetsFromBuildDate 通过多个 build_date(成立日期) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromBuildDate(buildDates []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBuildDates(buildDates))
	return
}

// GetFromRegisterMoney 通过单个 register_money(注册资本（单位：元，小数点后两位）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromRegisterMoney(registerMoney float64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterMoney(registerMoney))
	return
}

// GetsFromRegisterMoney 通过多个 register_money(注册资本（单位：元，小数点后两位）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromRegisterMoney(registerMoneys []float64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterMoneys(registerMoneys))
	return
}

// GetFromEmployeeNum 通过单个 employee_num(从业人数) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromEmployeeNum(employeeNum int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithEmployeeNum(employeeNum))
	return
}

// GetsFromEmployeeNum 通过多个 employee_num(从业人数) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromEmployeeNum(employeeNums []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithEmployeeNums(employeeNums))
	return
}

// GetFromRegisterAreaCode 通过单个 register_area_code(注册地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromRegisterAreaCode(registerAreaCode string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAreaCode(registerAreaCode))
	return
}

// GetsFromRegisterAreaCode 通过多个 register_area_code(注册地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromRegisterAreaCode(registerAreaCodes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAreaCodes(registerAreaCodes))
	return
}

// GetFromRegisterAddress 通过单个 register_address(注册地址详细信息) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromRegisterAddress(registerAddress string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAddress(registerAddress))
	return
}

// GetsFromRegisterAddress 通过多个 register_address(注册地址详细信息) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromRegisterAddress(registerAddresss []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterAddresss(registerAddresss))
	return
}

// GetFromBusinessAreaCode 通过单个 business_area_code(生产经营地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromBusinessAreaCode(businessAreaCode string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAreaCode(businessAreaCode))
	return
}

// GetsFromBusinessAreaCode 通过多个 business_area_code(生产经营地址行政区划代码（省市区）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromBusinessAreaCode(businessAreaCodes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAreaCodes(businessAreaCodes))
	return
}

// GetFromBusinessAddress 通过单个 business_address(生产经营地址详细信息) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromBusinessAddress(businessAddress string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAddress(businessAddress))
	return
}

// GetsFromBusinessAddress 通过多个 business_address(生产经营地址详细信息) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromBusinessAddress(businessAddresss []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessAddresss(businessAddresss))
	return
}

// GetFromCompanyType 通过单个 company_type(公司类型) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromCompanyType(companyType string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCompanyType(companyType))
	return
}

// GetsFromCompanyType 通过多个 company_type(公司类型) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromCompanyType(companyTypes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCompanyTypes(companyTypes))
	return
}

// GetFromBusinessScope 通过单个 business_scope(经营范围) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromBusinessScope(businessScope string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessScope(businessScope))
	return
}

// GetsFromBusinessScope 通过多个 business_scope(经营范围) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromBusinessScope(businessScopes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithBusinessScopes(businessScopes))
	return
}

// GetFromLegalName 通过单个 legal_name(法人姓名) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromLegalName(legalName string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalName(legalName))
	return
}

// GetsFromLegalName 通过多个 legal_name(法人姓名) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromLegalName(legalNames []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalNames(legalNames))
	return
}

// GetFromLegalCardType 通过单个 legal_card_type(法人身份证件种类) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromLegalCardType(legalCardType string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardType(legalCardType))
	return
}

// GetsFromLegalCardType 通过多个 legal_card_type(法人身份证件种类) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromLegalCardType(legalCardTypes []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardTypes(legalCardTypes))
	return
}

// GetFromLegalCardNo 通过单个 legal_card_no(法人身份证号) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromLegalCardNo(legalCardNo string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardNo(legalCardNo))
	return
}

// GetsFromLegalCardNo 通过多个 legal_card_no(法人身份证号) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromLegalCardNo(legalCardNos []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalCardNos(legalCardNos))
	return
}

// GetFromLegalPhone 通过单个 legal_phone(法人手机号码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromLegalPhone(legalPhone string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalPhone(legalPhone))
	return
}

// GetsFromLegalPhone 通过多个 legal_phone(法人手机号码) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromLegalPhone(legalPhones []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithLegalPhones(legalPhones))
	return
}

// GetFromMail 通过单个 mail(企业邮箱) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromMail(mail string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithMail(mail))
	return
}

// GetsFromMail 通过多个 mail(企业邮箱) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromMail(mails []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithMails(mails))
	return
}

// GetFromPhone 通过单个 phone(企业电话) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromPhone(phone string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithPhone(phone))
	return
}

// GetsFromPhone 通过多个 phone(企业电话) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromPhone(phones []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithPhones(phones))
	return
}

// GetFromContactPersonName 通过单个 contact_person_name(企业联系人姓名) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromContactPersonName(contactPersonName string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonName(contactPersonName))
	return
}

// GetsFromContactPersonName 通过多个 contact_person_name(企业联系人姓名) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromContactPersonName(contactPersonNames []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonNames(contactPersonNames))
	return
}

// GetFromContactPersonPhone 通过单个 contact_person_phone(联系人电话) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromContactPersonPhone(contactPersonPhone string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonPhone(contactPersonPhone))
	return
}

// GetsFromContactPersonPhone 通过多个 contact_person_phone(联系人电话) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromContactPersonPhone(contactPersonPhones []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithContactPersonPhones(contactPersonPhones))
	return
}

// GetFromPlatformURL 通过单个 platform_url(平台网址) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromPlatformURL(platformURL string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithPlatformURL(platformURL))
	return
}

// GetsFromPlatformURL 通过多个 platform_url(平台网址) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromPlatformURL(platformURLs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithPlatformURLs(platformURLs))
	return
}

// GetFromManageTax 通过单个 manage_tax(主管税务局) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromManageTax(manageTax string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithManageTax(manageTax))
	return
}

// GetsFromManageTax 通过多个 manage_tax(主管税务局) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromManageTax(manageTaxs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxs(manageTaxs))
	return
}

// GetFromManageTaxOrg 通过单个 manage_tax_org(主管税务机关（科、分局）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromManageTaxOrg(manageTaxOrg string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxOrg(manageTaxOrg))
	return
}

// GetsFromManageTaxOrg 通过多个 manage_tax_org(主管税务机关（科、分局）) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromManageTaxOrg(manageTaxOrgs []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithManageTaxOrgs(manageTaxOrgs))
	return
}

// GetFromQualificationURL1 通过单个 qualification_url1(资质文件1) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromQualificationURL1(qualificationURL1 string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL1(qualificationURL1))
	return
}

// GetsFromQualificationURL1 通过多个 qualification_url1(资质文件1) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromQualificationURL1(qualificationURL1s []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL1s(qualificationURL1s))
	return
}

// GetFromQualificationURL2 通过单个 qualification_url2(资质文件2) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromQualificationURL2(qualificationURL2 string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL2(qualificationURL2))
	return
}

// GetsFromQualificationURL2 通过多个 qualification_url2(资质文件2) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromQualificationURL2(qualificationURL2s []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL2s(qualificationURL2s))
	return
}

// GetFromQualificationURL3 通过单个 qualification_url3(资质文件3) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromQualificationURL3(qualificationURL3 string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL3(qualificationURL3))
	return
}

// GetsFromQualificationURL3 通过多个 qualification_url3(资质文件3) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromQualificationURL3(qualificationURL3s []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL3s(qualificationURL3s))
	return
}

// GetFromQualificationURL4 通过单个 qualification_url4(资质文件4) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromQualificationURL4(qualificationURL4 string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL4(qualificationURL4))
	return
}

// GetsFromQualificationURL4 通过多个 qualification_url4(资质文件4) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromQualificationURL4(qualificationURL4s []string) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithQualificationURL4s(qualificationURL4s))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromCreatedAt(createdAt int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromCreatedAt(createdAts []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromUpdatedAt(updatedAt int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetFromDeletedAt(deletedAt int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierFirstDAO) GetsFromDeletedAt(deletedAts []int64) (results []*SupplierFirst, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 org_id 字段值，获取单条记录
func (obj *SupplierFirstDAO) FetchByPrimaryKey(orgID int64) (result *SupplierFirst, err error) {
	result, err = obj.GetByOption(obj.WithOrgID(orgID))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
