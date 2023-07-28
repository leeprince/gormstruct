package model

// 发票导入任务表
type ImportInvoiceTask struct {
	ID             int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                             // 主键
	OrgID          int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                         // 租户ID
	ClientID       int64  `gorm:"column:client_id;type:int(11);not null;DEFAULT '0'" json:"client_id"`                   // 纳税人ID
	EnterpriseID   string `gorm:"column:enterprise_id;type:varchar(20);not null;DEFAULT ''" json:"enterprise_id"`        // 企业ID
	TaxCode        string `gorm:"column:tax_code;type:varchar(18);not null;" json:"tax_code"`                            // 税号
	FileMd5        string `gorm:"column:file_md5;type:varchar(64);not null;DEFAULT ''" json:"file_md5"`                  // 源文件md5
	SourceFileName string `gorm:"column:source_file_name;type:varchar(255);not null;DEFAULT ''" json:"source_file_name"` // 源文件名称
	FileCosURL     string `gorm:"column:file_cos_url;type:varchar(255);not null;DEFAULT ''" json:"file_cos_url"`         // 源文件cos地址
	TaskStatus     int8   `gorm:"column:task_status;type:tinyint(1);not null;DEFAULT '1'" json:"task_status"`            // 任务状态：1(等待中)；2(成功)；3(失败)
	FailType       int8   `gorm:"column:fail_type;type:tinyint(1);not null;DEFAULT '0'" json:"fail_type"`                // 失败的原因类型；1：文件错误；2：识别/解析错误，允许重试
	FailContent    string `gorm:"column:fail_content;type:varchar(255);not null;DEFAULT ''" json:"fail_content"`         // 失败原因
	RetryNum       int64  `gorm:"column:retry_num;type:int(11);not null;" json:"retry_num"`                              // 重试次数
	CreatedAt      int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                 // 创建时间
	UpdatedAt      int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                 // 更新时间
}

// 获取结构体对应的表名方法
func (m *ImportInvoiceTask) TableName() string {
	return "import_invoice_task"
}

// 实例化结构体对象
func NewImportInvoiceTask() *ImportInvoiceTask {
	return &ImportInvoiceTask{}
}

// 获取主键的对应字段
func (m *ImportInvoiceTask) PrimaryKey() string {
	return ImportInvoiceTaskColumns.ID
}

// 获取主键值
func (m *ImportInvoiceTask) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var ImportInvoiceTaskColumns = struct {
	ID             string
	OrgID          string
	ClientID       string
	EnterpriseID   string
	TaxCode        string
	FileMd5        string
	SourceFileName string
	FileCosURL     string
	TaskStatus     string
	FailType       string
	FailContent    string
	RetryNum       string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	OrgID:          "org_id",
	ClientID:       "client_id",
	EnterpriseID:   "enterprise_id",
	TaxCode:        "tax_code",
	FileMd5:        "file_md5",
	SourceFileName: "source_file_name",
	FileCosURL:     "file_cos_url",
	TaskStatus:     "task_status",
	FailType:       "fail_type",
	FailContent:    "fail_content",
	RetryNum:       "retry_num",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// 包含所有表字段的切片
var ImportInvoiceTaskAllColumns = []string{
	ImportInvoiceTaskColumns.ID,             // 主键
	ImportInvoiceTaskColumns.OrgID,          // 租户ID
	ImportInvoiceTaskColumns.ClientID,       // 纳税人ID
	ImportInvoiceTaskColumns.EnterpriseID,   // 企业ID
	ImportInvoiceTaskColumns.TaxCode,        // 税号
	ImportInvoiceTaskColumns.FileMd5,        // 源文件md5
	ImportInvoiceTaskColumns.SourceFileName, // 源文件名称
	ImportInvoiceTaskColumns.FileCosURL,     // 源文件cos地址
	ImportInvoiceTaskColumns.TaskStatus,     // 任务状态：1(等待中)；2(成功)；3(失败)
	ImportInvoiceTaskColumns.FailType,       // 失败的原因类型；1：文件错误；2：识别/解析错误，允许重试
	ImportInvoiceTaskColumns.FailContent,    // 失败原因
	ImportInvoiceTaskColumns.RetryNum,       // 重试次数
	ImportInvoiceTaskColumns.CreatedAt,      // 创建时间
	ImportInvoiceTaskColumns.UpdatedAt,      // 更新时间

}

// 设置值：主键
func (m *ImportInvoiceTask) SetID(v int64) {
	m.ID = v
}

// 设置值：租户ID
func (m *ImportInvoiceTask) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：纳税人ID
func (m *ImportInvoiceTask) SetClientID(v int64) {
	m.ClientID = v
}

// 设置值：企业ID
func (m *ImportInvoiceTask) SetEnterpriseID(v string) {
	m.EnterpriseID = v
}

// 设置值：税号
func (m *ImportInvoiceTask) SetTaxCode(v string) {
	m.TaxCode = v
}

// 设置值：源文件md5
func (m *ImportInvoiceTask) SetFileMd5(v string) {
	m.FileMd5 = v
}

// 设置值：源文件名称
func (m *ImportInvoiceTask) SetSourceFileName(v string) {
	m.SourceFileName = v
}

// 设置值：源文件cos地址
func (m *ImportInvoiceTask) SetFileCosURL(v string) {
	m.FileCosURL = v
}

// 设置值：任务状态：1(等待中)；2(成功)；3(失败)
func (m *ImportInvoiceTask) SetTaskStatus(v int8) {
	m.TaskStatus = v
}

// 设置值：失败的原因类型；1：文件错误；2：识别/解析错误，允许重试
func (m *ImportInvoiceTask) SetFailType(v int8) {
	m.FailType = v
}

// 设置值：失败原因
func (m *ImportInvoiceTask) SetFailContent(v string) {
	m.FailContent = v
}

// 设置值：重试次数
func (m *ImportInvoiceTask) SetRetryNum(v int64) {
	m.RetryNum = v
}

// 设置值：创建时间
func (m *ImportInvoiceTask) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *ImportInvoiceTask) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 获取值：主键
func (m *ImportInvoiceTask) GetID() int64 {
	return m.ID
}

// 获取值：租户ID
func (m *ImportInvoiceTask) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：纳税人ID
func (m *ImportInvoiceTask) GetClientID() int64 {
	return m.ClientID
}

// 获取值：企业ID
func (m *ImportInvoiceTask) GetEnterpriseID() string {
	return m.EnterpriseID
}

// 获取值：税号
func (m *ImportInvoiceTask) GetTaxCode() string {
	return m.TaxCode
}

// 获取值：源文件md5
func (m *ImportInvoiceTask) GetFileMd5() string {
	return m.FileMd5
}

// 获取值：源文件名称
func (m *ImportInvoiceTask) GetSourceFileName() string {
	return m.SourceFileName
}

// 获取值：源文件cos地址
func (m *ImportInvoiceTask) GetFileCosURL() string {
	return m.FileCosURL
}

// 获取值：任务状态：1(等待中)；2(成功)；3(失败)
func (m *ImportInvoiceTask) GetTaskStatus() int8 {
	return m.TaskStatus
}

// 获取值：失败的原因类型；1：文件错误；2：识别/解析错误，允许重试
func (m *ImportInvoiceTask) GetFailType() int8 {
	return m.FailType
}

// 获取值：失败原因
func (m *ImportInvoiceTask) GetFailContent() string {
	return m.FailContent
}

// 获取值：重试次数
func (m *ImportInvoiceTask) GetRetryNum() int64 {
	return m.RetryNum
}

// 获取值：创建时间
func (m *ImportInvoiceTask) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *ImportInvoiceTask) GetUpdatedAt() int64 {
	return m.UpdatedAt
}
