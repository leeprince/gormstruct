package model

// 银行账号历史月度余额表
type BankHistoryMonthBalance struct {
	ID              uint64 `gorm:"column:id;primaryKey;type:bigint(20) unsigned;autoIncrement" json:"id"`                  // 主键
	OrgID           int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                          // 租户ID
	BankNumber      string `gorm:"column:bank_number;type:varchar(50);not null;DEFAULT ''" json:"bank_number"`             // 银行卡号
	BankAccountName string `gorm:"column:bank_account_name;type:varchar(64);not null;DEFAULT ''" json:"bank_account_name"` // 银行账户名称
	StatisticsDate  int64  `gorm:"column:statistics_date;type:int(11);not null;DEFAULT '0'" json:"statistics_date"`        // 统计:年月,202308
	Balance         int64  `gorm:"column:balance;type:bigint(20);not null;DEFAULT '0'" json:"balance"`                     // 资产总额（余额），单位分
	CreatedAt       int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                  // 创建时间
}

// 获取结构体对应的表名方法
func (m *BankHistoryMonthBalance) TableName() string {
	return "bank_history_month_balance"
}

// 实例化结构体对象
func NewBankHistoryMonthBalance() *BankHistoryMonthBalance {
	return &BankHistoryMonthBalance{}
}

// 获取主键的对应字段
func (m *BankHistoryMonthBalance) PrimaryKey() string {
	return BankHistoryMonthBalanceColumns.ID
}

// 获取主键值
func (m *BankHistoryMonthBalance) PrimaryKeyValue() uint64 {
	return m.ID
}

// 表字段的映射
var BankHistoryMonthBalanceColumns = struct {
	ID              string
	OrgID           string
	BankNumber      string
	BankAccountName string
	StatisticsDate  string
	Balance         string
	CreatedAt       string
}{
	ID:              "id",
	OrgID:           "org_id",
	BankNumber:      "bank_number",
	BankAccountName: "bank_account_name",
	StatisticsDate:  "statistics_date",
	Balance:         "balance",
	CreatedAt:       "created_at",
}

// 包含所有表字段的切片
var BankHistoryMonthBalanceAllColumns = []string{
	BankHistoryMonthBalanceColumns.ID,              // 主键
	BankHistoryMonthBalanceColumns.OrgID,           // 租户ID
	BankHistoryMonthBalanceColumns.BankNumber,      // 银行卡号
	BankHistoryMonthBalanceColumns.BankAccountName, // 银行账户名称
	BankHistoryMonthBalanceColumns.StatisticsDate,  // 统计:年月,202308
	BankHistoryMonthBalanceColumns.Balance,         // 资产总额（余额），单位分
	BankHistoryMonthBalanceColumns.CreatedAt,       // 创建时间

}

// 设置值：主键
func (m *BankHistoryMonthBalance) SetID(v uint64) {
	m.ID = v
}

// 设置值：租户ID
func (m *BankHistoryMonthBalance) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：银行卡号
func (m *BankHistoryMonthBalance) SetBankNumber(v string) {
	m.BankNumber = v
}

// 设置值：银行账户名称
func (m *BankHistoryMonthBalance) SetBankAccountName(v string) {
	m.BankAccountName = v
}

// 设置值：统计:年月,202308
func (m *BankHistoryMonthBalance) SetStatisticsDate(v int64) {
	m.StatisticsDate = v
}

// 设置值：资产总额（余额），单位分
func (m *BankHistoryMonthBalance) SetBalance(v int64) {
	m.Balance = v
}

// 设置值：创建时间
func (m *BankHistoryMonthBalance) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 获取值：主键
func (m *BankHistoryMonthBalance) GetID() uint64 {
	return m.ID
}

// 获取值：租户ID
func (m *BankHistoryMonthBalance) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：银行卡号
func (m *BankHistoryMonthBalance) GetBankNumber() string {
	return m.BankNumber
}

// 获取值：银行账户名称
func (m *BankHistoryMonthBalance) GetBankAccountName() string {
	return m.BankAccountName
}

// 获取值：统计:年月,202308
func (m *BankHistoryMonthBalance) GetStatisticsDate() int64 {
	return m.StatisticsDate
}

// 获取值：资产总额（余额），单位分
func (m *BankHistoryMonthBalance) GetBalance() int64 {
	return m.Balance
}

// 获取值：创建时间
func (m *BankHistoryMonthBalance) GetCreatedAt() int64 {
	return m.CreatedAt
}
