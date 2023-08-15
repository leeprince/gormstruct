package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-15 10:39:21
 * @Desc:   fund_change_event 表
 */

// 资金变更事件表
type FundChangeEvent struct {
	ID                         int64  `gorm:"column:id;primaryKey;type:int(11);" json:"id"`                                                                   // 主键
	EventID                    string `gorm:"column:event_id;type:varchar(64);not null;DEFAULT ''" json:"event_id"`                                           // 事件ID
	EventType                  string `gorm:"column:event_type;type:varchar(64);not null;DEFAULT ''" json:"event_type"`                                       // 事件类型，区分大小写：Income(收入)、Expenditure(支出)
	EventMsg                   string `gorm:"column:event_msg;type:varchar(255);not null;DEFAULT ''" json:"event_msg"`                                        // 事件说明
	OrderID                    string `gorm:"column:order_id;type:varchar(255);not null;DEFAULT ''" json:"order_id"`                                          // 订单ID
	OrgID                      int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                                                  // 租户ID
	ExpenditureBankNumber      string `gorm:"column:expenditure_bank_number;type:varchar(50);not null;DEFAULT ''" json:"expenditure_bank_number"`             // 付款银行卡号
	ExpenditureBankAccountName string `gorm:"column:expenditure_bank_account_name;type:varchar(64);not null;DEFAULT ''" json:"expenditure_bank_account_name"` // 付款银行账户名称
	ExpenditureBankName        string `gorm:"column:expenditure_bank_name;type:varchar(32);not null;DEFAULT ''" json:"expenditure_bank_name"`                 // 付款银行名称
	IncomeBankNumber           string `gorm:"column:income_bank_number;type:varchar(50);not null;DEFAULT ''" json:"income_bank_number"`                       // 收款银行卡号
	IncomeBankAccountName      string `gorm:"column:income_bank_account_name;type:varchar(64);not null;DEFAULT ''" json:"income_bank_account_name"`           // 收款银行账户名称
	IncomeBankName             string `gorm:"column:income_bank_name;type:varchar(32);not null;DEFAULT ''" json:"income_bank_name"`                           // 收款银行名称
	TradingAmount              int64  `gorm:"column:trading_amount;type:bigint(20);not null;DEFAULT '0'" json:"trading_amount"`                               // 交易金额，单位分
	TradingTime                int64  `gorm:"column:trading_time;type:bigint(20);not null;DEFAULT '0'" json:"trading_time"`                                   // 交易时间戳
	CreatedAt                  int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                                          // 事件创建时间戳
}

// 获取结构体对应的表名方法
func (m *FundChangeEvent) TableName() string {
	return "fund_change_event"
}

// 实例化结构体对象
func NewFundChangeEvent() *FundChangeEvent {
	return &FundChangeEvent{}
}

// 获取主键的对应字段
func (m *FundChangeEvent) PrimaryKey() string {
	return FundChangeEventColumns.ID
}

// 获取主键值
func (m *FundChangeEvent) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var FundChangeEventColumns = struct {
	ID                         string
	EventID                    string
	EventType                  string
	EventMsg                   string
	OrderID                    string
	OrgID                      string
	ExpenditureBankNumber      string
	ExpenditureBankAccountName string
	ExpenditureBankName        string
	IncomeBankNumber           string
	IncomeBankAccountName      string
	IncomeBankName             string
	TradingAmount              string
	TradingTime                string
	CreatedAt                  string
}{
	ID:                         "id",
	EventID:                    "event_id",
	EventType:                  "event_type",
	EventMsg:                   "event_msg",
	OrderID:                    "order_id",
	OrgID:                      "org_id",
	ExpenditureBankNumber:      "expenditure_bank_number",
	ExpenditureBankAccountName: "expenditure_bank_account_name",
	ExpenditureBankName:        "expenditure_bank_name",
	IncomeBankNumber:           "income_bank_number",
	IncomeBankAccountName:      "income_bank_account_name",
	IncomeBankName:             "income_bank_name",
	TradingAmount:              "trading_amount",
	TradingTime:                "trading_time",
	CreatedAt:                  "created_at",
}

// 包含所有表字段的切片
var FundChangeEventAllColumns = []string{
	FundChangeEventColumns.ID,                         // 主键
	FundChangeEventColumns.EventID,                    // 事件ID
	FundChangeEventColumns.EventType,                  // 事件类型，区分大小写：Income(收入)、Expenditure(支出)
	FundChangeEventColumns.EventMsg,                   // 事件说明
	FundChangeEventColumns.OrderID,                    // 订单ID
	FundChangeEventColumns.OrgID,                      // 租户ID
	FundChangeEventColumns.ExpenditureBankNumber,      // 付款银行卡号
	FundChangeEventColumns.ExpenditureBankAccountName, // 付款银行账户名称
	FundChangeEventColumns.ExpenditureBankName,        // 付款银行名称
	FundChangeEventColumns.IncomeBankNumber,           // 收款银行卡号
	FundChangeEventColumns.IncomeBankAccountName,      // 收款银行账户名称
	FundChangeEventColumns.IncomeBankName,             // 收款银行名称
	FundChangeEventColumns.TradingAmount,              // 交易金额，单位分
	FundChangeEventColumns.TradingTime,                // 交易时间戳
	FundChangeEventColumns.CreatedAt,                  // 事件创建时间戳

}

// 设置值：主键
func (m *FundChangeEvent) SetID(v int64) {
	m.ID = v
}

// 设置值：事件ID
func (m *FundChangeEvent) SetEventID(v string) {
	m.EventID = v
}

// 设置值：事件类型，区分大小写：Income(收入)、Expenditure(支出)
func (m *FundChangeEvent) SetEventType(v string) {
	m.EventType = v
}

// 设置值：事件说明
func (m *FundChangeEvent) SetEventMsg(v string) {
	m.EventMsg = v
}

// 设置值：订单ID
func (m *FundChangeEvent) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：租户ID
func (m *FundChangeEvent) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：付款银行卡号
func (m *FundChangeEvent) SetExpenditureBankNumber(v string) {
	m.ExpenditureBankNumber = v
}

// 设置值：付款银行账户名称
func (m *FundChangeEvent) SetExpenditureBankAccountName(v string) {
	m.ExpenditureBankAccountName = v
}

// 设置值：付款银行名称
func (m *FundChangeEvent) SetExpenditureBankName(v string) {
	m.ExpenditureBankName = v
}

// 设置值：收款银行卡号
func (m *FundChangeEvent) SetIncomeBankNumber(v string) {
	m.IncomeBankNumber = v
}

// 设置值：收款银行账户名称
func (m *FundChangeEvent) SetIncomeBankAccountName(v string) {
	m.IncomeBankAccountName = v
}

// 设置值：收款银行名称
func (m *FundChangeEvent) SetIncomeBankName(v string) {
	m.IncomeBankName = v
}

// 设置值：交易金额，单位分
func (m *FundChangeEvent) SetTradingAmount(v int64) {
	m.TradingAmount = v
}

// 设置值：交易时间戳
func (m *FundChangeEvent) SetTradingTime(v int64) {
	m.TradingTime = v
}

// 设置值：事件创建时间戳
func (m *FundChangeEvent) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 获取值：主键
func (m *FundChangeEvent) GetID() int64 {
	return m.ID
}

// 获取值：事件ID
func (m *FundChangeEvent) GetEventID() string {
	return m.EventID
}

// 获取值：事件类型，区分大小写：Income(收入)、Expenditure(支出)
func (m *FundChangeEvent) GetEventType() string {
	return m.EventType
}

// 获取值：事件说明
func (m *FundChangeEvent) GetEventMsg() string {
	return m.EventMsg
}

// 获取值：订单ID
func (m *FundChangeEvent) GetOrderID() string {
	return m.OrderID
}

// 获取值：租户ID
func (m *FundChangeEvent) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：付款银行卡号
func (m *FundChangeEvent) GetExpenditureBankNumber() string {
	return m.ExpenditureBankNumber
}

// 获取值：付款银行账户名称
func (m *FundChangeEvent) GetExpenditureBankAccountName() string {
	return m.ExpenditureBankAccountName
}

// 获取值：付款银行名称
func (m *FundChangeEvent) GetExpenditureBankName() string {
	return m.ExpenditureBankName
}

// 获取值：收款银行卡号
func (m *FundChangeEvent) GetIncomeBankNumber() string {
	return m.IncomeBankNumber
}

// 获取值：收款银行账户名称
func (m *FundChangeEvent) GetIncomeBankAccountName() string {
	return m.IncomeBankAccountName
}

// 获取值：收款银行名称
func (m *FundChangeEvent) GetIncomeBankName() string {
	return m.IncomeBankName
}

// 获取值：交易金额，单位分
func (m *FundChangeEvent) GetTradingAmount() int64 {
	return m.TradingAmount
}

// 获取值：交易时间戳
func (m *FundChangeEvent) GetTradingTime() int64 {
	return m.TradingTime
}

// 获取值：事件创建时间戳
func (m *FundChangeEvent) GetCreatedAt() int64 {
	return m.CreatedAt
}
