package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:47:03
 * @Desc:   supplier_second 表
 */

// 二级供应商信息表
type SupplierSecond struct {
	ID                 int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                 // 主键id
	OpenOrgIDs         string `gorm:"column:open_org_ids;type:varchar(255);not null;DEFAULT ''" json:"open_org_ids"`             // 一级供应商（平台企业）租户id集合
	SupplierID         string `gorm:"column:supplier_id;type:varchar(50);not null;DEFAULT ''" json:"supplier_id"`                // 供应商在平台侧的唯一标识 主键-按照它来更新或插入
	RegisterDate       int64  `gorm:"column:register_date;type:int(11);not null;DEFAULT '0'" json:"register_date"`               // 供应商注册时间
	State              int    `gorm:"column:state;type:int(1);not null;DEFAULT '0'" json:"state"`                                // 供应商状态：0-异常 1-正常 2-退出
	StateDesc          string `gorm:"column:state_desc;type:varchar(255);not null;DEFAULT ''" json:"state_desc"`                 // 供应商状态说明
	Name               string `gorm:"column:name;type:varchar(50);not null;DEFAULT ''" json:"name"`                              // 供应商姓名
	CardType           string `gorm:"column:card_type;type:varchar(10);not null;DEFAULT ''" json:"card_type"`                    // 供应商身份证件种类
	CardNo             string `gorm:"column:card_no;type:varchar(50);not null;DEFAULT ''" json:"card_no"`                        // 身份证号
	Phone              string `gorm:"column:phone;type:varchar(11);not null;DEFAULT ''" json:"phone"`                            // 手机号码
	IDentityType       string `gorm:"column:identity_type;type:varchar(10);not null;DEFAULT ''" json:"identity_type"`            // 供应商业务身份标识
	CardFaceURL        string `gorm:"column:card_face_url;type:varchar(255);not null;DEFAULT ''" json:"card_face_url"`           // 身份证人像面
	CardEmblemURL      string `gorm:"column:card_emblem_url;type:varchar(255);not null;DEFAULT ''" json:"card_emblem_url"`       // 身份证国徽面
	IsIDentity         int    `gorm:"column:is_identity;type:int(1);not null;DEFAULT '0'" json:"is_identity"`                    // 是否完成实名认证：0-否 1-是
	IDentityDate       int64  `gorm:"column:identity_date;type:int(11);not null;DEFAULT '0'" json:"identity_date"`               // 实名认证时间
	IDentityState      int    `gorm:"column:identity_state;type:int(1);not null;DEFAULT '0'" json:"identity_state"`              // 实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态
	IDentityDesc       string `gorm:"column:identity_desc;type:varchar(50);not null;DEFAULT ''" json:"identity_desc"`            // 实名认证结果说明
	IDentityUpdateTime int64  `gorm:"column:identity_update_time;type:int(11);not null;DEFAULT '0'" json:"identity_update_time"` // 实名认证更新时间，时间戳
	CleanProtocalFile  int    `gorm:"column:clean_protocal_file;type:int(1);not null;DEFAULT '0'" json:"clean_protocal_file"`    // 是否清理协议文件：1：是 非1：否
	CleanAccounts      int    `gorm:"column:clean_accounts;type:int(1);not null;DEFAULT '0'" json:"clean_accounts"`              // 是否清理账户信息：1：是 非1：否
	CreatedAt          int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                     // 创建时间
	UpdatedAt          int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                     // 更新时间
	DeletedAt          int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                     // 删除时间
}

// 获取结构体对应的表名方法
func (m *SupplierSecond) TableName() string {
	return "supplier_second"
}

// 实例化结构体对象
func NewSupplierSecond() *SupplierSecond {
	return &SupplierSecond{}
}

// 获取主键的对应字段
func (m *SupplierSecond) PrimaryKey() string {
	return SupplierSecondColumns.ID
}

// 获取主键值
func (m *SupplierSecond) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var SupplierSecondColumns = struct {
	ID                 string
	OpenOrgIDs         string
	SupplierID         string
	RegisterDate       string
	State              string
	StateDesc          string
	Name               string
	CardType           string
	CardNo             string
	Phone              string
	IDentityType       string
	CardFaceURL        string
	CardEmblemURL      string
	IsIDentity         string
	IDentityDate       string
	IDentityState      string
	IDentityDesc       string
	IDentityUpdateTime string
	CleanProtocalFile  string
	CleanAccounts      string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
}{
	ID:                 "id",
	OpenOrgIDs:         "open_org_ids",
	SupplierID:         "supplier_id",
	RegisterDate:       "register_date",
	State:              "state",
	StateDesc:          "state_desc",
	Name:               "name",
	CardType:           "card_type",
	CardNo:             "card_no",
	Phone:              "phone",
	IDentityType:       "identity_type",
	CardFaceURL:        "card_face_url",
	CardEmblemURL:      "card_emblem_url",
	IsIDentity:         "is_identity",
	IDentityDate:       "identity_date",
	IDentityState:      "identity_state",
	IDentityDesc:       "identity_desc",
	IDentityUpdateTime: "identity_update_time",
	CleanProtocalFile:  "clean_protocal_file",
	CleanAccounts:      "clean_accounts",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// 包含所有表字段的切片
var SupplierSecondAllColumns = []string{
	SupplierSecondColumns.ID,                 // 主键id
	SupplierSecondColumns.OpenOrgIDs,         // 一级供应商（平台企业）租户id集合
	SupplierSecondColumns.SupplierID,         // 供应商在平台侧的唯一标识 主键-按照它来更新或插入
	SupplierSecondColumns.RegisterDate,       // 供应商注册时间
	SupplierSecondColumns.State,              // 供应商状态：0-异常 1-正常 2-退出
	SupplierSecondColumns.StateDesc,          // 供应商状态说明
	SupplierSecondColumns.Name,               // 供应商姓名
	SupplierSecondColumns.CardType,           // 供应商身份证件种类
	SupplierSecondColumns.CardNo,             // 身份证号
	SupplierSecondColumns.Phone,              // 手机号码
	SupplierSecondColumns.IDentityType,       // 供应商业务身份标识
	SupplierSecondColumns.CardFaceURL,        // 身份证人像面
	SupplierSecondColumns.CardEmblemURL,      // 身份证国徽面
	SupplierSecondColumns.IsIDentity,         // 是否完成实名认证：0-否 1-是
	SupplierSecondColumns.IDentityDate,       // 实名认证时间
	SupplierSecondColumns.IDentityState,      // 实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态
	SupplierSecondColumns.IDentityDesc,       // 实名认证结果说明
	SupplierSecondColumns.IDentityUpdateTime, // 实名认证更新时间，时间戳
	SupplierSecondColumns.CleanProtocalFile,  // 是否清理协议文件：1：是 非1：否
	SupplierSecondColumns.CleanAccounts,      // 是否清理账户信息：1：是 非1：否
	SupplierSecondColumns.CreatedAt,          // 创建时间
	SupplierSecondColumns.UpdatedAt,          // 更新时间
	SupplierSecondColumns.DeletedAt,          // 删除时间

}

// 设置值：主键id
func (m *SupplierSecond) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id集合
func (m *SupplierSecond) SetOpenOrgIDs(v string) {
	m.OpenOrgIDs = v
}

// 设置值：供应商在平台侧的唯一标识 主键-按照它来更新或插入
func (m *SupplierSecond) SetSupplierID(v string) {
	m.SupplierID = v
}

// 设置值：供应商注册时间
func (m *SupplierSecond) SetRegisterDate(v int64) {
	m.RegisterDate = v
}

// 设置值：供应商状态：0-异常 1-正常 2-退出
func (m *SupplierSecond) SetState(v int) {
	m.State = v
}

// 设置值：供应商状态说明
func (m *SupplierSecond) SetStateDesc(v string) {
	m.StateDesc = v
}

// 设置值：供应商姓名
func (m *SupplierSecond) SetName(v string) {
	m.Name = v
}

// 设置值：供应商身份证件种类
func (m *SupplierSecond) SetCardType(v string) {
	m.CardType = v
}

// 设置值：身份证号
func (m *SupplierSecond) SetCardNo(v string) {
	m.CardNo = v
}

// 设置值：手机号码
func (m *SupplierSecond) SetPhone(v string) {
	m.Phone = v
}

// 设置值：供应商业务身份标识
func (m *SupplierSecond) SetIDentityType(v string) {
	m.IDentityType = v
}

// 设置值：身份证人像面
func (m *SupplierSecond) SetCardFaceURL(v string) {
	m.CardFaceURL = v
}

// 设置值：身份证国徽面
func (m *SupplierSecond) SetCardEmblemURL(v string) {
	m.CardEmblemURL = v
}

// 设置值：是否完成实名认证：0-否 1-是
func (m *SupplierSecond) SetIsIDentity(v int) {
	m.IsIDentity = v
}

// 设置值：实名认证时间
func (m *SupplierSecond) SetIDentityDate(v int64) {
	m.IDentityDate = v
}

// 设置值：实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态
func (m *SupplierSecond) SetIDentityState(v int) {
	m.IDentityState = v
}

// 设置值：实名认证结果说明
func (m *SupplierSecond) SetIDentityDesc(v string) {
	m.IDentityDesc = v
}

// 设置值：实名认证更新时间，时间戳
func (m *SupplierSecond) SetIDentityUpdateTime(v int64) {
	m.IDentityUpdateTime = v
}

// 设置值：是否清理协议文件：1：是 非1：否
func (m *SupplierSecond) SetCleanProtocalFile(v int) {
	m.CleanProtocalFile = v
}

// 设置值：是否清理账户信息：1：是 非1：否
func (m *SupplierSecond) SetCleanAccounts(v int) {
	m.CleanAccounts = v
}

// 设置值：创建时间
func (m *SupplierSecond) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *SupplierSecond) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *SupplierSecond) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *SupplierSecond) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id集合
func (m *SupplierSecond) GetOpenOrgIDs() string {
	return m.OpenOrgIDs
}

// 获取值：供应商在平台侧的唯一标识 主键-按照它来更新或插入
func (m *SupplierSecond) GetSupplierID() string {
	return m.SupplierID
}

// 获取值：供应商注册时间
func (m *SupplierSecond) GetRegisterDate() int64 {
	return m.RegisterDate
}

// 获取值：供应商状态：0-异常 1-正常 2-退出
func (m *SupplierSecond) GetState() int {
	return m.State
}

// 获取值：供应商状态说明
func (m *SupplierSecond) GetStateDesc() string {
	return m.StateDesc
}

// 获取值：供应商姓名
func (m *SupplierSecond) GetName() string {
	return m.Name
}

// 获取值：供应商身份证件种类
func (m *SupplierSecond) GetCardType() string {
	return m.CardType
}

// 获取值：身份证号
func (m *SupplierSecond) GetCardNo() string {
	return m.CardNo
}

// 获取值：手机号码
func (m *SupplierSecond) GetPhone() string {
	return m.Phone
}

// 获取值：供应商业务身份标识
func (m *SupplierSecond) GetIDentityType() string {
	return m.IDentityType
}

// 获取值：身份证人像面
func (m *SupplierSecond) GetCardFaceURL() string {
	return m.CardFaceURL
}

// 获取值：身份证国徽面
func (m *SupplierSecond) GetCardEmblemURL() string {
	return m.CardEmblemURL
}

// 获取值：是否完成实名认证：0-否 1-是
func (m *SupplierSecond) GetIsIDentity() int {
	return m.IsIDentity
}

// 获取值：实名认证时间
func (m *SupplierSecond) GetIDentityDate() int64 {
	return m.IDentityDate
}

// 获取值：实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态
func (m *SupplierSecond) GetIDentityState() int {
	return m.IDentityState
}

// 获取值：实名认证结果说明
func (m *SupplierSecond) GetIDentityDesc() string {
	return m.IDentityDesc
}

// 获取值：实名认证更新时间，时间戳
func (m *SupplierSecond) GetIDentityUpdateTime() int64 {
	return m.IDentityUpdateTime
}

// 获取值：是否清理协议文件：1：是 非1：否
func (m *SupplierSecond) GetCleanProtocalFile() int {
	return m.CleanProtocalFile
}

// 获取值：是否清理账户信息：1：是 非1：否
func (m *SupplierSecond) GetCleanAccounts() int {
	return m.CleanAccounts
}

// 获取值：创建时间
func (m *SupplierSecond) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *SupplierSecond) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *SupplierSecond) GetDeletedAt() int64 {
	return m.DeletedAt
}
