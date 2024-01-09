package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:34
 * @Desc:   purchaser 表
 */

// 采购商表
type Purchaser struct {
	ID                    int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                             // 主键ID
	OpenOrgIDs            string `gorm:"column:open_org_ids;type:varchar(255);not null;DEFAULT ''" json:"open_org_ids"`                         // 平台企业租户id集合
	PurchaserID           string `gorm:"column:purchaser_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_id"`                          // 企业id（采购商ID）
	RegisterDate          int64  `gorm:"column:register_date;type:int(11);not null;DEFAULT '0'" json:"register_date"`                           // 企业注册时间
	State                 int64  `gorm:"column:state;type:int(11);not null;DEFAULT '0'" json:"state"`                                           // 采购商状态 0-异常 1-正常 2-退出
	StateDesc             string `gorm:"column:state_desc;type:varchar(255);not null;DEFAULT ''" json:"state_desc"`                             // 采购商状态说明
	PurchaserName         string `gorm:"column:purchaser_name;type:varchar(50);not null;DEFAULT ''" json:"purchaser_name"`                      // 企业（采购商）名称
	CreditNo              string `gorm:"column:credit_no;type:varchar(18);not null;DEFAULT ''" json:"credit_no"`                                // 统一社会信用代码
	BusinessLicenseImgURL string `gorm:"column:business_license_img_url;type:varchar(255);not null;DEFAULT ''" json:"business_license_img_url"` // 营业执照
	IndustryCode          string `gorm:"column:industry_code;type:varchar(50);not null;DEFAULT ''" json:"industry_code"`                        // 行业代码
	RegistrationOrg       string `gorm:"column:registration_org;type:varchar(50);not null;DEFAULT ''" json:"registration_org"`                  // 登记机关
	BuildDate             int64  `gorm:"column:build_date;type:int(11);not null;DEFAULT '0'" json:"build_date"`                                 // 成立日期
	RegisterMoney         int64  `gorm:"column:register_money;type:int(11);not null;DEFAULT '0'" json:"register_money"`                         // 注册资本（单位分）
	EmployeeNum           int64  `gorm:"column:employee_num;type:int(11);not null;DEFAULT '0'" json:"employee_num"`                             // 从业人数
	ProcureAmt            int64  `gorm:"column:procure_amt;type:int(11);not null;DEFAULT '0'" json:"procure_amt"`                               // 月均采购预算
	RegisterAreaCode      string `gorm:"column:register_area_code;type:varchar(10);not null;DEFAULT ''" json:"register_area_code"`              // 注册地址行政区划代码（省市区）
	RegisterAddress       string `gorm:"column:register_address;type:varchar(50);not null;DEFAULT ''" json:"register_address"`                  // 注册地址详细信息
	BusinessAreaCode      string `gorm:"column:business_area_code;type:varchar(10);not null;DEFAULT ''" json:"business_area_code"`              // 生产经营地址行政区划代码（省市区）
	BusinessAddress       string `gorm:"column:business_address;type:varchar(50);not null;DEFAULT ''" json:"business_address"`                  // 生产经营地址详细信息
	EntType               int64  `gorm:"column:ent_type;type:int(11);not null;DEFAULT '0'" json:"ent_type"`                                     // 公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社
	BusinessScope         string `gorm:"column:business_scope;type:varchar(255);not null;DEFAULT ''" json:"business_scope"`                     // 经营范围
	LegalName             string `gorm:"column:legal_name;type:varchar(50);not null;DEFAULT ''" json:"legal_name"`                              // 法人姓名
	LegalCardType         string `gorm:"column:legal_card_type;type:varchar(10);not null;DEFAULT ''" json:"legal_card_type"`                    // 法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证
	LegalCardNo           string `gorm:"column:legal_card_no;type:varchar(50);not null;DEFAULT ''" json:"legal_card_no"`                        // 法人身份证号
	LegalPhone            string `gorm:"column:legal_phone;type:varchar(11);not null;DEFAULT ''" json:"legal_phone"`                            // 法人手机号码
	Mail                  string `gorm:"column:mail;type:varchar(50);not null;DEFAULT ''" json:"mail"`                                          // 企业邮箱
	Phone                 string `gorm:"column:phone;type:varchar(50);not null;DEFAULT ''" json:"phone"`                                        // 企业电话
	ContactPersonName     string `gorm:"column:contact_person_name;type:varchar(50);not null;DEFAULT ''" json:"contact_person_name"`            // 企业联系人姓名
	ContactPersonPhone    string `gorm:"column:contact_person_phone;type:varchar(50);not null;DEFAULT ''" json:"contact_person_phone"`          // 联系人电话
	ManageTax             string `gorm:"column:manage_tax;type:varchar(50);not null;DEFAULT ''" json:"manage_tax"`                              // 主管税务局
	ManageTaxOrg          string `gorm:"column:manage_tax_org;type:varchar(50);not null;DEFAULT ''" json:"manage_tax_org"`                      // 主管税务机关（科、分局）
	CleanProtocalFile     int64  `gorm:"column:clean_protocal_file;type:int(11);not null;DEFAULT '0'" json:"clean_protocal_file"`               // 是否清理协议文件 1：是；其他否
	CreatedAt             int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                                 // 创建时间
	UpdatedAt             int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                                 // 更新时间
	DeletedAt             int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                                 // 删除时间
}

// 获取结构体对应的表名方法
func (m *Purchaser) TableName() string {
	return "purchaser"
}

// 实例化结构体对象
func NewPurchaser() *Purchaser {
	return &Purchaser{}
}

// 获取主键的对应字段
func (m *Purchaser) PrimaryKey() string {
	return PurchaserColumns.ID
}

// 获取主键值
func (m *Purchaser) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserColumns = struct {
	ID                    string
	OpenOrgIDs            string
	PurchaserID           string
	RegisterDate          string
	State                 string
	StateDesc             string
	PurchaserName         string
	CreditNo              string
	BusinessLicenseImgURL string
	IndustryCode          string
	RegistrationOrg       string
	BuildDate             string
	RegisterMoney         string
	EmployeeNum           string
	ProcureAmt            string
	RegisterAreaCode      string
	RegisterAddress       string
	BusinessAreaCode      string
	BusinessAddress       string
	EntType               string
	BusinessScope         string
	LegalName             string
	LegalCardType         string
	LegalCardNo           string
	LegalPhone            string
	Mail                  string
	Phone                 string
	ContactPersonName     string
	ContactPersonPhone    string
	ManageTax             string
	ManageTaxOrg          string
	CleanProtocalFile     string
	CreatedAt             string
	UpdatedAt             string
	DeletedAt             string
}{
	ID:                    "id",
	OpenOrgIDs:            "open_org_ids",
	PurchaserID:           "purchaser_id",
	RegisterDate:          "register_date",
	State:                 "state",
	StateDesc:             "state_desc",
	PurchaserName:         "purchaser_name",
	CreditNo:              "credit_no",
	BusinessLicenseImgURL: "business_license_img_url",
	IndustryCode:          "industry_code",
	RegistrationOrg:       "registration_org",
	BuildDate:             "build_date",
	RegisterMoney:         "register_money",
	EmployeeNum:           "employee_num",
	ProcureAmt:            "procure_amt",
	RegisterAreaCode:      "register_area_code",
	RegisterAddress:       "register_address",
	BusinessAreaCode:      "business_area_code",
	BusinessAddress:       "business_address",
	EntType:               "ent_type",
	BusinessScope:         "business_scope",
	LegalName:             "legal_name",
	LegalCardType:         "legal_card_type",
	LegalCardNo:           "legal_card_no",
	LegalPhone:            "legal_phone",
	Mail:                  "mail",
	Phone:                 "phone",
	ContactPersonName:     "contact_person_name",
	ContactPersonPhone:    "contact_person_phone",
	ManageTax:             "manage_tax",
	ManageTaxOrg:          "manage_tax_org",
	CleanProtocalFile:     "clean_protocal_file",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
	DeletedAt:             "deleted_at",
}

// 包含所有表字段的切片
var PurchaserAllColumns = []string{
	PurchaserColumns.ID,                    // 主键ID
	PurchaserColumns.OpenOrgIDs,            // 平台企业租户id集合
	PurchaserColumns.PurchaserID,           // 企业id（采购商ID）
	PurchaserColumns.RegisterDate,          // 企业注册时间
	PurchaserColumns.State,                 // 采购商状态 0-异常 1-正常 2-退出
	PurchaserColumns.StateDesc,             // 采购商状态说明
	PurchaserColumns.PurchaserName,         // 企业（采购商）名称
	PurchaserColumns.CreditNo,              // 统一社会信用代码
	PurchaserColumns.BusinessLicenseImgURL, // 营业执照
	PurchaserColumns.IndustryCode,          // 行业代码
	PurchaserColumns.RegistrationOrg,       // 登记机关
	PurchaserColumns.BuildDate,             // 成立日期
	PurchaserColumns.RegisterMoney,         // 注册资本（单位分）
	PurchaserColumns.EmployeeNum,           // 从业人数
	PurchaserColumns.ProcureAmt,            // 月均采购预算
	PurchaserColumns.RegisterAreaCode,      // 注册地址行政区划代码（省市区）
	PurchaserColumns.RegisterAddress,       // 注册地址详细信息
	PurchaserColumns.BusinessAreaCode,      // 生产经营地址行政区划代码（省市区）
	PurchaserColumns.BusinessAddress,       // 生产经营地址详细信息
	PurchaserColumns.EntType,               // 公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社
	PurchaserColumns.BusinessScope,         // 经营范围
	PurchaserColumns.LegalName,             // 法人姓名
	PurchaserColumns.LegalCardType,         // 法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证
	PurchaserColumns.LegalCardNo,           // 法人身份证号
	PurchaserColumns.LegalPhone,            // 法人手机号码
	PurchaserColumns.Mail,                  // 企业邮箱
	PurchaserColumns.Phone,                 // 企业电话
	PurchaserColumns.ContactPersonName,     // 企业联系人姓名
	PurchaserColumns.ContactPersonPhone,    // 联系人电话
	PurchaserColumns.ManageTax,             // 主管税务局
	PurchaserColumns.ManageTaxOrg,          // 主管税务机关（科、分局）
	PurchaserColumns.CleanProtocalFile,     // 是否清理协议文件 1：是；其他否
	PurchaserColumns.CreatedAt,             // 创建时间
	PurchaserColumns.UpdatedAt,             // 更新时间
	PurchaserColumns.DeletedAt,             // 删除时间

}

// 设置值：主键ID
func (m *Purchaser) SetID(v int64) {
	m.ID = v
}

// 设置值：平台企业租户id集合
func (m *Purchaser) SetOpenOrgIDs(v string) {
	m.OpenOrgIDs = v
}

// 设置值：企业id（采购商ID）
func (m *Purchaser) SetPurchaserID(v string) {
	m.PurchaserID = v
}

// 设置值：企业注册时间
func (m *Purchaser) SetRegisterDate(v int64) {
	m.RegisterDate = v
}

// 设置值：采购商状态 0-异常 1-正常 2-退出
func (m *Purchaser) SetState(v int64) {
	m.State = v
}

// 设置值：采购商状态说明
func (m *Purchaser) SetStateDesc(v string) {
	m.StateDesc = v
}

// 设置值：企业（采购商）名称
func (m *Purchaser) SetPurchaserName(v string) {
	m.PurchaserName = v
}

// 设置值：统一社会信用代码
func (m *Purchaser) SetCreditNo(v string) {
	m.CreditNo = v
}

// 设置值：营业执照
func (m *Purchaser) SetBusinessLicenseImgURL(v string) {
	m.BusinessLicenseImgURL = v
}

// 设置值：行业代码
func (m *Purchaser) SetIndustryCode(v string) {
	m.IndustryCode = v
}

// 设置值：登记机关
func (m *Purchaser) SetRegistrationOrg(v string) {
	m.RegistrationOrg = v
}

// 设置值：成立日期
func (m *Purchaser) SetBuildDate(v int64) {
	m.BuildDate = v
}

// 设置值：注册资本（单位分）
func (m *Purchaser) SetRegisterMoney(v int64) {
	m.RegisterMoney = v
}

// 设置值：从业人数
func (m *Purchaser) SetEmployeeNum(v int64) {
	m.EmployeeNum = v
}

// 设置值：月均采购预算
func (m *Purchaser) SetProcureAmt(v int64) {
	m.ProcureAmt = v
}

// 设置值：注册地址行政区划代码（省市区）
func (m *Purchaser) SetRegisterAreaCode(v string) {
	m.RegisterAreaCode = v
}

// 设置值：注册地址详细信息
func (m *Purchaser) SetRegisterAddress(v string) {
	m.RegisterAddress = v
}

// 设置值：生产经营地址行政区划代码（省市区）
func (m *Purchaser) SetBusinessAreaCode(v string) {
	m.BusinessAreaCode = v
}

// 设置值：生产经营地址详细信息
func (m *Purchaser) SetBusinessAddress(v string) {
	m.BusinessAddress = v
}

// 设置值：公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社
func (m *Purchaser) SetEntType(v int64) {
	m.EntType = v
}

// 设置值：经营范围
func (m *Purchaser) SetBusinessScope(v string) {
	m.BusinessScope = v
}

// 设置值：法人姓名
func (m *Purchaser) SetLegalName(v string) {
	m.LegalName = v
}

// 设置值：法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证
func (m *Purchaser) SetLegalCardType(v string) {
	m.LegalCardType = v
}

// 设置值：法人身份证号
func (m *Purchaser) SetLegalCardNo(v string) {
	m.LegalCardNo = v
}

// 设置值：法人手机号码
func (m *Purchaser) SetLegalPhone(v string) {
	m.LegalPhone = v
}

// 设置值：企业邮箱
func (m *Purchaser) SetMail(v string) {
	m.Mail = v
}

// 设置值：企业电话
func (m *Purchaser) SetPhone(v string) {
	m.Phone = v
}

// 设置值：企业联系人姓名
func (m *Purchaser) SetContactPersonName(v string) {
	m.ContactPersonName = v
}

// 设置值：联系人电话
func (m *Purchaser) SetContactPersonPhone(v string) {
	m.ContactPersonPhone = v
}

// 设置值：主管税务局
func (m *Purchaser) SetManageTax(v string) {
	m.ManageTax = v
}

// 设置值：主管税务机关（科、分局）
func (m *Purchaser) SetManageTaxOrg(v string) {
	m.ManageTaxOrg = v
}

// 设置值：是否清理协议文件 1：是；其他否
func (m *Purchaser) SetCleanProtocalFile(v int64) {
	m.CleanProtocalFile = v
}

// 设置值：创建时间
func (m *Purchaser) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *Purchaser) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *Purchaser) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID
func (m *Purchaser) GetID() int64 {
	return m.ID
}

// 获取值：平台企业租户id集合
func (m *Purchaser) GetOpenOrgIDs() string {
	return m.OpenOrgIDs
}

// 获取值：企业id（采购商ID）
func (m *Purchaser) GetPurchaserID() string {
	return m.PurchaserID
}

// 获取值：企业注册时间
func (m *Purchaser) GetRegisterDate() int64 {
	return m.RegisterDate
}

// 获取值：采购商状态 0-异常 1-正常 2-退出
func (m *Purchaser) GetState() int64 {
	return m.State
}

// 获取值：采购商状态说明
func (m *Purchaser) GetStateDesc() string {
	return m.StateDesc
}

// 获取值：企业（采购商）名称
func (m *Purchaser) GetPurchaserName() string {
	return m.PurchaserName
}

// 获取值：统一社会信用代码
func (m *Purchaser) GetCreditNo() string {
	return m.CreditNo
}

// 获取值：营业执照
func (m *Purchaser) GetBusinessLicenseImgURL() string {
	return m.BusinessLicenseImgURL
}

// 获取值：行业代码
func (m *Purchaser) GetIndustryCode() string {
	return m.IndustryCode
}

// 获取值：登记机关
func (m *Purchaser) GetRegistrationOrg() string {
	return m.RegistrationOrg
}

// 获取值：成立日期
func (m *Purchaser) GetBuildDate() int64 {
	return m.BuildDate
}

// 获取值：注册资本（单位分）
func (m *Purchaser) GetRegisterMoney() int64 {
	return m.RegisterMoney
}

// 获取值：从业人数
func (m *Purchaser) GetEmployeeNum() int64 {
	return m.EmployeeNum
}

// 获取值：月均采购预算
func (m *Purchaser) GetProcureAmt() int64 {
	return m.ProcureAmt
}

// 获取值：注册地址行政区划代码（省市区）
func (m *Purchaser) GetRegisterAreaCode() string {
	return m.RegisterAreaCode
}

// 获取值：注册地址详细信息
func (m *Purchaser) GetRegisterAddress() string {
	return m.RegisterAddress
}

// 获取值：生产经营地址行政区划代码（省市区）
func (m *Purchaser) GetBusinessAreaCode() string {
	return m.BusinessAreaCode
}

// 获取值：生产经营地址详细信息
func (m *Purchaser) GetBusinessAddress() string {
	return m.BusinessAddress
}

// 获取值：公司类型 0 有限责任公司 1 股份有限公司 2 个人独资企业 3 合伙企业 4 全民所有制企业 5 集体所有制企业 6 农民专业合作社
func (m *Purchaser) GetEntType() int64 {
	return m.EntType
}

// 获取值：经营范围
func (m *Purchaser) GetBusinessScope() string {
	return m.BusinessScope
}

// 获取值：法人姓名
func (m *Purchaser) GetLegalName() string {
	return m.LegalName
}

// 获取值：法人身份证件种类 0 居民身份证 1 护照 2 港澳居民居住证 3 台湾居民居住证 4 临时居民身份证 5 香港、澳门身份证 6 军官证 7 警官证 8 士兵证 9 学生证 10 教师证 11 驾驶证 12 医保卡 13 社保卡 14 残疾证
func (m *Purchaser) GetLegalCardType() string {
	return m.LegalCardType
}

// 获取值：法人身份证号
func (m *Purchaser) GetLegalCardNo() string {
	return m.LegalCardNo
}

// 获取值：法人手机号码
func (m *Purchaser) GetLegalPhone() string {
	return m.LegalPhone
}

// 获取值：企业邮箱
func (m *Purchaser) GetMail() string {
	return m.Mail
}

// 获取值：企业电话
func (m *Purchaser) GetPhone() string {
	return m.Phone
}

// 获取值：企业联系人姓名
func (m *Purchaser) GetContactPersonName() string {
	return m.ContactPersonName
}

// 获取值：联系人电话
func (m *Purchaser) GetContactPersonPhone() string {
	return m.ContactPersonPhone
}

// 获取值：主管税务局
func (m *Purchaser) GetManageTax() string {
	return m.ManageTax
}

// 获取值：主管税务机关（科、分局）
func (m *Purchaser) GetManageTaxOrg() string {
	return m.ManageTaxOrg
}

// 获取值：是否清理协议文件 1：是；其他否
func (m *Purchaser) GetCleanProtocalFile() int64 {
	return m.CleanProtocalFile
}

// 获取值：创建时间
func (m *Purchaser) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *Purchaser) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *Purchaser) GetDeletedAt() int64 {
	return m.DeletedAt
}
