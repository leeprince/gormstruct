package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 16:48:53
 * @Desc:   supplier_first 表
 */

// 一级供应商（平台企业）表
type SupplierFirst struct {
	OrgID                 int64   `gorm:"column:org_id;primaryKey;type:int(11);" json:"org_id"`                                                  // 一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入
	OpenOrgID             string  `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"`                            // 一级供应商高灯云开放租户ID
	Name                  string  `gorm:"column:name;type:varchar(50);not null;DEFAULT ''" json:"name"`                                          // 平台方企业名称
	CreditNo              string  `gorm:"column:credit_no;type:varchar(18);not null;DEFAULT ''" json:"credit_no"`                                // 统一社会信用代码
	BusinessLicenseImgURL string  `gorm:"column:business_license_img_url;type:varchar(255);not null;DEFAULT ''" json:"business_license_img_url"` // 营业执照
	IndustryCode          string  `gorm:"column:industry_code;type:varchar(10);not null;DEFAULT ''" json:"industry_code"`                        // 行业代码
	RegistrationOrg       string  `gorm:"column:registration_org;type:varchar(50);not null;DEFAULT ''" json:"registration_org"`                  // 登记机关
	BuildDate             int64   `gorm:"column:build_date;type:int(11);not null;DEFAULT '0'" json:"build_date"`                                 // 成立日期
	RegisterMoney         float64 `gorm:"column:register_money;type:decimal(11,2);not null;DEFAULT '0.00'" json:"register_money"`                // 注册资本（单位：元，小数点后两位）
	EmployeeNum           int64   `gorm:"column:employee_num;type:int(11);not null;DEFAULT '0'" json:"employee_num"`                             // 从业人数
	RegisterAreaCode      string  `gorm:"column:register_area_code;type:varchar(10);not null;DEFAULT ''" json:"register_area_code"`              // 注册地址行政区划代码（省市区）
	RegisterAddress       string  `gorm:"column:register_address;type:varchar(50);not null;DEFAULT ''" json:"register_address"`                  // 注册地址详细信息
	BusinessAreaCode      string  `gorm:"column:business_area_code;type:varchar(10);not null;DEFAULT ''" json:"business_area_code"`              // 生产经营地址行政区划代码（省市区）
	BusinessAddress       string  `gorm:"column:business_address;type:varchar(50);not null;DEFAULT ''" json:"business_address"`                  // 生产经营地址详细信息
	CompanyType           string  `gorm:"column:company_type;type:varchar(10);not null;DEFAULT ''" json:"company_type"`                          // 公司类型
	BusinessScope         string  `gorm:"column:business_scope;type:varchar(255);not null;DEFAULT ''" json:"business_scope"`                     // 经营范围
	LegalName             string  `gorm:"column:legal_name;type:varchar(50);not null;DEFAULT ''" json:"legal_name"`                              // 法人姓名
	LegalCardType         string  `gorm:"column:legal_card_type;type:varchar(10);not null;DEFAULT ''" json:"legal_card_type"`                    // 法人身份证件种类
	LegalCardNo           string  `gorm:"column:legal_card_no;type:varchar(50);not null;DEFAULT ''" json:"legal_card_no"`                        // 法人身份证号
	LegalPhone            string  `gorm:"column:legal_phone;type:varchar(11);not null;DEFAULT ''" json:"legal_phone"`                            // 法人手机号码
	Mail                  string  `gorm:"column:mail;type:varchar(50);not null;DEFAULT ''" json:"mail"`                                          // 企业邮箱
	Phone                 string  `gorm:"column:phone;type:varchar(50);not null;DEFAULT ''" json:"phone"`                                        // 企业电话
	ContactPersonName     string  `gorm:"column:contact_person_name;type:varchar(50);not null;DEFAULT ''" json:"contact_person_name"`            // 企业联系人姓名
	ContactPersonPhone    string  `gorm:"column:contact_person_phone;type:varchar(50);not null;DEFAULT ''" json:"contact_person_phone"`          // 联系人电话
	PlatformURL           string  `gorm:"column:platform_url;type:varchar(255);not null;DEFAULT ''" json:"platform_url"`                         // 平台网址
	ManageTax             string  `gorm:"column:manage_tax;type:varchar(50);not null;DEFAULT ''" json:"manage_tax"`                              // 主管税务局
	ManageTaxOrg          string  `gorm:"column:manage_tax_org;type:varchar(50);not null;DEFAULT ''" json:"manage_tax_org"`                      // 主管税务机关（科、分局）
	QualificationURL1     string  `gorm:"column:qualification_url1;type:varchar(255);not null;DEFAULT ''" json:"qualification_url1"`             // 资质文件1
	QualificationURL2     string  `gorm:"column:qualification_url2;type:varchar(255);not null;DEFAULT ''" json:"qualification_url2"`             // 资质文件2
	QualificationURL3     string  `gorm:"column:qualification_url3;type:varchar(255);not null;DEFAULT ''" json:"qualification_url3"`             // 资质文件3
	QualificationURL4     string  `gorm:"column:qualification_url4;type:varchar(255);not null;DEFAULT ''" json:"qualification_url4"`             // 资质文件4
	CreatedAt             int64   `gorm:"column:created_at;type:int(11);not null;" json:"created_at"`                                            // 创建时间
	UpdatedAt             int64   `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`                                            // 更新时间
	DeletedAt             int64   `gorm:"column:deleted_at;type:int(11);not null;" json:"deleted_at"`                                            // 删除时间
}

// 获取结构体对应的表名方法
func (m *SupplierFirst) TableName() string {
	return "supplier_first"
}

// 实例化结构体对象
func NewSupplierFirst() *SupplierFirst {
	return &SupplierFirst{}
}

// 获取主键的对应字段
func (m *SupplierFirst) PrimaryKey() string {
	return SupplierFirstColumns.OrgID
}

// 获取主键值
func (m *SupplierFirst) PrimaryKeyValue() int64 {
	return m.OrgID
}

// 表字段的映射
var SupplierFirstColumns = struct {
	OrgID                 string
	OpenOrgID             string
	Name                  string
	CreditNo              string
	BusinessLicenseImgURL string
	IndustryCode          string
	RegistrationOrg       string
	BuildDate             string
	RegisterMoney         string
	EmployeeNum           string
	RegisterAreaCode      string
	RegisterAddress       string
	BusinessAreaCode      string
	BusinessAddress       string
	CompanyType           string
	BusinessScope         string
	LegalName             string
	LegalCardType         string
	LegalCardNo           string
	LegalPhone            string
	Mail                  string
	Phone                 string
	ContactPersonName     string
	ContactPersonPhone    string
	PlatformURL           string
	ManageTax             string
	ManageTaxOrg          string
	QualificationURL1     string
	QualificationURL2     string
	QualificationURL3     string
	QualificationURL4     string
	CreatedAt             string
	UpdatedAt             string
	DeletedAt             string
}{
	OrgID:                 "org_id",
	OpenOrgID:             "open_org_id",
	Name:                  "name",
	CreditNo:              "credit_no",
	BusinessLicenseImgURL: "business_license_img_url",
	IndustryCode:          "industry_code",
	RegistrationOrg:       "registration_org",
	BuildDate:             "build_date",
	RegisterMoney:         "register_money",
	EmployeeNum:           "employee_num",
	RegisterAreaCode:      "register_area_code",
	RegisterAddress:       "register_address",
	BusinessAreaCode:      "business_area_code",
	BusinessAddress:       "business_address",
	CompanyType:           "company_type",
	BusinessScope:         "business_scope",
	LegalName:             "legal_name",
	LegalCardType:         "legal_card_type",
	LegalCardNo:           "legal_card_no",
	LegalPhone:            "legal_phone",
	Mail:                  "mail",
	Phone:                 "phone",
	ContactPersonName:     "contact_person_name",
	ContactPersonPhone:    "contact_person_phone",
	PlatformURL:           "platform_url",
	ManageTax:             "manage_tax",
	ManageTaxOrg:          "manage_tax_org",
	QualificationURL1:     "qualification_url1",
	QualificationURL2:     "qualification_url2",
	QualificationURL3:     "qualification_url3",
	QualificationURL4:     "qualification_url4",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	DeletedAt:             "deleted_at",
}

// 包含所有表字段的切片
var SupplierFirstAllColumns = []string{
	SupplierFirstColumns.OrgID,                 // 一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入
	SupplierFirstColumns.OpenOrgID,             // 一级供应商高灯云开放租户ID
	SupplierFirstColumns.Name,                  // 平台方企业名称
	SupplierFirstColumns.CreditNo,              // 统一社会信用代码
	SupplierFirstColumns.BusinessLicenseImgURL, // 营业执照
	SupplierFirstColumns.IndustryCode,          // 行业代码
	SupplierFirstColumns.RegistrationOrg,       // 登记机关
	SupplierFirstColumns.BuildDate,             // 成立日期
	SupplierFirstColumns.RegisterMoney,         // 注册资本（单位：元，小数点后两位）
	SupplierFirstColumns.EmployeeNum,           // 从业人数
	SupplierFirstColumns.RegisterAreaCode,      // 注册地址行政区划代码（省市区）
	SupplierFirstColumns.RegisterAddress,       // 注册地址详细信息
	SupplierFirstColumns.BusinessAreaCode,      // 生产经营地址行政区划代码（省市区）
	SupplierFirstColumns.BusinessAddress,       // 生产经营地址详细信息
	SupplierFirstColumns.CompanyType,           // 公司类型
	SupplierFirstColumns.BusinessScope,         // 经营范围
	SupplierFirstColumns.LegalName,             // 法人姓名
	SupplierFirstColumns.LegalCardType,         // 法人身份证件种类
	SupplierFirstColumns.LegalCardNo,           // 法人身份证号
	SupplierFirstColumns.LegalPhone,            // 法人手机号码
	SupplierFirstColumns.Mail,                  // 企业邮箱
	SupplierFirstColumns.Phone,                 // 企业电话
	SupplierFirstColumns.ContactPersonName,     // 企业联系人姓名
	SupplierFirstColumns.ContactPersonPhone,    // 联系人电话
	SupplierFirstColumns.PlatformURL,           // 平台网址
	SupplierFirstColumns.ManageTax,             // 主管税务局
	SupplierFirstColumns.ManageTaxOrg,          // 主管税务机关（科、分局）
	SupplierFirstColumns.QualificationURL1,     // 资质文件1
	SupplierFirstColumns.QualificationURL2,     // 资质文件2
	SupplierFirstColumns.QualificationURL3,     // 资质文件3
	SupplierFirstColumns.QualificationURL4,     // 资质文件4
	SupplierFirstColumns.CreatedAt,             // 创建时间
	SupplierFirstColumns.UpdatedAt,             // 更新时间
	SupplierFirstColumns.DeletedAt,             // 删除时间

}

// 设置值：一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入
func (m *SupplierFirst) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：一级供应商高灯云开放租户ID
func (m *SupplierFirst) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：平台方企业名称
func (m *SupplierFirst) SetName(v string) {
	m.Name = v
}

// 设置值：统一社会信用代码
func (m *SupplierFirst) SetCreditNo(v string) {
	m.CreditNo = v
}

// 设置值：营业执照
func (m *SupplierFirst) SetBusinessLicenseImgURL(v string) {
	m.BusinessLicenseImgURL = v
}

// 设置值：行业代码
func (m *SupplierFirst) SetIndustryCode(v string) {
	m.IndustryCode = v
}

// 设置值：登记机关
func (m *SupplierFirst) SetRegistrationOrg(v string) {
	m.RegistrationOrg = v
}

// 设置值：成立日期
func (m *SupplierFirst) SetBuildDate(v int64) {
	m.BuildDate = v
}

// 设置值：注册资本（单位：元，小数点后两位）
func (m *SupplierFirst) SetRegisterMoney(v float64) {
	m.RegisterMoney = v
}

// 设置值：从业人数
func (m *SupplierFirst) SetEmployeeNum(v int64) {
	m.EmployeeNum = v
}

// 设置值：注册地址行政区划代码（省市区）
func (m *SupplierFirst) SetRegisterAreaCode(v string) {
	m.RegisterAreaCode = v
}

// 设置值：注册地址详细信息
func (m *SupplierFirst) SetRegisterAddress(v string) {
	m.RegisterAddress = v
}

// 设置值：生产经营地址行政区划代码（省市区）
func (m *SupplierFirst) SetBusinessAreaCode(v string) {
	m.BusinessAreaCode = v
}

// 设置值：生产经营地址详细信息
func (m *SupplierFirst) SetBusinessAddress(v string) {
	m.BusinessAddress = v
}

// 设置值：公司类型
func (m *SupplierFirst) SetCompanyType(v string) {
	m.CompanyType = v
}

// 设置值：经营范围
func (m *SupplierFirst) SetBusinessScope(v string) {
	m.BusinessScope = v
}

// 设置值：法人姓名
func (m *SupplierFirst) SetLegalName(v string) {
	m.LegalName = v
}

// 设置值：法人身份证件种类
func (m *SupplierFirst) SetLegalCardType(v string) {
	m.LegalCardType = v
}

// 设置值：法人身份证号
func (m *SupplierFirst) SetLegalCardNo(v string) {
	m.LegalCardNo = v
}

// 设置值：法人手机号码
func (m *SupplierFirst) SetLegalPhone(v string) {
	m.LegalPhone = v
}

// 设置值：企业邮箱
func (m *SupplierFirst) SetMail(v string) {
	m.Mail = v
}

// 设置值：企业电话
func (m *SupplierFirst) SetPhone(v string) {
	m.Phone = v
}

// 设置值：企业联系人姓名
func (m *SupplierFirst) SetContactPersonName(v string) {
	m.ContactPersonName = v
}

// 设置值：联系人电话
func (m *SupplierFirst) SetContactPersonPhone(v string) {
	m.ContactPersonPhone = v
}

// 设置值：平台网址
func (m *SupplierFirst) SetPlatformURL(v string) {
	m.PlatformURL = v
}

// 设置值：主管税务局
func (m *SupplierFirst) SetManageTax(v string) {
	m.ManageTax = v
}

// 设置值：主管税务机关（科、分局）
func (m *SupplierFirst) SetManageTaxOrg(v string) {
	m.ManageTaxOrg = v
}

// 设置值：资质文件1
func (m *SupplierFirst) SetQualificationURL1(v string) {
	m.QualificationURL1 = v
}

// 设置值：资质文件2
func (m *SupplierFirst) SetQualificationURL2(v string) {
	m.QualificationURL2 = v
}

// 设置值：资质文件3
func (m *SupplierFirst) SetQualificationURL3(v string) {
	m.QualificationURL3 = v
}

// 设置值：资质文件4
func (m *SupplierFirst) SetQualificationURL4(v string) {
	m.QualificationURL4 = v
}

// 设置值：创建时间
func (m *SupplierFirst) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *SupplierFirst) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *SupplierFirst) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：一级供应商高灯云租户ID。完成高灯云入驻后获得，主键-按照它来更新或插入
func (m *SupplierFirst) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：一级供应商高灯云开放租户ID
func (m *SupplierFirst) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：平台方企业名称
func (m *SupplierFirst) GetName() string {
	return m.Name
}

// 获取值：统一社会信用代码
func (m *SupplierFirst) GetCreditNo() string {
	return m.CreditNo
}

// 获取值：营业执照
func (m *SupplierFirst) GetBusinessLicenseImgURL() string {
	return m.BusinessLicenseImgURL
}

// 获取值：行业代码
func (m *SupplierFirst) GetIndustryCode() string {
	return m.IndustryCode
}

// 获取值：登记机关
func (m *SupplierFirst) GetRegistrationOrg() string {
	return m.RegistrationOrg
}

// 获取值：成立日期
func (m *SupplierFirst) GetBuildDate() int64 {
	return m.BuildDate
}

// 获取值：注册资本（单位：元，小数点后两位）
func (m *SupplierFirst) GetRegisterMoney() float64 {
	return m.RegisterMoney
}

// 获取值：从业人数
func (m *SupplierFirst) GetEmployeeNum() int64 {
	return m.EmployeeNum
}

// 获取值：注册地址行政区划代码（省市区）
func (m *SupplierFirst) GetRegisterAreaCode() string {
	return m.RegisterAreaCode
}

// 获取值：注册地址详细信息
func (m *SupplierFirst) GetRegisterAddress() string {
	return m.RegisterAddress
}

// 获取值：生产经营地址行政区划代码（省市区）
func (m *SupplierFirst) GetBusinessAreaCode() string {
	return m.BusinessAreaCode
}

// 获取值：生产经营地址详细信息
func (m *SupplierFirst) GetBusinessAddress() string {
	return m.BusinessAddress
}

// 获取值：公司类型
func (m *SupplierFirst) GetCompanyType() string {
	return m.CompanyType
}

// 获取值：经营范围
func (m *SupplierFirst) GetBusinessScope() string {
	return m.BusinessScope
}

// 获取值：法人姓名
func (m *SupplierFirst) GetLegalName() string {
	return m.LegalName
}

// 获取值：法人身份证件种类
func (m *SupplierFirst) GetLegalCardType() string {
	return m.LegalCardType
}

// 获取值：法人身份证号
func (m *SupplierFirst) GetLegalCardNo() string {
	return m.LegalCardNo
}

// 获取值：法人手机号码
func (m *SupplierFirst) GetLegalPhone() string {
	return m.LegalPhone
}

// 获取值：企业邮箱
func (m *SupplierFirst) GetMail() string {
	return m.Mail
}

// 获取值：企业电话
func (m *SupplierFirst) GetPhone() string {
	return m.Phone
}

// 获取值：企业联系人姓名
func (m *SupplierFirst) GetContactPersonName() string {
	return m.ContactPersonName
}

// 获取值：联系人电话
func (m *SupplierFirst) GetContactPersonPhone() string {
	return m.ContactPersonPhone
}

// 获取值：平台网址
func (m *SupplierFirst) GetPlatformURL() string {
	return m.PlatformURL
}

// 获取值：主管税务局
func (m *SupplierFirst) GetManageTax() string {
	return m.ManageTax
}

// 获取值：主管税务机关（科、分局）
func (m *SupplierFirst) GetManageTaxOrg() string {
	return m.ManageTaxOrg
}

// 获取值：资质文件1
func (m *SupplierFirst) GetQualificationURL1() string {
	return m.QualificationURL1
}

// 获取值：资质文件2
func (m *SupplierFirst) GetQualificationURL2() string {
	return m.QualificationURL2
}

// 获取值：资质文件3
func (m *SupplierFirst) GetQualificationURL3() string {
	return m.QualificationURL3
}

// 获取值：资质文件4
func (m *SupplierFirst) GetQualificationURL4() string {
	return m.QualificationURL4
}

// 获取值：创建时间
func (m *SupplierFirst) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *SupplierFirst) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *SupplierFirst) GetDeletedAt() int64 {
	return m.DeletedAt
}
