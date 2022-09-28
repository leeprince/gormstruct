package model

// 到账消息记录表
type AccountMessage struct {
	ID               int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                      // 主键ID,消息ID
	OrgID            int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                  // 租户ID
	ClientID         int64  `gorm:"column:client_id;type:int(11);not null;" json:"client_id"`                       // 商户ID,租户ID
	ClientName       string `gorm:"column:client_name;type:varchar(64);not null;DEFAULT ''" json:"client_name"`     // 商户名称
	OrderID          string `gorm:"column:order_id;type:varchar(32);not null;DEFAULT ''" json:"order_id"`           // 高灯云订单号
	MessageType      int8   `gorm:"column:message_type;type:tinyint(1);not null;DEFAULT '1'" json:"message_type"`   // 消息类型(1:到账; 2:系统; 3:活动)；固定为1
	RealAmount       int64  `gorm:"column:real_amount;type:int(11);not null;" json:"real_amount"`                   // 实付金额。单位分
	Status           int8   `gorm:"column:status;type:tinyint(1);not null;" json:"status"`                          // 打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功
	FailReason       string `gorm:"column:fail_reason;type:varchar(255);not null;" json:"fail_reason"`              // 失败原因
	PayeeBankAccount string `gorm:"column:payee_bank_account;type:varchar(32);not null;" json:"payee_bank_account"` // 打款银行卡号
	IsRead           int8   `gorm:"column:is_read;type:tinyint(1);not null;DEFAULT '0'" json:"is_read"`             // 消息是否已读；0：未读；1：已读
	CreatedAt        int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`          // 创建时间
	UpdatedAt        int64  `gorm:"column:updated_at;type:int(11);not null;" json:"updated_at"`                     // 更新时间
	DeletedAt        int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`          // 删除时间
}

// 获取结构体对应的表名方法
func (m *AccountMessage) TableName() string {
	return "account_message"
}

// 实例化结构体对象
func NewAccountMessage() *AccountMessage {
	return &AccountMessage{}
}

// 获取主键的对应字段
func (m *AccountMessage) PrimaryKey() string {
	return AccountMessageColumns.ID
}

// 获取主键值
func (m *AccountMessage) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var AccountMessageColumns = struct {
	ID               string
	OrgID            string
	ClientID         string
	ClientName       string
	OrderID          string
	MessageType      string
	RealAmount       string
	Status           string
	FailReason       string
	PayeeBankAccount string
	IsRead           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
}{
	ID:               "id",
	OrgID:            "org_id",
	ClientID:         "client_id",
	ClientName:       "client_name",
	OrderID:          "order_id",
	MessageType:      "message_type",
	RealAmount:       "real_amount",
	Status:           "status",
	FailReason:       "fail_reason",
	PayeeBankAccount: "payee_bank_account",
	IsRead:           "is_read",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// 包含所有表字段的切片
var AccountMessageAllColumns = []string{
	AccountMessageColumns.ID,               // 主键ID,消息ID
	AccountMessageColumns.OrgID,            // 租户ID
	AccountMessageColumns.ClientID,         // 商户ID,租户ID
	AccountMessageColumns.ClientName,       // 商户名称
	AccountMessageColumns.OrderID,          // 高灯云订单号
	AccountMessageColumns.MessageType,      // 消息类型(1:到账; 2:系统; 3:活动)；固定为1
	AccountMessageColumns.RealAmount,       // 实付金额。单位分
	AccountMessageColumns.Status,           // 打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功
	AccountMessageColumns.FailReason,       // 失败原因
	AccountMessageColumns.PayeeBankAccount, // 打款银行卡号
	AccountMessageColumns.IsRead,           // 消息是否已读；0：未读；1：已读
	AccountMessageColumns.CreatedAt,        // 创建时间
	AccountMessageColumns.UpdatedAt,        // 更新时间
	AccountMessageColumns.DeletedAt,        // 删除时间

}

// 设置值：主键ID,消息ID
func (m *AccountMessage) SetID(v int64) {
	m.ID = v
}

// 设置值：租户ID
func (m *AccountMessage) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：商户ID,租户ID
func (m *AccountMessage) SetClientID(v int64) {
	m.ClientID = v
}

// 设置值：商户名称
func (m *AccountMessage) SetClientName(v string) {
	m.ClientName = v
}

// 设置值：高灯云订单号
func (m *AccountMessage) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：消息类型(1:到账; 2:系统; 3:活动)；固定为1
func (m *AccountMessage) SetMessageType(v int8) {
	m.MessageType = v
}

// 设置值：实付金额。单位分
func (m *AccountMessage) SetRealAmount(v int64) {
	m.RealAmount = v
}

// 设置值：打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功
func (m *AccountMessage) SetStatus(v int8) {
	m.Status = v
}

// 设置值：失败原因
func (m *AccountMessage) SetFailReason(v string) {
	m.FailReason = v
}

// 设置值：打款银行卡号
func (m *AccountMessage) SetPayeeBankAccount(v string) {
	m.PayeeBankAccount = v
}

// 设置值：消息是否已读；0：未读；1：已读
func (m *AccountMessage) SetIsRead(v int8) {
	m.IsRead = v
}

// 设置值：创建时间
func (m *AccountMessage) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *AccountMessage) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *AccountMessage) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键ID,消息ID
func (m *AccountMessage) GetID() int64 {
	return m.ID
}

// 获取值：租户ID
func (m *AccountMessage) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：商户ID,租户ID
func (m *AccountMessage) GetClientID() int64 {
	return m.ClientID
}

// 获取值：商户名称
func (m *AccountMessage) GetClientName() string {
	return m.ClientName
}

// 获取值：高灯云订单号
func (m *AccountMessage) GetOrderID() string {
	return m.OrderID
}

// 获取值：消息类型(1:到账; 2:系统; 3:活动)；固定为1
func (m *AccountMessage) GetMessageType() int8 {
	return m.MessageType
}

// 获取值：实付金额。单位分
func (m *AccountMessage) GetRealAmount() int64 {
	return m.RealAmount
}

// 获取值：打款状态。-1：打款失败； 0：打款中； 1：打款成功； 2：打款退票； 3：冻结出款；仅记录失败和成功
func (m *AccountMessage) GetStatus() int8 {
	return m.Status
}

// 获取值：失败原因
func (m *AccountMessage) GetFailReason() string {
	return m.FailReason
}

// 获取值：打款银行卡号
func (m *AccountMessage) GetPayeeBankAccount() string {
	return m.PayeeBankAccount
}

// 获取值：消息是否已读；0：未读；1：已读
func (m *AccountMessage) GetIsRead() int8 {
	return m.IsRead
}

// 获取值：创建时间
func (m *AccountMessage) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *AccountMessage) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *AccountMessage) GetDeletedAt() int64 {
	return m.DeletedAt
}
